package ueditor_go

type UeditorNotSupportedRespData struct {
	State string `json:"state"`
}

type UeditorNotSupportedHandler struct {
	*UEditorHandler
}

func (u *UeditorNotSupportedHandler) Process() {
	TD := &UeditorNotSupportedRespData{State: "action 参数为空或者 action 不被支持。"}
	u.WriteJson(TD)
}
