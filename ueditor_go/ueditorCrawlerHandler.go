package ueditor_go

type UeditorCrawlerHandler struct {
	Sources  []string
	Crawlers []UeditorCrawler
	*UEditorHandler
}

func (u *UeditorCrawlerHandler) Process() {
	source := u.Ctx.FormValue("source[]")
	if source != nil && len(source) > 0 {
		for _, v := range u.Crawlers {
			v.Fetch()
		}
		var TD struct {
			State string           `json:"state"`
			List  []UeditorCrawler `json:"list"`
		}
		TD.State = "SUCCESS"
		TD.List = u.Crawlers
		u.WriteJson(TD)
	} else {
		var TD struct{ State string }
		TD.State = "参数错误：没有指定抓取源"
		u.WriteJson(TD)
	}
}
