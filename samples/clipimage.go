package samples

import (
	"math"
	cairo "github.com/dtromb/gocairo"
)

type ClipImageSample struct {
	CairoSampleImpl
}

func (as *ClipImageSample) Run() {
	
	ctx := as.CairoContext()
	ctx.Arc(128.0, 128.0, 76.8, 0, 2*math.Pi)
	ctx.Clip()
	ctx.NewPath()
	
	image := cairo.ImageSurfaceCreateFromPng("../romedalen.png")
	if image.Status() != cairo.StatusSuccess {
		panic(image.Status().String())
	}
	w := float64(image.GetWidth())
	h := float64(image.GetHeight())
	ctx.Scale(256.0/w, 256.0/h)
	ctx.SetSourceSurface(image, 0, 0)
	ctx.Paint()
}