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
		typedef void *FcStrSet;
		typedef void *FcStrList;
		typedef int FcBool;
		FcStrList * FcStrListCreate(FcStrSet *set);
		void FcStrListFirst(FcStrList *list);
		FcChar8 * FcStrListNext(FcStrList *list);
		void FcStrListDone(FcStrList *list);
		FcBool FcStrSetAdd(FcStrSet *set, const FcChar8 *s);
		FcBool FcStrSetAddFilename(FcStrSet *set, const FcChar8 *s);
		FcBool FcStrSetDel(FcStrSet *set, const FcChar8 *s);
		void FcStrSetDestroy(FcStrSet *set);
		FcBool FcStrSetMember(FcStrSet *set, const FcChar8 *s);
		FcBool FcStrSetEqual(FcStrSet *set_a, FcStrSet *set_b);
	#endif
*/
import "C"

type StrSet interface {
	Member(str string) bool
	Equal(ss StrSet) bool
	Add(str string) bool
	AddFilename(str string) bool
	Del(str string) bool
	List() StrList
}

type StrList interface {
	First()
	Next() (string, bool)
}

type stdStrSet struct {
	hnd *C.FcStrSet
}

type stdStrList struct {
	hnd *C.FcStrList
	set *stdStrSet
}

func destroyStrSet(x StrSet) {
	if f, ok := x.(cairo.Finalizable); ok {
		f.Finalize(x)
	}
	if ss, ok := x.(*stdStrSet); ok {
		C.FcStrSetDestroy(ss.hnd)
	}
}

func destroyStrList(x StrList) {
	if f, ok := x.(cairo.Finalizable); ok {
		f.Finalize(x)
	}
	if ss, ok := x.(*stdStrList); ok {
		C.FcStrListDone(ss.hnd)
		ss.hnd = nil
		ss.set = nil
	}
}

func blessStrSet(x *C.FcStrSet) StrSet {
	ss := &stdStrSet{hnd: x}
	runtime.SetFinalizer(ss, destroyStrSet)
	return ss
}

func blessStrList(x *C.FcStrList, set *stdStrSet) StrList {
	ss := &stdStrList{hnd: x, set: set}
	runtime.SetFinalizer(ss, destroyStrList)
	return ss
}

func (ss *stdStrSet) Member(str string) bool {
	return C.FcStrSetMember(ss.hnd, (*C.FcChar8)(unsafe.Pointer(C.CString(str)))) > 0
}

func (ss *stdStrSet) Equal(s StrSet) bool {
	if sd, ok := s.(*stdStrSet); ok {
		return C.FcStrSetEqual(ss.hnd, sd.hnd) > 0
	}
	panic("stdStrSet.Equal(s) unimplemented for non-standard string set arguments")
}

func (ss *stdStrSet) Add(str string) bool {
	return C.FcStrSetAdd(ss.hnd, (*C.FcChar8)(unsafe.Pointer(C.CString(str)))) > 0
}

func (ss *stdStrSet) AddFilename(str string) bool {
	return C.FcStrSetAddFilename(ss.hnd, (*C.FcChar8)(unsafe.Pointer(C.CString(str)))) > 0
}

func (ss *stdStrSet) Del(str string) bool {
	return C.FcStrSetDel(ss.hnd, (*C.FcChar8)(unsafe.Pointer(C.CString(str)))) > 0
}

func (ss *stdStrSet) List() StrList {
	clst := C.FcStrListCreate(ss.hnd)
	return blessStrList(clst, ss)
}

func (ss *stdStrList) First() {
	C.FcStrListFirst(ss.hnd)
}

func (ss *stdStrList) Next() (string, bool) {
	str := C.FcStrListNext(ss.hnd)
	if str == nil {
		return "", false
	}
	return C.GoString((*C.char)(unsafe.Pointer(str))), true
}
