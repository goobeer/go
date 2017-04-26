package controllers

import (
	"fasthttpweb/area"
	av "fasthttpweb/area/admin/views/index"

	//	"github.com/kataras/go-sessions"
)

type IndexController struct {
	*area.BaseController
}

func init() {
	c := &IndexController{BaseController: &area.BaseController{}}
	c.RegistRoutes(areaName, c)
	c.RegistRoute("get", "/admin/index/index", "/admin", "/admin/index")
}

func (c *IndexController) Index() {
	bpd := area.NewBasePageData(areaName, "article-title", "article-kwd", c)
	bp := area.NewBasePage(c.Ctx, bpd)
	ip := &av.IndexPage{bp}
	c.View(ip, "text/html")
}
