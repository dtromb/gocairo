package cairo

/*
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>

*/
import "C"

type FillRule uint32

const (
	FillRuleWinding FillRule = C.CAIRO_FILL_RULE_WINDING
	FillRuleEvenOdd          = C.CAIRO_FILL_RULE_EVEN_ODD
)
