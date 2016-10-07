package cairo

import (
	"errors"
	"reflect"
	"runtime"
	"unsafe"
)

/*
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
	#include <inttypes.h>
	#include <stdlib.h>

	extern cairo_user_data_key_t *cgo_get_cairo_userdata_key(int32_t keyid);
	extern uint32_t cgo_get_refkey(void *cref);
	extern void* cgo_get_keyref(uint32_t key);

	void cgo_cairo_font_face_userdata_destroy_c(void* x) {
		FreeFontFaceNotify((cairo_font_face_t*)x);
	}
	cairo_destroy_func_t cgo_cairo_font_face_userdata_destroy = cgo_cairo_font_face_userdata_destroy_c;

	void cgo_cairo_scaled_font_userdata_destroy_c(void* x) {
		FreeScaledFontNotify((cairo_scaled_font_t*)x);
	}
	cairo_destroy_func_t cgo_cairo_scaled_font_userdata_destroy = cgo_cairo_scaled_font_userdata_destroy_c;

	double cgo_cairo_font_extents_get_ascent(cairo_font_extents_t *extents) {
		return extents->ascent;
	}

	double cgo_cairo_font_extents_get_descent(cairo_font_extents_t *extents) {
		return extents->descent;
	}

	double cgo_cairo_font_extents_get_height(cairo_font_extents_t *extents) {
		return extents->height;
	}

	double cgo_cairo_font_extents_get_max_x_advance(cairo_font_extents_t *extents) {
		return extents->max_x_advance;
	}

	double cgo_cairo_font_extents_get_max_y_advance(cairo_font_extents_t *extents) {
		return extents->max_y_advance;
	}

	double cgo_cairo_text_extents_get_x_bearing(cairo_text_extents_t *extents) {
		return extents->x_bearing;
	}

	double cgo_cairo_text_extents_get_y_bearing(cairo_text_extents_t *extents) {
		return extents->y_bearing;
	}

	double cgo_cairo_text_extents_get_width(cairo_text_extents_t *extents) {
		return extents->width;
	}

	double cgo_cairo_text_extents_get_height(cairo_text_extents_t *extents) {
		return extents->height;
	}

	double cgo_cairo_text_extents_get_x_advance(cairo_text_extents_t *extents) {
		return extents->x_advance;
	}

	double cgo_cairo_text_extents_get_y_advance(cairo_text_extents_t *extents) {
		return extents->y_advance;
	}

	uint32_t cgo_cairo_glyph_get_index(cairo_glyph_t *glyph) {
		return glyph->index;
	}

	double cgo_cairo_glyph_get_x(cairo_glyph_t *glyph) {
		return glyph->x;
	}

	double cgo_cairo_glyph_get_y(cairo_glyph_t *glyph) {
		return glyph->y;
	}

	cairo_glyph_t *cgo_glyph_array_index(cairo_glyph_t *array, int index) {
		return &array[index];
	}

	int cgo_cairo_text_cluster_get_num_bytes(cairo_text_cluster_t *cluster) {
		return cluster->num_bytes;
	}

	int cgo_cairo_text_cluster_get_num_glyphs(cairo_text_cluster_t *cluster) {
		return cluster->num_glyphs;
	}

	cairo_text_cluster_t *cgo_text_cluster_array_index(cairo_text_cluster_t *array, int index) {
		return &array[index];
	}

	void cgo_cairo_init_glyph(cairo_glyph_t *glyph, int index, double x, double y) {
		glyph->index = index;
		glyph->x = x;
		glyph->y = y;
	}

	void cgo_init_glyph_array(cairo_glyph_t *glyphs, int i, int index, double x, double y) {
		glyphs[i].index = index;
		glyphs[i].x = x;
		glyphs[i].y = y;
	}

	void cgo_init_text_cluster_array(cairo_text_cluster_t *cluster, int i, int num_bytes, int num_glyphs) {
		cluster[i].num_bytes = num_bytes;
		cluster[i].num_glyphs = num_glyphs;
	}
*/
import "C"

type FontType uint32

