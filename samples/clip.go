package samples

import (
	"math"
)

type ClipSample struct {
	CairoSampleImpl
}

func (as *ClipSample) Run() {

	ctx := as.CairoContext()

	ctx.Arc(128.0, 128.0, 76.8, 0, 2*math.Pi)
	ctx.Clip()
	ctx.NewPath()
	ctx.Rectangle(0, 0, 256, 256)
	ctx.Fill()
	ctx.SetSourceRgb(0, 1, 0)
	ctx.MoveTo(0, 0)
	ctx.LineTo(256, 256)
	ctx.MoveTo(256, 0)
	ctx.LineTo(0, 256)
	ctx.SetLineWidth(10.0)
	ctx.Stroke()
}
