package cairo

/*
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
	#include <inttypes.h>
	#include <stdlib.h>

	unsigned long cgo_cairo_glyph_index(cairo_glyph_t *glyph) {
		return glyph->index;
	}

	double cgo_cairo_glyph_x(cairo_glyph_t *glyph) {
		return glyph->x;
	}

	double cgo_cairo_glyph_y(cairo_glyph_t *glyph) {
		return glyph->y;
	}

*/
import "C"

type Glyph C.cairo_glyph_t

/*
typedef struct {
    unsigned long        index;
    double               x;
    double               y;
} cairo_glyph_t;
*/

// XXX - These need write accessors.

func (g Glyph) Index() uint {
	return uint(C.cgo_cairo_glyph_index((*C.cairo_glyph_t)(&g)))
}

func (g Glyph) X() float64 {
	return float64(C.cgo_cairo_glyph_x((*C.cairo_glyph_t)(&g)))
}

func (g Glyph) Y() float64 {
	return float64(C.cgo_cairo_glyph_y((*C.cairo_glyph_t)(&g)))
}
