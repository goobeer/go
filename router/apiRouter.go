package router

import (
	"fasthttpweb/common"
	"image"
	"image/color"
	"image/png"

	"github.com/kataras/go-sessions"
	"github.com/valyala/fasthttp"
)

func init() {
	R.GET("/api/verifyimg", getVerifyImg)
}

func getVerifyImg(ctx *fasthttp.RequestCtx) {
	verifyCode, err := common.CreateRandomFromSeed(5, common.RandSeed)
	if err != nil {
		common.PanicError(err)
	}
	sess := sessions.StartFasthttp(ctx)
	sess.Set("vcode", verifyCode)
	conf := &common.ImageVerfiyConfig{FontFile: "./public/simhei.ttf", FontSize: 16, DPI: 72, W: 80, H: 35}
	conf.Fg = &image.Uniform{color.Black}
	conf.Bg = &image.Uniform{color.White}
	rgba, err := common.NewVerifyImg(conf, verifyCode)
	if err != nil {
		common.PanicError(err)
	}

	err = png.Encode(ctx.Response.BodyWriter(), rgba)
	if err != nil {
		common.PanicError(err)
	}

	ctx.SetContentType("image/png")
}
