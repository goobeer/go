package controllers

import (
	"fasthttpweb/area"
	view "fasthttpweb/area/admin/views/article"
	"fasthttpweb/model"
	"fmt"
	"strconv"
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
	var err error
	pageNumber, pageSize := 0, 10
	page := c.Ctx.FormValue("pageNumber")
	if page != nil {
		pageNumber, _ = strconv.Atoi(string(page))
	}

	art := &model.Article{}
	var arts []model.Article
	err = art.GetList(&arts, uint(pageNumber), uint(pageSize), nil)
	if err != nil {
		bp.BPD.Data["error"] = err.Error()
	} else {
		bp.BPD.Data["Model"] = arts
	}

	ip := &view.IndexPage{bp}

	c.View(ip, "text/html")
}

func (c *ArticleController) Create() {
	bpd := c.InitBasePageData(areaName, "article-create", "article-create", "article-create")
	bp := &area.BasePage{CTX: c.Ctx, BPD: bpd}

	page := &view.CreatePage{bp}
	c.View(page, "text/html")
}

func (c *ArticleController) PostCreate() {
	if c.Ctx.IsPost() {
		bpd := c.InitBasePageData(areaName, "Index", "article-title", "article-kwd")
		bp := &area.BasePage{CTX: c.Ctx, BPD: bpd}

		title := c.Ctx.FormValue("Title")
		content := c.Ctx.FormValue("Content")
		fmt.Println("daung", title, "\r\n", content)
		if title != nil && len(title) > 0 && content != nil && len(content) > 0 {
			fmt.Println(string(title), string(content))
			art := &model.Article{Title: string(title), Content: string(content)}
			_, err := art.Add(art)
			if err != nil {
				bpd.Data["error"] = err.Error()
				page := &view.CreatePage{bp}
				c.View(page, "text/html")
			} else {
				c.Ctx.Redirect("/admin/article", 302)
			}
		}
	} else {

	}
}
