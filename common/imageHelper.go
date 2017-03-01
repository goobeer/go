package common

import (
	"image"
	"image/draw"
	"io/ioutil"
	"math"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

//var (
//	//none | full
////	hinting string = "none"

//	//line spacing (e.g. 2 means double spaced)
//	//	spacing float64 = 1

//	//	title string = "卡的"

////	text []string = []string{"abc", "汉子"}
//)

type ImageVerfiyConfig struct {
	FontFile      string         //filename of the ttf font
	FontSize      float64        //font size in points
	DPI           float64        //screen resolution in Dots Per Inch
	W             int            //img width
	H             int            //img height
	Fg            *image.Uniform //fg
	Bg            *image.Uniform //bg
	ImgHitingMode font.Hinting   //ImgHitingMode
	LineSpacing   float64        //line spacing (e.g. 2 means double spaced)
}

//验证图片
func NewVerifyImg(conf *ImageVerfiyConfig, verifyCode ...string) (rgba *image.RGBA, err error) {
	// Read the font data.
	fontBytes, err := ioutil.ReadFile(conf.FontFile)
	if err != nil {
		return
	}

	ttf, err := truetype.Parse(fontBytes)
	if err != nil {
		return
	}

	// Draw the background and the guidelines.
	//	ruler := color.RGBA{0xdd, 0xdd, 0xdd, 0xff}

	rgba = image.NewRGBA(image.Rect(0, 0, conf.W, conf.H))
	draw.Draw(rgba, rgba.Bounds(), conf.Bg, image.ZP, draw.Src)
	//	for i := 0; i < 200; i++ {
	//		rgba.Set(10, 10+i, ruler)
	//		rgba.Set(10+i, 10, ruler)
	//	}

	// Draw the text.

	fontDrawer := &font.Drawer{
		Dst: rgba,
		Src: conf.Fg,
		Face: truetype.NewFace(ttf, &truetype.Options{
			Size:    conf.FontSize,
			DPI:     conf.DPI,
			Hinting: conf.ImgHitingMode,
		}),
	}

	y := 10 + int(math.Ceil(conf.FontSize*conf.DPI/72))
	//字体尺寸+行间距
	dy := int(math.Ceil(conf.FontSize * conf.LineSpacing * conf.DPI / 72))

	//	fontDrawer.Dot = fixed.Point26_6{
	//		X: (fixed.I(conf.W) - fontDrawer.MeasureString(title)) / 2,
	//		Y: fixed.I(y),
	//	}

	//	fontDrawer.DrawString(title)

	y += dy
	for _, s := range verifyCode {
		fontDrawer.Dot = fixed.P(5, y)
		fontDrawer.DrawString(s)
		y += dy
	}

	return
}
