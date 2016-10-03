package cairo

/*
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
	#include <inttypes.h>
	#include <stdlib.h>
*/
import "C"

type Extend uint32
const (
	ExtendNone			Extend 	= C.CAIRO_EXTEND_NONE
	ExtendRepeat				= C.CAIRO_EXTEND_REPEAT
	ExtendReflect				= C.CAIRO_EXTEND_REFLECT
	ExtendPad					= C.CAIRO_EXTEND_PAD
)