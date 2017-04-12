package ueditor_go

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type IUeditorHandler interface {
	Process()
	WriteJson(response interface{})
}

type UEditorHandler struct {
	Ctx fasthttp.RequestCtx
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
