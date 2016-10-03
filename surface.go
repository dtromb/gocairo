package cairo

import (
	"reflect"
	"unsafe"
	"runtime"
	"errors"
)

/*
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
	#include <inttypes.h>
	#include <stdlib.h>
	
	
	extern cairo_user_data_key_t *cgo_get_cairo_userdata_key(int32_t keyid);
	extern uint32_t cgo_get_refkey(void *cref);
	extern void* cgo_get_keyref(uint32_t key);
	
	struct cgo_cairo_surface_double_pair {
		double x;
		double y;
	};
	
	int cgo_cairo_device_isnull(cairo_device_t *x) {
		return x == NULL;
	}
	
	void cgo_cairo_get_device_offset(cairo_surface_t *s, struct cgo_cairo_surface_double_pair *pair) {
		cairo_surface_get_device_offset(s, &pair->x, &pair->y);
	}
	
	void cgo_cairo_get_device_scale(cairo_surface_t *s, struct cgo_cairo_surface_double_pair *pair) {
		cairo_surface_get_device_scale(s, &pair->x, &pair->y);
	}
	
	void cgo_cairo_get_fallback_resolution(cairo_surface_t *s, struct cgo_cairo_surface_double_pair *pair) {
		cairo_surface_get_fallback_resolution(s, &pair->x, &pair->y);
	}
		
	void cgo_cairo_surface_userdata_destroy_c(void *sptr) {
		FreeSurfaceNotify((cairo_surface_t*)sptr);
	}
	cairo_destroy_func_t cgo_cairo_surface_destroy = cgo_cairo_surface_userdata_destroy_c;

	void cgo_cairo_surface_mime_data_destroy_c(void *kptr) {
		FreeMimeDataNotify(kptr);
	}
	cairo_destroy_func_t cgo_cairo_surface_mime_data_destroy = cgo_cairo_surface_mime_data_destroy_c;
*/
import "C"

type Content uint32 
const (
	ContentColor		Content		= C.CAIRO_CONTENT_COLOR
	ContentAlpha					= C.CAIRO_CONTENT_ALPHA
	ContentColorAlpha				= C.CAIRO_CONTENT_COLOR_ALPHA
)

type Surface interface {
	CreateSimilar(content Content, width int, height int) (Surface	,error)
	CreateSimilarImage(format Format, width int, height int) (Surface,error)
	CreateForRectangle(x float64, y float64, width float64, height float64) (Surface,error)
	Status() Status
	Finish()
	Flush()
	GetDevice() Device
	GetFontOptions() FontOptions
 	GetContent() Content
	MarkDirty()
	MarkDirtyRectangle(x, y, width, height int)
	SetDeviceOffset(x, y float64)
	GetDeviceOffset() (float64, float64)
	SetDeviceScale(x, y float64) 
	GetDeviceScale() (float64, float64)
	SetFallbackResolution(x, y float64) 
	GetFallbackResolution() (float64, float64)
	GetType() SurfaceType
	CopyPage()
	ShowPage()
	SetMimeData(mimeType MimeType, data []byte)
	SupportsMimeType(mimeType MimeType) bool
}

type StandardSurface interface {
	GetStandardSurface() *stdSurface
}

