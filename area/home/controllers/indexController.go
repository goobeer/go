package controllers

import (
	"fasthttpweb/area"
	hv "fasthttpweb/area/home/views/index"

	"fasthttpweb/common"
	"fasthttpweb/model"
	//	"fasthttpweb/router"

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

	//auto regist route

	//func intercept
}

func (c *IndexController) Index() {
	ctx := c.Ctx
	sess := c.StartSession()
	verfyVal := sess.Get("verfy")
	if verfyVal == nil {
		ctx.Redirect(c.ParseFunc(areaName, "Verify", c), fasthttp.StatusFound)
		return
	}
	vfv := verfyVal.(int)
	if vfv != common.LoginVerfied {
		ctx.Redirect(c.ParseFunc(areaName, "login", c), fasthttp.StatusFound)
		return
	}

	user := sess.Get("user")
	bpd := area.NewBasePageData(areaName, "", "", c)
	bp := area.NewBasePage(c.Ctx, bpd)
	bpd.Data["user"] = user
	ip := &hv.IndexPage{BasePage: bp}

	c.View(ip, "text/html")
}

func (c *IndexController) Verify() {
	bpd := area.NewBasePageData(areaName, "", "", c)
	bp := area.NewBasePage(c.Ctx, bpd)
	ip := &hv.VerfyPage{BasePage: bp}

	c.View(ip, "text/html")
}

func (c *IndexController) PostVerify() {
	ctx := c.Ctx
	if ctx.IsPost() {
		phoneBytes := ctx.FormValue("phone")
		vcodeBytes := ctx.FormValue("vcode")
		sess := c.StartSession()
		vcode := sess.Get("vcode")
		if len(phoneBytes) > 0 && string(phoneBytes) == phoneNumber && len(vcodeBytes) > 0 && string(vcodeBytes) == vcode.(string) {
			sess.Set("verfy", common.Verfied)

			ctx.Redirect(c.ParseFunc(areaName, "login", c), fasthttp.StatusFound)
			return
		}
	}
	c.getErrorHomeVerfyPage("错误的手机号或错误的验证码!")
}

func (c *IndexController) Login() {
	ctx := c.Ctx
	sess := c.StartSession()
	verfyVal := sess.Get("verfy")
	if verfyVal == nil {
		ctx.Redirect("/home/index/verify", fasthttp.StatusFound)
		return
	}
	vfv := verfyVal.(int)
	if vfv == common.LoginVerfied {
		ctx.Redirect("/home/index", fasthttp.StatusFound)
		return
	}

	bpd := area.NewBasePageData(areaName, "", "", c)
	bp := area.NewBasePage(c.Ctx, bpd)

	page := &hv.LoginPage{bp}

	c.View(page, "text/html")
}

func (c *IndexController) PostLogin() {
	ctx := c.Ctx
	if ctx.IsPost() {
		unameBytes := ctx.FormValue("uname")
		pwdBytes := ctx.FormValue("pwd")
		vcodeBytes := ctx.FormValue("vcode")
		sess := c.StartSession()
		vcode := sess.Get("vcode")
		if len(unameBytes) > 0 && len(pwdBytes) > 0 && len(vcodeBytes) > 0 && string(vcodeBytes) == vcode.(string) {
			uname := string(unameBytes)
			pwd := string(pwdBytes)
			user := new(model.Users)

			has, err := user.VerfyUser(uname, pwd)
			if err != nil {
				panic(err)
			}
			if !has {
				c.getErrorHomeLoginPage("用户名或密码错误!")
				return
			}

			sess.Set("verfy", common.LoginVerfied)
			sess.Set("user", user)
			ctx.Redirect("/home/index/index", fasthttp.StatusContinue)
			return
		}
	}
	c.getErrorHomeLoginPage("请检查输入信息是否正确!")
}

func (c *IndexController) Logout() {
	sess := c.StartSession()
	sess.Delete("vcode")
	sess.Delete("verfy")

	c.Ctx.Redirect("/home/index/Verify", fasthttp.StatusContinue)
}

func (c *IndexController) getErrorHomeVerfyPage(errMsg string) {
	bpd := area.NewBasePageData(areaName, "", "", c)
	bp := area.NewBasePage(c.Ctx, bpd)
	bpd.Data["error"] = errMsg
	phoneBytes := c.Ctx.FormValue("phone")
	bpd.Data["phone"] = string(phoneBytes)
	ip := &hv.VerfyPage{bp}

	c.View(ip, "text/html")
}

func (c *IndexController) getErrorHomeLoginPage(errMsg string) {
	bpd := area.NewBasePageData(areaName, "", "", c)
	bp := area.NewBasePage(c.Ctx, bpd)
	bpd.Data["ErrMsg"] = errMsg
	page := &hv.LoginPage{bp}

	c.View(page, "text/html")
}
