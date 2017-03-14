package filter

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

const (
	Get    = iota
	Post   = iota
	Put    = iota
	Delete = iota
)

var (
	MapFilter = make(map[string]*RouteFilter)
)

type RouteFilter struct {
	BeforeRequestHandlers []fasthttp.RequestHandler
	AfterRequestHandlers  []fasthttp.RequestHandler
	Uri                   string
	HttpVerb              string
}

func init() {
	logFilter := NewRouteFilter("/home/index/index")
	logFilter.BeforeRequestHandlers = append(logFilter.BeforeRequestHandlers, logFilter.Log(nil))
	MapFilter["/home/index/index"] = logFilter
}

func NewRouteFilter(uri string) *RouteFilter {
	return &RouteFilter{Uri: uri, BeforeRequestHandlers: []fasthttp.RequestHandler{}, AfterRequestHandlers: []fasthttp.RequestHandler{}}
}

func (r *RouteFilter) Log(data interface{}) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		fmt.Println(data, string(ctx.RequestURI()))
	}
}

func (r *RouteFilter) Err(err error, data interface{}) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if x := recover(); x != nil {

			}
		}()
	}
}

func (r *RouteFilter) Auth() {

}
