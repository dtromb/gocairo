package cairo

import (
	"reflect"
	"unsafe"
	"runtime"
)

/*
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
	#include <inttypes.h>
	
	extern cairo_user_data_key_t *cgo_get_cairo_userdata_key(int32_t keyid);
	extern uint32_t cgo_get_refkey(void *cref);
	extern void* cgo_get_keyref(uint32_t key);
	
	void cgo_cairo_userdata_destroy_c(void *ptr) {
		FreeCairoNotify((cairo_t*)ptr);
	}
	cairo_destroy_func_t cgo_cairo_userdata_destroy = cgo_cairo_userdata_destroy_c;
*/
import "C"

type RectangleInt struct {
	X, Y int
	Width, Height int
}

type Finalizable interface {
	Finalize(x interface{})
}

type Status uint32
const (
	StatusSuccess					Status	= C.CAIRO_STATUS_SUCCESS
	StatusNoMemory							= C.CAIRO_STATUS_NO_MEMORY
	StatusInvalidRestore					= C.CAIRO_STATUS_INVALID_RESTORE
	StatusInvalidPopCount					= C.CAIRO_STATUS_INVALID_POP_GROUP
	StatusNoCurrentPoint					= C.CAIRO_STATUS_NO_CURRENT_POINT
	StatusInvalidMatrix						= C.CAIRO_STATUS_INVALID_MATRIX
	StatusInvalidStatus						= C.CAIRO_STATUS_INVALID_STATUS
	StatusNullPointer						= C.CAIRO_STATUS_NULL_POINTER
	StatusInvalidString						= C.CAIRO_STATUS_INVALID_STRING
	StatusInvalidPathData					= C.CAIRO_STATUS_INVALID_PATH_DATA
	StatusReadError							= C.CAIRO_STATUS_READ_ERROR
	StatusWriteError						= C.CAIRO_STATUS_WRITE_ERROR
	StatusSurfaceFinished					= C.CAIRO_STATUS_SURFACE_FINISHED
	StatusSurfaceTypeMismatch				= C.CAIRO_STATUS_SURFACE_TYPE_MISMATCH
	StatusPatternTypeMismatch				= C.CAIRO_STATUS_PATTERN_TYPE_MISMATCH
	StatusInvalidContent					= C.CAIRO_STATUS_INVALID_CONTENT
	StatusInvalidFormat						= C.CAIRO_STATUS_INVALID_FORMAT
	StatusInvalidVisual						= C.CAIRO_STATUS_INVALID_VISUAL
	StatusFileNotFound						= C.CAIRO_STATUS_FILE_NOT_FOUND
	StatusInvalidDash						= C.CAIRO_STATUS_INVALID_DASH
	StatusInvalidDscComment					= C.CAIRO_STATUS_INVALID_DSC_COMMENT
	StatusInvalidIndex						= C.CAIRO_STATUS_INVALID_INDEX
	StatusClipNotRepresentable				= C.CAIRO_STATUS_CLIP_NOT_REPRESENTABLE
	StatusTempFileError						= C.CAIRO_STATUS_TEMP_FILE_ERROR
	StatusInvalidStride						= C.CAIRO_STATUS_INVALID_STRIDE
	StatusFontTypeMismatch					= C.CAIRO_STATUS_FONT_TYPE_MISMATCH
	StatusUserFontImmutable					= C.CAIRO_STATUS_USER_FONT_IMMUTABLE
	StatusUserFontError						= C.CAIRO_STATUS_USER_FONT_ERROR
	StatusNegativeCount						= C.CAIRO_STATUS_NEGATIVE_COUNT
	StatusInvalidClusters					= C.CAIRO_STATUS_INVALID_CLUSTERS
	StatusInvalidSlant						= C.CAIRO_STATUS_INVALID_SLANT
	StatusInvalidWeight						= C.CAIRO_STATUS_INVALID_WEIGHT
	StatusInvalidSize						= C.CAIRO_STATUS_INVALID_SIZE
	StatusUserFontNotImplemented			= C.CAIRO_STATUS_USER_FONT_NOT_IMPLEMENTED
	StatusDeviceTypeMismatch				= C.CAIRO_STATUS_DEVICE_TYPE_MISMATCH
	StatusDeviceError						= C.CAIRO_STATUS_DEVICE_ERROR
	StatusInvalidMeshConstruction			= C.CAIRO_STATUS_INVALID_MESH_CONSTRUCTION
	StatusDeviceFinished					= C.CAIRO_STATUS_DEVICE_FINISHED
	StatusJbig2GlobalMissing				= C.CAIRO_STATUS_JBIG2_GLOBAL_MISSING
)

