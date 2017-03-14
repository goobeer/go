package area

import (
	"fasthttpweb/filter"
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

var (
	RouteTables = make(map[string]func(ctx *fasthttp.RequestCtx))
)

type IRouter interface {
	RegistRoutes()
}

type BaseController struct {
	Ctx *fasthttp.RequestCtx
}

func injectFilter(rtFilter *filter.RouteFilter, handle func(ctx *fasthttp.RequestCtx)) func(ctx *fasthttp.RequestCtx) {
	if rtFilter == nil {
		return handle
	}
	return func(ctx *fasthttp.RequestCtx) {
		if rtFilter.BeforeRequestHandlers != nil && len(rtFilter.BeforeRequestHandlers) > 0 {
			for _, v := range rtFilter.BeforeRequestHandlers {
				v(ctx)
			}
		}
		handle(ctx)
		if rtFilter.AfterRequestHandlers != nil && len(rtFilter.AfterRequestHandlers) > 0 {
			for _, v := range rtFilter.AfterRequestHandlers {
				v(ctx)
			}
		}
	}
}

func (c *BaseController) InitBasePageData(areaName, ctrlName, title, kwd string) *BasePageData {
	return &BasePageData{AreaName: areaName, CtrlName: ctrlName, Title: title, Kwd: kwd, Data: make(map[string]interface{})}
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

		var url, preFix string
		var routeAction func(ctx *fasthttp.RequestCtx)

		if strings.HasPrefix(fName, "get") {
			fName = strings.TrimPrefix(fName, "get")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "get"

			routeAction = func(ctx *fasthttp.RequestCtx) {
				rVal := reflect.ValueOf(ci)
				rVal.Elem().FieldByName("BaseController").Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx))
				rVal.MethodByName(m.Name).Call(nil)
			}

			routeAction = injectFilter(filter.MapFilter[url], routeAction)

			router.R.GET(url, routeAction)
		} else if strings.HasPrefix(fName, "post") {
			fName = strings.TrimPrefix(fName, "post")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "post"
			routeAction = func(ctx *fasthttp.RequestCtx) {
				rVal := reflect.ValueOf(ci)
				rVal.Elem().FieldByName("BaseController").Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx))
				rVal.MethodByName(m.Name).Call(nil)
			}
			routeAction = injectFilter(filter.MapFilter[url], routeAction)
			router.R.POST(url, routeAction)
		} else if strings.HasPrefix(fName, "head") {
			fName = strings.TrimPrefix(fName, "head")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "head"
			routeAction = func(ctx *fasthttp.RequestCtx) {
				rVal := reflect.ValueOf(ci)
				rVal.Elem().FieldByName("BaseController").Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx))
				rVal.MethodByName(m.Name).Call(nil)
			}
			routeAction = injectFilter(filter.MapFilter[url], routeAction)
			router.R.HEAD(url, routeAction)
		} else if strings.HasPrefix(fName, "put") {
			fName = strings.TrimPrefix(fName, "put")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "put"
			routeAction = func(ctx *fasthttp.RequestCtx) {
				rVal := reflect.ValueOf(ci)
				rVal.Elem().FieldByName("BaseController").Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx))
				rVal.MethodByName(m.Name).Call(nil)
			}
			routeAction = injectFilter(filter.MapFilter[url], routeAction)
			router.R.PUT(url, routeAction)
		} else if strings.HasPrefix(fName, "options") {
			fName = strings.TrimPrefix(fName, "options")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "options"
			routeAction = func(ctx *fasthttp.RequestCtx) {
				rVal := reflect.ValueOf(ci)
				rVal.Elem().FieldByName("BaseController").Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx))
				rVal.MethodByName(m.Name).Call(nil)
			}
			routeAction = injectFilter(filter.MapFilter[url], routeAction)
			router.R.OPTIONS(url, routeAction)
		} else if strings.HasPrefix(fName, "delete") {
			fName = strings.TrimPrefix(fName, "delete")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "delete"
			routeAction = func(ctx *fasthttp.RequestCtx) {
				rVal := reflect.ValueOf(ci)
				rVal.Elem().FieldByName("BaseController").Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx))
				rVal.MethodByName(m.Name).Call(nil)
			}
			routeAction = injectFilter(filter.MapFilter[url], routeAction)
			router.R.DELETE(url, routeAction)
		} else if strings.HasPrefix(fName, "patch") {
			fName = strings.TrimPrefix(fName, "patch")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "patch"
			routeAction = func(ctx *fasthttp.RequestCtx) {
				rVal := reflect.ValueOf(ci)
				rVal.Elem().FieldByName("BaseController").Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx))
				rVal.MethodByName(m.Name).Call(nil)
			}
			routeAction = injectFilter(filter.MapFilter[url], routeAction)
			router.R.PATCH(url, routeAction)
		} else {
			fName = strings.TrimPrefix(fName, "get")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "get"
			routeAction = func(ctx *fasthttp.RequestCtx) {
				rVal := reflect.ValueOf(ci)
				rVal.Elem().FieldByName("BaseController").Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx))
				rVal.MethodByName(m.Name).Call(nil)
			}
			routeAction = injectFilter(filter.MapFilter[url], routeAction)
			router.R.GET(url, routeAction)
		}

		RouteTables[fmt.Sprintf("%s_%s", preFix, strings.ToLower(url))] = routeAction
		fmt.Println(url)
	}
}

func (c *BaseController) RegistRoute(newPath, httpVerb, oldPath string) {
	if len(httpVerb) == 0 {
		httpVerb = "get"
	}
	newPathTemp := fmt.Sprintf("%s_%s", httpVerb, strings.ToLower(newPath))
	oldPath = fmt.Sprintf("%s_%s", httpVerb, strings.ToLower(oldPath))
	if len(newPathTemp) > 0 && len(oldPath) > 0 && RouteTables[newPathTemp] == nil && RouteTables[oldPath] != nil {
		RouteTables[newPathTemp] = RouteTables[oldPath]
		switch httpVerb {
		case "get":
			router.R.GET(newPath, RouteTables[oldPath])
		case "post":
			router.R.POST(newPath, RouteTables[oldPath])
		case "head":
			router.R.HEAD(newPath, RouteTables[oldPath])
		case "put":
			router.R.PUT(newPath, RouteTables[oldPath])
		case "options":
			router.R.OPTIONS(newPath, RouteTables[oldPath])
		case "delete":
			router.R.DELETE(newPath, RouteTables[oldPath])
		case "patch":
			router.R.PATCH(newPath, RouteTables[oldPath])
		}
	} else {
		fmt.Printf("url:[%s] is already mapped to url:[%s].", newPath, oldPath)
	}
}

func (c *BaseController) ParseFunc(areaName, actionName string, ci interface{}) string {
	ciType := reflect.TypeOf(ci)
	t := strings.Split(strings.ToLower(ciType.String()), ".")
	typeStr := t[len(t)-1]
	typeStr = strings.TrimSuffix(typeStr, controllerFlag)
	return fmt.Sprintf("/%s/%s/%s", areaName, typeStr, actionName)
}
