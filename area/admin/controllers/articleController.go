package controllers

import (
	"errors"
	"fasthttpweb/area"
	view "fasthttpweb/area/admin/views/article"
	"fasthttpweb/model"
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
	bpd := c.InitBasePageData(areaName, "article", "article-title", "article-kwd")
	bp := &area.BasePage{CTX: c.Ctx, BPD: bpd}
	var err error
	pageNumber, pageSize := 0, 10
	page := c.Ctx.FormValue("pageNumber")
	if page != nil {
		pageNumber, _ = strconv.Atoi(string(page))
		if pageNumber > 0 {
			pageNumber--
		}
	}

	art := &model.Article{}
	var arts []model.Article

	err = art.GetListWithCols(&arts, "ID,Title,CreateTime,Used", "", uint(pageNumber), uint(pageSize), nil)
	if err != nil {
		c.ErrorView(areaName, "article", "article-index", err)
		return
	} else {
		bp.BPD.Data["Model"] = arts
	}

	ip := &view.IndexPage{bp}
	c.View(ip, "text/html")
}

func (c *ArticleController) Create() {
	bpd := c.InitBasePageData(areaName, "article", "article-create", "article-create")
	bp := &area.BasePage{CTX: c.Ctx, BPD: bpd}

	page := &view.CreatePage{bp}
	c.View(page, "text/html")
}

func (c *ArticleController) PostCreate() {
	bpd := c.InitBasePageData(areaName, "article", "article-title", "article-kwd")
	bp := &area.BasePage{CTX: c.Ctx, BPD: bpd}

	title := c.Ctx.FormValue("Title")
	content := c.Ctx.FormValue("Content")

	if title != nil && len(title) > 0 && content != nil && len(content) > 0 {
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
}

func (c *ArticleController) Detail() {
	bpd := c.InitBasePageData(areaName, "article", "article-title", "article-kwd")
	bp := &area.BasePage{CTX: c.Ctx, BPD: bpd}
	id := c.Ctx.FormValue("id")
	idv := -1
	var err error
	if id != nil && len(id) > 0 {
		idv, err = strconv.Atoi(string(id))
		if err != nil {
			bpd.Data["error"] = err.Error()
		} else {
			art := &model.Article{}
			var ok bool
			ok, err = art.Get(art, "id=?", idv)
			if ok {
				bpd.Data["Model"] = art
				bpd.Title = art.Title
			} else {
				bpd.Data["error"] = err.Error()
			}
		}
	}
	page := &view.DetailPage{bp}
	c.View(page, "text/html")
}

func (c *ArticleController) Edit() {
	id := c.Ctx.FormValue("id")
	if id != nil && len(id) > 0 {
		idv, err := strconv.Atoi(string(id))
		if err != nil {
			c.ErrorView(areaName, "article", "article-edit", errors.New("invilated param"))
			return
		}
		bpd := c.InitBasePageData(areaName, "article", "article-title", "article-kwd")
		bp := &area.BasePage{CTX: c.Ctx, BPD: bpd}
		art := &model.Article{}
		_, err = art.Get(art, "id=?", int(idv))
		if err != nil {
			c.ErrorView(areaName, "article", "article-edit", err)
		} else {
			bpd.Data["Model"] = art
		}
		page := &view.EditPage{bp}
		c.View(page, "text/html")
	} else {
		c.ErrorView(areaName, "article", "article-edit", errors.New("invilated param"))
	}

}

func (c *ArticleController) PostEdit() {
	id := c.Ctx.FormValue("ID")
	title := c.Ctx.FormValue("Title")
	content := c.Ctx.FormValue("Content")
	//	used := c.Ctx.FormValue("Used")
	var err error
	if id != nil && len(id) > 0 && title != nil && len(title) > 0 && content != nil && len(content) > 0 {

		var idv int
		idv, err = strconv.Atoi(string(id))
		if err == nil {
			art := &model.Article{BaseModel: &model.BaseModel{ID: int64(idv)}, Title: string(title), Content: string(content)}
			_, err = art.UpdateByID(art, "", "CreateTime,Used")
			if err == nil {
				c.Index()
				return
			}
		}
	}

	if err == nil {
		err = errors.New("invalidate param")
	}
	c.ErrorView(areaName, "article", "article-edit", err)
}

func (c *ArticleController) MDelete() {
	id := c.Ctx.FormValue("id")
	if id == nil {
		c.ErrorView(areaName, "article", "article-delete", errors.New("invalidate param"))
	} else {
		var err error
		var idv int
		idv, err = strconv.Atoi(string(id))
		if err == nil {
			art := &model.Article{}
			_, err = art.Delete(art, "id=?", idv)
			if err == nil {
				c.Ctx.Redirect("/admin/article", 302)
			}
		}

		c.ErrorView(areaName, "article", "article-delete", err)
		return
	}
}
