package router

import (
	"encoding/hex"
	"fasthttpweb/common"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"strings"

	"github.com/kataras/go-sessions"
	"github.com/valyala/fasthttp"
)

func init() {
	R.GET("/api/verifyimg", getVerifyImg)
}

func getVerifyImg(ctx *fasthttp.RequestCtx) {
	var err error
	var verifyCode string

	key := ctx.QueryArgs().QueryString()
	if !(key != nil && len(key) > 0) {
		return
	}
	fmt.Println(string(key))
	keyVal := string(key)
	if strings.Index(keyVal, "&") > 0 {
		key = []byte(strings.Split(keyVal, "&")[0])
	}
	dst := make([]byte, hex.DecodedLen(len(key)))
	_, err = hex.Decode(dst, key)
	if err != nil {
		return
	}
	fmt.Println("[sb]", dst)
	dst, err = common.TripleDesDecrypt(dst, nil)
	if err != nil {
		return
	}

	verifyCode, err = common.CreateRandomFromSeed(5, common.RandSeed)
	if err != nil {
		panic(err.Error())
	}
	sess := sessions.StartFasthttp(ctx)

	sess.Set(string(dst), verifyCode)

	conf := &common.ImageVerfiyConfig{FontFile: "./public/simhei.ttf", FontSize: 16, DPI: 72, W: 80, H: 35}
	conf.Fg = &image.Uniform{color.Black}
	conf.Bg = &image.Uniform{color.White}
	rgba, err := common.NewVerifyImg(conf, verifyCode)
	if err != nil {
		panic(err.Error())
	}

	err = png.Encode(ctx.Response.BodyWriter(), rgba)
	if err != nil {
		panic(err.Error())
	}

	ctx.SetContentType("image/png")
}
