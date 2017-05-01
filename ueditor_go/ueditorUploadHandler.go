package ueditor_go

import (
	"encoding/base64"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/valyala/fasthttp"
)

const (
	Success         = 0
	SizeLimitExceed = -1
	TypeNotAllow    = -2
	FileAccessError = -3
	NetworkError    = -4
	InvalidParam    = -5
	AuthorizError   = -6
	IOError         = -7
	PathNotFound    = -8
	Unknown         = 1
)

type UeditorUploadConfig struct {
	PathFormat      string
	UploadFieldName string
	SizeLimit       int64
	AllowExtensions []string
	Base64          bool
	Base64Filename  string
}

type UeditorUploadResult struct {
	State          int
	Url            string
	OriginFileName string
	ErrorMessage   string
}

type UeditorUploadHandler struct {
	*UeditorUploadConfig
	*UeditorUploadResult
	*UEditorHandler
}

func NewUeditorUploadHandler(ctx *fasthttp.RequestCtx) *UeditorUploadHandler {
	uploadHandler := &UeditorUploadHandler{}
	return uploadHandler
}

func (u *UeditorUploadHandler) Process() {
	var uploadFileBytes []byte
	var uploadFileName string
	TD := &UeditorUploadImageResponseData{}
	if u.UeditorUploadConfig.Base64 {
		uploadFileName = u.Base64Filename
		fileBytes := u.Ctx.FormValue(u.UploadFieldName)
		if !(fileBytes != nil && len(fileBytes) > 0) {
			u.State = InvalidParam

			TD.State = u.getStateMessage(u.State)
			TD.Url = string(u.Ctx.RequestURI())
			TD.Original = uploadFileName
			TD.Title = uploadFileName
			u.WriteJson(TD)
			return
		}
		base64.StdEncoding.Decode(uploadFileBytes, fileBytes)
	} else {
		fh, _ := u.Ctx.FormFile(u.UploadFieldName)
		if fh == nil {
			u.State = InvalidParam

			TD.State = u.getStateMessage(u.State)
			TD.Url = string(u.Ctx.RequestURI())
			TD.Original = uploadFileName
			TD.Title = uploadFileName
			u.WriteJson(TD)
			return
		}
		uploadFileName = fh.Filename

		if !u.checkFileType(uploadFileName) {
			u.State = TypeNotAllow

			TD.State = u.getStateMessage(u.State)
			TD.Url = string(u.Ctx.RequestURI())
			TD.Original = uploadFileName
			TD.Title = uploadFileName
			u.WriteJson(TD)
			return
		}

		TD.Original = uploadFileName
		TD.Title = uploadFileName
		savePath := UeditorPathFormatter(uploadFileName, u.PathFormat)
		baseDir := path.Dir(savePath)
		_, err := os.Stat(baseDir)

		mutex := &sync.RWMutex{}
		mutex.Lock()
		if err != nil && os.IsNotExist(err) {
			err = os.MkdirAll(baseDir, 0666)
		}
		mutex.Unlock()

		if err != nil {
			TD.Error = err.Error()
			TD.State = u.getStateMessage(FileAccessError)
		} else {
			err = fasthttp.SaveMultipartFile(fh, savePath)
			if err == nil {
				TD.Url = savePath
				TD.State = u.getStateMessage(Success)
			} else {
				TD.Error = err.Error()
				TD.State = u.getStateMessage(FileAccessError)
			}
		}

		u.WriteJson(TD)
	}
}

func (u *UeditorUploadHandler) getStateMessage(state int) (stateStr string) {
	switch state {
	case Success:
		stateStr = "SUCCESS"
	case FileAccessError:
		stateStr = "文件访问出错，请检查写入权限"
	case SizeLimitExceed:
		stateStr = "文件大小超出服务器限制"
	case TypeNotAllow:
		stateStr = "不允许的文件格式"
	case NetworkError:
		stateStr = "网络错误"
	default:
		stateStr = "未知错误"
	}
	return
}

func (u *UeditorUploadHandler) checkFileType(filename string) bool {
	ext := strings.ToLower(path.Ext(filename))
	for _, v := range u.UeditorUploadConfig.AllowExtensions {
		if strings.ToLower(v) == ext {
			return true
		}
	}
	return false
}