const (
	FontTypeToy    FontType = C.CAIRO_FONT_TYPE_TOY
	FontTypeFt              = C.CAIRO_FONT_TYPE_FT
	FontTypeWin32           = C.CAIRO_FONT_TYPE_WIN32
	FontTypeQuartz          = C.CAIRO_FONT_TYPE_QUARTZ
	FontTypeUser            = C.CAIRO_FONT_TYPE_USER
)

type FontFace interface {
	Status() Status
	GetType() FontType
	SetUserData(key string, data interface{})
	GetUserData(key string) (interface{}, bool)
}

type ToyFontFace interface {
	FontFace
	GetFamily() string
	GetSlant() FontSlant
	GetWeight() FontWeight
}

type FontSlant uint32

const (
	FontSlantNormal  FontSlant = C.CAIRO_FONT_SLANT_NORMAL
	FontSlantItalic            = C.CAIRO_FONT_SLANT_ITALIC
	FontSlantOblique           = C.CAIRO_FONT_SLANT_OBLIQUE
)

type FontWeight uint32

const (
	FontWeightNormal FontWeight = C.CAIRO_FONT_WEIGHT_NORMAL
	FontWeightBold              = C.CAIRO_FONT_WEIGHT_BOLD
)

type ClusterFlags uint32

const (
	ClusterFlagBackward ClusterFlags = C.CAIRO_TEXT_CLUSTER_FLAG_BACKWARD
)

type TextCluster interface {
	NumBytes() int
	NumGlyphs() int
}

type Glyph interface {
	Index() uint32
	X() float64
	Y() float64
}

type GlyphString interface {
	NumGlyphs() uint32
	Glyph(idx int) Glyph
	Get(buf []Glyph, offset, length int)
	NumClusters() uint32
	Cluster(idx int) TextCluster
	GetClusters(buf []TextCluster, offset, length int)
}

type ScaledFont interface {
	Status() Status
	Extents() FontExtents
	TextExtents(text string) TextExtents
	GlyphExtents(glyphs []Glyph) TextExtents
	TextToGlyphs(x, y float64, text string) (GlyphString, error)
	GetFontFace() FontFace
	GetFontOptions() FontOptions
	GetFontMatrix() Matrix
	GetCtm() Matrix
	GetScaleMatrix() Matrix
	GetType() FontType
	SetUserData(key string, data interface{})
	GetUserData(key string) (interface{}, bool)
}

type FontExtents interface {
	Ascent() float64
	Descent() float64
	Height() float64
	MaxXAdvance() float64
	MaxYAdvance() float64
}

type TextExtents interface {
	XBearing() float64
	YBearing() float64
	Width() float64
	Height() float64
	XAdvance() float64
	YAdvance() float64
}

func destroyFontFace(f FontFace) {
	if fn, ok := f.(Finalizable); ok {
		fn.Finalize(f)
	}
	if sf, ok := f.(*StdFontFace); ok {
		C.cairo_font_face_destroy(sf.hnd)
		sf.hnd = nil
	}
}

func referenceFontFace(hnd *C.cairo_font_face_t) FontFace {
	ff := blessFontFace(hnd, true)
	return ff
}

func blessFontFace(hnd *C.cairo_font_face_t, addRef bool) FontFace {
	s := &StdFontFace{
		hnd: hnd,
	}
	if addRef {
		s.hnd = C.cairo_font_face_reference(s.hnd)
	}
	runtime.SetFinalizer(s, destroyFontFace)
	return s
}

type StdFontFace struct {
	hnd        *C.cairo_font_face_t
	userdata_r Reference
}

func (f *StdFontFace) Status() Status {
	return Status(C.cairo_font_face_status(f.hnd))
}

func (f *StdFontFace) GetType() FontType {
	return FontType(C.cairo_font_face_get_type(f.hnd))
}

