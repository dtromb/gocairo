package main

// Run all of the cairo samples, generating PDF output.

import (
	"reflect"
	cairo "github.com/dtromb/gocairo"
	"github.com/dtromb/gocairo/samples"
)

func main() {
	
	for _, impl := range samples.GetSampleTypes() {
		surface := cairo.PdfSurfaceCreate(impl.Name()+".pdf", 256, 256)
		sample := reflect.New(impl)
		sample.Elem().FieldByName("CairoSampleImpl").FieldByName("Ctx").Set(reflect.ValueOf(cairo.Create(surface)))
		sample.Interface().(samples.CairoSample).Run()
		surface.Flush()
		surface.Finish()
	}
	
}