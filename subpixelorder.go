package cairo

/*
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
	#include <inttypes.h>
	#include <stdlib.h>
*/
import "C"

type SubpixelOrder uint32

const (
	SubpixelOrderDefault SubpixelOrder = C.CAIRO_SUBPIXEL_ORDER_DEFAULT
	SubpixelOrderRGB                   = C.CAIRO_SUBPIXEL_ORDER_RGB
	SubpixelOrderBGR                   = C.CAIRO_SUBPIXEL_ORDER_BGR
	SubpixelOrderVRGB                  = C.CAIRO_SUBPIXEL_ORDER_VRGB
	SubpixelOrderVBGR                  = C.CAIRO_SUBPIXEL_ORDER_VBGR
)
