package cairo

import (
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


//export FreeSurfaceNotify
func FreeSurfaceNotify(csurf *C.cairo_surface_t) {
	refkey := C.cgo_get_refkey(C.cairo_surface_get_user_data(csurf, C.cgo_get_cairo_userdata_key(GO_DATAKEY_KEY)))
	if refkey != 0 {
		r, ok := LookupGlobalReference(uint32(refkey))
		if ok {
			DecrementGlobalReferenceCount(r)
		}
	}
}

//export FreeMimeDataNotify
func FreeMimeDataNotify(key uintptr) {
	refkey := C.cgo_get_refkey(unsafe.Pointer(key))
	if refkey != 0 {
		r, ok := LookupGlobalReference(uint32(refkey))
		if ok {
			DecrementGlobalReferenceCount(r)
		}
	}
}