func (f *StdFontFace) SetUserData(key string, data interface{}) {
	if f.userdata_r == nil {
		// Attempt to load the Go ref table from the C-held ref key.
		refkey := C.cgo_get_refkey(C.cairo_device_get_user_data(f.hnd, C.cgo_get_cairo_userdata_key(GO_DATAKEY_KEY)))
		if refkey != 0 {
			var ok bool
			f.userdata_r, ok = LookupGlobalReference(uint32(refkey))
			if !ok {
				panic("missing global reference to font face userdata")
			}
		} else {
			userdataMap := make(map[string]interface{})
			f.userdata_r = MakeGlobalReference(userdataMap)
			C.cairo_device_set_user_data(f.hnd, C.cgo_get_cairo_userdata_key(GO_DATAKEY_KEY),
				C.cgo_get_keyref(C.uint32_t(f.userdata_r.Key())), C.cgo_cairo_font_face_userdata_destroy)
			IncrementGlobalReferenceCount(f.userdata_r)
		}
	}
	userdata := f.userdata_r.Ref().(map[string]interface{})
	userdata[key] = data
}

func (f *StdFontFace) Hnd() uintptr {
	return uintptr(unsafe.Pointer(f.hnd))
}

func (f *StdFontFace) GetUserData(key string) (interface{}, bool) {
	if f.userdata_r == nil {
		// Attempt to load the Go ref table from the C-held ref key.
		refkey := C.cgo_get_refkey(C.cairo_font_face_get_user_data(f.hnd, C.cgo_get_cairo_userdata_key(GO_DATAKEY_KEY)))
		if refkey != 0 {
			var ok bool
			f.userdata_r, ok = LookupGlobalReference(uint32(refkey))
			if !ok {
				panic("missing global reference to font face userdata")
			}
		} else {
			return nil, false
		}
	}
	userdata := f.userdata_r.Ref().(map[string]interface{})
	val, has := userdata[key]
	return val, has
}

type stdScaledFont struct {
	hnd        *C.cairo_scaled_font_t
	userdata_r Reference
}

func destroyScaledFont(x ScaledFont) {
	if f, ok := x.(Finalizable); ok {
		f.Finalize(x)
	}
	if ssf, ok := x.(*stdScaledFont); ok {
		C.cairo_scaled_font_destroy(ssf.hnd)
		ssf.hnd = nil
	}
}

func blessScaledFont(hnd *C.cairo_scaled_font_t, addRef bool) ScaledFont {
	ssf := &stdScaledFont{hnd: hnd}
	if addRef {
		C.cairo_scaled_font_reference(hnd)
	}
	runtime.SetFinalizer(ssf, destroyScaledFont)
	return ssf
}

func (ssf *stdScaledFont) Status() Status {
	return Status(C.cairo_scaled_font_status(ssf.hnd))
}

func (ssf *stdScaledFont) Extents() FontExtents {
	extents := stdFontExtents{}
	C.cairo_scaled_font_extents(ssf.hnd, &extents.extents)
	return &extents
}

func (ssf *stdScaledFont) TextExtents(text string) TextExtents {
	extents := stdTextExtents{}
	C.cairo_scaled_font_text_extents(ssf.hnd, C.CString(text), &extents.extents)
	return &extents
}

func (ssf *stdScaledFont) GlyphExtents(glyphs []Glyph) TextExtents {
	arg := make([]stdGlyph, len(glyphs))
	for i, g := range glyphs {
		if sg, ok := g.(*stdGlyph); ok {
			arg[i] = *sg
		} else {
			arg[i] = *NewGlyph(g.Index(), g.X(), g.Y()).(*stdGlyph)
		}
	}
	extents := stdTextExtents{}
	hdr := *(*reflect.SliceHeader)(unsafe.Pointer(&arg))
	C.cairo_scaled_font_glyph_extents(ssf.hnd, (*C.cairo_glyph_t)(unsafe.Pointer(hdr.Data)), C.int(hdr.Len), &extents.extents)
	return &extents
}

