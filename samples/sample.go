package samples

import (
	cairo "github.com/dtromb/gocairo"
	"reflect"
)

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

func GetSampleTypes() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf([]ArcSample{}).Elem(),
		reflect.TypeOf([]ArcNegativeSample{}).Elem(),
		reflect.TypeOf([]ClipSample{}).Elem(),
		reflect.TypeOf([]ClipImageSample{}).Elem(),
		reflect.TypeOf([]CurveRectangleSample{}).Elem(),
		reflect.TypeOf([]CurveToSample{}).Elem(),
		reflect.TypeOf([]DashSample{}).Elem(),
		reflect.TypeOf([]FillAndStroke2Sample{}).Elem(),
		reflect.TypeOf([]FillStyleSample{}).Elem(),
		reflect.TypeOf([]GradientSample{}).Elem(),
		reflect.TypeOf([]ImageSample{}).Elem(),
		reflect.TypeOf([]ImagePatternSample{}).Elem(),
		reflect.TypeOf([]MultiSegmentCapsSample{}).Elem(),
		reflect.TypeOf([]RoundedRectangleSample{}).Elem(),
		reflect.TypeOf([]SetLineCapSample{}).Elem(),
		reflect.TypeOf([]SetLineJoinSample{}).Elem(),
		reflect.TypeOf([]TextSample{}).Elem(),
		reflect.TypeOf([]TextAlignCenterSample{}).Elem(),
		reflect.TypeOf([]TextExtentsSample{}).Elem(),
	}
}
