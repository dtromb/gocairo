package cairo

import (
	"errors"
	"runtime"
)

/*
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
    #include <cairo/cairo-pdf.h>
	#include <inttypes.h>
	#include <stdlib.h>
		
	extern cairo_user_data_key_t *cgo_get_cairo_userdata_key(int32_t keyid);
	extern uint32_t cgo_get_refkey(void *cref);
	extern void* cgo_get_keyref(uint32_t key);
	
	void cgo_cairo_pattern_userdata_destroy_c(void *sptr) {
		FreePatternNotify((cairo_pattern_t*)sptr);
	}
	cairo_destroy_func_t cgo_cairo_pattern_userdata_destroy = cgo_cairo_pattern_userdata_destroy_c;

*/
import "C"

type PatternType uint32 
const(
	PatternTypeSolid		PatternType	= C.CAIRO_PATTERN_TYPE_SOLID
	PatternTypeSurface					= C.CAIRO_PATTERN_TYPE_SURFACE
	PatternTypeLinear					= C.CAIRO_PATTERN_TYPE_LINEAR
	PatternTypeRadial					= C.CAIRO_PATTERN_TYPE_RADIAL
	PatternTypeMesh						= C.CAIRO_PATTERN_TYPE_MESH
	PatternTypeRasterSource				= C.CAIRO_PATTERN_TYPE_RASTER_SOURCE
)

type Pattern interface {
	AddColorStopRgb(offset float64, r, g, b float64)
	AddColorStopRgba(offset float64, r, g, b, a float64)
	GetColorStopCount() (int,error)
	GetColorStopRgba(idx int) ([]float64,error)
	GetRgba() ([]float64,error)
	GetSurface() (Surface,error)
	GetLinearPoints() ([]float64, error)
	GetRadialCircles() ([]float64, error)
	Status() Status
	SetExtend(ex Extend) 
	GetExtend() Extend 
	SetFilter(ex Filter) 
	GetFilter() Filter 
	SetMatrix(matrix Matrix)
	GetMatrix() Matrix
	SetUserData(key string, data interface{})
	GetUserData(key string) (interface{},bool)
	GetType() PatternType
}

type MeshPattern interface {
	Pattern
	BeginPatch()
	EndPatch()
	MoveTo(x, y float64)
	LineTo(x, y float64)
	CurveTo(x1, y1, x2, y2, x3, y3 float64)
	SetControlPoint(idx int, x, y float64)
	SetCornerColorRgb(idx int, r, g, b float64)
	SetCornerColorRgba(idx int, r, g, b, a float64)
	GetPatchCount() (uint,error)
	GetPath(idx int) Path
	GetControlPoint(patch int, path int) ([]float64, error)
	GetCornerColorRgba(patch int, corneridx int) ([]float64, error)
}

type stdPattern struct {
	hnd *C.cairo_pattern_t
	userdata_r Reference
}

func referencePattern(cref *C.cairo_pattern_t) Pattern {
	ptype := C.cairo_pattern_get_type(cref)
	switch(ptype) {
		case PatternTypeMesh: {
			return blessMeshPattern(cref, true)
		}
	}
	return blessPattern(cref, true)
}

func destroyPattern(p Pattern) {
	if fn, ok := p.(Finalizable); ok {
		fn.Finalize(p)
	}
	if sp, ok := p.(*stdPattern); ok {
		C.cairo_pattern_destroy(sp.hnd)
		sp.hnd = nil
	}
}

func blessMeshPattern(hnd *C.cairo_pattern_t, addRef bool) MeshPattern {
	sp := blessPattern(hnd,addRef)
	mp := &stdMeshPattern{stdPattern:sp.(*stdPattern)}
	return mp
}

func blessPattern(hnd *C.cairo_pattern_t, addRef bool) Pattern {
	p := &stdPattern{
		hnd: hnd,
	}
	if addRef {
		p.hnd = C.cairo_pattern_reference(p.hnd)
	}
	runtime.SetFinalizer(p, destroyPattern)
	return p
}

func (sp *stdPattern) AddColorStopRgb(offset float64, r, g, b float64) {
	C.cairo_pattern_add_color_stop_rgb(sp.hnd, C.double(offset), C.double(r), C.double(g), C.double(b))
}

func (sp *stdPattern) AddColorStopRgba(offset float64, r, g, b, a float64) {
	C.cairo_pattern_add_color_stop_rgba(sp.hnd, C.double(offset), C.double(r), C.double(g), C.double(b), C.double(a))
}

