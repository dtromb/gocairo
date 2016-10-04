package main

// Run all of the cairo samples, generating PDF output.

import (
	"reflect"
	cairo "github.com/dtromb/gocairo"
	"github.com/dtromb/gocairo/samples"
)

func main() {

	for _, impl := range samples.GetSampleTypes() {
		surface := cairo.ImageSurfaceCreate(cairo.FormatRgb16, 256, 256)
		sample := reflect.New(impl)
		ctx := cairo.Create(surface)
		ctx.Save()
		ctx.SetLineWidth(0.5)
		ctx.SetSourceRgb(0.8,0.8,0.8)
		ctx.Rectangle(0,0,256,256)
		ctx.Fill()
		ctx.Restore()
		sample.Elem().FieldByName("CairoSampleImpl").FieldByName("Ctx").Set(reflect.ValueOf(ctx))
		sample.Interface().(samples.CairoSample).Run()
		surface.WriteToPng(impl.Name()+".png")
		surface.Flush()
		surface.Finish()
	}
	
}