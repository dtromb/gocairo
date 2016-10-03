package cairo

import (
	"runtime"
)

/*
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
    #include <cairo/cairo-pdf.h>
	#include <inttypes.h>
	#include <stdlib.h>
*/ 
import "C"

// XXX - Write data accessors in C and expose them on this object.

type RectangleList interface {
}

type stdRectangleList struct {
	hnd *C.cairo_rectangle_list_t
}

func destroyRectangleList(rl RectangleList) {
	if fn, ok := rl.(Finalizable); ok {
		fn.Finalize(rl)
	}
	if srl, ok := rl.(*stdRectangleList); ok {
		C.cairo_rectangle_list_destroy(srl.hnd)
		srl.hnd = nil
	}
}

func blessRectangleList(hnd *C.cairo_rectangle_list_t) RectangleList {
	rl := &stdRectangleList{
		hnd: hnd,
	}
	runtime.SetFinalizer(rl, destroyRectangleList)
	return rl
}