func (sp *stdPattern) GetColorStopCount() (int,error) {
	var x C.int
	status := Status(C.cairo_pattern_get_color_stop_count(sp.hnd, &x))
	if status != StatusSuccess {
		return 0, errors.New(status.String())
	}
	return int(x), nil
}

func (sp *stdPattern) GetColorStopRgba(idx int) ([]float64,error) {
	var ofs, r, g, b, a C.double
	status := Status(C.cairo_pattern_get_color_stop_rgba(sp.hnd, C.int(idx),
														   &ofs, &r, &g, &b, &a))
	if status != StatusSuccess {
		return nil, errors.New(status.String())
	}
	return []float64{float64(ofs),float64(r),float64(g),float64(b),float64(a)}, nil
}

func (sp *stdPattern) GetRgba() ([]float64,error) {
	var r, g, b, a C.double
	status := Status(C.cairo_pattern_get_rgba(sp.hnd, &r, &g, &b, &a))
	if status != StatusSuccess {
		return nil, errors.New(status.String())
	}
	return []float64{float64(r),float64(g),float64(b),float64(a)}, nil
}

func (sp *stdPattern) GetSurface() (Surface,error) {
	var csurf *C.cairo_surface_t
	status := Status(C.cairo_pattern_get_surface(sp.hnd,&csurf))
	if status != StatusSuccess {
		return nil, errors.New(status.String())
	}
	return referenceSurface(csurf), nil
}

func (sp *stdPattern) GetLinearPoints() ([]float64, error) {
	var x1, y1, x2, y2 C.double
	status := Status(C.cairo_pattern_get_linear_points(sp.hnd, &x1, &y1, &x2, &y2))
	if status != StatusSuccess {
		return nil, errors.New(status.String())
	}
	return []float64{float64(x1),float64(y1),float64(x2),float64(y2)}, nil
}

func (sp *stdPattern) GetRadialCircles() ([]float64, error) {
	var x1, y1, r1, x2, y2, r2 C.double
	status := Status(C.cairo_pattern_get_radial_circles(sp.hnd, &x1, &y1, &r1, &x2, &y2, &r2))
	if status != StatusSuccess {
		return nil, errors.New(status.String())
	}
	return []float64{float64(x1),float64(y1),float64(r1),float64(x2),float64(y2),float64(r2)}, nil
}

func (sp *stdPattern) Status() Status {
	return Status(C.cairo_pattern_status(sp.hnd))
}

func (sp *stdPattern) SetExtend(ex Extend) {
	C.cairo_pattern_set_extend(sp.hnd, C.cairo_extend_t(ex))
}

func (sp *stdPattern) GetExtend() Extend {
	return Extend(C.cairo_pattern_get_extend(sp.hnd))
}

func (sp *stdPattern) SetFilter(filter Filter) {
	C.cairo_pattern_set_filter(sp.hnd, C.cairo_filter_t(filter))
}

func (sp *stdPattern) GetFilter() Filter {
	return Filter(C.cairo_pattern_get_filter(sp.hnd))
}

func (sp *stdPattern) SetMatrix(matrix Matrix) {
	C.cairo_pattern_set_matrix(sp.hnd, matrix.dataref())
}

func (sp *stdPattern) GetMatrix() Matrix {
	m := NewMatrix()
	C.cairo_pattern_get_matrix(sp.hnd, m.dataref())
	return m
}

func (sp *stdPattern) SetUserData(key string, data interface{}) {
	if sp.userdata_r == nil {
		// Attempt to load the Go ref table from the C-held ref key.
		refkey := C.cgo_get_refkey(C.cairo_pattern_get_user_data(sp.hnd, C.cgo_get_cairo_userdata_key(GO_DATAKEY_KEY)))
	 	if refkey != 0 {
			var ok bool
			sp.userdata_r, ok = LookupGlobalReference(uint32(refkey))
			if !ok {
				panic("missing global reference to device userdata")
			}
		} else {
			userdataMap := make(map[string]interface{})
			sp.userdata_r = MakeGlobalReference(userdataMap)
			C.cairo_pattern_set_user_data(sp.hnd, C.cgo_get_cairo_userdata_key(GO_DATAKEY_KEY), 
			                              C.cgo_get_keyref(C.uint32_t(sp.userdata_r.Key())), C.cgo_cairo_pattern_userdata_destroy)
			IncrementGlobalReferenceCount(sp.userdata_r)
		}
	}
	userdata := sp.userdata_r.Ref().(map[string]interface{})
	userdata[key] = data  
}