type SurfaceType uint32
const (
	SurfaceTypeImage		SurfaceType		= C.CAIRO_SURFACE_TYPE_IMAGE
	SurfaceTypePdf							= C.CAIRO_SURFACE_TYPE_PDF
	SurfaceTypePs							= C.CAIRO_SURFACE_TYPE_PS
	SurfaceTypeXlib							= C.CAIRO_SURFACE_TYPE_XLIB
	SurfaceTypeXcb							= C.CAIRO_SURFACE_TYPE_XCB
	SurfaceTypeGlitz						= C.CAIRO_SURFACE_TYPE_GLITZ
	SurfaceTypeQuartz						= C.CAIRO_SURFACE_TYPE_QUARTZ
	SurfaceTypeWin32						= C.CAIRO_SURFACE_TYPE_WIN32
	SurfaceTypeBeos							= C.CAIRO_SURFACE_TYPE_BEOS
	SurfaceTypeDirectFb						= C.CAIRO_SURFACE_TYPE_DIRECTFB
	SurfaceTypeSvg							= C.CAIRO_SURFACE_TYPE_SVG
	SurfaceTypeOS2							= C.CAIRO_SURFACE_TYPE_OS2
	SurfaceTypePrinting						= C.CAIRO_SURFACE_TYPE_WIN32_PRINTING
	SurfaceTypeQuartzImage					= C.CAIRO_SURFACE_TYPE_QUARTZ_IMAGE
	SurfaceTypeScript						= C.CAIRO_SURFACE_TYPE_SCRIPT
	SurfaceTypeQt							= C.CAIRO_SURFACE_TYPE_QT
	SurfaceTypeRecording					= C.CAIRO_SURFACE_TYPE_RECORDING
	SurfaceTypeVg							= C.CAIRO_SURFACE_TYPE_VG
	SurfaceTypeGl 							= C.CAIRO_SURFACE_TYPE_GL
	SurfaceTypeDrm 							= C.CAIRO_SURFACE_TYPE_DRM
	SurfaceTypeTee 							= C.CAIRO_SURFACE_TYPE_TEE
	SurfaceTypeXml 							= C.CAIRO_SURFACE_TYPE_XML
	SurfaceTypeSkia 						= C.CAIRO_SURFACE_TYPE_SKIA
	SurfaceTypeSubsurface 					= C.CAIRO_SURFACE_TYPE_SUBSURFACE
	SurfaceTypeCogl 						= C.CAIRO_SURFACE_TYPE_COGL
)

type stdSurface struct {
	hnd *C.cairo_surface_t
	device Device
	userdata_r Reference
}

func referenceSurface(cref *C.cairo_surface_t) Surface {
	stype := SurfaceType(C.cairo_surface_get_type(cref))
	baseSurface := blessSurface(cref, true)
	ss := baseSurface.(*stdSurface)
	switch(stype) {
		case SurfaceTypeImage: {
			return &stdImageSurface{stdSurface:ss}
		}
		case SurfaceTypePdf: {
			return &stdPdfSurface{stdSurface:ss}
		}
	}
	return baseSurface
}


func destroySurface(s Surface) {
	if fn, ok := s.(Finalizable); ok {
		fn.Finalize(s)
	}
	if ss, ok := s.(*stdSurface); ok {
		C.cairo_surface_destroy(ss.hnd)
		ss.hnd = nil
	}
}

func blessSurface(hnd *C.cairo_surface_t, addRef bool) Surface {
	s := &stdSurface{
		hnd: hnd,
	}
	if addRef {
		s.hnd = C.cairo_surface_reference(s.hnd)
	}
	runtime.SetFinalizer(s, destroySurface)
	return s
}

func (s *stdSurface) GetStandardSurface() *stdSurface {
	return s
}

func (s *stdSurface) CreateSimilar(content Content, width int, height int) (Surface, error) {
	// Request the new surface.
	nsurf := C.cairo_surface_create_similar(s.hnd, C.cairo_content_t(content), 
	                                        C.int(width), C.int(height))
	// Check for error condition.
	rc := uint(C.cairo_surface_get_reference_count(nsurf))
	if rc == 0 {
		status := Status(C.cairo_surface_status(s.hnd))
		C.cairo_surface_destroy(nsurf)
		return nil, errors.New(status.String())
	}
	// Return the new Surface.
	return blessSurface(nsurf,false), nil
}

