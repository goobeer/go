package area

import (
	"fasthttpweb/router"
	"fmt"
	"reflect"
	"strings"

	"github.com/kataras/go-sessions"
	"github.com/valyala/fasthttp"
)

const (
	controllerFlag = "controller"
)

type IRouter interface {
	RegistRoutes()
}

type BaseController struct {
	Ctx *fasthttp.RequestCtx
}

func (c *BaseController) StartSession() sessions.Session {
	return sessions.StartFasthttp(c.Ctx)
}

func (c *BaseController) View(page Page, contentType string) {
	WritePageTemplate(c.Ctx, page)
	c.Ctx.SetContentType(contentType)
}

func (c *BaseController) RegistRoutes(areaFlag string, ci interface{}) {
	rType := reflect.TypeOf(c)
	ciType := reflect.TypeOf(ci)

	t := strings.Split(strings.ToLower(ciType.String()), ".")

	typeStr := t[len(t)-1]
	typeStr = strings.TrimSuffix(typeStr, controllerFlag)

	for i := 0; i < ciType.NumMethod(); i++ {
		m := ciType.Method(i)
		if _, ok := rType.MethodByName(m.Name); ok {
			continue
		}

		fName := strings.ToLower(m.Name)
		var url string
		if strings.HasPrefix(fName, "get") {
			fName = strings.TrimPrefix(fName, "get")
			url = fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName)
			router.R.GET(url, func(ctx *fasthttp.RequestCtx) {
				rVal := reflect.ValueOf(ci)
				rVal.Elem().FieldByName("BaseController").Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx))
				rVal.MethodByName(m.Name).Call(nil)
			})
		} else if strings.HasPrefix(fName, "post") {
			fName = strings.TrimPrefix(fName, "post")
			url = fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName)
			router.R.POST(url, func(ctx *fasthttp.RequestCtx) {
				rVal := reflect.ValueOf(ci)
				rVal.Elem().FieldByName("BaseController").Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx))
				rVal.MethodByName(m.Name).Call(nil)
			})
		} else if strings.HasPrefix(fName, "head") {
			fName = strings.TrimPrefix(fName, "head")
			url = fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName)
			router.R.HEAD(url, func(ctx *fasthttp.RequestCtx) {
				rVal := reflect.ValueOf(ci)
				rVal.Elem().FieldByName("BaseController").Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx))
				rVal.MethodByName(m.Name).Call(nil)
			})
		} else if strings.HasPrefix(fName, "put") {
			fName = strings.TrimPrefix(fName, "put")
			url = fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName)
			router.R.PUT(url, func(ctx *fasthttp.RequestCtx) {
				rVal := reflect.ValueOf(ci)
				rVal.Elem().FieldByName("BaseController").Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx))
				rVal.MethodByName(m.Name).Call(nil)
			})
		} else if strings.HasPrefix(fName, "options") {
			fName = strings.TrimPrefix(fName, "options")
			url = fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName)
			router.R.OPTIONS(url, func(ctx *fasthttp.RequestCtx) {
				rVal := reflect.ValueOf(ci)
				rVal.Elem().FieldByName("BaseController").Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx))
				rVal.MethodByName(m.Name).Call(nil)
			})
		} else if strings.HasPrefix(fName, "delete") {
			fName = strings.TrimPrefix(fName, "delete")
			url = fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName)
			router.R.DELETE(url, func(ctx *fasthttp.RequestCtx) {
				rVal := reflect.ValueOf(ci)
				rVal.Elem().FieldByName("BaseController").Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx))
				rVal.MethodByName(m.Name).Call(nil)
			})
		} else if strings.HasPrefix(fName, "patch") {
			fName = strings.TrimPrefix(fName, "patch")
			url = fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName)
			router.R.PATCH(url, func(ctx *fasthttp.RequestCtx) {
				rVal := reflect.ValueOf(ci)
				rVal.Elem().FieldByName("BaseController").Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx))
				rVal.MethodByName(m.Name).Call(nil)
			})
		} else {
			fName = strings.TrimPrefix(fName, "get")
			url = fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName)
			router.R.GET(url, func(ctx *fasthttp.RequestCtx) {
				rVal := reflect.ValueOf(ci)
				rVal.Elem().FieldByName("BaseController").Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx))
				rVal.MethodByName(m.Name).Call(nil)
			})
		}
		fmt.Println(url)
	}
}

func (c *BaseController) ParseFunc(areaName, actionName string, ci interface{}) string {
	ciType := reflect.TypeOf(ci)
	t := strings.Split(strings.ToLower(ciType.String()), ".")
	typeStr := t[len(t)-1]
	typeStr = strings.TrimSuffix(typeStr, controllerFlag)
	return fmt.Sprintf("/%s/%s/%s", areaName, typeStr, actionName)
}
