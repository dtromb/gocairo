package samples

import (
	cairo "github.com/dtromb/gocairo"
	"math"
)

type GradientSample struct {
	CairoSampleImpl
}

func (as *GradientSample) Run() {

	ctx := as.CairoContext()
	pat := cairo.PatternCreateLinear(0.0, 0.0, 0.0, 256.0)
	pat.AddColorStopRgba(1, 0, 0, 0, 1)
	pat.AddColorStopRgba(0, 1, 1, 1, 1)
	ctx.Rectangle(0, 0, 256, 256)
	ctx.SetSource(pat)
	ctx.Fill()

	pat = cairo.PatternCreateRadial(115.2, 102.4, 25.6, 102.4, 102.4, 128.0)
	pat.AddColorStopRgba(0, 1, 1, 1, 1)
	pat.AddColorStopRgba(1, 0, 0, 0, 1)
	ctx.SetSource(pat)
	ctx.Arc(128.0, 128.0, 76.8, 0, 2*math.Pi)
	ctx.Fill()
}
