package cairo

import (
	"runtime"
)

/*
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
	#include <inttypes.h>
	#include <stdlib.h>
*/
import "C"

type FontOptions interface {
	Copy() FontOptions
	Status() Status
	Merge(fromOptions FontOptions)
	Hash() uint64
	Equals(other FontOptions) bool
	SetAntialias(aaMode Antialias)
	GetAntialias() Antialias
	SetSubpixelOrder(spoMode SubpixelOrder)
	GetSubpixelOrder() SubpixelOrder
	SetHintStyle(hsMode HintStyle)
	GetHintStyle() HintStyle
	SetHintMetrics(hmMode HintMetrics)
	GetHintMetrics() HintMetrics
}

type stdFontOptions struct {
	hnd *C.cairo_font_options_t
}

func destroyFontOptions(fo FontOptions) {
	if fn, ok := fo.(Finalizable); ok {
		fn.Finalize(fo)
	}
	if sfo, ok := fo.(*stdFontOptions); ok {
		C.cairo_font_options_destroy(sfo.hnd)
		sfo.hnd = nil
	}
}

func NewFontOptions() FontOptions {
	nfo := &stdFontOptions{
		hnd: C.cairo_font_options_create(),
	}
	runtime.SetFinalizer(nfo, destroyFontOptions)
	return nfo
}

func (fo *stdFontOptions) Copy() FontOptions {
	nfo := &stdFontOptions{
		hnd: C.cairo_font_options_copy(fo.hnd),
	}
	runtime.SetFinalizer(nfo, destroyFontOptions)
	return nfo
}

func (fo *stdFontOptions) Status() Status {
	return Status(C.cairo_font_options_status(fo.hnd))
}

func (fo *stdFontOptions) Merge(fromOptions FontOptions) {
	if sfo, ok := fromOptions.(*stdFontOptions); ok {
		C.cairo_font_options_merge(fo.hnd, sfo.hnd)
	} else {
		sfo := NewFontOptions()
		sfo.SetAntialias(fromOptions.GetAntialias())
		sfo.SetHintMetrics(fromOptions.GetHintMetrics())
		sfo.SetHintStyle(fromOptions.GetHintStyle())
		sfo.SetSubpixelOrder(fromOptions.GetSubpixelOrder())
		fo.Merge(sfo)
	}
}

func (fo *stdFontOptions) Hash() uint64 {
	return uint64(C.cairo_font_options_hash(fo.hnd))
}

func (fo *stdFontOptions) Equals(other FontOptions) bool {
	if sfo, ok := other.(*stdFontOptions); ok {
		return int(C.cairo_font_options_equal(fo.hnd, sfo.hnd)) > 0
	} else {
		return fo.GetAntialias() == other.GetAntialias() &&
			fo.GetHintMetrics() == other.GetHintMetrics() &&
			fo.GetHintStyle() == other.GetHintStyle() &&
			fo.GetSubpixelOrder() == other.GetSubpixelOrder()
	}
}

func (fo *stdFontOptions) SetAntialias(aaMode Antialias) {
	C.cairo_font_options_set_antialias(fo.hnd, C.cairo_antialias_t(aaMode))
}

func (fo *stdFontOptions) GetAntialias() Antialias {
	return Antialias(C.cairo_font_options_get_antialias(fo.hnd))
}

func (fo *stdFontOptions) SetSubpixelOrder(spoMode SubpixelOrder) {
	C.cairo_font_options_set_subpixel_order(fo.hnd, C.cairo_subpixel_order_t(spoMode))
}

func (fo *stdFontOptions) GetSubpixelOrder() SubpixelOrder {
	return SubpixelOrder(C.cairo_font_options_get_subpixel_order(fo.hnd))
}

func (fo *stdFontOptions) SetHintStyle(hsMode HintStyle) {
	C.cairo_font_options_set_hint_style(fo.hnd, C.cairo_hint_style_t(hsMode))
}

func (fo *stdFontOptions) GetHintStyle() HintStyle {
	return HintStyle(C.cairo_font_options_get_hint_style(fo.hnd))
}

func (fo *stdFontOptions) SetHintMetrics(hmMode HintMetrics) {
	C.cairo_font_options_set_hint_metrics(fo.hnd, C.cairo_hint_metrics_t(hmMode))
}

func (fo *stdFontOptions) GetHintMetrics() HintMetrics {
	return HintMetrics(C.cairo_font_options_get_hint_metrics(fo.hnd))
}
