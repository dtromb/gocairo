package samples

import (
	"math"
	cairo "github.com/dtromb/gocairo"
)

type TextExtentsSample struct {
	CairoSampleImpl
}

func (as *TextExtentsSample) Run() {
	ctx := as.CairoContext()

	utf8 := "cairo"
	ctx.SelectFontFace("Sans", cairo.FontSlantNormal, cairo.FontWeightNormal)
	ctx.SetFontSize(100.0)
	extents := ctx.TextExtents(utf8)
	
	x := 25.0
	y := 150.0
	
	ctx.MoveTo(x, y)
	ctx.ShowText(utf8)

	/* draw helping lines */
	ctx.SetSourceRgba(1, 0.2, 0.2, 0.6)
	ctx.SetLineWidth(6.0)
	ctx.Arc(x, y, 10.0, 0, 2*math.Pi)
	ctx.Fill()
	
	ctx.MoveTo(x, y)
	ctx.RelLineTo(0, -extents.Height())
	ctx.RelLineTo(extents.Width(), 0)
	ctx.RelLineTo(extents.XBearing(), -extents.YBearing())
	ctx.Stroke()
}