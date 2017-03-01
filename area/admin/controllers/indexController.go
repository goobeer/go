package controllers

import (
	"fasthttpweb/area"
	av "fasthttpweb/area/admin/views/index"

	//	"github.com/kataras/go-sessions"
	"github.com/valyala/fasthttp"
)

type IndexController struct {
}

func getIndex(ctx *fasthttp.RequestCtx) {
	ip := &av.IndexPage{CTX: ctx}
	area.WritePageTemplate(ctx, ip)

	ctx.SetContentType("text/html")
}
