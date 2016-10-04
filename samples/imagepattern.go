package samples

import (
	cairo "github.com/dtromb/gocairo"
	"math"
)

type ImagePatternSample struct {
	CairoSampleImpl
}

func (as *ImagePatternSample) Run() {

	ctx := as.CairoContext()

	image := cairo.ImageSurfaceCreateFromPng("../romedalen.png")
	w := float64(image.GetWidth())
	h := float64(image.GetHeight())

	pattern := cairo.PatternCreateForSurface(image)
	pattern.SetExtend(cairo.ExtendRepeat)

	ctx.Translate(128.0, 128.0)
	ctx.Rotate(math.Pi / 4)
	ctx.Scale(1/math.Sqrt2, 1/math.Sqrt2)
	ctx.Translate(-128.0, -128.0)

	matrix := cairo.NewMatrix()
	matrix.InitScale(w/256.0*5.0, h/256.0*5.0)
	pattern.SetMatrix(matrix)

	ctx.SetSource(pattern)
	ctx.Rectangle(0, 0, 256.0, 256.0)
	ctx.Fill()
}
