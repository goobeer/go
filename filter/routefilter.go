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

	Filters []BaseFilter

	FilterExp    string
	HttpVerbType int
	MatchType    int
	MatchOnOff   bool //匹配方式:true:匹配成功执行过滤器，false:匹配失败执行过滤器
}

func init() {
	comFilter := NewRouteFilter("^/home/index/verify", Get, MatchReg, false)
	phoneFilter := &PhoneVerifyFilter{FilterStateResult: &FilterStateResult{}}
	logFilter := &LogFilter{FilterStateResult: &FilterStateResult{}}
	performFilter := &PerformanceMonitorFilter{FilterStateResult: &FilterStateResult{}}
	comFilter.Filters = append(comFilter.Filters, phoneFilter, logFilter, performFilter)

	comFilter1 := NewRouteFilter("[^/home/index/verify|/home/index/login]", Get, MatchReg, true)
	loginFilter := &LoginAuthFilter{FilterStateResult: &FilterStateResult{}}
	comFilter1.Filters = append(comFilter1.Filters, loginFilter)

	MapFilter[comFilter.generateFilterKey()] = comFilter
	MapFilter[comFilter1.generateFilterKey()] = comFilter1
}

func NewRouteFilter(filterExp string, httpVerbType, matchType int, matchOnOff bool) *RouteFilter {
	return &RouteFilter{Filters: make([]BaseFilter, 0), FilterExp: filterExp, HttpVerbType: httpVerbType, MatchType: matchType, MatchOnOff: matchOnOff}
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
			if rf.MatchOnOff {
				isMatched = uri == rf.FilterExp
			} else {
				isMatched = uri != rf.FilterExp
			}

		case MatchReg:
			reg, err := regexp.Compile(rf.FilterExp)
			if err != nil {
				panic(err)
			}
			if rf.MatchOnOff {
				isMatched = reg.MatchString(uri)
			} else {
				isMatched = !reg.MatchString(uri)
			}

		default:
			isMatched = false
		}
	}
	return isMatched
}
