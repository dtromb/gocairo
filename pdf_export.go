package cairo

import (
	"reflect"
	"unsafe"
)

/*
    #include <cairo/cairo.h>
	#include <inttypes.h>

	extern cairo_user_data_key_t *cgo_get_cairo_userdata_key(int32_t keyid);
	extern uint32_t cgo_get_refkey(void *cref);
	extern void* cgo_get_keyref(uint32_t key);
*/
import "C"

//export WriteStdPdfSurface
func WriteStdPdfSurface(sptr unsafe.Pointer, dptr unsafe.Pointer, length C.uint) C.cairo_status_t {
	ps := (*stdPdfSurface)(sptr)
	if ps.pdfOut != nil {
		hdr := reflect.SliceHeader{
			Len:  int(length),
			Cap:  int(length),
			Data: uintptr(dptr),
		}
		slice := *(*[]byte)(unsafe.Pointer(&hdr))
		n, err := ps.pdfOut.Write(slice)
		if err != nil || n != int(length) {
			return C.CAIRO_STATUS_WRITE_ERROR
		}
		return C.CAIRO_STATUS_SUCCESS
	}
	return C.CAIRO_STATUS_WRITE_ERROR
}
