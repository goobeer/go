package main

import (
	"fasthttpweb/router"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"reflect"
	"runtime"

	_ "fasthttpweb/area/admin/controllers"
	_ "fasthttpweb/area/home/controllers"

	"github.com/valyala/fasthttp"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func main() {
	test()
}

func test() {
	router.R.ServeFiles("/public/*filepath", "./public")
	fasthttp.ListenAndServe(":8080", router.R.Handler)
}

type Putpixel func(x, y int)

type A struct {
}

type BA struct {
	A
}

func (a A) Show(msg string) {
	fmt.Println(">>>", msg)
}

func (a A) Hello() {
	rType := reflect.TypeOf(a)
	p, f, l, ok := runtime.Caller(0)
	fmt.Println(p, f, l, ok, runtime.FuncForPC(p).Name())
	p, f, l, ok = runtime.Caller(1)
	fmt.Println(p, f, l, ok, runtime.FuncForPC(p).Name())
	p, f, l, ok = runtime.Caller(2)
	fmt.Println(p, f, l, ok, runtime.FuncForPC(p).Name())
	fmt.Println(rType, rType.NumMethod())
}

func (b BA) Woda() {

}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func drawline(x0, y0, x1, y1 int, brush Putpixel) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	sx, sy := 1, 1
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}
	err := dx - dy

	for {
		brush(x0, y0)
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

func ExampleDraw() {
	rect := image.Rect(0, 0, 100, 100)
	nrgba := image.NewRGBA(rect)

	addLabel(nrgba, 20, 30, "xiexx")

	file, _ := os.Create("a.png")
	defer file.Close()
	err := png.Encode(file, nrgba)

	if err != nil {
		fmt.Println("err focus")
		return
	}
}
