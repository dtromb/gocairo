package samples

import (
	cairo "github.com/dtromb/gocairo"
)

type MultiSegmentCapsSample struct {
	CairoSampleImpl
}

func (as *MultiSegmentCapsSample) Run() {

	ctx := as.CairoContext()

	ctx.MoveTo(50.0, 75.0)
	ctx.LineTo(200.0, 75.0)

	ctx.MoveTo(50.0, 125.0)
	ctx.LineTo(200.0, 125.0)

	ctx.MoveTo(50.0, 175.0)
	ctx.LineTo(200.0, 175.0)

	ctx.SetLineWidth(30.0)
	ctx.SetLineCap(cairo.LineCapRound)
	ctx.Stroke()

}
