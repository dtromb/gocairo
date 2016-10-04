package cairo

import (
	"io"
	"reflect"
	"unsafe"
)

/*
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
	#include <inttypes.h>
	#include <stdlib.h>

	cairo_status_t cgo_cairo_surface_png_read_c(void *cl, unsigned char *data, unsigned int len) {
		return SurfacePngRead(cl,data,len);
	}
	cairo_read_func_t cgo_cairo_surface_png_read = cgo_cairo_surface_png_read_c;

	cairo_status_t cgo_cairo_surface_png_write_c(void *cl, const unsigned char *data, unsigned int len) {
		return SurfacePngWrite(cl,data,len);
	}
	cairo_write_func_t cgo_cairo_surface_png_write = cgo_cairo_surface_png_write_c;
*/
import "C"

type Format int32

const (
	InvalidFormat Format = C.CAIRO_FORMAT_INVALID
	FormatArgb32         = C.CAIRO_FORMAT_ARGB32
	FormatRgb24          = C.CAIRO_FORMAT_RGB24
	ForamatA8            = C.CAIRO_FORMAT_A8
	FormatA1             = C.CAIRO_FORMAT_A1
	FormatRgb16          = C.CAIRO_FORMAT_RGB16_565
	FormatRgb30          = C.CAIRO_FORMAT_RGB30
)

func (f Format) StrideForWidth(width int) int {
	return int(C.cairo_format_stride_for_width(C.cairo_format_t(f), C.int(width)))
}

type ImageSurface interface {
	Surface
	GetData() []byte
	GetFormat() Format
	GetWidth() int
	GetHeight() int
	GetStride() int
}

func blessImageSurface(hnd *C.cairo_surface_t, addRef bool) ImageSurface {
	ss := stdImageSurface{}
	ss.stdSurface = blessSurface(hnd, addRef).(*stdSurface)
	return &ss
}

func ImageSurfaceCreate(format Format, width, height int) ImageSurface {
	csurf := C.cairo_image_surface_create(C.cairo_format_t(format), C.int(width), C.int(height))
	return blessImageSurface(csurf, false)
}

func ImageSurfaceCreateForData(data []byte, format Format, width, height, stride int) ImageSurface {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	csurf := C.cairo_image_surface_create_for_data((*C.uchar)(unsafe.Pointer(hdr.Data)), C.cairo_format_t(format),
		C.int(width), C.int(height), C.int(stride))
	ss := blessImageSurface(csurf, false)
	ss.(*stdImageSurface).data = data
	return ss
}

func ImageSurfaceCreateFromPng(filename string) ImageSurface {
	csurf := C.cairo_image_surface_create_from_png(C.CString(filename))
	return blessImageSurface(csurf, false)
}

func ImageSurfaceCreateFromPngStream(in io.Reader) ImageSurface {
	ref := &InterfaceRef{x: in}
	csurf := C.cairo_image_surface_create_from_png_stream(C.cgo_cairo_surface_png_read,
		unsafe.Pointer(ref))
	return blessImageSurface(csurf, false)
}

type stdImageSurface struct {
	*stdSurface
	data []byte
}

func (ss *stdImageSurface) GetData() []byte {
	if ss.data == nil {
		dptr := C.cairo_image_surface_get_data(ss.hnd)
		dlen := ss.GetHeight() * ss.GetStride()
		hdr := &reflect.SliceHeader{
			Len:  dlen,
			Cap:  dlen,
			Data: uintptr(unsafe.Pointer(dptr)),
		}
		ss.data = *(*[]byte)(unsafe.Pointer(hdr))
	}
	return ss.data
}

func (ss *stdImageSurface) GetFormat() Format {
	return Format(C.cairo_image_surface_get_format(ss.hnd))
}

func (ss *stdImageSurface) GetWidth() int {
	return int(C.cairo_image_surface_get_width(ss.hnd))
}

func (ss *stdImageSurface) GetHeight() int {
	return int(C.cairo_image_surface_get_height(ss.hnd))
}

func (ss *stdImageSurface) GetStride() int {
	return int(C.cairo_image_surface_get_stride(ss.hnd))
}

func (ss *stdSurface) WriteToPng(filename string) {
	C.cairo_surface_write_to_png(ss.hnd, C.CString(filename))
}

func (ss *stdSurface) WriteToPngStream(out io.Writer) {
	ref := &InterfaceRef{x: out}
	C.cairo_surface_write_to_png_stream(ss.hnd, C.cgo_cairo_surface_png_write, unsafe.Pointer(ref))
}
