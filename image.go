package cairo

import (
	"reflect"
	"unsafe"
)

/*
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
	#include <inttypes.h>
	#include <stdlib.h>
*/
import "C"


type Format int32
const (
	InvalidFormat	Format 	= C.CAIRO_FORMAT_INVALID
	ARGB32					= C.CAIRO_FORMAT_ARGB32
	RGB24					= C.CAIRO_FORMAT_RGB24
	A8						= C.CAIRO_FORMAT_A8
	A1						= C.CAIRO_FORMAT_A1
	RGB_16_565				= C.CAIRO_FORMAT_RGB16_565
	RGB30					= C.CAIRO_FORMAT_RGB30
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
	ss.stdSurface = blessSurface(hnd,addRef).(*stdSurface)
	return &ss;
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

type stdImageSurface struct {
	*stdSurface
	data []byte
}


func (ss *stdImageSurface) GetData() []byte {
	if ss.data == nil {
		dptr := C.cairo_image_surface_get_data(ss.hnd)
		dlen := ss.GetHeight() * ss.GetStride()
		hdr := &reflect.SliceHeader{
			Len: dlen,
			Cap: dlen,
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