func (ssf *stdScaledFont) TextToGlyphs(x, y float64, text string) (GlyphString, error) {
	var glyphs, clusters uintptr
	var numGlyphs, numClusters C.int
	var clusterFlags C.cairo_text_cluster_flags_t
	status := Status(C.cairo_scaled_font_text_to_glyphs(ssf.hnd, C.double(x), C.double(y),
		C.CString(text), C.int(len(text)),
		(**C.cairo_glyph_t)(unsafe.Pointer(&glyphs)), &numGlyphs,
		(**C.cairo_text_cluster_t)(unsafe.Pointer(&clusters)), &numClusters,
		&clusterFlags))
	if status != StatusSuccess {
		return nil, errors.New(status.String())
	}
	result := &stdGlyphString{
		glyphs:       (*C.cairo_glyph_t)(unsafe.Pointer(glyphs)),
		numGlyphs:    uint32(numGlyphs),
		clusters:     (*C.cairo_text_cluster_t)(unsafe.Pointer(clusters)),
		numClusters:  uint32(numClusters),
		clusterFlags: ClusterFlags(clusterFlags),
	}
	runtime.SetFinalizer(result, destroyGlyphString)
	return result, nil
}

func (ssf *stdScaledFont) GetFontFace() FontFace {
	return referenceFontFace(C.cairo_scaled_font_get_font_face(ssf.hnd))
}

func (ssf *stdScaledFont) GetFontOptions() FontOptions {
	opts := NewFontOptions().(*stdFontOptions)
	C.cairo_scaled_font_get_font_options(ssf.hnd, opts.hnd)
	return opts
}

func (ssf *stdScaledFont) GetFontMatrix() Matrix {
	matrix := NewMatrix()
	C.cairo_scaled_font_get_font_matrix(ssf.hnd, matrix.DataRef())
	return matrix
}

func (ssf *stdScaledFont) GetCtm() Matrix {
	matrix := NewMatrix()
	C.cairo_scaled_font_get_ctm(ssf.hnd, matrix.DataRef())
	return matrix
}

func (ssf *stdScaledFont) GetScaleMatrix() Matrix {
	matrix := NewMatrix()
	C.cairo_scaled_font_get_scale_matrix(ssf.hnd, matrix.DataRef())
	return matrix
}

func (ssf *stdScaledFont) GetType() FontType {
	return FontType(C.cairo_scaled_font_get_type(ssf.hnd))
}

func (ssf *stdScaledFont) SetUserData(key string, data interface{}) {
	if ssf.userdata_r == nil {
		// Attempt to load the Go ref table from the C-held ref key.
		refkey := C.cgo_get_refkey(C.cairo_scaled_font_get_user_data(ssf.hnd, C.cgo_get_cairo_userdata_key(GO_DATAKEY_KEY)))
		if refkey != 0 {
			var ok bool
			ssf.userdata_r, ok = LookupGlobalReference(uint32(refkey))
			if !ok {
				panic("missing global reference to scaled font userdata")
			}
		} else {
			userdataMap := make(map[string]interface{})
			ssf.userdata_r = MakeGlobalReference(userdataMap)
			C.cairo_device_set_user_data(ssf.hnd, C.cgo_get_cairo_userdata_key(GO_DATAKEY_KEY),
				C.cgo_get_keyref(C.uint32_t(ssf.userdata_r.Key())), C.cgo_cairo_scaled_font_userdata_destroy)
			IncrementGlobalReferenceCount(ssf.userdata_r)
		}
	}
	userdata := ssf.userdata_r.Ref().(map[string]interface{})
	userdata[key] = data
}

func (ssf *stdScaledFont) GetUserData(key string) (interface{}, bool) {
	if ssf.userdata_r == nil {
		// Attempt to load the Go ref table from the C-held ref key.
		refkey := C.cgo_get_refkey(C.cairo_scaled_font_get_user_data(ssf.hnd, C.cgo_get_cairo_userdata_key(GO_DATAKEY_KEY)))
		if refkey != 0 {
			var ok bool
			ssf.userdata_r, ok = LookupGlobalReference(uint32(refkey))
			if !ok {
				panic("missing global reference to scaled font userdata")
			}
		} else {
			return nil, false
		}
	}
	userdata := ssf.userdata_r.Ref().(map[string]interface{})
	val, has := userdata[key]
	return val, has
}

type stdFontExtents struct {
	extents C.cairo_font_extents_t
}

