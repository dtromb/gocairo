package fontconfig

import (
	cairo "github.com/dtromb/gocairo"
	"runtime"
	"unsafe"
)

/*
	#cgo LDFLAGS: -lcairo -lfontconfig
    #include <cairo/cairo.h>
	#include <inttypes.h>
	#include <stdlib.h>

	extern int cgo_fontconfig_supported();

	#ifdef CAIRO_HAS_FC_FONT
		#include <fontconfig/fontconfig.h>
	#else
		typedef char FcChar8;
		typedef void *FcLangSet;
		typedef struct {} FcStrSet;
		typedef struct {} FcCharSet;
		typedef int FcBool;
		typedef int FcLangResult;
		typedef int FcChar32;
		#define FcLangDifferentLang 0
		#define FcLangDifferentTerritory 0
		#define FcLangEqual 0
		FcLangSet * FcLangSetCreate(void);
		void FcLangSetDestroy(FcLangSet *ls);
		FcLangSet * FcLangSetCopy(const FcLangSet *ls);
		FcBool FcLangSetAdd(FcLangSet *ls, const FcChar8 *lang);
		FcBool FcLangSetDel(FcLangSet *ls, const FcChar8 *lang);
		FcLangSet * FcLangSetUnion(const FcLangSet *ls_a, const FcLangSet *ls_b);
		FcLangSet * FcLangSetSubtract(const FcLangSet *ls_a, const FcLangSet *ls_b);
		FcLangResult FcLangSetCompare(const FcLangSet *ls_a, const FcLangSet *ls_b);
		FcBool FcLangSetContains(const FcLangSet *ls_a, const FcLangSet *ls_b);
		FcBool FcLangSetEqual(const FcLangSet *ls_a, const FcLangSet *ls_b);
		FcChar32 FcLangSetHash(const FcLangSet *ls);
		FcLangResult FcLangSetHasLang(const FcLangSet *ls, const FcChar8 *lang);
		FcStrSet * FcGetDefaultLangs(void);
		FcStrSet * FcLangSetGetLangs(const FcLangSet *ls);
		FcStrSet * FcGetLangs(void);
		FcChar8 * FcLangNormalize(const FcChar8 *lang);
		const FcCharSet * FcLangGetCharSet(const FcChar8 *lang);
	#endif
*/
import "C"

type LangResult uint32

const (
	LangResultEqual              LangResult = C.FcLangEqual
	LangResultDifferentTerritory            = C.FcLangDifferentTerritory
	LangResultDifferentLang                 = C.FcLangDifferentLang
)

type LangSet interface {
	//FcLangSet * FcLangSetCopy(const FcLangSet *ls);
	Copy() LangSet
	//FcBool FcLangSetAdd(FcLangSet *ls, const FcChar8 *lang);
	Add(lang string) bool
	//FcBool FcLangSetDel(FcLangSet *ls, const FcChar8 *lang);
	Del(lang string) bool
	//FcLangSet * FcLangSetUnion(const FcLangSet *ls_a, const FcLangSet *ls_b);
	Union(ls LangSet) LangSet
	//FcLangSet * FcLangSetSubtract(const FcLangSet *ls_a, const FcLangSet *ls_b);
	Subtract(ls LangSet) LangSet
	//FcLangResult FcLangSetCompare(const FcLangSet *ls_a, const FcLangSet *ls_b);
	Compare(ls LangSet) LangResult
	//FcBool FcLangSetContains(const FcLangSet *ls_a, const FcLangSet *ls_b);
	Contains(ls LangSet) bool
	//FcBool FcLangSetEqual(const FcLangSet *ls_a, const FcLangSet *ls_b);
	Equal(ls LangSet) bool
	//FcChar32 FcLangSetHash(const FcLangSet *ls);
	Hash() uint32
	//FcLangResult FcLangSetHasLang(const FcLangSet *ls, const FcChar8 *lang);
	HasLang(lang string) LangResult
	//FcStrSet * FcLangSetGetLangs(const FcLangSet *ls);
	GetLangs() StrSet
}