func (s *stdSurface) CreateSimilarImage(format Format, width int, height int) (Surface, error) {
	// Request the new surface.
	nsurf := C.cairo_surface_create_similar_image(s.hnd, C.cairo_format_t(format), 
	                                              C.int(width), C.int(height))
	// Check for error condition.
	rc := uint(C.cairo_surface_get_reference_count(nsurf))
	if rc == 0 {
		status := Status(C.cairo_surface_status(s.hnd))
		C.cairo_surface_destroy(nsurf)
		return nil, errors.New(status.String())
	}
	// Return the new Surface.
	return blessSurface(nsurf,false), nil
}

func (s *stdSurface) CreateForRectangle(x float64, y float64, width float64, height float64) (Surface,error) {
	// Request the new surface.
	nsurf := C.cairo_surface_create_for_rectangle(s.hnd, C.double(x), C.double(y), C.double(width), C.double(height))
	// Check for error condition.
	rc := uint(C.cairo_surface_get_reference_count(nsurf))
	if rc == 0 {
		status := Status(C.cairo_surface_status(s.hnd))
		C.cairo_surface_destroy(nsurf)
		return nil, errors.New(status.String())
	}
	// Return the new Surface.
	return blessSurface(nsurf,false), nil
}	

func (s *stdSurface) Status() Status {
	return Status(C.cairo_surface_status(s.hnd))
}

func (s *stdSurface) Finish() {
	C.cairo_surface_finish(s.hnd)
}

func (s *stdSurface) Flush() {
	C.cairo_surface_flush(s.hnd)
}

func (s *stdSurface) GetDevice() Device {
	if s.device == nil {
		ndev := C.cairo_surface_get_device(s.hnd)
		if C.cgo_cairo_device_isnull(ndev) > 0 {
			return nil
		}
		// XXX - The docs do not explicitly say that cairo_surface_get_device()
		// increments the device's reference count, check that this is true.
		s.device = blessDevice(ndev,false)
	}
	return s.device
}

func (s *stdSurface) GetFontOptions() FontOptions {
	opts := NewFontOptions()
	stdOpts := opts.(*stdFontOptions)
	C.cairo_surface_get_font_options(s.hnd, stdOpts.hnd)
	return opts	
}

func (s *stdSurface) GetContent() Content {
	return Content(C.cairo_surface_get_content(s.hnd))
}

func (s *stdSurface) MarkDirty() {
	C.cairo_surface_mark_dirty(s.hnd)
}

func (s *stdSurface) MarkDirtyRectangle(x, y, width, height int) {
	C.cairo_surface_mark_dirty_rectangle(s.hnd, C.int(x), C.int(y), C.int(width), C.int(height))
}

func (s *stdSurface) SetDeviceOffset(x, y float64) {
	C.cairo_surface_set_device_offset(s.hnd, C.double(x), C.double(y))
}

func (s *stdSurface) GetDeviceOffset() (float64, float64) {
	var cpair C.struct_cgo_cairo_surface_double_pair
	C.cgo_cairo_get_device_offset(s.hnd, &cpair)
	return float64(cpair.x), float64(cpair.y)
}

func (s *stdSurface) SetDeviceScale(x, y float64) {
	C.cairo_surface_set_device_scale(s.hnd, C.double(x), C.double(y))
}

func (s *stdSurface) GetDeviceScale() (float64, float64) {
	var cpair C.struct_cgo_cairo_surface_double_pair
	C.cgo_cairo_get_device_scale(s.hnd, &cpair)
	return float64(cpair.x), float64(cpair.y)
}

func (s *stdSurface) SetFallbackResolution(x, y float64) {
	C.cairo_surface_set_fallback_resolution(s.hnd, C.double(x), C.double(y))
}

func (s *stdSurface) GetFallbackResolution() (float64, float64) {
	var cpair C.struct_cgo_cairo_surface_double_pair
	C.cgo_cairo_get_fallback_resolution(s.hnd, &cpair)
	return float64(cpair.x), float64(cpair.y)
}

func (s *stdSurface) GetType() SurfaceType {
	return SurfaceType(C.cairo_surface_get_type(s.hnd))
}

