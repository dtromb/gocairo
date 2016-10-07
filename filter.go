package cairo

/*
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
	#include <inttypes.h>
	#include <stdlib.h>
*/
import "C"

type Filter uint32

const (
	FilterFast     Filter = C.CAIRO_FILTER_FAST
	FilterGood            = C.CAIRO_FILTER_GOOD
	FilterBest            = C.CAIRO_FILTER_BEST
	FilterNearest         = C.CAIRO_FILTER_NEAREST
	FilterBilinear        = C.CAIRO_FILTER_BILINEAR
	FilterGaussian        = C.CAIRO_FILTER_GAUSSIAN
)