func (st Status) String() string {
	return C.GoString(C.cairo_status_to_string(C.cairo_status_t(st)))
}

func Debug_ResetStaticData() {
	C.cairo_debug_reset_static_data()
}

func Version() int {
	return int(C.cairo_version())
}

func VersionString() string {
	return C.GoString(C.cairo_version_string())
}

type Cairo interface {
	Status() Status
	Save()
	Restore()
	GetTarget() Surface
	PushGroup()
	PushGroupWithContent(c Content)
	PopGroup()
	PopGroupToSource()
	GetGroupTarget() Surface
	SetSourceRgb(r, g, b float64)
	SetSourceRgba(r, g, b, a float64)
	SetSource(source Pattern)
	SetSourceSurface(s Surface, x, y float64)
	GetSource() Pattern
	SetAntialias(aaMode Antialias)
	GetAntialias() Antialias
	SetDash(dash []float64, offset float64)
	GetDashCount() int
	GetDash() ([]float64,float64)
	SetFillRule(rule FillRule)
	GetFillRule() FillRule
	SetLineCap(linecap LineCap)
	GetLineCap() LineCap
	SetLineJoin(linejoin LineJoin)
	GetLineJoin() LineJoin
	SetLineWidth(width float64)
	GetLineWidth() float64
	SetMiterLimit(limit float64)
	GetMiterLimit() float64
	SetOperator(op Operator)
	GetOperator() Operator
	SetTolerance(tol float64)
	GetTolerance() float64
	Clip()
	ClipPreserve()
	ClipExtents() []float64
	InClip(x, y float64) bool
	ResetClip()
	CopyClipRectangleList() RectangleList
	Fill()
	FillPreserve()
	FillExtents() []float64
	InFill(x, y float64) bool
	Mask(mask Pattern)
	MaskSurface(s Surface, x, y float64)
	Paint()
	PaintWithAlpha(alpha float64)
	Stroke()
	StrokePreserve()
	StrokeExtents() []float64
	InStroke(x, y float64) bool
	CopyPage()
	ShowPage()
	SetUserData(key string, data interface{})
	GetUserData(key string) (interface{},bool)
}

type stdCairo struct {
	hnd *C.cairo_t
	userdata_r Reference
}

func destroyCairo(c Cairo) {
	if fn, ok := c.(Finalizable); ok {
		fn.Finalize(c)
	}
	if sc, ok := c.(*stdCairo); ok {
		C.cairo_destroy(sc.hnd)
		sc.hnd = nil
	}
}

func blessCairo(hnd *C.cairo_t, addRef bool) Cairo {
	c := &stdCairo{
		hnd: hnd,
	}
	if addRef {
		c.hnd = C.cairo_reference(c.hnd)
	}
	runtime.SetFinalizer(c, destroyCairo)
	return c
}

func (sc *stdCairo) Status() Status {
	return Status(C.cairo_status(sc.hnd))
}

func (sc *stdCairo) Save() {
	C.cairo_save(sc.hnd)
}

func (sc *stdCairo) Restore() {
	C.cairo_restore(sc.hnd)
}

func (sc *stdCairo) GetTarget() Surface {
	csurf := C.cairo_get_target(sc.hnd)
	return referenceSurface(csurf)
}

func (sc *stdCairo) PushGroup() {
	C.cairo_push_group(sc.hnd)
}

func (sc *stdCairo) PushGroupWithContent(c Content) {
	C.cairo_push_group_with_content(sc.hnd, C.cairo_content_t(c))
}

func (sc *stdCairo) PopGroup() {
	C.cairo_pop_group(sc.hnd)
}

func (sc *stdCairo) PopGroupToSource() {
	C.cairo_pop_group_to_source(sc.hnd)
}

func (sc *stdCairo) GetGroupTarget() Surface {
	csurf := C.cairo_get_group_target(sc.hnd)
	return referenceSurface(csurf)
}

func (sc *stdCairo) SetSourceRgb(r, g, b float64) {
	C.cairo_set_source_rgb(sc.hnd, C.double(r), C.double(g), C.double(b))
}

func (sc *stdCairo) SetSourceRgba(r, g, b, a float64) {
	C.cairo_set_source_rgba(sc.hnd, C.double(r), C.double(g), C.double(b), C.double(a))
}

