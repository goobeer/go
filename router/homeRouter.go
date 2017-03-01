package router

//	"fasthttpweb/area"
//	hv "fasthttpweb/area/home/views"
//	"fasthttpweb/common"
//	"fasthttpweb/model"

//	"github.com/kataras/go-sessions"
//	"github.com/valyala/fasthttp"

const (
	Verfied      = iota
	LoginVerfied = iota
	phoneNumber  = "13121327365"
)

func init() {
	/*
		R.GET("/", getHomeVerfyPage)
		R.GET("/home/verfy", getHomeVerfyPage)
		R.POST("/home/verfy", postHomeVerfyPage)
		R.GET("/home/login", getHomeLoginPage)
		R.POST("/home/login", postHomeLoginPage)

		R.GET("/home", getHomeIndexPage)
		R.GET("/home/index", getHomeIndexPage)
		R.GET("/woda", func(ctx *fasthttp.RequestCtx) {
			ctx.WriteString("woda")
		})*/

}

/*
func getHomeIndexPage(ctx *fasthttp.RequestCtx) {
	sess := sessions.StartFasthttp(ctx)
	verfyVal := sess.Get("verfy")
	if verfyVal == nil {
		ctx.Redirect("/home/verfy", fasthttp.StatusFound)
		return
	}
	vfv := verfyVal.(int)
	if vfv != LoginVerfied {
		ctx.Redirect("/home/login", fasthttp.StatusFound)
		return
	}

	ip := &hv.IndexPage{CTX: ctx, Titles: "homeIndex"}
	area.WritePageTemplate(ctx, ip)
	ctx.SetContentType("text/html")
}

func getHomeVerfyPage(ctx *fasthttp.RequestCtx) {
	ip := &hv.VerfyPage{CTX: ctx, Titles: "homeVerfy"}
	area.WritePageTemplate(ctx, ip)
	ctx.SetContentType("text/html")
}

func getErrorHomeVerfyPage(ctx *fasthttp.RequestCtx, errMsg string) {
	ip := &hv.VerfyPage{CTX: ctx, Titles: "homeVerfyError", ErrorMsg: errMsg}
	area.WritePageTemplate(ctx, ip)
	ctx.SetContentType("text/html")
}

func postHomeVerfyPage(ctx *fasthttp.RequestCtx) {
	if ctx.IsPost() {
		phoneBytes := ctx.FormValue("phone")
		vcodeBytes := ctx.FormValue("vcode")
		sess := sessions.StartFasthttp(ctx)
		vcode := sess.Get("vcode")
		if len(phoneBytes) > 0 && string(phoneBytes) == phoneNumber && len(vcodeBytes) > 0 && string(vcodeBytes) == vcode.(string) {
			sess.Set("verfy", Verfied)
			ctx.Redirect("/home/login", fasthttp.StatusFound)
			return
		}
	}
	getErrorHomeVerfyPage(ctx, "错误的手机号或错误的验证码!")
}

func getHomeLoginPage(ctx *fasthttp.RequestCtx) {
	sess := sessions.StartFasthttp(ctx)
	verfyVal := sess.Get("verfy")
	if verfyVal == nil {
		ctx.Redirect("/home/verfy", fasthttp.StatusFound)
		return
	}
	vfv := verfyVal.(int)
	if vfv == LoginVerfied {
		ctx.Redirect("/home/index", fasthttp.StatusFound)
		return
	}

	page := &hv.LoginPage{CTX: ctx, Titles: "login"}
	area.WritePageTemplate(ctx, page)
	ctx.SetContentType("text/html")
}

func getErrorHomeLoginPage(ctx *fasthttp.RequestCtx, errMsg string) {
	page := &hv.LoginPage{CTX: ctx, Titles: "loginError", ErrorMsg: errMsg}
	area.WritePageTemplate(ctx, page)
	ctx.SetContentType("text/html")
}

func postHomeLoginPage(ctx *fasthttp.RequestCtx) {
	if ctx.IsPost() {
		unameBytes := ctx.FormValue("uname")
		pwdBytes := ctx.FormValue("pwd")
		vcodeBytes := ctx.FormValue("vcode")
		sess := sessions.StartFasthttp(ctx)
		vcode := sess.Get("vcode")
		if len(unameBytes) > 0 && len(pwdBytes) > 0 && len(vcodeBytes) > 0 && string(vcodeBytes) == vcode.(string) {
			uname := string(unameBytes)
			pwd := string(pwdBytes)
			user := new(model.Users)

			has, err := user.VerfyUser(uname, pwd)
			common.PanicError(err)
			if !has {
				getErrorHomeLoginPage(ctx, "用户名或密码错误!")
				return
			}

			sess.Set("verfy", LoginVerfied)
			ctx.Redirect("/home/index", fasthttp.StatusContinue)
			return
		}
	}
	getErrorHomeLoginPage(ctx, "请检查输入信息是否正确!")
}
*/