type stdLangSet struct {
	hnd *C.FcLangSet
}

func LangNormalize(lang string) string {
	return C.GoString((*C.char)(unsafe.Pointer(C.FcLangNormalize((*C.FcChar8)(unsafe.Pointer(C.CString(lang)))))))
}

func LangGetCharSet(lang string) CharSet {
	cset := C.FcLangGetCharSet((*C.FcChar8)(unsafe.Pointer(C.CString(lang))))
	return blessCharSet(cset)
}

func GetLangs() StrSet {
	cset := C.FcGetLangs()
	return blessStrSet(cset)
}

func GetDefaultLangs() StrSet {
	return blessStrSet(C.FcGetDefaultLangs())
}

func NewLangSet() LangSet {
	return blessLangSet(C.FcLangSetCreate())
}

func destroyLangSet(ls LangSet) {
	if f, ok := ls.(cairo.Finalizable); ok {
		f.Finalize(ls)
	}
	if sls, ok := ls.(*stdLangSet); ok {
		C.FcLangSetDestroy(sls.hnd)
	}
}

func blessLangSet(x *C.FcLangSet) LangSet {
	ls := &stdLangSet{hnd: x}
	runtime.SetFinalizer(ls, destroyLangSet)
	return ls
}

func (sls *stdLangSet) Copy() LangSet {
	return blessLangSet(C.FcLangSetCopy(sls.hnd))
}

func (sls *stdLangSet) Add(lang string) bool {
	return C.FcLangSetAdd(sls.hnd, (*C.FcChar8)(unsafe.Pointer(C.CString(lang)))) > 0
}

func (sls *stdLangSet) Del(lang string) bool {
	return C.FcLangSetDel(sls.hnd, (*C.FcChar8)(unsafe.Pointer(C.CString(lang)))) > 0
}

func (sls *stdLangSet) Union(ls LangSet) LangSet {
	if s, ok := ls.(*stdLangSet); ok {
		return blessLangSet(C.FcLangSetUnion(sls.hnd, s.hnd))
	} else {
		panic("stdLangSet.Union(ls) unsupported for non-standard lang set arguments")
	}
}

func (sls *stdLangSet) Subtract(ls LangSet) LangSet {
	if s, ok := ls.(*stdLangSet); ok {
		return blessLangSet(C.FcLangSetSubtract(sls.hnd, s.hnd))
	} else {
		panic("stdLangSet.Subtract(ls) unsupported for non-standard lang set arguments")
	}
}

func (sls *stdLangSet) Compare(ls LangSet) LangResult {
	if s, ok := ls.(*stdLangSet); ok {
		return LangResult(C.FcLangSetCompare(sls.hnd, s.hnd))
	} else {
		panic("stdLangSet.Compare(ls) unsupported for non-standard lang set arguments")
	}
}

func (sls *stdLangSet) Contains(ls LangSet) bool {
	if s, ok := ls.(*stdLangSet); ok {
		return C.FcLangSetContains(sls.hnd, s.hnd) > 0
	} else {
		panic("stdLangSet.Contains(ls) unsupported for non-standard lang set arguments")
	}
}

func (sls *stdLangSet) Equal(ls LangSet) bool {
	if s, ok := ls.(*stdLangSet); ok {
		return C.FcLangSetEqual(sls.hnd, s.hnd) > 0
	} else {
		panic("stdLangSet.Equal(ls) unsupported for non-standard lang set arguments")
	}
}

func (sls *stdLangSet) Hash() uint32 {
	return uint32(C.FcLangSetHash(sls.hnd))
}

func (sls *stdLangSet) HasLang(lang string) LangResult {
	return LangResult(C.FcLangSetHasLang(sls.hnd, (*C.FcChar8)(unsafe.Pointer(C.CString(lang)))))
}

func (sls *stdLangSet) GetLangs() StrSet {
	return blessStrSet(C.FcLangSetGetLangs(sls.hnd))
}
