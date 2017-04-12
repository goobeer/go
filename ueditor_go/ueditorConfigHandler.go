package ueditor_go

type UeditorConfigHandler struct {
	*UEditorHandler
}

func (h *UeditorConfigHandler) Process() {
	responseData, _ := BuildItems()
	h.WriteJson(responseData)
}
