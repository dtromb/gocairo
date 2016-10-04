package samples

import (
	"math"
)

type RoundedRectangleSample struct {
	CairoSampleImpl
}

func (as *RoundedRectangleSample) Run() {

	ctx := as.CairoContext()

	x := 25.6
	y := 25.6
	width := 204.8
	height := 204.8
	aspect := 1.0
	cornerRadius := height / 10.0

	radius := cornerRadius / aspect
	degrees := math.Pi / 180.0

	ctx.NewSubPath()
	ctx.Arc(x+width-radius, y+radius, radius, -90*degrees, 0*degrees)
	ctx.Arc(x+width-radius, y+height-radius, radius, 0*degrees, 90*degrees)
	ctx.Arc(x+radius, y+height-radius, radius, 90*degrees, 180*degrees)
	ctx.Arc(x+radius, y+radius, radius, 180*degrees, 270*degrees)
	ctx.ClosePath()

	ctx.SetSourceRgb(0.5, 0.5, 1)
	ctx.FillPreserve()
	ctx.SetSourceRgba(0.5, 0, 0, 0.5)
	ctx.SetLineWidth(10.0)
	ctx.Stroke()
}
