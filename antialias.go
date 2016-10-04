package cairo

/*
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
	#include <inttypes.h>
	#include <stdlib.h>

*/
import "C"

type Antialias uint32

const (
	AntialiasDefault  Antialias = C.CAIRO_ANTIALIAS_DEFAULT
	AntialiasNone               = C.CAIRO_ANTIALIAS_NONE
	AntialiasGray               = C.CAIRO_ANTIALIAS_GRAY
	AntialiasSubpixel           = C.CAIRO_ANTIALIAS_SUBPIXEL
	AntialiasFast               = C.CAIRO_ANTIALIAS_FAST
	AntialiasGood               = C.CAIRO_ANTIALIAS_GOOD
	AntialiasBest               = C.CAIRO_ANTIALIAS_BEST
)
