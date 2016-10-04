package samples

import (
	"math"
)

type ArcNegativeSample struct {
	CairoSampleImpl
}

func (ans *ArcNegativeSample) Run() {

	xc := 128.0
	yc := 128.0
	radius := 100.0
	angle1 := 45.0 * (math.Pi / 180.0)
	angle2 := 180.0 * (math.Pi / 180.0)

	ctx := ans.CairoContext()

	ctx.SetLineWidth(10.0)
	ctx.ArcNegative(xc, yc, radius, angle1, angle2)
	ctx.Stroke()

	/* draw helping lines */
	ctx.SetSourceRgba(1, 0.2, 0.2, 0.6)
	ctx.SetLineWidth(6.0)
	ctx.Arc(xc, yc, 10.0, 0, 2*math.Pi)
	ctx.Fill()

	ctx.Arc(xc, yc, radius, angle1, angle1)
	ctx.LineTo(xc, yc)
	ctx.Arc(xc, yc, radius, angle2, angle2)
	ctx.LineTo(xc, yc)
	ctx.Stroke()
}
