package samples

import (
	"math"
	cairo "github.com/dtromb/gocairo"
)


type FillStyleSample struct {
	CairoSampleImpl
}

func (as *FillStyleSample) Run() {
	
	ctx := as.CairoContext()
	
	ctx.SetLineWidth(6)
	
	ctx.Rectangle(12, 12, 232, 70)
	ctx.NewSubPath()
	ctx.Arc(64, 64, 40, 0, 2*math.Pi)
	ctx.NewSubPath()
	ctx.ArcNegative(192, 64, 40, 0, -2*math.Pi)
	
	ctx.SetFillRule(cairo.FillRuleEvenOdd)
	ctx.SetSourceRgb(0, 0.7, 0)
	ctx.FillPreserve()
	ctx.SetSourceRgb(0, 0, 0)
	ctx.Stroke()
	
	ctx.Translate(0, 128)
	ctx.Rectangle(12, 12, 232, 70)
	ctx.NewSubPath()
	ctx.Arc(64, 64, 40, 0, 2*math.Pi)
	ctx.NewSubPath()
	ctx.ArcNegative(192, 64, 40, 0, -2*math.Pi)
	
	ctx.SetFillRule(cairo.FillRuleWinding)
	ctx.SetSourceRgb(0, 0, 0.9)
	ctx.FillPreserve()
	ctx.SetSourceRgb(0, 0, 0)
	ctx.Stroke()
}