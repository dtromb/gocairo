package samples

type CurveRectangleSample struct {
	CairoSampleImpl
}

func (as *CurveRectangleSample) Run() {

	ctx := as.CairoContext()

	x0 := 25.6
	y0 := 25.6
	rectWidth := 204.8
	rectHeight := 204.8
	radius := 102.4

	x1 := x0 + rectWidth
	y1 := y0 + rectHeight

	if rectWidth <= 0 || rectHeight <= 0 {
		return
	}

	if rectWidth/2 < radius {
		if rectHeight/2 < radius {
			ctx.MoveTo(x0, (y0+y1)/2)
			ctx.CurveTo(x0, y0, x0, y0, (x0+x1)/2, y0)
			ctx.CurveTo(x1, y0, x1, y0, x1, (y0+y1)/2)
			ctx.CurveTo(x1, y1, x1, y1, (x1+x0)/2, y1)
			ctx.CurveTo(x0, y1, x0, y1, x0, (y0+y1)/2)
		} else {
			ctx.MoveTo(x0, y0+radius)
			ctx.CurveTo(x0, y0, x0, y0, (x0+x1)/2, y0)
			ctx.CurveTo(x1, y0, x1, y0, x1, y0+radius)
			ctx.LineTo(x1, y1-radius)
			ctx.CurveTo(x1, y1, x1, y1, (x1+x0)/2, y1)
			ctx.CurveTo(x0, y1, x0, y1, x0, y1-radius)
		}
	} else {
		if rectHeight/2 < radius {
			ctx.MoveTo(x0, (y0+y1)/2)
			ctx.CurveTo(x0, y0, x0, y0, x0+radius, y0)
			ctx.LineTo(x1-radius, y0)
			ctx.CurveTo(x1, y0, x1, y0, x1, (y0+y1)/2)
			ctx.CurveTo(x1, y1, x1, y1, x1-radius, y1)
			ctx.LineTo(x0+radius, y1)
			ctx.CurveTo(x0, y1, x0, y1, x0, (y0+y1)/2)
		} else {
			ctx.MoveTo(x0, y0+radius)
			ctx.CurveTo(x0, y0, x0, y0, x0+radius, y0)
			ctx.LineTo(x1-radius, y0)
			ctx.CurveTo(x1, y0, x1, y0, x1, y0+radius)
			ctx.LineTo(x1, y1-radius)
			ctx.CurveTo(x1, y1, x1, y1, x1-radius, y1)
			ctx.LineTo(x0+radius, y1)
			ctx.CurveTo(x0, y1, x0, y1, x0, y1-radius)
		}
	}
	ctx.ClosePath()
	ctx.SetSourceRgb(0.5, 0.5, 1)
	ctx.FillPreserve()
	ctx.SetSourceRgba(0.5, 0, 0, 0.5)
	ctx.SetLineWidth(10.0)
	ctx.Stroke()
}
