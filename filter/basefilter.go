package filter

import (
	"fasthttpweb/common"
	"fasthttpweb/model"

	"github.com/kataras/go-sessions"
	"github.com/valyala/fasthttp"
)

type FilterStateResult struct {
	FilterState bool
	Data        interface{}
}

type BaseFilter interface {
	GetFilterState() *FilterStateResult
}

type ActionFilter interface {
	BaseFilter
	BeforeExecute(ctx *fasthttp.RequestCtx)
	AfterExecute(ctx *fasthttp.RequestCtx)
}

type AuthorizeFilter interface {
	BaseFilter
	Authorization(ctx *fasthttp.RequestCtx)
}

type ExceptionFilter interface {
	BaseFilter
	OnException(ctx *fasthttp.RequestCtx)
}

type PhoneVerifyFilter struct {
	*FilterStateResult
}

type LoginAuthFilter struct {
	*FilterStateResult
}

type LogFilter struct {
	*FilterStateResult
}

type ErrorFilter struct {
	*FilterStateResult
}

func (filter *PhoneVerifyFilter) Authorization(ctx *fasthttp.RequestCtx) {
	sess := sessions.StartFasthttp(ctx)
	verfy := sess.Get("verfy")
	if verfy == nil {
		ctx.Redirect("/home/index/Verify", fasthttp.StatusNonAuthoritativeInfo)
	}
}

func (filter *PhoneVerifyFilter) GetFilterState() *FilterStateResult {
	return filter.FilterStateResult
}

func (filter *LoginAuthFilter) Authorization(ctx *fasthttp.RequestCtx) {
	sess := sessions.StartFasthttp(ctx)
	verfy := sess.Get("verfy")
	if verfy != nil && verfy.(int) < common.LoginVerfied {
		ctx.Redirect("/home/index/login", fasthttp.StatusNonAuthoritativeInfo)
	}
}

func (filter *LoginAuthFilter) GetFilterState() *FilterStateResult {
	return filter.FilterStateResult
}

func (filter *LogFilter) Authorization(ctx *fasthttp.RequestCtx) {
	log := &model.LogInfo{Url: string(ctx.RequestURI()), IP: ctx.RemoteIP().String(), UserAgent: string(ctx.UserAgent())}
	data := filter.Data
	if data != nil {
		switch data.(type) {
		case string:
			log.Msg = data.(string)
		}
	}
	log.Add()
}

func (filter *LogFilter) GetFilterState() *FilterStateResult {
	return filter.FilterStateResult
}

func (filter *ErrorFilter) OnException(ctx *fasthttp.RequestCtx) {
	if err := filter.Data; filter.FilterState && err != nil {
		switch err.(type) {
		case error:
			log := &model.LogInfo{Url: string(ctx.RequestURI()), IP: ctx.RemoteIP().String(), UserAgent: string(ctx.UserAgent()), LogCatelog: "sys_error", Msg: string(ctx.Method())}
			log.ExpMsg = err.(error).Error()
			log.Add()
		}
	}
}

func (filter *ErrorFilter) GetFilterState() *FilterStateResult {
	return filter.FilterStateResult
}
