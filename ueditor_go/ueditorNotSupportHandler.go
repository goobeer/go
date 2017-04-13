package ueditor_go

type UeditorNotSupportedHandler struct {
	*UEditorHandler
}

func (u *UeditorNotSupportedHandler) Process() {
	var TD struct{ State string }
	TD.State = "action 参数为空或者 action 不被支持。"
	u.WriteJson(TD)
}