func (sc *stdCairo) SetSource(source Pattern) {
	if sp, ok := source.(*stdPattern); ok {
		C.cairo_set_source(sc.hnd, sp.hnd)
	} else {
		panic("stdCairo.setSource(p) unimplemented for non-standard pattern argument")
	}
}

func (sc *stdCairo) SetSourceSurface(s Surface, x, y float64) {
	if ss, ok := s.(*stdSurface); ok {
		C.cairo_set_source_surface(sc.hnd, ss.hnd, C.double(x), C.double(y))
	} else {
		panic("stdCairo.setSourceSurface(s) not implemented for non-standard surface argument")
	}
}

func (sc *stdCairo) GetSource() Pattern {
	return referencePattern(C.cairo_get_source(sc.hnd))
}

func (sc *stdCairo) SetAntialias(aaMode Antialias) {
	C.cairo_set_antialias(sc.hnd, C.cairo_antialias_t(aaMode))
}

func (sc *stdCairo) GetAntialias() Antialias {
	return Antialias(C.cairo_get_antialias(sc.hnd))
}

func (sc *stdCairo) SetDash(dash []float64, offset float64) {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&dash))
	C.cairo_set_dash(sc.hnd, (*C.double)(unsafe.Pointer(hdr.Data)), C.int(hdr.Len), C.double(offset))
}

func (sc *stdCairo) GetDashCount() int {
	return int(C.cairo_get_dash_count(sc.hnd))
}

func (sc *stdCairo) GetDash() ([]float64,float64) {
	var offset C.double
	dashes := make([]float64, sc.GetDashCount())
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&dashes))
	C.cairo_get_dash(sc.hnd, (*C.double)(unsafe.Pointer(hdr.Data)), &offset)
	return dashes, float64(offset)
}

func (sc *stdCairo) SetFillRule(rule FillRule) {
	C.cairo_set_fill_rule(sc.hnd, C.cairo_fill_rule_t(rule))
}

func (sc *stdCairo) GetFillRule() FillRule {
	return FillRule(C.cairo_get_fill_rule(sc.hnd))
}

func (sc *stdCairo) SetLineCap(linecap LineCap) {
	C.cairo_set_line_cap(sc.hnd, C.cairo_line_cap_t(linecap))
}

func (sc *stdCairo) GetLineCap() LineCap {
	return LineCap(C.cairo_get_line_cap(sc.hnd))
}

func (sc *stdCairo) SetLineJoin(linejoin LineJoin) {
	C.cairo_set_line_join(sc.hnd, C.cairo_line_join_t(linejoin)) 
}

func (sc *stdCairo) GetLineJoin() LineJoin {
	return LineJoin(C.cairo_get_line_join(sc.hnd))
}

func (sc *stdCairo) SetLineWidth(width float64) {
	C.cairo_set_line_width(sc.hnd, C.double(width))
}

func (sc *stdCairo) GetLineWidth() float64 {
	return float64(C.cairo_get_line_width(sc.hnd))
}

func (sc *stdCairo) SetMiterLimit(limit float64) {
	C.cairo_set_miter_limit(sc.hnd, C.double(limit))
}

func (sc *stdCairo) GetMiterLimit() float64 {
	return float64(C.cairo_get_miter_limit(sc.hnd))
}

func (sc *stdCairo) SetOperator(op Operator) {
	C.cairo_set_operator(sc.hnd, C.cairo_operator_t(op))
}

func (sc *stdCairo) GetOperator() Operator {
	return Operator(C.cairo_get_operator(sc.hnd))
}

func (sc *stdCairo) SetTolerance(tol float64) {
	C.cairo_set_tolerance(sc.hnd, C.double(tol))
}

func (sc *stdCairo) GetTolerance() float64 {
	return float64(C.cairo_get_tolerance(sc.hnd))
}

func (sc *stdCairo) Clip() {
	C.cairo_clip(sc.hnd)
}

func (sc *stdCairo) ClipPreserve() {
	C.cairo_clip_preserve(sc.hnd)
}

func (sc *stdCairo) ClipExtents() []float64 {
	var x1, y1, x2, y2 C.double
	C.cairo_clip_extents(sc.hnd, &x1, &y1, &x2, &y2)
	return []float64{float64(x1),float64(y1),float64(x2),float64(y2)}
}

func (sc *stdCairo) InClip(x, y float64) bool {
	return C.cairo_in_clip(sc.hnd, C.double(x), C.double(y)) > 0
}

func (sc *stdCairo) ResetClip() {
	C.cairo_reset_clip(sc.hnd)
}

