package fontconfig

import (
	"errors"
	"reflect"
	"runtime"
	"strconv"
	"unsafe"
	//"unsafe"
	//"reflect"
	//"errors"
	//"runtime"
	cairo "github.com/dtromb/gocairo"
)

/*
	#cgo CFLAGS: -I/usr/include/freetype2
	#cgo LDFLAGS: -lcairo -lfontconfig
    #include <cairo/cairo.h>
	#include <inttypes.h>
	#include <stdlib.h>

	extern int cgo_fontconfig_supported();

	#ifdef CAIRO_HAS_FC_FONT

		#include <fontconfig/fontconfig.h>

		// XXX - Fc is only used for FT support in this project.
		// We need to disable Fc here and shim if FT is not available.

		#include <freetype2/ft2build.h>
		#include FT_FREETYPE_H

		void cgo_fontconfig_set_fcvalue_int(FcValue *v, int x) {
			v->type = FcTypeInteger;
			v->u.i = x;
		}

		void cgo_fontconfig_set_fcvalue_double(FcValue *v, double x) {
			v->type = FcTypeDouble;
			v->u.d = x;
		}

		void cgo_fontconfig_set_fcvalue_string(FcValue *v, char *x) {
			v->type = FcTypeString;
			v->u.s = (FcChar8*)x;
		}

		void cgo_fontconfig_set_fcvalue_bool(FcValue *v, FcBool x) {
			v->type = FcTypeBool;
			v->u.b = x;
		}

		void cgo_fontconfig_set_fcvalue_matrix(FcValue *v, FcMatrix *x) {
			v->type = FcTypeMatrix;
			v->u.m = x;
		}

		void cgo_fontconfig_set_fcvalue_charset(FcValue *v, FcCharSet *x) {
			v->type = FcTypeCharSet;
			v->u.c = x;
		}

		void cgo_fontconfig_set_fcvalue_ftface(FcValue *v, void *x) {
			v->type = FcTypeFTFace;
			v->u.f = (FT_Face)x;
		}

		void cgo_fontconfig_set_fcvalue_langset(FcValue *v, FcLangSet *x) {
			v->type = FcTypeLangSet;
			v->u.l = x;
		}

		FcType cgo_fontconfig_value_get_type(FcValue *v) {
			return v->type;
		}

		const FcCharSet *cgo_fontconfig_value_get_char_set(FcValue *fcv) {
			return fcv->u.c;
		}

		const FcLangSet *cgo_fontconfig_value_get_lang_set(FcValue *fcv) {
			return fcv->u.l;
		}

		double cgo_fontconfig_value_get_matrix_xx(FcValue *fcv) {
			return fcv->u.m->xx;
		}

		double cgo_fontconfig_value_get_matrix_xy(FcValue *fcv) {
			return fcv->u.m->xy;
		}

		double cgo_fontconfig_value_get_matrix_yx(FcValue *fcv) {
			return fcv->u.m->yx;
		}

		double cgo_fontconfig_value_get_matrix_yy(FcValue *fcv) {
			return fcv->u.m->yy;
		}

		void *cgo_fontconfig_value_get_ft_face(FcValue *fcv) {
			return fcv->u.f;
		}

		double cgo_fontconfig_value_get_double(FcValue *fcv) {
			return fcv->u.d;
		}

		int cgo_fontconfig_value_get_integer(FcValue *fcv) {
			return fcv->u.i;
		}

		const char *cgo_fontconfig_value_get_string(FcValue *fcv) {
			return fcv->u.s;
		}

		FcBool cgo_fontconfig_value_get_bool(FcValue *fcv) {
			return fcv->u.b;
		}

		double cgo_fontconfig_fcmatrix_get_xx(FcMatrix *m) {
			return m->xx;
		}

		double cgo_fontconfig_fcmatrix_get_yx(FcMatrix *m) {
			return m->yx;
		}

		double cgo_fontconfig_fcmatrix_get_xy(FcMatrix *m) {
			return m->xy;
		}

		double cgo_fontconfig_fcmatrix_get_yy(FcMatrix *m) {
			return m->yy;
		}

	#else
	#endif
*/
import "C"

//  XXX - Fc docs do not say what the FcBool return args for the pattern
//  modify functions indicate.   So, we just ignore them.

