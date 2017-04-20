package ueditor_go

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

var (
	config = &UeditorConfig{}
)

type IUeditorHandler interface {
	Process()
	//	WriteJson(response interface{})
}

type UEditorHandler struct {
	Ctx fasthttp.RequestCtx
}

func (h *UEditorHandler) Process(ctx fasthttp.RequestCtx) {
	action := ctx.FormValue("action")
	var ueditorHandle IUeditorHandler
	if action != nil && len(action) > 0 {
		actionStr := string(action)
		switch actionStr {
		case "config":
			ueditorHandle = new(UeditorConfigHandler)
		case "uploadimage":
			ueditorHandle = &UeditorUploadHandler{UEditorHandler: &UEditorHandler{Ctx: ctx}, UeditorUploadConfig: &UeditorUploadConfig{AllowExtensions: config.ImageAllowFiles, PathFormat: config.ImagePathFormat, SizeLimit: config.ImageMaxSize, UploadFieldName: config.ImageFieldName}}
		case "uploadscrawl":
			ueditorHandle = &UeditorUploadHandler{UEditorHandler: &UEditorHandler{Ctx: ctx}, UeditorUploadConfig: &UeditorUploadConfig{AllowExtensions: []string{".png"}, PathFormat: config.ScrawlPathFormat, SizeLimit: config.ScrawlMaxSize, UploadFieldName: config.ScrawlFieldName, Base64: true, Base64Filename: "scrawl.png"}}
		case "uploadvideo":
			ueditorHandle = &UeditorUploadHandler{UEditorHandler: &UEditorHandler{Ctx: ctx}, UeditorUploadConfig: &UeditorUploadConfig{AllowExtensions: config.VideoAllowFiles, PathFormat: config.VideoPathFormat, SizeLimit: config.VideoMaxSize, UploadFieldName: config.VideoFieldName}}
		case "uploadfile":
			ueditorHandle = &UeditorUploadHandler{UEditorHandler: &UEditorHandler{Ctx: ctx}, UeditorUploadConfig: &UeditorUploadConfig{AllowExtensions: config.FileAllowFiles, PathFormat: config.FilePathFormat, SizeLimit: config.FileMaxSize, UploadFieldName: config.FileFieldName}}
		case "listimage":
			//			ueditorHandle = &UeditorL
		case "listfile":
		case "catchimage":
			ueditorHandle = &UeditorCrawler{UEditorHandler: &UEditorHandler{Ctx: ctx}}
		default:
			ueditorHandle = &UeditorNotSupportedHandler{UEditorHandler: &UEditorHandler{Ctx: ctx}}
		}
		ueditorHandle.Process()
	}
}

func (h *UEditorHandler) WriteJson(response interface{}) {
	jsonpCallback := h.Ctx.QueryArgs().Peek("callback")

	reponseJson, err := json.Marshal(response)
	if err != nil {
		h.Ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	if jsonpCallback != nil && len(jsonpCallback) > 0 {
		h.Ctx.SetContentType("text/plain")
		h.Ctx.Write(reponseJson)
	} else {
		h.Ctx.SetContentType("application/javascript")
		jsonpCallback = append(jsonpCallback, []byte("{")...)
		jsonpCallback = append(jsonpCallback, reponseJson...)
		jsonpCallback = append(jsonpCallback, []byte("}")...)
		h.Ctx.Write(jsonpCallback)
	}

}