func (sfe *stdFontExtents) Ascent() float64 {
	return float64(C.cgo_cairo_font_extents_get_ascent(&sfe.extents))
}

func (sfe *stdFontExtents) Descent() float64 {
	return float64(C.cgo_cairo_font_extents_get_descent(&sfe.extents))
}

func (sfe *stdFontExtents) Height() float64 {
	return float64(C.cgo_cairo_font_extents_get_height(&sfe.extents))
}

func (sfe *stdFontExtents) MaxXAdvance() float64 {
	return float64(C.cgo_cairo_font_extents_get_max_x_advance(&sfe.extents))
}

func (sfe *stdFontExtents) MaxYAdvance() float64 {
	return float64(C.cgo_cairo_font_extents_get_max_y_advance(&sfe.extents))
}

type stdTextExtents struct {
	extents C.cairo_text_extents_t
}

func (ste *stdTextExtents) XBearing() float64 {
	return float64(C.cgo_cairo_text_extents_get_x_bearing(&ste.extents))
}

func (ste *stdTextExtents) YBearing() float64 {
	return float64(C.cgo_cairo_text_extents_get_y_bearing(&ste.extents))
}

func (ste *stdTextExtents) Width() float64 {
	return float64(C.cgo_cairo_text_extents_get_width(&ste.extents))
}

func (ste *stdTextExtents) Height() float64 {
	return float64(C.cgo_cairo_text_extents_get_height(&ste.extents))
}

func (ste *stdTextExtents) XAdvance() float64 {
	return float64(C.cgo_cairo_text_extents_get_x_advance(&ste.extents))
}

func (ste *stdTextExtents) YAdvance() float64 {
	return float64(C.cgo_cairo_text_extents_get_y_advance(&ste.extents))
}

type stdGlyph struct {
	glyph C.cairo_glyph_t
}

func NewGlyph(index uint32, x, y float64) Glyph {
	g := stdGlyph{}
	return &g
}

func (g *stdGlyph) Index() uint32 {
	return uint32(C.cgo_cairo_glyph_get_index(&g.glyph))
}

func (g *stdGlyph) X() float64 {
	return float64(C.cgo_cairo_glyph_get_index(&g.glyph))
}

func (g *stdGlyph) Y() float64 {
	return float64(C.cgo_cairo_glyph_get_index(&g.glyph))
}

type stdTextCluster struct {
	cluster C.cairo_text_cluster_t
}

func (stc *stdTextCluster) NumBytes() int {
	return int(C.cgo_cairo_text_cluster_get_num_bytes(&stc.cluster))
}

func (stc *stdTextCluster) NumGlyphs() int {
	return int(C.cgo_cairo_text_cluster_get_num_glyphs(&stc.cluster))
}

type stdGlyphString struct {
	glyphs       *C.cairo_glyph_t
	clusters     *C.cairo_text_cluster_t
	numGlyphs    uint32
	numClusters  uint32
	clusterFlags ClusterFlags
}

func destroyGlyphString(gs GlyphString) {
	if f, ok := gs.(Finalizable); ok {
		f.Finalize(gs)
	}
	if sgs, ok := gs.(*stdGlyphString); ok {
		C.cairo_glyph_free(sgs.glyphs)
		C.cairo_text_cluster_free(sgs.clusters)
		sgs.glyphs = nil
		sgs.numGlyphs = 0
		sgs.clusters = nil
		sgs.numClusters = 0
	}
}

func importGlyphString(gs GlyphString) *stdGlyphString {
	if sgs, ok := gs.(*stdGlyphString); ok {
		return sgs
	}
	sgs := &stdGlyphString{}
	sgs.numClusters = gs.NumClusters()
	sgs.numGlyphs = gs.NumGlyphs()
	sgs.glyphs = C.cairo_glyph_allocate(C.int(sgs.numGlyphs))
	sgs.clusters = C.cairo_text_cluster_allocate(C.int(sgs.numClusters))
	for i := 0; i < int(gs.NumGlyphs()); i++ {
		g := gs.Glyph(i)
		C.cgo_init_glyph_array(sgs.glyphs, C.int(i), C.int(g.Index()), C.double(g.X()), C.double(g.Y()))
	}
	for i := 0; i < int(gs.NumClusters()); i++ {
		c := gs.Cluster(i)
		C.cgo_init_text_cluster_array(sgs.clusters, C.int(i), C.int(c.NumBytes()), C.int(c.NumGlyphs()))
	}
	runtime.SetFinalizer(sgs, destroyGlyphString)
	return sgs
}

