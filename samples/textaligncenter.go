package samples

import (
	"math"
	cairo "github.com/dtromb/gocairo"
)

type TextAlignCenterSample struct {
	CairoSampleImpl
}

func (as *TextAlignCenterSample) Run() {

	ctx := as.CairoContext()
	utf8 := "cairo"
	
	ctx.SelectFontFace("Sans", cairo.FontSlantNormal, cairo.FontWeightNormal)
	ctx.SetFontSize(52.0)
	extents := ctx.TextExtents(utf8)
	x := 128.0 - (extents.Width()/2 + extents.XBearing())
	y := 128.0 - (extents.Height()/2 + extents.YBearing())
	
	ctx.MoveTo(x, y)
	ctx.ShowText(utf8)
	
	/* draw helping lines */
	ctx.SetSourceRgba(1, 0.2, 0.2, 0.6)
	ctx.SetLineWidth(6.0)
	ctx.Arc(x, y, 10.0, 0, 2*math.Pi)
	ctx.Fill()
	ctx.MoveTo(128.0, 0)
	ctx.RelLineTo(0, 256)
	ctx.MoveTo(0, 128.0)
	ctx.RelLineTo(256, 0)
	ctx.Stroke()
}