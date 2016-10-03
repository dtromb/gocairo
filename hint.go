package cairo

/*
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
	#include <inttypes.h>
	#include <stdlib.h>
	
*/
import "C"

type HintStyle uint32
const (
	HintStyleDefault			HintStyle	= C.CAIRO_HINT_STYLE_DEFAULT
	HintStyleNone							= C.CAIRO_HINT_STYLE_NONE
	HintStyleSlight							= C.CAIRO_HINT_STYLE_SLIGHT
	HintStyleMedium							= C.CAIRO_HINT_STYLE_MEDIUM
	HintStyleFull							= C.CAIRO_HINT_STYLE_FULL
)	


type HintMetrics uint32
const (
	HintMetricsDefault			HintMetrics	= C.CAIRO_HINT_METRICS_DEFAULT
	HintMetricsOff							= C.CAIRO_HINT_METRICS_OFF
	HintMetricsOn							= C.CAIRO_HINT_METRICS_ON
)