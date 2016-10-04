package samples

type DashSample struct {
	CairoSampleImpl
}

func (as *DashSample) Run() {

	ctx := as.CairoContext()

	dashes := []float64{50.0, 10.0, 10.0, 10.0}
	offset := -50.0

	ctx.SetDash(dashes, offset)
	ctx.SetLineWidth(10.0)
	ctx.MoveTo(128.0, 25.6)
	ctx.LineTo(230.4, 230.4)
	ctx.RelLineTo(-102.4, 0.0)
	ctx.CurveTo(51.2, 230.4, 51.2, 128.0, 128.0, 128.0)
	ctx.Stroke()
}
