package samples


type FillAndStroke2Sample struct {
	CairoSampleImpl
}

func (as *FillAndStroke2Sample) Run() {
	
	ctx := as.CairoContext()
	
	ctx.MoveTo(128.0, 25.6)
	ctx.LineTo(230.4, 230.4)
	ctx.RelLineTo(-102.4, 0.0)
	ctx.CurveTo(51.2, 230.4, 51.2, 128.0, 128.0, 128.0)
	ctx.ClosePath()
	
	ctx.MoveTo(64.0, 25.6)
	ctx.RelLineTo(51.2, 51.2)
	ctx.RelLineTo(-51.2, 51.2)
	ctx.RelLineTo(-51.2, -51.2)
	ctx.ClosePath()
	
	ctx.SetLineWidth(10.0)
	ctx.SetSourceRgb(0, 0, 1)
	ctx.FillPreserve()
	ctx.SetSourceRgb(0, 0, 0)
	ctx.Stroke()
}