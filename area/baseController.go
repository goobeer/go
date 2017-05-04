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
					errFilter := &filter.ErrorFilter{FilterStateResult: &filter.FilterStateResult{FilterState: true, Data: err}}
					errFilter.OnException(ctx)
				}
			}()
			rVal := reflect.ValueOf(ci)
			rVal.Elem().FieldByName("BaseController").Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx))
			rVal.MethodByName(m.Name).Call(nil)
		}
		return
	}

	filters := make([]filter.BaseFilter, 0)

	for _, rtFilter := range filter.MapFilter {
		if !rtFilter.FilterMatch(uri, httpVerb) {
			continue
		}

		if rtFilter.Filters != nil && len(rtFilter.Filters) > 0 {
			filters = append(filters, rtFilter.Filters...)
		}
	}

	handle = func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if err := recover(); err != nil {
				errFilter := &filter.ErrorFilter{FilterStateResult: &filter.FilterStateResult{FilterState: true, Data: err}}
				errFilter.OnException(ctx)
			}
		}()

		var aftReqHandlers []func(ctx *fasthttp.RequestCtx) *filter.FilterStateResult
		var stateResult *filter.FilterStateResult
		for _, v := range filters {
			switch v.(type) {
			case filter.AuthorizeFilter:
				authorFilter := v.(filter.AuthorizeFilter)
				stateResult = authorFilter.Authorization(ctx)
				if !(stateResult != nil && stateResult.FilterState) {
					return
				}
			case filter.ActionFilter:
				actionFilter := v.(filter.ActionFilter)
				actionFilter.BeforeExecute(ctx)
				aftReqHandlers = append(aftReqHandlers, actionFilter.AfterExecute)
			case filter.BeforeActionFilter:
				beforeActionFilter := v.(filter.BeforeActionFilter)
				beforeActionFilter.BeforeExecute(ctx)
			case filter.AfterActionFilter:
				afterActionFilter := v.(filter.AfterActionFilter)
				aftReqHandlers = append(aftReqHandlers, afterActionFilter.AfterExecute)
			case filter.ExceptionFilter:
				expFilter := v.(filter.ExceptionFilter)
				stateResult = expFilter.OnException(ctx)
				if !(stateResult != nil && stateResult.FilterState) {
					return
				}
			}
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

func NewBasePage(ctx *fasthttp.RequestCtx, bpd *BasePageData) *BasePage {
	return &BasePage{CTX: ctx, BPD: bpd}
}

func NewBasePageData(areaName, title, kwd string, ciPtr interface{}) *BasePageData {
	ciType := reflect.TypeOf(ciPtr)
	t := strings.Split(strings.ToLower(ciType.String()), ".")
	ctrlName := t[len(t)-1]
	ctrlName = strings.TrimSuffix(ctrlName, controllerFlag)
	return &BasePageData{AreaName: areaName, CtrlName: ctrlName, Title: title, Kwd: kwd, Data: make(map[string]interface{})}
}

func Url(bp *BasePage, actionName string, data map[string]interface{}) string {
	url := fmt.Sprintf("/%s/%s/%s", bp.BPD.AreaName, bp.BPD.CtrlName, actionName)
	if data != nil && len(data) > 0 {
		var urlParams string
		i := 0
		for k, v := range data {
			if i == 0 {
				i++
			}
			if i > 1 {
				urlParams += fmt.Sprintf("&%s=%v", k, v)
			} else {
				urlParams = fmt.Sprintf("?%s=%v", k, v)
			}
		}
		url += urlParams
	}
	return url
}

func Url2(areaName, ctrlName, actionName string, data map[string]interface{}) string {
	url := fmt.Sprintf("/%s/%s/%s", areaName, ctrlName, actionName)
	if data != nil && len(data) > 0 {
		var urlParams string
		i := 0
		for k, v := range data {
			if i == 0 {
				i++
			}
			if i > 1 {
				urlParams += fmt.Sprintf("&%s=%v", k, v)
			} else {
				urlParams = fmt.Sprintf("?%s=%v", k, v)
			}
		}
		url += urlParams
	}
	return url
}

func (c *BaseController) StartSession() sessions.Session {
	return sessions.StartFasthttp(c.Ctx)
}

func (c *BaseController) View(page Page, contentType string) {
	WritePageTemplate(c.Ctx, page)
	c.Ctx.SetContentType(contentType)
}

func (c *BaseController) RegistRoutes(areaFlag string, ciPtr interface{}) {
	rType := reflect.TypeOf(c)
	ciType := reflect.TypeOf(ciPtr)

	t := strings.Split(strings.ToLower(ciType.String()), ".")

	typeStr := t[len(t)-1]
	typeStr = strings.TrimSuffix(typeStr, controllerFlag)

	for i := 0; i < ciType.NumMethod(); i++ {
		m := ciType.Method(i)
		if _, ok := rType.MethodByName(m.Name); ok {
			continue
		}

		fName := strings.ToLower(m.Name)
		if strings.HasPrefix(fName, "notmap") {
			continue
		}

		var url, preFix string
		var routeAction func(ctx *fasthttp.RequestCtx)

		if strings.HasPrefix(fName, "get") {
			fName = strings.TrimPrefix(fName, "get")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "get"

			routeAction = injectFilter(preFix, url, m, ciPtr)

			router.R.GET(url, routeAction)
		} else if strings.HasPrefix(fName, "post") {
			fName = strings.TrimPrefix(fName, "post")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "post"

			routeAction = injectFilter(preFix, url, m, ciPtr)

			router.R.POST(url, routeAction)
		} else if strings.HasPrefix(fName, "head") {
			fName = strings.TrimPrefix(fName, "head")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "head"
			routeAction = injectFilter(preFix, url, m, ciPtr)
			router.R.HEAD(url, routeAction)
		} else if strings.HasPrefix(fName, "put") {
			fName = strings.TrimPrefix(fName, "put")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "put"
			routeAction = injectFilter(preFix, url, m, ciPtr)
			router.R.PUT(url, routeAction)
		} else if strings.HasPrefix(fName, "options") {
			fName = strings.TrimPrefix(fName, "options")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "options"
			routeAction = injectFilter(preFix, url, m, ciPtr)
			router.R.OPTIONS(url, routeAction)
		} else if strings.HasPrefix(fName, "delete") {
			fName = strings.TrimPrefix(fName, "delete")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "delete"
			routeAction = injectFilter(preFix, url, m, ciPtr)
			router.R.DELETE(url, routeAction)
		} else if strings.HasPrefix(fName, "patch") {
			fName = strings.TrimPrefix(fName, "patch")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "patch"
			routeAction = injectFilter(preFix, url, m, ciPtr)
			router.R.PATCH(url, routeAction)
		} else {
			fName = strings.TrimPrefix(fName, "get")
			url = strings.ToLower(fmt.Sprintf("/%s/%s/%s", areaFlag, typeStr, fName))
			preFix = "get"
			routeAction = injectFilter(preFix, url, m, ciPtr)
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

	oldPath = fmt.Sprintf("%s_%s", httpVerb, strings.ToLower(oldPath))
	for _, newPath := range newPaths {

		newPathTemp := fmt.Sprintf("%s_%s", httpVerb, strings.ToLower(newPath))

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
			//			fmt.Printf("url:[%s] is already mapped to url:[%s].\r\n", newPath, oldPath)
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

func (c *BaseController) ErrorView(areaName, title string, ciPtr interface{}, err error) {
	bpd := NewBasePageData(areaName, title, "", ciPtr)
	bp := NewBasePage(c.Ctx, bpd)

	bpd.Data["error"] = err.Error()
	page := &ErrorPage{bp}
	c.View(page, "text/html")
}
