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

// XXX - Write path data accessors in C and expose them on this object.

// XXX - Memory management for these object is badly define in the API.   Sort this out.

type PathDataType uint32

const (
	PathDataTypeMoveTo    PathDataType = C.CAIRO_PATH_MOVE_TO
	PathDataTypeLineTo                 = C.CAIRO_PATH_LINE_TO
	PathDataTyoeCurveTo                = C.CAIRO_PATH_CURVE_TO
	PathDataTypeClosePath              = C.CAIRO_PATH_CLOSE_PATH
)

type Path interface {
}

type stdPath struct {
	hnd *C.cairo_path_t
}

func destroyPath(p Path) {
	if fn, ok := p.(Finalizable); ok {
		fn.Finalize(p)
	}
	if sp, ok := p.(*stdPath); ok {
		C.cairo_path_destroy(sp.hnd)
		sp.hnd = nil
	}
}

func blessPath(hnd *C.cairo_path_t) Path {
	p := &stdPath{
		hnd: hnd,
	}
	runtime.SetFinalizer(p, destroyPath)
	return p
}
