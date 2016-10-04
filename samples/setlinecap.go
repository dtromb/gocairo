package samples

import (
	cairo "github.com/dtromb/gocairo"
)

type SetLineCapSample struct {
	CairoSampleImpl
}

func (as *SetLineCapSample) Run() {

	ctx := as.CairoContext()

	ctx.SetLineWidth(30.0)
	ctx.SetLineCap(cairo.LineCapButt)
	ctx.MoveTo(64.0, 50.0)
	ctx.LineTo(64.0, 200.0)
	ctx.Stroke()
	ctx.SetLineCap(cairo.LineCapRound)
	ctx.MoveTo(128.0, 50.0)
	ctx.LineTo(128.0, 200.0)
	ctx.Stroke()
	ctx.SetLineCap(cairo.LineCapSquare)
	ctx.MoveTo(192.0, 50.0)
	ctx.LineTo(192.0, 200.0)
	ctx.Stroke()

	/* draw helping lines */
	ctx.SetSourceRgb(1, 0.2, 0.2)
	ctx.SetLineWidth(2.56)
	ctx.MoveTo(64.0, 50.0)
	ctx.LineTo(64.0, 200.0)
	ctx.MoveTo(128.0, 50.0)
	ctx.LineTo(128.0, 200.0)
	ctx.MoveTo(192.0, 50.0)
	ctx.LineTo(192.0, 200.0)
	ctx.Stroke()
}
