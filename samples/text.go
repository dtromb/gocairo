package samples

import (
	"math"
	cairo "github.com/dtromb/gocairo"
)

type TextSample struct {
	CairoSampleImpl
}

func (as *TextSample) Run() {

	ctx := as.CairoContext()
	
	ctx.SelectFontFace("Sans", cairo.FontSlantNormal, cairo.FontWeightBold)
	ctx.SetFontSize(90.0)
	ctx.MoveTo(10.0, 135.0)
	ctx.ShowText("Hello")
	
	ctx.MoveTo(70.0, 165.0)
	ctx.TextPath("void")
	ctx.SetSourceRgb(0.5, 0.5, 1)
	ctx.FillPreserve()
	ctx.SetSourceRgb(0, 0, 0)
	ctx.SetLineWidth(2.56)
	ctx.Stroke()
	
	/* draw helping lines */
	ctx.SetSourceRgba(1, 0.2, 0.2, 0.6)
	ctx.Arc(10.0, 135.0, 5.12, 0, 2*math.Pi)
	ctx.ClosePath()
	ctx.Arc(70.0, 165.0, 5.12, 0, 2*math.Pi)
	ctx.Fill()
}