func (sgs *stdGlyphString) NumGlyphs() uint32 {
	return sgs.numGlyphs
}

func (sgs *stdGlyphString) Glyph(idx int) Glyph {
	if idx < 0 || idx >= int(sgs.numGlyphs) {
		return nil
	}
	return &stdGlyph{glyph: *C.cgo_glyph_array_index(sgs.glyphs, C.int(idx))}
}

func (sgs *stdGlyphString) Get(buf []Glyph, offset, length int) {
	if offset < 0 || offset >= int(sgs.numGlyphs) {
		return
	}
	if length > len(buf) {
		length = len(buf)
	}
	if offset+length > int(sgs.numGlyphs) {
		length = int(sgs.numGlyphs) - offset
	}
	for k := offset; k < offset+length; k++ {
		buf[k-offset] = sgs.Glyph(k)
	}
}

func (sgs *stdGlyphString) NumClusters() uint32 {
	return sgs.numClusters
}

func (sgs *stdGlyphString) Cluster(idx int) TextCluster {
	if idx < 0 || idx >= int(sgs.numClusters) {
		return nil
	}
	return &stdTextCluster{cluster: *C.cgo_text_cluster_array_index(sgs.clusters, C.int(idx))}
}

func (sgs *stdGlyphString) GetClusters(buf []TextCluster, offset, length int) {
	if offset < 0 || offset >= int(sgs.numClusters) {
		return
	}
	if length > len(buf) {
		length = len(buf)
	}
	if offset+length > int(sgs.numClusters) {
		length = int(sgs.numClusters) - offset
	}
	for k := offset; k < offset+length; k++ {
		buf[k-offset] = sgs.Cluster(k)
	}
}

func (ctx *stdCairo) SelectFontFace(face string, slant FontSlant, weight FontWeight) {
	C.cairo_select_font_face(ctx.hnd, C.CString(face), C.cairo_font_slant_t(slant), C.cairo_font_weight_t(weight))
}

func (ctx *stdCairo) SetFontSize(size float64) {
	C.cairo_set_font_size(ctx.hnd, C.double(size))
}

func (ctx *stdCairo) SetFontMatrix(matrix Matrix) {
	C.cairo_set_font_matrix(ctx.hnd, matrix.DataRef())
}

func (ctx *stdCairo) GetFontMatrix() Matrix {
	matrix := NewMatrix()
	C.cairo_get_font_matrix(ctx.hnd, matrix.DataRef())
	return matrix
}

func (ctx *stdCairo) SetFontOptions(opts FontOptions) {
	if sfo, ok := opts.(*stdFontOptions); ok {
		C.cairo_set_font_options(ctx.hnd, sfo.hnd)
	} else {
		panic("stdCairo.SetFontOptions(opts) unimplemented for non-standard font options arguments")
	}
}

func (ctx *stdCairo) GetFontOptions() FontOptions {
	opts := NewFontOptions().(*stdFontOptions)
	C.cairo_get_font_options(ctx.hnd, opts.hnd)
	return opts
}

func (ctx *stdCairo) SetFontFace(face FontFace) {
	if cp, ok := face.(CPeer); ok {
		hnd := (*C.cairo_font_face_t)(unsafe.Pointer(cp.Hnd()))
		C.cairo_set_font_face(ctx.hnd, hnd)
	} else {
		panic("FontFace implementation did not have a C language peer")
	}
}

func (ctx *stdCairo) GetFontFace() FontFace {
	return blessFontFace(C.cairo_get_font_face(ctx.hnd), true)
}

