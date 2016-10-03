package samples

import cairo "github.com/dtromb/gocairo"

type CairoSample interface {
	CairoContext() cairo.Cairo
	Run()
}

type CairoSampleImpl struct {
	Ctx cairo.Cairo
}

func (impl *CairoSampleImpl) CairoContext() cairo.Cairo {
	return impl.Ctx
}


