package ueditor_go

import (
	"encoding/base64"
	"path"
	"strings"
)

const (
	Success         = 0
	SizeLimitExceed = -1
	TypeNotAllow    = -2
	FileAccessError = -3
	NetworkError    = -4
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

func (u *UeditorUploadHandler) Process() {
	var uploadFileBytes []byte
	var uploadFileName string
	if u.UeditorUploadConfig.Base64 {
		uploadFileName = u.Base64Filename
		fileBytes := u.Ctx.FormValue(u.UploadFieldName)
		base64.StdEncoding.Decode(uploadFileBytes, fileBytes)
	} else {
		//TODO xxx
	}
	//        if (UploadConfig.Base64)
	//        {
	//            uploadFileName = UploadConfig.Base64Filename;
	//            uploadFileBytes = Convert.FromBase64String(Request[UploadConfig.UploadFieldName]);
	//        }
	//        else
	//        {
	//            var file = Request.Files[UploadConfig.UploadFieldName];
	//            uploadFileName = file.FileName;

	//            if (!CheckFileType(uploadFileName))
	//            {
	//                Result.State = UploadState.TypeNotAllow;
	//                WriteResult();
	//                return;
	//            }
	//            if (!CheckFileSize(file.ContentLength))
	//            {
	//                Result.State = UploadState.SizeLimitExceed;
	//                WriteResult();
	//                return;
	//            }

	//            uploadFileBytes = new byte[file.ContentLength];
	//            try
	//            {
	//                file.InputStream.Read(uploadFileBytes, 0, file.ContentLength);
	//            }
	//            catch (Exception)
	//            {
	//                Result.State = UploadState.NetworkError;
	//                WriteResult();
	//            }
	//        }

	//        Result.OriginFileName = uploadFileName;

	//        var savePath = PathFormatter.Format(uploadFileName, UploadConfig.PathFormat);
	//        var localPath = Server.MapPath(savePath);
	//        try
	//        {
	//            if (!Directory.Exists(Path.GetDirectoryName(localPath)))
	//            {
	//                Directory.CreateDirectory(Path.GetDirectoryName(localPath));
	//            }
	//            File.WriteAllBytes(localPath, uploadFileBytes);
	//            Result.Url = savePath;
	//            Result.State = UploadState.Success;
	//        }
	//        catch (Exception e)
	//        {
	//            Result.State = UploadState.FileAccessError;
	//            Result.ErrorMessage = e.Message;
	//        }
	//        finally
	//        {
	//            WriteResult();
	//        }
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

func (u *UeditorUploadHandler) checkFileSize(size int64) bool {
	return size < u.UeditorUploadConfig.SizeLimit
}
