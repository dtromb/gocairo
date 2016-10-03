package main

// Run all of the cairo samples, generating PDF output.

import (
	"reflect"
	cairo "github.com/dtromb/gocairo"
	"github.com/dtromb/gocairo/samples"
)

func main() {
	
	exampleImpls := []reflect.Type{
		reflect.TypeOf([]samples.ArcSample{}).Elem(),
		reflect.TypeOf([]samples.ArcNegativeSample{}).Elem(),
	}
	
	for _, impl := range exampleImpls {
		surface := cairo.PdfSurfaceCreate(impl.Name()+".pdf", 400, 400)
		sample := reflect.New(impl)
		sample.Elem().FieldByName("CairoSampleImpl").FieldByName("Ctx").Set(reflect.ValueOf(cairo.Create(surface)))
		sample.Interface().(samples.CairoSample).Run()
		surface.Flush()
		surface.Finish()
	}
	
}