package controllers

import (
	"fasthttpweb/area"
	view "fasthttpweb/area/home/views/index"
	"fmt"
	"strconv"

	"fasthttpweb/common"
	"fasthttpweb/model"

	"github.com/valyala/fasthttp"
)

const (
	phoneNumber    = "13121327365"
	controllerFlag = "controller"
)

type IndexController struct {
	*area.BaseController
}

func init() {
	ic := &IndexController{BaseController: &area.BaseController{}}
	ic.RegistRoutes(areaName, ic)
	ic.RegistRoute("get", "/home/index/index", "/", "/home", "/home/index")
}

func (c *IndexController) Index() {
	sess := c.StartSession()
	verfyVal := sess.Get("verfy")
	if verfyVal == nil {
		c.Ctx.Redirect(c.ParseFunc(areaName, "Verify", c), fasthttp.StatusFound)
		return
	}
	vfv := verfyVal.(int)
	if vfv != common.LoginVerfied {
		c.Ctx.Redirect(c.ParseFunc(areaName, "login", c), fasthttp.StatusFound)
		return
	}

	user := sess.Get("user")
	bpd := area.NewBasePageData(areaName, "", "", c)
	bp := area.NewBasePage(c.Ctx, bpd)
	bpd.Data["user"] = user
	ip := &view.IndexPage{BasePage: bp}

	c.View(ip, "text/html")
}

func (c *IndexController) Verify() {
	fmt.Println(string(c.Ctx.QueryArgs().QueryString()))
	bpd := area.NewBasePageData(areaName, "", "", c)
	bp := area.NewBasePage(c.Ctx, bpd)
	ip := &view.VerfyPage{BasePage: bp}
	c.View(ip, "text/html")
}

func (c *IndexController) PostVerify() {
	phoneBytes := c.Ctx.FormValue("phone")
	vcodeBytes := c.Ctx.FormValue("vcode")
	sess := c.StartSession()
	vcode := sess.Get("vcode")
	var phone string
	if phoneBytes != nil && len(phoneBytes) > 0 {
		phone = string(phoneBytes)
	}
	if len(phoneBytes) > 0 && phone == phoneNumber && len(vcodeBytes) > 0 && string(vcodeBytes) == vcode.(string) {
		sess.Set("verfy", common.Verfied)
		c.Ctx.Redirect(c.ParseFunc(areaName, "login", c), fasthttp.StatusFound)
		return
	}

	bpd := area.NewBasePageData(areaName, "", "", c)
	bp := area.NewBasePage(c.Ctx, bpd)
	bpd.Data["error"] = "错误的手机号或错误的验证码!"
	bpd.Data["phone"] = phone
	ip := &view.VerfyPage{bp}

	c.View(ip, "text/html")
}

func (c *IndexController) Login() {
	sess := c.StartSession()
	verfyVal := sess.Get("verfy")
	if verfyVal == nil {
		c.Ctx.Redirect("/home/index/verify", fasthttp.StatusFound)
		return
	}
	vfv := verfyVal.(int)
	if vfv == common.LoginVerfied {
		c.Ctx.Redirect("/home/index", fasthttp.StatusFound)
		return
	}

	bpd := area.NewBasePageData(areaName, "", "", c)
	bp := area.NewBasePage(c.Ctx, bpd)

	page := &view.LoginPage{bp}
	c.View(page, "text/html")
}

func (c *IndexController) PostLogin() {
	unameBytes := c.Ctx.FormValue("uname")
	pwdBytes := c.Ctx.FormValue("pwd")
	vcodeBytes := c.Ctx.FormValue("vcode")
	sess := c.StartSession()
	vcode := sess.Get("vcode")
	var errMsg string
	if len(unameBytes) > 0 && len(pwdBytes) > 0 && len(vcodeBytes) > 0 && string(vcodeBytes) == vcode.(string) {
		uname := string(unameBytes)
		pwd := string(pwdBytes)
		user := new(model.Users)

		has, err := user.VerfyUser(uname, pwd)
		if err != nil {
			panic(err)
		}
		if has {
			sess.Set("verfy", common.LoginVerfied)
			sess.Set("user", user)
			c.Ctx.Redirect("/home/index/index", fasthttp.StatusFound)
			return
		} else {
			errMsg = "用户名或密码错误!"
		}
	} else {
		errMsg = "用户名或密码不能为空!"
	}

	bpd := area.NewBasePageData(areaName, "", "", c)
	bp := area.NewBasePage(c.Ctx, bpd)
	bpd.Data["ErrMsg"] = errMsg
	page := &view.LoginPage{bp}
	c.View(page, "text/html")
}

func (c *IndexController) Logout() {
	sess := c.StartSession()
	sess.Clear()
	c.Ctx.Redirect("/home/index/Verify", fasthttp.StatusContinue)
}

func (c *IndexController) PageTest() {
	bpd := area.NewBasePageData(areaName, "", "", c)
	bp := area.NewBasePage(c.Ctx, bpd)
	currentPage := c.Ctx.FormValue("page")
	currentPageVal := 1
	if currentPage != nil && len(currentPage) > 0 {
		currentPageVal, _ = strconv.Atoi(string(currentPage))
		if currentPageVal < 1 {
			currentPageVal = 1
		}
	}
	bpd.Data["page"] = currentPageVal
	page := &view.PageTest{bp}
	c.View(page, "text/html")
}
