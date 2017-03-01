package common

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func TestNewVerifyImg(t *testing.T) {
	conf := &ImageVerfiyConfig{FontFile: "../public/simhei.ttf", FontSize: 12, DPI: 72, W: 80, H: 60}
	conf.Fg = &image.Uniform{color.Black}
	conf.Bg = &image.Uniform{color.White}
	t.Log(conf)
	rgba, err := NewVerifyImg(conf, "中国中国中国中国中国")
	if err != nil {
		t.Error(err)
	} else {
		t.Log("success", rgba)

		// Save that RGBA image to disk.
		outFile, err := os.Create("out.png")
		if err != nil {
			return
		}
		defer outFile.Close()
		b := bufio.NewWriter(outFile)

		err = png.Encode(b, rgba)
		if err != nil {
			return
		}
		err = b.Flush()
		if err != nil {
			return
		}
	}
}