func (sp *stdPattern) GetUserData(key string) (interface{},bool) {
	if sp.userdata_r == nil {
		// Attempt to load the Go ref table from the C-held ref key.
		refkey := C.cgo_get_refkey(C.cairo_pattern_get_user_data(sp.hnd, C.cgo_get_cairo_userdata_key(GO_DATAKEY_KEY)))
		if refkey != 0 {
			var ok bool
			sp.userdata_r, ok = LookupGlobalReference(uint32(refkey))
			if !ok {
				panic("missing global reference to pattern userdata")
			}
		} else {
			return nil, false
		}
	}
	userdata := sp.userdata_r.Ref().(map[string]interface{})
	val, has := userdata[key]
	return val, has
}		

func (mp *stdPattern) GetType() PatternType {
	return PatternType(C.cairo_pattern_get_type(mp.hnd))
}

type stdMeshPattern struct {
	*stdPattern
}		

func (mp *stdMeshPattern) BeginPatch() {
	C.cairo_mesh_pattern_begin_patch(mp.stdPattern.hnd)
}

func (mp *stdMeshPattern) EndPatch() {
	C.cairo_mesh_pattern_end_patch(mp.stdPattern.hnd)
}

func (mp *stdMeshPattern) MoveTo(x, y float64) {
	C.cairo_mesh_pattern_move_to(mp.stdPattern.hnd, C.double(x), C.double(y))
}

func (mp *stdMeshPattern) LineTo(x, y float64) {
	C.cairo_mesh_pattern_line_to(mp.stdPattern.hnd, C.double(x), C.double(y))
}

func (mp *stdMeshPattern) CurveTo(x1, y1, x2, y2, x3, y3 float64) {
	C.cairo_mesh_pattern_curve_to(mp.stdPattern.hnd, 	C.double(x1), C.double(y1), 
													  	C.double(x2), C.double(y2), 
													  	C.double(x3), C.double(y3))
}

func (mp *stdMeshPattern) SetControlPoint(idx int, x, y float64) {
	C.cairo_mesh_pattern_set_control_point(mp.stdPattern.hnd, C.uint(idx), C.double(x), C.double(y))
}

func (mp *stdMeshPattern) SetCornerColorRgb(idx int, r, g, b float64) {
	C.cairo_mesh_pattern_set_corner_color_rgb(mp.stdPattern.hnd, C.uint(idx), C.double(r), C.double(g), C.double(b))
}

func (mp *stdMeshPattern) SetCornerColorRgba(idx int, r, g, b, a float64) {
	C.cairo_mesh_pattern_set_corner_color_rgba(mp.stdPattern.hnd, C.uint(idx), C.double(r), C.double(g), C.double(b), C.double(a))
}

func (mp *stdMeshPattern) GetPatchCount() (uint,error) {
	var c C.uint
	status := Status(C.cairo_mesh_pattern_get_patch_count(mp.stdPattern.hnd, &c))
	if status != StatusSuccess {
		return 0, errors.New(status.String())
	}
	return uint(c), nil
}

func (mp *stdMeshPattern) GetPath(idx int) Path {
	cpath := C.cairo_mesh_pattern_get_path(mp.stdPattern.hnd, C.uint(idx))
	return blessPath(cpath)
}

func (mp *stdMeshPattern) GetControlPoint(patch int, path int) ([]float64, error) {
	var x, y C.double
	status := Status(C.cairo_mesh_pattern_get_control_point(mp.stdPattern.hnd, C.uint(patch), C.uint(path), &x, &y))
	if status != StatusSuccess {
		return nil, errors.New(status.String())
	}
	return []float64{float64(x),float64(y)}, nil
}

func (mp *stdMeshPattern) GetCornerColorRgba(patch int, corneridx int) ([]float64, error) {
	var r, g, b, a C.double
	status := Status(C.cairo_mesh_pattern_get_corner_color_rgba(mp.stdPattern.hnd, C.uint(patch), C.uint(corneridx), &r, &g, &b, &a))
	if status != StatusSuccess {
		return nil, errors.New(status.String())
	}
	return []float64{float64(r),float64(g),float64(b),float64(a)}, nil
}