type Pattern interface {
	Duplicate() Pattern
	Equal(p Pattern) bool
	EqualSubset(p Pattern, os ObjectSet) bool
	Filter(os ObjectSet) Pattern
	Hash() uint32
	Add(name string, value interface{}, addAppend bool)
	AddWeak(name string, value interface{}, addAppend bool)
	Get(name string, id int) (interface{}, bool, error)
	GetInteger(name string, n int) (int, bool, error)
	GetDouble(name string, n int) (float64, bool, error)
	GetString(name string, n int) (string, bool, error)
	GetMatrix(name string, n int) (cairo.Matrix, bool, error)
	GetCharSet(name string, n int) (CharSet, bool, error)
	GetBool(name string, n int) (bool, bool, error)
	GetFtFace(name string, n int) (cairo.FtFace, bool, error)
	GetLangSet(name string, n int) (LangSet, bool, error)
	GetRange(name string, n int) (Range, bool, error)
	Del(name string)
	Remove(name string, n int)
	DefaultSubstitute()
	NameUnparse() string
	Format(fmt string) string
}

type stdPattern struct {
	hnd *C.FcPattern
}

func destroyPattern(p Pattern) {
	if f, ok := p.(cairo.Finalizable); ok {
		f.Finalize(p)
	}
	if sp, ok := p.(*stdPattern); ok {
		C.FcPatternDestroy(sp.hnd)
	}
}

func blessPattern(hnd *C.FcPattern, addRef bool) Pattern {
	pattern := &stdPattern{hnd: hnd}
	if addRef {
		C.FcPatternReference(hnd)
	}
	runtime.SetFinalizer(pattern, destroyPattern)
	return pattern
}

func referencePattern(hnd *C.FcPattern) Pattern {
	return blessPattern(hnd, true)
}

func FcPatternCreate() Pattern {
	pattern := blessPattern(C.FcPatternCreate(), false)
	return pattern
}
func FcNameParse(name string) Pattern {
	// XXX - Incref here or not??   Docs are silent on which of these reference... =/
	return blessPattern(C.FcNameParse((*C.FcChar8)(unsafe.Pointer(C.CString(name)))), false)
}

type FcResult uint32

const (
	FcResultMatch        FcResult = C.FcResultMatch
	FcResultNoMatch               = C.FcResultNoMatch
	FcResultTypeMismatch          = C.FcResultTypeMismatch
	FcResultNoId                  = C.FcResultNoId
	FcResultOutOfMemory           = C.FcResultOutOfMemory
)

func (fcr FcResult) String() string {
	switch fcr {
	case FcResultMatch:
		return "success"
	case FcResultNoMatch:
		return "object not found"
	case FcResultTypeMismatch:
		return "object type mismatch"
	case FcResultNoId:
		return "object id out of range"
	case FcResultOutOfMemory:
		return "out of memory"
	}
	return "?"
}

type FcType int8

const (
	FcTypeUnknown FcType = C.FcTypeUnknown
	FcTypeVoid           = C.FcTypeVoid
	FcTypeInteger        = C.FcTypeInteger
	FcTypeDouble         = C.FcTypeDouble
	FcTypeString         = C.FcTypeString
	FcTypeBool           = C.FcTypeBool
	FcTypeMatrix         = C.FcTypeMatrix
	FcTypeCharSet        = C.FcTypeCharSet
	FcTypeFtFace         = C.FcTypeFTFace
	FcTypeLangSet        = C.FcTypeLangSet
)

func fcTypeForInterface(obj interface{}) FcType {
	val := reflect.ValueOf(obj)
	switch val.Type().Kind() {
	case reflect.Bool:
		return FcTypeBool
	case reflect.Int:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Uint:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
		return FcTypeInteger
	case reflect.Float32:
	case reflect.Float64:
		return FcTypeDouble
	}
	if _, ok := obj.(string); ok {
		return FcTypeString
	}
	if _, ok := obj.(cairo.FtFace); ok {
		return FcTypeFtFace
	}
	if _, ok := obj.(cairo.Matrix); ok {
		return FcTypeMatrix
	}
	if _, ok := obj.(CharSet); ok {
		return FcTypeCharSet
	}
	if _, ok := obj.(LangSet); ok {
		return FcTypeLangSet
	}
	return FcTypeUnknown
}

