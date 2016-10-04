package samples

import (
	cairo "github.com/dtromb/gocairo"
	"math"
)

type ImageSample struct {
	CairoSampleImpl
}

func (as *ImageSample) Run() {

	ctx := as.CairoContext()

	image := cairo.ImageSurfaceCreateFromPng("../romedalen.png")
	w := float64(image.GetWidth())
	h := float64(image.GetHeight())
	ctx.Translate(128.0, 128.0)
	ctx.Rotate(45 * (math.Pi / 180.0))
	ctx.Scale(256.0/w, 256.0/h)
	ctx.Translate(-0.5*w, -0.5*h)
	ctx.SetSourceSurface(image, 0, 0)
	ctx.Paint()
}
