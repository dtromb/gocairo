package samples

type CurveToSample struct {
	CairoSampleImpl
}

func (as *CurveToSample) Run() {

	ctx := as.CairoContext()

	x := 25.6
	y := 128.0
	x1 := 102.4
	y1 := 230.4
	x2 := 153.6
	y2 := 25.6
	x3 := 230.4
	y3 := 128.0

	ctx.MoveTo(x, y)
	ctx.CurveTo(x1, y1, x2, y2, x3, y3)
	ctx.SetLineWidth(10.0)
	ctx.Stroke()
	ctx.SetSourceRgba(1, 0.2, 0.2, 0.6)
	ctx.SetLineWidth(6.0)
	ctx.MoveTo(x, y)
	ctx.LineTo(x1, y1)
	ctx.MoveTo(x2, y2)
	ctx.LineTo(x3, y3)
	ctx.Stroke()
}