func (sc *stdCairo) CopyClipRectangleList() RectangleList {
	return blessRectangleList(C.cairo_copy_clip_rectangle_list(sc.hnd))
}

func (sc *stdCairo) Fill() {
	C.cairo_fill(sc.hnd)
}

func (sc *stdCairo) FillPreserve() {
	C.cairo_fill_preserve(sc.hnd)
}

func (sc *stdCairo) FillExtents() []float64 {
	var x1, y1, x2, y2 C.double
	C.cairo_fill_extents(sc.hnd, &x1, &y1, &x2, &y2)
	return []float64{float64(x1),float64(y1),float64(x2),float64(y2)}
}

func (sc *stdCairo) InFill(x, y float64) bool {
	return C.cairo_in_fill(sc.hnd, C.double(x), C.double(y)) > 0
}

func (sc *stdCairo) Mask(mask Pattern) {
	if sp, ok := mask.(*stdPattern); ok {
		C.cairo_mask(sc.hnd, sp.hnd)
	} else {
		panic("stdCairo.Mask(p) unimplemented for non-standard pattern arguments")
	}
}

func (sc *stdCairo) MaskSurface(s Surface, x, y float64) {
	if ss, ok := s.(*stdSurface); ok {
		C.cairo_mask_surface(sc.hnd, ss.hnd, C.double(x), C.double(y))
	} else {
		panic("stdCairo.MaskSurface(s) unimplemented for non-standard surface arguments")
	}
}

func (sc *stdCairo) Paint() {
	C.cairo_paint(sc.hnd)
}

func (sc *stdCairo) PaintWithAlpha(alpha float64) {
	C.cairo_paint_with_alpha(sc.hnd, C.double(alpha))
}

func (sc *stdCairo) Stroke() {
	C.cairo_stroke(sc.hnd)
}

func (sc *stdCairo) StrokePreserve() {
	C.cairo_stroke_preserve(sc.hnd)
}

func (sc *stdCairo) StrokeExtents() []float64 {
	var x1, y1, x2, y2 C.double
	C.cairo_stroke_extents(sc.hnd, &x1, &y1, &x2, &y2)
	return []float64{float64(x1),float64(y1),float64(x2),float64(y2)}
}

func (sc *stdCairo) InStroke(x, y float64) bool {
	return C.cairo_in_stroke(sc.hnd, C.double(x), C.double(y)) > 0
}

func (sc *stdCairo) CopyPage() {
	C.cairo_copy_page(sc.hnd)
}

func (sc *stdCairo) ShowPage() {
	C.cairo_show_page(sc.hnd)
}

func (sc *stdCairo) SetUserData(key string, data interface{}) {
	if sc.userdata_r == nil {
		// Attempt to load the Go ref table from the C-held ref key.
		refkey := C.cgo_get_refkey(C.cairo_get_user_data(sc.hnd, C.cgo_get_cairo_userdata_key(GO_DATAKEY_KEY)))
	 	if refkey != 0 {
			var ok bool
			sc.userdata_r, ok = LookupGlobalReference(uint32(refkey))
			if !ok {
				panic("missing global reference to device userdata")
			}
		} else {
			userdataMap := make(map[string]interface{})
			sc.userdata_r = MakeGlobalReference(userdataMap)
			C.cairo_device_set_user_data(sc.hnd, C.cgo_get_cairo_userdata_key(GO_DATAKEY_KEY), 
			                             C.cgo_get_keyref(C.uint32_t(sc.userdata_r.Key())), C.cgo_cairo_userdata_destroy)
			IncrementGlobalReferenceCount(sc.userdata_r)
		}
	}
	userdata := sc.userdata_r.Ref().(map[string]interface{})
	userdata[key] = data  
}

func (sc *stdCairo) GetUserData(key string) (interface{},bool) {
	if sc.userdata_r == nil {
		// Attempt to load the Go ref table from the C-held ref key.
		refkey := C.cgo_get_refkey(C.cairo_device_get_user_data(sc.hnd, C.cgo_get_cairo_userdata_key(GO_DATAKEY_KEY)))
	 	if refkey != 0 {
			var ok bool
			sc.userdata_r, ok = LookupGlobalReference(uint32(refkey))
			if !ok {
				panic("missing global reference to device userdata")
			}
		} else {
			return nil, false
		}
	}
	userdata := sc.userdata_r.Ref().(map[string]interface{})
	val, has := userdata[key]
	return val, has
}