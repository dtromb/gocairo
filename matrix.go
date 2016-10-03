package cairo

import (
	"unsafe"
	"reflect"
)


/*
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
	#include <inttypes.h>
	#include <stdlib.h>
*/
import "C"

type Matrix []float64

func NewMatrix() Matrix {
	return Matrix{0,0,0,0,0,0}
}

func (m Matrix) dataref() *C.cairo_matrix_t {
	slice := []float64(m)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	return (*C.cairo_matrix_t)(unsafe.Pointer(hdr.Data))
}

func (m Matrix) Init(xx, yx, xy, yy, x0, y0 float64) {
	copy(m[:],[]float64{xx,yx,xy,yy,x0,y0})
}

func (m Matrix) InitIdentity() {
	copy(m[:],[]float64{1,0,0,1,0,0})
}

func (m Matrix) InitTranslate(x, y float64) {
	copy(m[:],[]float64{0,0,0,0,x,y})
}

func (m Matrix) InitScale(x, y float64) {
	copy(m[:],[]float64{x,0,0,y,0,0})
}

func (m Matrix) InitRotate(rad float64) {
	ref := m.dataref()
	C.cairo_matrix_init_rotate(ref, C.double(rad))
}

func (m Matrix) Translate(x, y float64) {
	ref := m.dataref()
	C.cairo_matrix_translate(ref, C.double(x), C.double(y))
}

func (m Matrix) Scale(x, y float64) {
	ref := m.dataref()
	C.cairo_matrix_scale(ref, C.double(x), C.double(y))
}

func (m Matrix) Rotate(rad float64) {
	ref := m.dataref()
	C.cairo_matrix_rotate(ref, C.double(rad))
}

func (m Matrix) Invert() Status {
	return Status(C.cairo_matrix_invert(m.dataref()))
}

func MatrixMultiply(r, a, b Matrix) {
	C.cairo_matrix_multiply(r.dataref(), a.dataref(), b.dataref())
}

func (m Matrix) TransformDistance(x, y float64) (float64,float64) {
	cx := C.double(x)
	cy := C.double(y)
	C.cairo_matrix_transform_distance(m.dataref(),&cx,&cy)
	return float64(cx), float64(cy)
}

func (m Matrix) TransformPoint(x, y float64) (float64,float64) {
	cx := C.double(x)
	cy := C.double(y)
	C.cairo_matrix_transform_point(m.dataref(),&cx,&cy)
	return float64(cx), float64(cy)
}
