package cairo

/*
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>

*/
import "C"

type LineCap uint32

const (
	LineCapButt   LineCap = C.CAIRO_LINE_CAP_BUTT
	LineCapRound          = C.CAIRO_LINE_CAP_ROUND
	LineCapSquare         = C.CAIRO_LINE_CAP_SQUARE
)

type LineJoin uint32

const (
	LineJoinMiter LineJoin = C.CAIRO_LINE_JOIN_MITER
	LineJoinRound          = C.CAIRO_LINE_JOIN_ROUND
	LineJoinBevel          = C.CAIRO_LINE_JOIN_BEVEL
)