func (s *stdSurface) SetUserData(key string, data interface{}) {
    if s.userdata_r == nil {
		// Attempt to load the Go ref table from the C-held ref key.
		refkey := C.cgo_get_refkey(C.cairo_surface_get_user_data(s.hnd, C.cgo_get_cairo_userdata_key(GO_DATAKEY_KEY)))
	 	if refkey != 0 {
			var ok bool
			s.userdata_r, ok = LookupGlobalReference(uint32(refkey))
			if !ok {
				panic("missing global reference to surface userdata")
			}
		} else {
			userdataMap := make(map[string]interface{})
			s.userdata_r = MakeGlobalReference(userdataMap)
			C.cairo_surface_set_user_data(s.hnd, C.cgo_get_cairo_userdata_key(GO_DATAKEY_KEY), 
			                              C.cgo_get_keyref(C.uint32_t(s.userdata_r.Key())), C.cgo_cairo_surface_destroy)
			IncrementGlobalReferenceCount(s.userdata_r)
		}
	}
	userdata := s.userdata_r.Ref().(map[string]interface{})
	userdata[key] = data  
}

func (s *stdSurface) GetUserData(key string) (interface{},bool) {
    if s.userdata_r == nil {
		// Attempt to load the Go ref table from the C-held ref key.
		refkey := C.cgo_get_refkey(C.cairo_device_get_user_data(s.hnd, C.cgo_get_cairo_userdata_key(GO_DATAKEY_KEY)))
	 	if refkey != 0 {
			var ok bool
			s.userdata_r, ok = LookupGlobalReference(uint32(refkey))
			if !ok {
				panic("missing global reference to device userdata")
			}
		} else {
			return nil, false
		}
	}
	userdata := s.userdata_r.Ref().(map[string]interface{})
	val, has := userdata[key]
	return val, has
}

func (s *stdSurface) CopyPage() {
	C.cairo_surface_copy_page(s.hnd)
}

func (s *stdSurface) ShowPage() {
	C.cairo_surface_show_page(s.hnd)
}

func (s *stdSurface) HasShowTextGlyphs() bool {
	return C.cairo_surface_has_show_text_glyphs(s.hnd) > 0
}

func (s *stdSurface) SetMimeData(mimeType MimeType, data []byte) {
	ref := MakeGlobalReference(data)
	datahdr := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	C.cairo_surface_set_mime_data(s.hnd, C.CString(string(mimeType)), 
						  		   (*C.uchar)(unsafe.Pointer(datahdr.Data)), C.ulong(datahdr.Len),
								   C.cgo_cairo_surface_mime_data_destroy, 
								   C.cgo_get_keyref(C.uint32_t(ref.Key())))
}

func (s *stdSurface) GetMimeData(mimeType MimeType) ([]byte,bool) {
	// XXX - How do alloc semantics for this work?   Completely
	// undocumented in the reference!
	panic("stdSurface.GetMimeData() unimplemented")
}

func (s *stdSurface) SupportsMimeType(mimeType MimeType) bool {
	return C.cairo_surface_supports_mime_type(s.hnd, C.CString(string(mimeType))) > 0;
}

func (s *stdSurface) MapToImage(rect RectangleInt) Surface {
	var r C.cairo_rectangle_int_t
	r.x = C.int(rect.X) 
	r.y = C.int(rect.Y) 
	r.width = C.int(rect.Width)
	r.height = C.int(rect.Height)
	nsurf := C.cairo_surface_map_to_image(s.hnd, &r)
	return blessSurface(nsurf, false)
}

func (s *stdSurface) UnmapImage(img Surface) {
	if ss, ok := img.(*stdSurface); ok {
		C.cairo_surface_unmap_image(s.hnd, ss.hnd)
	} else {
		panic("stdSurface.UnmapImage() argument is not the result of a call to stdSurface.MapToImage()")
	}
}
