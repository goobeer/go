package controllers

import (
	"fasthttpweb/area"
	view "fasthttpweb/area/admin/views/articletype"
	"fasthttpweb/model"
	"strconv"
)

type ArticleTypeController struct {
	*area.BaseController
}

func init() {
	c := &ArticleTypeController{BaseController: &area.BaseController{}}
	c.RegistRoutes(areaName, c)
	c.RegistRoute("get", "/admin/articletype/index", "/admin/articletype")
}

func (c *ArticleTypeController) Index() {
	bpd := area.NewBasePageData(areaName, "articletype", "articleType", "")
	bp := area.NewBasePage(c.Ctx, bpd)
	pageNumber := c.Ctx.FormValue("pageNumber")
	pageNumVal, pageSize := 1, 10
	if pageNumber != nil && len(pageNumber) > 0 {
		pageNumVal, _ = strconv.Atoi(string(pageNumber))
	}
	if pageNumVal > 0 {
		pageNumVal--
	}
	artType := &model.ArticleType{}
	var artTypes []model.ArticleType
	err := artType.GetList(&artTypes, uint(pageNumVal), uint(pageSize), nil)
	if err == nil {
		bpd.Data["Model"] = artTypes
	} else {
		c.ErrorView(areaName, "articletype", "articletype-list", err)
		return
	}

	page := &view.IndexPage{bp}
	c.View(page, "text/html")
}

func (c *ArticleTypeController) Create() {

}
func (c *ArticleTypeController) PostCreate() {

}
func (c *ArticleTypeController) Edit() {

}

func (c *ArticleTypeController) PostEdit() {

}

func (c *ArticleTypeController) Detail() {

}

func (c *ArticleTypeController) MDelete() {

}