func (pt *stdPattern) Add(name string, value interface{}, addAppend bool) {
	propType := fcTypeForInterface(value)
	var fcv C.FcValue
	switch propType {
	case FcTypeInteger:
		{
			C.cgo_fontconfig_set_fcvalue_int(&fcv, C.int(reflect.ValueOf(value).Int()))
		}
	case FcTypeDouble:
		{
			C.cgo_fontconfig_set_fcvalue_double(&fcv, C.double(reflect.ValueOf(value).Float()))
		}
	case FcTypeString:
		{
			C.cgo_fontconfig_set_fcvalue_string(&fcv, C.CString(reflect.ValueOf(value).String()))
		}
	case FcTypeBool:
		{
			bv := reflect.ValueOf(value).Bool()
			k := 0
			if bv {
				k = 1
			}
			C.cgo_fontconfig_set_fcvalue_bool(&fcv, C.FcBool(k))
		}
	case FcTypeMatrix:
		{
			C.cgo_fontconfig_set_fcvalue_matrix(&fcv, (*C.FcMatrix)(unsafe.Pointer(value.(cairo.Matrix).DataRef())))
		}
	case FcTypeCharSet:
		{
			if scs, ok := value.(*stdCharSet); ok {
				C.cgo_fontconfig_set_fcvalue_charset(&fcv, scs.hnd)
			} else {
				panic("stPattern.AddWeak(name,cs,append) unsupported for non-standard CharSet arguments")
			}
		}
	case FcTypeFtFace:
		{
			if sff, ok := value.(*cairo.StdFontFace); ok {
				C.cgo_fontconfig_set_fcvalue_ftface(&fcv, unsafe.Pointer(sff.Hnd()))
			} else {
				panic("stPattern.AddWeak(name,font,append) unsupported for non-standard FontFace arguments")
			}
		}
	case FcTypeLangSet:
		{
			if sls, ok := value.(*stdLangSet); ok {
				C.cgo_fontconfig_set_fcvalue_langset(&fcv, sls.hnd)
			} else {
				panic("stPattern.Add(name,ls,append) unsupported for non-standard LAngSet arguments")
			}
		}
	default:
		{
			panic("data type " + reflect.TypeOf(value).String() + " not supported by stdPattern.Add()")
		}
	}
	fcb := 0
	if addAppend {
		fcb = 1
	}
	C.FcPatternAdd(pt.hnd, C.CString(name), fcv, C.FcBool(fcb))
}

func (pt *stdPattern) AddWeak(name string, value interface{}, addAppend bool) {
	propType := fcTypeForInterface(value)
	var fcv C.FcValue
	switch propType {
	case FcTypeInteger:
		{
			C.cgo_fontconfig_set_fcvalue_int(&fcv, C.int(reflect.ValueOf(value).Int()))
		}
	case FcTypeDouble:
		{
			C.cgo_fontconfig_set_fcvalue_double(&fcv, C.double(reflect.ValueOf(value).Float()))
		}
	case FcTypeString:
		{
			C.cgo_fontconfig_set_fcvalue_string(&fcv, C.CString(reflect.ValueOf(value).String()))
		}
	case FcTypeBool:
		{
			bv := reflect.ValueOf(value).Bool()
			k := 0
			if bv {
				k = 1
			}
			C.cgo_fontconfig_set_fcvalue_bool(&fcv, C.FcBool(k))
		}
	case FcTypeMatrix:
		{
			C.cgo_fontconfig_set_fcvalue_matrix(&fcv, (*C.FcMatrix)(unsafe.Pointer(value.(cairo.Matrix).DataRef())))
		}
	case FcTypeCharSet:
		{
			if scs, ok := value.(*stdCharSet); ok {
				C.cgo_fontconfig_set_fcvalue_charset(&fcv, scs.hnd)
			} else {
				panic("stPattern.AddWeak(name,cs,append) unsupported for non-standard CharSet arguments")
			}
		}
	case FcTypeFtFace:
		{
			if sff, ok := value.(*cairo.StdFontFace); ok {
				C.cgo_fontconfig_set_fcvalue_ftface(&fcv, unsafe.Pointer(sff.Hnd()))
			} else {
				panic("stPattern.AddWeak(name,font,append) unsupported for non-standard FontFace arguments")
			}
		}
	case FcTypeLangSet:
		{
			if sls, ok := value.(*stdLangSet); ok {
				C.cgo_fontconfig_set_fcvalue_langset(&fcv, sls.hnd)
			} else {
				panic("stPattern.Add(name,ls,append) unsupported for non-standard LAngSet arguments")
			}
		}
	default:
		{
			panic("data type " + reflect.TypeOf(value).String() + " not supported by stdPattern.Add()")
		}
	}
	fcb := 0
	if addAppend {
		fcb = 1
	}
	C.FcPatternAddWeak(pt.hnd, C.CString(name), fcv, C.FcBool(fcb))
}

func (pt *stdPattern) Duplicate() Pattern {
	return blessPattern(C.FcPatternDuplicate(pt.hnd), false)
}

