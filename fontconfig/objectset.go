package fontconfig

import (
	cairo "github.com/dtromb/gocairo"
	"runtime"
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
		typedef void *FcObjectSet;
		typedef int FcBool;
		FcBool FcObjectSetAdd(FcObjectSet *os, const char *object);
		FcObjectSet * FcObjectSetCreate(void);
		void FcObjectSetDestroy(FcObjectSet *os);
	#endif
*/
import "C"

type ObjectSet interface {
	//FcBool FcObjectSetAdd(FcObjectSet *os, const char *object);
	Add(object string) bool
}

type stdObjectSet struct {
	hnd *C.FcObjectSet
}

func destroyObjectSet(x ObjectSet) {
	if f, ok := x.(cairo.Finalizable); ok {
		f.Finalize(x)
	}
	if s, ok := x.(*stdObjectSet); ok {
		C.FcObjectSetDestroy(s.hnd)
	}
}

func blessObjectSet(hnd *C.FcObjectSet) ObjectSet {
	os := &stdObjectSet{hnd: hnd}
	runtime.SetFinalizer(os, destroyObjectSet)
	return os
}

func (ss *stdObjectSet) Add(object string) bool {
	return C.FcObjectSetAdd(ss.hnd, C.CString(object)) > 0
}

func NewObjectSet() ObjectSet {
	cset := C.FcObjectSetCreate()
	return blessObjectSet(cset)
}

func ObjectSetBuild(objects ...string) ObjectSet {
	os := NewObjectSet()
	for _, o := range objects {
		os.Add(o)
	}
	return os
}
