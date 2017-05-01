package ueditor_go

import (
	"encoding/json"
	"fasthttpweb/router"

	"github.com/valyala/fasthttp"
)

var (
	config, _ = BuildItems()
)

type UeditorUploadImageResponseData struct {
	State    string `json:"state"`
	Url      string `json:"url"`
	Title    string `json:"title"`
	Original string `json:"original"`
	Error    string `json:"error"`
}

type IUeditorHandler interface {
	Process()
}

type UEditorHandler struct {
	Ctx *fasthttp.RequestCtx
}

func init() {
	//注册 url

	urlUeditor := "/ueditor/net/controller.ashx"
	router.R.GET(urlUeditor, ueditorRequestHandle)

	router.R.POST(urlUeditor, ueditorRequestHandle)
}

func ueditorRequestHandle(ctx *fasthttp.RequestCtx) {
	ueditor := &UEditorHandler{Ctx: ctx}
	ueditor.Process(ctx)
}

func (h *UEditorHandler) Process(ctx *fasthttp.RequestCtx) {
	action := ctx.FormValue("action")
	var ueditorHandle IUeditorHandler
	if action != nil && len(action) > 0 {
		actionStr := string(action)
		switch actionStr {
		case "config":
			ueditorHandle = &UeditorConfigHandler{UEditorHandler: &UEditorHandler{Ctx: ctx}}
		case "uploadimage":
			ueditorHandle = &UeditorUploadHandler{UEditorHandler: &UEditorHandler{Ctx: ctx}, UeditorUploadConfig: &UeditorUploadConfig{AllowExtensions: config.ImageAllowFiles, PathFormat: config.ImagePathFormat, SizeLimit: config.ImageMaxSize, UploadFieldName: config.ImageFieldName}, UeditorUploadResult: &UeditorUploadResult{}}
		case "uploadscrawl":
			ueditorHandle = &UeditorUploadHandler{UEditorHandler: &UEditorHandler{Ctx: ctx}, UeditorUploadConfig: &UeditorUploadConfig{AllowExtensions: []string{".png"}, PathFormat: config.ScrawlPathFormat, SizeLimit: config.ScrawlMaxSize, UploadFieldName: config.ScrawlFieldName, Base64: true, Base64Filename: "scrawl.png"}, UeditorUploadResult: &UeditorUploadResult{}}
		case "uploadvideo":
			ueditorHandle = &UeditorUploadHandler{UEditorHandler: &UEditorHandler{Ctx: ctx}, UeditorUploadConfig: &UeditorUploadConfig{AllowExtensions: config.VideoAllowFiles, PathFormat: config.VideoPathFormat, SizeLimit: config.VideoMaxSize, UploadFieldName: config.VideoFieldName}, UeditorUploadResult: &UeditorUploadResult{}}
		case "uploadfile":
			ueditorHandle = &UeditorUploadHandler{UEditorHandler: &UEditorHandler{Ctx: ctx}, UeditorUploadConfig: &UeditorUploadConfig{AllowExtensions: config.FileAllowFiles, PathFormat: config.FilePathFormat, SizeLimit: config.FileMaxSize, UploadFieldName: config.FileFieldName}, UeditorUploadResult: &UeditorUploadResult{}}
		case "listimage":
			ueditorHandle = &UeditorListFileManager{UEditorHandler: &UEditorHandler{Ctx: ctx}, PathToList: config.ImageManagerListPath, SearchExtensions: config.ImageManagerAllowFiles}
		case "listfile":
			ueditorHandle = &UeditorListFileManager{UEditorHandler: &UEditorHandler{Ctx: ctx}, PathToList: config.FileManagerListPath, SearchExtensions: config.FileManagerAllowFiles}
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
		jsonpCallback = append(jsonpCallback, reponseJson...)
		h.Ctx.Write(jsonpCallback)
	}

}
