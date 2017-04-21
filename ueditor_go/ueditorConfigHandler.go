package ueditor_go

type UeditorConfigHandler struct {
	*UEditorHandler
}

func (h *UeditorConfigHandler) Process() {
	responseData, err := BuildItems()
	if err != nil {
		panic(err)
	}
	h.WriteJson(responseData)
}
