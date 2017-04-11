package controllers

import (
	"fasthttpweb/area"
	au "fasthttpweb/area/admin/views/user"

	//	"github.com/kataras/go-sessions"
)

type userController struct {
	*area.BaseController
}

func init() {
	c := &userController{BaseController: &area.BaseController{}}
	c.RegistRoutes(areaName, c)
	c.RegistRoute("get", "/admin/user/index", "/admin/user")
}

func (c *userController) Index() {
	bpd := c.InitBasePageData(areaName, "Index", "admin-user-title", "admin-user-kwd")
	bp := &area.BasePage{CTX: c.Ctx, BPD: bpd}
	ip := &au.IndexPage{bp}
	c.View(ip, "text/html")
}