func (ctx *stdCairo) SetScaledFont(font ScaledFont) {
	if ssf, ok := font.(*stdScaledFont); ok {
		C.cairo_set_scaled_font(ctx.hnd, ssf.hnd)
	} else {
		panic("stdCairo.SetScaledFont(font) unimplemented for non-standard scaled font arguments")
	}
}

func (ctx *stdCairo) GetScaledFont() ScaledFont {
	return blessScaledFont(C.cairo_get_scaled_font(ctx.hnd), true)
}

func (ctx *stdCairo) ShowText(str string) {
	C.cairo_show_text(ctx.hnd, C.CString(str))
}

func (ctx *stdCairo) ShowGlyphs(glyphs []Glyph) {
	if len(glyphs) == 0 {
		return
	}
	sglyphs := make([]C.cairo_glyph_t, len(glyphs))
	for i, g := range glyphs {
		if sg, ok := g.(*stdGlyph); ok {
			sglyphs[i] = sg.glyph
		} else {
			C.cgo_cairo_init_glyph(&sglyphs[i], C.int(g.Index()), C.double(g.X()), C.double(g.Y()))
		}
	}
	hdr := *(*reflect.SliceHeader)(unsafe.Pointer(&sglyphs))
	C.cairo_show_glyphs(ctx.hnd, (*C.cairo_glyph_t)(unsafe.Pointer(hdr.Data)), C.int(hdr.Len))
}

func (ctx *stdCairo) ShowTextGlyphs(text string, glyphs GlyphString) {
	var sgs *stdGlyphString
	var ok bool
	if sgs, ok = glyphs.(*stdGlyphString); !ok {
		sgs = importGlyphString(glyphs)
	}
	C.cairo_show_text_glyphs(ctx.hnd, C.CString(text), C.int(len(text)),
		sgs.glyphs, C.int(sgs.numGlyphs),
		sgs.clusters, C.int(sgs.numClusters),
		C.cairo_text_cluster_flags_t(sgs.clusterFlags))
}

func (ctx *stdCairo) FontExtents() FontExtents {
	exts := stdFontExtents{}
	C.cairo_font_extents(ctx.hnd, &exts.extents)
	return &exts
}

func (ctx *stdCairo) TextExtents(str string) TextExtents {
	exts := stdTextExtents{}
	C.cairo_text_extents(ctx.hnd, C.CString(str), &exts.extents)
	return &exts
}

func (ctx *stdCairo) GlyphExtents(glyphs []Glyph) TextExtents {
	exts := stdTextExtents{}
	sglyphs := make([]C.cairo_glyph_t, len(glyphs))
	for i, g := range glyphs {
		if sg, ok := g.(*stdGlyph); ok {
			sglyphs[i] = sg.glyph
		} else {
			C.cgo_cairo_init_glyph(&sglyphs[i], C.int(g.Index()), C.double(g.X()), C.double(g.Y()))
		}
	}
	hdr := *(*reflect.SliceHeader)(unsafe.Pointer(&sglyphs))
	C.cairo_glyph_extents(ctx.hnd, (*C.cairo_glyph_t)(unsafe.Pointer(hdr.Data)), C.int(hdr.Len), &exts.extents)
	return &exts
}

type stdToyFontFace struct {
	StdFontFace
}

func (ctx *stdCairo) ToyFontFaceCreate(face string, slant FontSlant, weight FontWeight) ToyFontFace {
	hnd := C.cairo_toy_font_face_create(C.CString(face), C.cairo_font_slant_t(slant), C.cairo_font_weight_t(weight))
	stf := &stdToyFontFace{StdFontFace: StdFontFace{hnd: hnd}}
	blessFontFace(stf.hnd, true) // We add a ref because the return value here is ignored and will finalize...
	return stf
}

func (stf *stdToyFontFace) GetFamily() string {
	return C.GoString(C.cairo_toy_font_face_get_family(stf.hnd))
}

func (stf *stdToyFontFace) GetSlant() FontSlant {
	return FontSlant(C.cairo_toy_font_face_get_slant(stf.hnd))
}

func (stf *stdToyFontFace) GetWeight() FontWeight {
	return FontWeight(C.cairo_toy_font_face_get_weight(stf.hnd))
}
