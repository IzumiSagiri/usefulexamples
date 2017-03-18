package main

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"log"
)

func main() {
	dest := image.NewRGBA(image.Rect(0, 0, 10, 10))
	gc := draw2dimg.NewGraphicContext(dest)
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetLineWidth(5)

	gc.MoveTo(0, 0)
	gc.LineTo(10, 0)
	gc.LineTo(10, 10)
	gc.LineTo(0, 0)
	gc.FillStroke()

	log.Print(dest)
}
