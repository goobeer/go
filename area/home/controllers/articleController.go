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
	bpd := area.NewBasePageData(areaName, "article-title", "article-kwd", c)
	bp := area.NewBasePage(c.Ctx, bpd)

	ip := &view.IndexPage{bp}

	c.View(ip, "text/html")
}

func (c *ArticleController) Detail() {

}
