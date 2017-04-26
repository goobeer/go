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
	bpd := area.NewBasePageData(areaName, "articleType", "", c)
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
	err := artType.DbListCommandSession(uint(pageNumVal), uint(pageSize), nil).Desc("ord").Find(&artTypes)
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
	bpd := area.NewBasePageData(areaName, "articleType-create", "articleType-create", c)
	bp := area.NewBasePage(c.Ctx, bpd)

	page := &view.CreatePage{bp}
	c.View(page, "text/html")
}
func (c *ArticleTypeController) PostCreate() {
	name := c.Ctx.FormValue("Name")
	ord := c.Ctx.FormValue("Ord")
	ordVal := 0
	var err error
	if name != nil && len(name) > 0 && ord != nil && len(ord) > 0 {
		ordVal, err = strconv.Atoi(string(ord))
		if err != nil {
			c.ErrorView(areaName, "articletype-create", c, err)
			return
		}
	}
	artType := &model.ArticleType{Name: string(name), Ord: ordVal}
	_, err = artType.Add(artType)
	if err != nil {
		c.ErrorView(areaName, "save-articletype-create", c, err)
		return
	}
	c.Ctx.Redirect("/admin/articletype", 302)
}
func (c *ArticleTypeController) Edit() {
	bpd := area.NewBasePageData(areaName, "articleType-create", "articleType-create", c)
	bp := area.NewBasePage(c.Ctx, bpd)

	page := &view.EditPage{bp}
	c.View(page, "text/html")
}

func (c *ArticleTypeController) PostEdit() {

}

func (c *ArticleTypeController) Detail() {
	bpd := area.NewBasePageData(areaName, "articleType-create", "articleType-create", c)
	bp := area.NewBasePage(c.Ctx, bpd)

	page := &view.DetailPage{bp}
	c.View(page, "text/html")
}

func (c *ArticleTypeController) MDelete() {

}
