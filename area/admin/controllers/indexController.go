package controllers

import (
	"fasthttpweb/area"
	av "fasthttpweb/area/admin/views/index"

	//	"github.com/kataras/go-sessions"
)

const (
	areaName = "Admin"
)

type IndexController struct {
	*area.BaseController
}

func init() {
	c := &IndexController{BaseController: &area.BaseController{}}
	c.RegistRoutes(areaName, c)
	c.RegistRoute("/admin", "get", "/admin/index/index")
	c.RegistRoute("/admin/index", "get", "/admin/index/index")
}

func (c *IndexController) Index() {
	bpd := c.InitBasePageData(areaName, "Index", "admin-title", "admin-kwd")
	bp := &area.BasePage{CTX: c.Ctx, BPD: bpd}
	ip := &av.IndexPage{bp}
	c.View(ip, "text/html")
}
