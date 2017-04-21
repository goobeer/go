package controllers

import (
	"fasthttpweb/area"
	view "fasthttpweb/area/admin/views/article"
	"fmt"
)

type ArticleController struct {
	*area.BaseController
}

func init() {
	c := &ArticleController{BaseController: &area.BaseController{}}
	c.RegistRoutes(areaName, c)
	c.RegistRoute("get", "/admin/article/index", "/admin/article")
}

func (c *ArticleController) Index() {
	bpd := c.InitBasePageData(areaName, "Index", "article-title", "article-kwd")
	bp := &area.BasePage{CTX: c.Ctx, BPD: bpd}
	ip := &view.IndexPage{bp}
	c.View(ip, "text/html")
}

func (c *ArticleController) PostCreate() {
	if c.Ctx.IsPost() {
		val := c.Ctx.FormValue("ueditor")
		if val != nil && len(val) > 0 {
			fmt.Println(string(val))
		}
	} else {

	}
}
