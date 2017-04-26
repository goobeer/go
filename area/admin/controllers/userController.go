package controllers

import (
	"errors"
	"fasthttpweb/area"
	view "fasthttpweb/area/admin/views/user"
	"fasthttpweb/model"
	"strconv"

	"github.com/valyala/fasthttp"
)

type UserController struct {
	*area.BaseController
}

func init() {
	c := &UserController{BaseController: &area.BaseController{}}
	c.RegistRoutes(areaName, c)
	c.RegistRoute("get", "/admin/user/index", "/admin/user")
}

func (c *UserController) Index() {
	bpd := area.NewBasePageData(areaName, "admin-user-title", "admin-user-kwd", c)
	bp := area.NewBasePage(c.Ctx, bpd)
	user := &model.Users{}
	var users []model.Users
	pageNumber, pageSize := 0, 10
	err := user.GetList(&users, uint(pageNumber), uint(pageSize), nil)
	if err != nil {
		bpd.Data["error"] = err.Error()
	} else {
		bpd.Data["Model"] = users
	}
	ip := &view.IndexPage{bp}
	c.View(ip, "text/html")
}

func (c *UserController) Create() {
	bpd := area.NewBasePageData(areaName, "admin-user-title", "admin-user-kwd", c)
	bp := area.NewBasePage(c.Ctx, bpd)
	page := &view.CreatePage{bp}
	c.View(page, "text/html")
}

func (c *UserController) PostCreate() {
	name := c.Ctx.FormValue("Name")
	pwd := c.Ctx.FormValue("Pwd")
	var err error
	if name != nil && len(name) > 0 && pwd != nil && len(pwd) > 0 {
		user := &model.Users{Name: string(name), Pwd: string(pwd)}
		user.GeneratePwd()
		if !user.RepeatName() {
			_, err = user.Add(user)
			if err == nil {
				c.Ctx.Redirect("/admin/user", 302)
				return
			}
		}
	}
	if err == nil {
		err = errors.New("invilidate param")
	}
	c.ErrorView(areaName, "User", "user-create", err)
}

func (c *UserController) Edit() {
	id := c.Ctx.FormValue("id")
	var err error
	if id != nil && len(id) > 0 {
		var idv int
		idv, err = strconv.Atoi(string(id))
		if err == nil {
			user := &model.Users{BaseModel: &model.BaseModel{}}
			_, err = user.Get(user, "id=?", idv)
			if err == nil {
				bpd := area.NewBasePageData(areaName, "admin-user-title", "admin-user-kwd", c)
				bp := area.NewBasePage(c.Ctx, bpd)
				bpd.Data["Model"] = user
				page := &view.EditPage{bp}
				c.View(page, "text/html")
				return
			}
		}
	}
	if err == nil {
		err = errors.New("invilidate param")
	}
	c.ErrorView(areaName, "User", "User-edit", err)
}

func (c *UserController) PostEdit() {
	id := c.Ctx.FormValue("ID")
	pwd := c.Ctx.FormValue("Pwd")
	var err error
	user := &model.Users{BaseModel: &model.BaseModel{}}
	if pwd != nil && len(pwd) > 0 && id != nil && len(id) > 0 {

		var idv int
		idv, err = strconv.Atoi(string(id))
		if err == nil {
			user.Pwd = string(pwd)
			user.ID = int64(idv)
			user.GeneratePwd()

			_, err = user.UpdateByID(user, "Pwd", "")
			if err == nil {
				c.Ctx.Redirect("/admin/user", fasthttp.StatusContinue)
				return
			}
		}
	}
	if err == nil {
		err = errors.New("invilidate param")
	}
	c.ErrorView(areaName, "User", "user-create", err)
}

func (c *UserController) Detail() {
	id := c.Ctx.FormValue("id")
	var err error
	if id != nil && len(id) > 0 {
		var idv int
		idv, err = strconv.Atoi(string(id))
		if err == nil {
			user := &model.Users{BaseModel: &model.BaseModel{}}
			_, err = user.Get(user, "id=?", idv)
			if err == nil {
				bpd := area.NewBasePageData(areaName, "User", "admin-user-title", "admin-user-kwd")
				bp := area.NewBasePage(c.Ctx, bpd)
				bpd.Data["Model"] = user
				page := &view.DetailPage{bp}
				c.View(page, "text/html")
				return
			}
		}
	}
	if err == nil {
		err = errors.New("invilidate param")
	}
	c.ErrorView(areaName, "User", "User-edit", err)
}

func (c *UserController) MDelete() {
	id := c.Ctx.FormValue("id")
	var err error
	if id != nil && len(id) > 0 {
		var idv int
		idv, err = strconv.Atoi(string(id))
		if err == nil {
			user := &model.Users{}
			r, err := user.Delete(user, "id=?", idv)
			if r != 0 && err == nil {
				c.Ctx.Redirect("/admin/user", fasthttp.StatusContinue)
			}
		}
	}
	if err == nil {
		err = errors.New("invilidate param")
	}
	c.ErrorView(areaName, "User", "User-delete", err)
}

func (c *UserController) Logout() {
	sess := c.StartSession()
	sess.Clear()
	c.Ctx.Redirect("/home/index/verify", fasthttp.StatusContinue)
}