func (pt *stdPattern) Equal(p Pattern) bool {
	if spt, ok := p.(*stdPattern); ok {
		return C.FcPatternEqual(pt.hnd, spt.hnd) > 0
	}
	panic("stdPattern.Equal(p) unimplemented for non-standard Pattern arguments")
}

func (pt *stdPattern) EqualSubset(p Pattern, os ObjectSet) bool {
	if spt, ok := p.(*stdPattern); ok {
		if sos, ok := os.(*stdObjectSet); ok {
			return C.FcPatternEqualSubset(pt.hnd, spt.hnd, sos.hnd) > 0
		}
		panic("stdPattern.EqualSubset(p, os) unimplemented for non-standard ObjectSet arguments")
	}
	panic("stdPattern.EqualSubset(p, os) unimplemented for non-standard Pattern arguments")
}

func (pt *stdPattern) Filter(os ObjectSet) Pattern {
	if sos, ok := os.(*stdObjectSet); ok {
		return blessPattern(C.FcPatternFilter(pt.hnd, sos.hnd), false)
	}
	panic("stdPattern.Filter(os) unimplemented for non-standard ObjectSet arguments")
}

func (pt *stdPattern) Hash() uint32 {
	return uint32(C.FcPatternHash(pt.hnd))
}

func (pt *stdPattern) Get(name string, id int) (interface{}, bool, error) {
	var fcv C.FcValue
	fcr := FcResult(C.FcPatternGet(pt.hnd, C.CString(name), C.int(id), &fcv))
	if fcr != FcResultMatch {
		if fcr == FcResultNoMatch {
			return nil, false, nil
		}
		return nil, false, errors.New(fcr.String())
	}
	var x interface{}
	switch FcType(C.cgo_fontconfig_value_get_type(&fcv)) {
	case FcTypeInteger:
		{
			x = int64(C.cgo_fontconfig_value_get_integer(&fcv))
		}
	case FcTypeDouble:
		{
			x = float64(C.cgo_fontconfig_value_get_double(&fcv))
		}
	case FcTypeString:
		{
			x = C.GoString(C.cgo_fontconfig_value_get_string(&fcv))
		}
	case FcTypeBool:
		{
			if C.cgo_fontconfig_value_get_bool(&fcv) > 0 {
				x = true
			} else {
				x = false
			}
		}
	case FcTypeMatrix:
		{
			matrix := cairo.NewMatrix()
			matrix.Init(float64(C.cgo_fontconfig_value_get_matrix_xx(&fcv)),
				float64(C.cgo_fontconfig_value_get_matrix_yx(&fcv)),
				float64(C.cgo_fontconfig_value_get_matrix_xy(&fcv)),
				float64(C.cgo_fontconfig_value_get_matrix_yy(&fcv)),
				0, 0)
			x = matrix
		}
	case FcTypeCharSet:
		{
			x = blessCharSet(C.cgo_fontconfig_value_get_char_set(&fcv))
		}
	case FcTypeFtFace:
		{
			x = cairo.UnsafeBlessFtFace(uintptr(C.cgo_fontconfig_value_get_ft_face(&fcv)), true)
		}
	case FcTypeLangSet:
		{
			x = blessLangSet(C.cgo_fontconfig_value_get_lang_set(&fcv))
		}
	default:
		{
			return nil, false, errors.New("pattern object had unsupported type")
		}
	}
	return x, true, nil
}

func (pt *stdPattern) GetInteger(name string, n int) (int, bool, error) {
	var i C.int
	fcr := FcResult(C.FcPatternGetInteger(pt.hnd, C.CString(name), C.int(n), &i))
	switch fcr {
	case FcResultMatch:
		return int(i), true, nil
	case FcResultNoMatch:
		return 0, false, nil
	default:
		return 0, false, errors.New(fcr.String())
	}
}

func (pt *stdPattern) GetDouble(name string, n int) (float64, bool, error) {
	var d C.double
	fcr := FcResult(C.FcPatternGetDouble(pt.hnd, C.CString(name), C.int(n), &d))
	switch fcr {
	case FcResultMatch:
		return float64(d), true, nil
	case FcResultNoMatch:
		return 0, false, nil
	default:
		return 0, false, errors.New(fcr.String())
	}
}

func (pt *stdPattern) GetString(name string, n int) (string, bool, error) {
	var s *C.FcChar8
	fcr := FcResult(C.FcPatternGetString(pt.hnd, C.CString(name), C.int(n), &s))
	switch fcr {
	case FcResultMatch:
		return C.GoString((*C.char)(unsafe.Pointer(s))), true, nil
	case FcResultNoMatch:
		return "", false, nil
	default:
		return "", false, errors.New(fcr.String())
	}
}

