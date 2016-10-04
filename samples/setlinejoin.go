package samples

import (
	cairo "github.com/dtromb/gocairo"
)

type SetLineJoinSample struct {
	CairoSampleImpl
}

func (as *SetLineJoinSample) Run() {

	ctx := as.CairoContext()

	ctx.SetLineWidth(40.96)

	ctx.MoveTo(76.8, 84.48)
	ctx.RelLineTo(51.2, -51.2)
	ctx.RelLineTo(51.2, 51.2)
	ctx.SetLineJoin(cairo.LineJoinMiter)
	ctx.Stroke()

	ctx.MoveTo(76.8, 161.28)
	ctx.RelLineTo(51.2, -51.2)
	ctx.RelLineTo(51.2, 51.2)
	ctx.SetLineJoin(cairo.LineJoinBevel)
	ctx.Stroke()

	ctx.MoveTo(76.8, 238.08)
	ctx.RelLineTo(51.2, -51.2)
	ctx.RelLineTo(51.2, 51.2)
	ctx.SetLineJoin(cairo.LineJoinRound)
	ctx.Stroke()
}
