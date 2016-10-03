package cairo

import (
	"runtime"
)

/*
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
	#include <inttypes.h>
	#include <stdlib.h>
	
	
	extern cairo_user_data_key_t *cgo_get_cairo_userdata_key(int32_t keyid);
	extern uint32_t cgo_get_refkey(void *cref);
	extern void* cgo_get_keyref(uint32_t key);
	
	void cgo_cairo_device_userdata_destroy_c(void* x) {
		FreeDeviceNotify((cairo_device_t*)x);
	}
	cairo_destroy_func_t cgo_cairo_device_userdata_destroy = cgo_cairo_device_userdata_destroy_c;
	
	
*/
import "C"

type Device interface {
	Finish()
	Flush()
	GetUserData(key string) (interface{}, bool)
	ObserverElapsed() float64 
	ObserverFillElapsed() float64 
	ObserverPaintElapsed() float64 
	ObserverMaskElapsed() float64 
	ObserverGlyphsElapsed() float64
	SetUserData(key string, data interface{})
	Status() Status
}

type stdDevice struct {
	hnd *C.cairo_device_t
	userdata_r Reference
}

type DeviceType uint32
const (
	DeviceDRM		DeviceType 	=		C.CAIRO_DEVICE_TYPE_DRM
	DeviceGL				   	=		C.CAIRO_DEVICE_TYPE_GL
	DeviceScript				=		C.CAIRO_DEVICE_TYPE_SCRIPT
	DeviceXCB					=		C.CAIRO_DEVICE_TYPE_XCB
	DeviceXLib					=		C.CAIRO_DEVICE_TYPE_XLIB
	DeviceXML					=		C.CAIRO_DEVICE_TYPE_XML
	DeviceCOGL					=		C.CAIRO_DEVICE_TYPE_COGL
	DeviceWin32					=		C.CAIRO_DEVICE_TYPE_WIN32
	DeviceInvalidDevice			=		C.CAIRO_DEVICE_TYPE_INVALID
)

func destroyDevice(d Device) {
	if fn, ok := d.(Finalizable); ok {
		fn.Finalize(d)
	}
	if sd, ok := d.(*stdDevice); ok {
		C.cairo_device_destroy(sd.hnd)
		sd.hnd = nil
	}
}

func blessDevice(hnd *C.cairo_device_t, addRef bool) Device {
	s := &stdDevice{
		hnd: hnd,
	}
	if addRef {
		s.hnd = C.cairo_device_reference(s.hnd)
	}
	runtime.SetFinalizer(s, destroyDevice)
	return s
}

func (d *stdDevice) Status() Status {
	return Status(C.cairo_device_status(d.hnd))
}

func (d *stdDevice) Finish() {
	C.cairo_device_finish(d.hnd)
}

func (d *stdDevice) Flush() {
	C.cairo_device_flush(d.hnd)
}

func (d *stdDevice) SetUserData(key string, data interface{}) {
    if d.userdata_r == nil {
		// Attempt to load the Go ref table from the C-held ref key.
		refkey := C.cgo_get_refkey(C.cairo_device_get_user_data(d.hnd, C.cgo_get_cairo_userdata_key(GO_DATAKEY_KEY)))
	 	if refkey != 0 {
			var ok bool
			d.userdata_r, ok = LookupGlobalReference(uint32(refkey))
			if !ok {
				panic("missing global reference to device userdata")
			}
		} else {
			userdataMap := make(map[string]interface{})
			d.userdata_r = MakeGlobalReference(userdataMap)
			C.cairo_device_set_user_data(d.hnd, C.cgo_get_cairo_userdata_key(GO_DATAKEY_KEY), 
			                             C.cgo_get_keyref(C.uint32_t(d.userdata_r.Key())), C.cgo_cairo_device_userdata_destroy)
			IncrementGlobalReferenceCount(d.userdata_r)
		}
	}
	userdata := d.userdata_r.Ref().(map[string]interface{})
	userdata[key] = data  
}

func (d *stdDevice) GetUserData(key string) (interface{},bool) {
    if d.userdata_r == nil {
		// Attempt to load the Go ref table from the C-held ref key.
		refkey := C.cgo_get_refkey(C.cairo_device_get_user_data(d.hnd, C.cgo_get_cairo_userdata_key(GO_DATAKEY_KEY)))
	 	if refkey != 0 {
			var ok bool
			d.userdata_r, ok = LookupGlobalReference(uint32(refkey))
			if !ok {
				panic("missing global reference to device userdata")
			}
		} else {
			return nil, false
		}
	}
	userdata := d.userdata_r.Ref().(map[string]interface{})
	val, has := userdata[key]
	return val, has
}

func (d *stdDevice) ObserverElapsed() float64 {
	return float64(C.cairo_device_observer_elapsed(d.hnd))
}

func (d *stdDevice) ObserverFillElapsed() float64 {
	return float64(C.cairo_device_observer_fill_elapsed(d.hnd))
}

func (d *stdDevice) ObserverPaintElapsed() float64 {
	return float64(C.cairo_device_observer_paint_elapsed(d.hnd))
}

func (d *stdDevice) ObserverMaskElapsed() float64 {
	return float64(C.cairo_device_observer_mask_elapsed(d.hnd))
}

func (d *stdDevice) ObserverGlyphsElapsed() float64 {
	return float64(C.cairo_device_observer_glyphs_elapsed(d.hnd))
}