func (pt *stdPattern) GetMatrix(name string, n int) (cairo.Matrix, bool, error) {
	m := cairo.NewMatrix()
	var fcm *C.FcMatrix
	fcr := FcResult(C.FcPatternGetMatrix(pt.hnd, C.CString(name), C.int(n), &fcm))
	switch fcr {
	case FcResultMatch:
		{
			m.Init(float64(C.cgo_fontconfig_fcmatrix_get_xx(fcm)),
				float64(C.cgo_fontconfig_fcmatrix_get_yx(fcm)),
				float64(C.cgo_fontconfig_fcmatrix_get_xy(fcm)),
				float64(C.cgo_fontconfig_fcmatrix_get_yy(fcm)), 0, 0)
			return m, true, nil
		}
	case FcResultNoMatch:
		return nil, false, nil
	default:
		return nil, false, errors.New(fcr.String())
	}
}

func (pt *stdPattern) GetCharSet(name string, n int) (CharSet, bool, error) {
	var c *C.FcCharSet
	fcr := FcResult(C.FcPatternGetCharSet(pt.hnd, C.CString(name), C.int(n), &c))
	switch fcr {
	case FcResultMatch:
		{
			return blessCharSet(C.FcCharSetCopy(c)), true, nil
		}
	case FcResultNoMatch:
		return nil, false, nil
	default:
		return nil, false, errors.New(fcr.String())
	}
}

func (pt *stdPattern) GetBool(name string, n int) (bool, bool, error) {
	var b C.FcBool
	fcr := FcResult(C.FcPatternGetBool(pt.hnd, C.CString(name), C.int(n), &b))
	switch fcr {
	case FcResultMatch:
		return b > 0, true, nil
	case FcResultNoMatch:
		return false, false, nil
	default:
		return false, false, errors.New(fcr.String())
	}
}

func (pt *stdPattern) GetLangSet(name string, n int) (LangSet, bool, error) {
	var l *C.FcLangSet
	fcr := FcResult(C.FcPatternGetLangSet(pt.hnd, C.CString(name), C.int(n), &l))
	switch fcr {
	case FcResultMatch:
		{
			return blessLangSet(C.FcLangSetCopy(l)), true, nil
		}
	case FcResultNoMatch:
		return nil, false, nil
	default:
		return nil, false, errors.New(fcr.String())
	}
}

// Documented functions with no version information in the docs, totally
// missing from the headers/libs.
func (pt *stdPattern) GetRange(name string, n int) (Range, bool, error) {
	return nil, false, errors.New("stdPattern.GetRange() unimplemented")
}
func (pt *stdPattern) GetFtFace(name string, n int) (cairo.FtFace, bool, error) {
	return nil, false, errors.New("stdPattern.GetFtFace() unimplemented")
}

func (pt *stdPattern) Del(name string) {
	C.FcPatternDel(pt.hnd, C.CString(name))
}

func (pt *stdPattern) Remove(name string, n int) {
	C.FcPatternRemove(pt.hnd, C.CString(name), C.int(n))
}

func (pt *stdPattern) DefaultSubstitute() {
	C.FcDefaultSubstitute(pt.hnd)
}

func (pt *stdPattern) NameUnparse() string {
	return C.GoString((*C.char)(unsafe.Pointer(C.FcNameUnparse(pt.hnd))))
}

func (pt *stdPattern) Format(fmt string) string {
	return C.GoString((*C.char)(unsafe.Pointer(C.FcPatternFormat(pt.hnd, (*C.FcChar8)(unsafe.Pointer(C.CString(fmt)))))))
}

//pattern = FcPatternBuild (0, FC_FAMILY, FcTypeString, "Times", (char *) 0);
func FcPatternBuild(props ...interface{}) Pattern {
	pattern := FcPatternCreate()
	if len(props)%3 != 0 {
		panic("FcPatternBuild() arguments must be grouped in triples")
	}
	i := 0
	c := 1
	var ok bool
	var propName string
	var propType FcType
	var x interface{}
	for i < len(props) {
		if propName, ok = props[i].(string); !ok {
			panic("first member of FcPatternBuild() argument triple must be a string")
		}
		if propType, ok = props[i+1].(FcType); !ok {
			panic("first member of FcPatternBuild() argument triple must be an FcType")
		}
		x = props[i+2]
		if fcTypeForInterface(x) != propType {
			panic("FcPatternBuild() argument type mismatched (arg=" + strconv.Itoa(c) + ")")
		}
		pattern.Add(propName, x, false)
		i += 3
		c += 1
	}
	return pattern
}
