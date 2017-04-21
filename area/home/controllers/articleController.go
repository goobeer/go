package controllers

import (
	"fasthttpweb/area"

	view "fasthttpweb/area/admin/views/article"
)

type ArticleController struct {
	*area.BaseController
}

func init() {
	c := &ArticleController{BaseController: &area.BaseController{}}
	c.RegistRoutes(areaName, c)
}

func (c *ArticleController) Index() {
	bpd := c.InitBasePageData(areaName, "Index", "article-title", "article-kwd")
	bp := &area.BasePage{CTX: c.Ctx, BPD: bpd}

	ip := &view.IndexPage{bp}

	c.View(ip, "text/html")
}

func (c *ArticleController) Detail() {

}
