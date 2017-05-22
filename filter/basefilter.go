package filter

import (
	"fasthttpweb/common"
	"fasthttpweb/model"
	"regexp"

	"github.com/kataras/go-sessions"
	"github.com/valyala/fasthttp"
)

var (
	spaceReg, _ = regexp.Compile(">[.\\s\r\n\t]*?<")
	replaceTag  = []byte{'>', '<'}
)

type FilterStateResult struct {
	FilterState bool
	Data        interface{}
}

type BaseFilter interface{}

type BeforeActionFilter interface {
	BeforeExecute(ctx *fasthttp.RequestCtx) *FilterStateResult
}

type AfterActionFilter interface {
	AfterExecute(ctx *fasthttp.RequestCtx) *FilterStateResult
}

type ActionFilter interface {
	BaseFilter
	BeforeActionFilter
	AfterActionFilter
}

type AuthorizeFilter interface {
	BaseFilter
	Authorization(ctx *fasthttp.RequestCtx) *FilterStateResult
}

type ExceptionFilter interface {
	BaseFilter
	OnException(ctx *fasthttp.RequestCtx) *FilterStateResult
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

type CompressFilter struct {
	*FilterStateResult
}

func (filter *PhoneVerifyFilter) Authorization(ctx *fasthttp.RequestCtx) *FilterStateResult {
	sess := sessions.StartFasthttp(ctx)
	verfy := sess.Get("verfy")
	if verfy == nil {
		ctx.Redirect("/home/index/Verify", fasthttp.StatusNonAuthoritativeInfo)
	} else {
		filter.FilterState = true
	}
	return filter.FilterStateResult
}

func (filter *LoginAuthFilter) Authorization(ctx *fasthttp.RequestCtx) *FilterStateResult {
	sess := sessions.StartFasthttp(ctx)
	verfy := sess.Get("verfy")
	if verfy != nil && verfy.(int) < common.LoginVerfied {
		ctx.Redirect("/home/index/login", fasthttp.StatusNonAuthoritativeInfo)
	} else {
		filter.FilterState = true
	}
	return filter.FilterStateResult
}

func (filter *LogFilter) BeforeExecute(ctx *fasthttp.RequestCtx) *FilterStateResult {
	log := &model.LogInfo{Url: string(ctx.RequestURI()), IP: ctx.RemoteIP().String(), UserAgent: string(ctx.UserAgent())}
	data := filter.Data
	if data != nil {
		switch data.(type) {
		case string:
			log.Msg = data.(string)
			filter.FilterState = true
		}
	}
	log.Add()
	return filter.FilterStateResult
}

func (filter *ErrorFilter) OnException(ctx *fasthttp.RequestCtx) *FilterStateResult {
	if err := filter.Data; filter.FilterState && err != nil {
		switch err.(type) {
		case error:
			log := &model.LogInfo{Url: string(ctx.RequestURI()), IP: ctx.RemoteIP().String(), UserAgent: string(ctx.UserAgent()), LogCatelog: "sys_error", Msg: string(ctx.Method())}
			log.ExpMsg = err.(error).Error()
			log.Add()
			filter.FilterState = true
		}
	}
	return filter.FilterStateResult
}

func (filter *CompressFilter) AfterExecute(ctx *fasthttp.RequestCtx) *FilterStateResult {
	body := ctx.Response.Body()
	body = spaceReg.ReplaceAll(body, replaceTag)

	ctx.ResetBody()
	ctx.Response.SetBody(body)
	return filter.FilterStateResult
}
