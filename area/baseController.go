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

func injectFilter(httpVerb, uri string, m reflect.Method, ci interface{}) (handle func(ctx *fasthttp.RequestCtx)) {
	if !(filter.MapFilter != nil && len(filter.MapFilter) > 0) {
		handle = func(ctx *fasthttp.RequestCtx) {
			defer func() {
				if err := recover(); err != nil {
					filter.ErrLog(err, ctx)
				}
			}()
			rVal := reflect.ValueOf(ci)
			rVal.Elem().FieldByName("BaseController").Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx))
			rVal.MethodByName(m.Name).Call(nil)
		}
		return
	}

	var befReqHandlers, aftReqHandlers []fasthttp.RequestHandler

	for _, rtFilter := range filter.MapFilter {
		if !rtFilter.FilterMatch(uri, httpVerb) {
			continue
		}

		if rtFilter.BeforeRequestHandlers != nil && len(rtFilter.BeforeRequestHandlers) > 0 {
			befReqHandlers = append(befReqHandlers, rtFilter.BeforeRequestHandlers...)
		}

		if rtFilter.AfterRequestHandlers != nil && len(rtFilter.AfterRequestHandlers) > 0 {
			aftReqHandlers = append(aftReqHandlers, rtFilter.AfterRequestHandlers...)
		}
	}

	handle = func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if err := recover(); err != nil {
				filter.ErrLog(err, ctx)
			}
		}()

		for _, v := range befReqHandlers {
			v(ctx)
		}

		rVal := reflect.ValueOf(ci)
		rVal.Elem().FieldByName("BaseController").Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx))
		rVal.MethodByName(m.Name).Call(nil)

		for _, v := range aftReqHandlers {
			v(ctx)
		}
	}
	return
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

			routeAction = injectFilter(preFix, url, m, ci)

			router.R.GET(url, routeAction)
		} else if strings.HasPrefix(fName, "post") {
			fName = strings.TrimPrefix(fName, "post")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "post"

			routeAction = injectFilter(preFix, url, m, ci)

			router.R.POST(url, routeAction)
		} else if strings.HasPrefix(fName, "head") {
			fName = strings.TrimPrefix(fName, "head")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "head"
			routeAction = injectFilter(preFix, url, m, ci)
			router.R.HEAD(url, routeAction)
		} else if strings.HasPrefix(fName, "put") {
			fName = strings.TrimPrefix(fName, "put")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "put"
			routeAction = injectFilter(preFix, url, m, ci)
			router.R.PUT(url, routeAction)
		} else if strings.HasPrefix(fName, "options") {
			fName = strings.TrimPrefix(fName, "options")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "options"
			routeAction = injectFilter(preFix, url, m, ci)
			router.R.OPTIONS(url, routeAction)
		} else if strings.HasPrefix(fName, "delete") {
			fName = strings.TrimPrefix(fName, "delete")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "delete"
			routeAction = injectFilter(preFix, url, m, ci)
			router.R.DELETE(url, routeAction)
		} else if strings.HasPrefix(fName, "patch") {
			fName = strings.TrimPrefix(fName, "patch")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "patch"
			routeAction = injectFilter(preFix, url, m, ci)
			router.R.PATCH(url, routeAction)
		} else {
			fName = strings.TrimPrefix(fName, "get")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "get"
			routeAction = injectFilter(preFix, url, m, ci)
			router.R.GET(url, routeAction)
		}

		RouteTables[fmt.Sprintf("%s_%s", preFix, strings.ToLower(url))] = routeAction
		//		fmt.Println(url)
	}
}

func (c *BaseController) RegistRoute(httpVerb, oldPath string, newPaths ...string) {
	if len(httpVerb) == 0 {
		httpVerb = "get"
	}
	if len(newPaths) == 0 {
		return
	}

	for _, newPath := range newPaths {

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
			fmt.Printf("url:[%s] is already mapped to url:[%s].\r\n", newPath, oldPath)
		}
	}
}

func (c *BaseController) ParseFunc(areaName, actionName string, ci interface{}) string {
	ciType := reflect.TypeOf(ci)
	t := strings.Split(strings.ToLower(ciType.String()), ".")
	typeStr := t[len(t)-1]
	typeStr = strings.TrimSuffix(typeStr, controllerFlag)
	return fmt.Sprintf("/%s/%s/%s", areaName, typeStr, actionName)
}
