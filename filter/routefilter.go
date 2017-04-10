package filter

import (
	"fmt"
	"regexp"

	"github.com/valyala/fasthttp"
)

const (
	//httpverb 类型
	Get     = 1
	Post    = 2
	Put     = 4
	Delete  = 8
	Head    = 16
	Patch   = 32
	Options = 64
	Normal  = Get | Post
	All     = Get | Post | Put | Delete | Head | Patch | Options

	//过滤匹配 模型[支持 完全匹配、正则匹配]
	MatchFull = iota
	MatchReg
)

var (
	MapFilter = make(map[string]*RouteFilter)
)

type RouteFilter struct {
	BeforeRequestHandlers []fasthttp.RequestHandler
	AfterRequestHandlers  []fasthttp.RequestHandler
	FilterExp             string
	HttpVerbType          int
	MatchType             int
}

func init() {
	comFilter := NewRouteFilter("[^/home/index/verify]", All, MatchReg)
	comFilter.BeforeRequestHandlers = append(comFilter.BeforeRequestHandlers, comFilter.Log(nil))

	MapFilter[comFilter.generateFilterKey()] = comFilter
}

func NewRouteFilter(filterExp string, httpVerbType, matchType int) *RouteFilter {
	return &RouteFilter{FilterExp: filterExp, HttpVerbType: httpVerbType, MatchType: matchType, BeforeRequestHandlers: []fasthttp.RequestHandler{}, AfterRequestHandlers: []fasthttp.RequestHandler{}}
}

func (rf *RouteFilter) generateFilterKey() (filterKey string) {
	filterKey = fmt.Sprintf("%s%d_%d", rf.FilterExp, rf.HttpVerbType, rf.MatchType)
	return
}

func (rf *RouteFilter) FilterMatch(uri, httpVerb string) bool {
	isMatched := false
	switch rf.HttpVerbType {
	case Get:
		isMatched = httpVerb == "get"
	case Post:
		isMatched = httpVerb == "post"
	case Put:
		isMatched = httpVerb == "put"
	case Delete:
		isMatched = httpVerb == "delete"
	case Head:
		isMatched = httpVerb == "head"
	case Patch:
		isMatched = httpVerb == "patch"
	case Options:
		isMatched = httpVerb == "options"
	case Normal:
		isMatched = httpVerb == "get" || httpVerb == "post"
	case All:
		isMatched = httpVerb == "get" || httpVerb == "post" || httpVerb == "put" || httpVerb == "delete" || httpVerb == "head" || httpVerb == "patch" || httpVerb == "options"
	}

	if isMatched {
		switch rf.MatchType {
		case MatchFull:
			isMatched = uri == rf.FilterExp
		case MatchReg:
			isMatched, _ = regexp.MatchString(rf.FilterExp, uri)

		default:
			isMatched = false
		}
	}
	return isMatched
}

func (r *RouteFilter) Log(data interface{}) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		fmt.Println(data, string(ctx.Method()), string(ctx.RequestURI()), ctx.RemoteIP().String(), string(ctx.UserAgent()))
	}
}

func (r *RouteFilter) Auth(data interface{}) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		fmt.Println("auth processing")
	}
}
