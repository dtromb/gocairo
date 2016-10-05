package fontconfig

import (
	"errors"
	cairo "github.com/dtromb/gocairo"
	"reflect"
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
		typedef int FcBool;
		typedef uint32_t FcChar32;
		typedef void *FcCharSet;
		#define FC_CHARSET_MAP_SIZE		1
		#define FC_CHARSET_DONE		1
		FcCharSet * FcCharSetCreate(void) { return NULL; }
		void FcCharSetDestroy(FcCharSet *fcs) {}
		FcBool FcCharSetAddChar(FcCharSet *fcs, FcChar32 ucs4) { return 0; }
		FcBool FcCharSetDelChar(FcCharSet *fcs, FcChar32 ucs4) { return 0; }
		FcCharSet * FcCharSetCopy(FcCharSet *src) { return NULL; }
		FcBool FcCharSetEqual(const FcCharSet *a, const FcCharSet *b) { return 0; }
		FcCharSet * FcCharSetIntersect(const FcCharSet *a, const FcCharSet *b) { return NULL; }
		FcCharSet * FcCharSetUnion(const FcCharSet *a, const FcCharSet *b) { return NULL; }
		FcCharSet * FcCharSetSubtract(const FcCharSet *a, const FcCharSet *b) { return NULL; }
		FcBool FcCharSetMerge(FcCharSet *a, const FcCharSet *b, FcBool *changed) { return 0; }
		FcBool FcCharSetHasChar(const FcCharSet *fcs, FcChar32 ucs4) { return 0; }
		FcChar32 FcCharSetCount(const FcCharSet *a) { return 0; }
		FcChar32 FcCharSetIntersectCount(const FcCharSet *a, const FcCharSet *b) { return 0; }
		FcChar32 FcCharSetSubtractCount(const FcCharSet *a, const FcCharSet *b) { return 0; }
		FcBool FcCharSetIsSubset(const FcCharSet *a, const FcCharSet *b) { return 0; }
		FcChar32 FcCharSetFirstPage(const FcCharSet *a, FcChar32 map[FC_CHARSET_MAP_SIZE] , FcChar32 *next) { return 0; }
		FcChar32 FcCharSetNextPage(const FcCharSet *a, FcChar32 map[FC_CHARSET_MAP_SIZE] , FcChar32 *next) { return 0; }
	#endif
*/
import "C"

type CharSet interface {
	//FcBool FcCharSetAddChar(FcCharSet *fcs, FcChar32 ucs4);
	AddChar(ucs4 rune) error
	//FcBool FcCharSetDelChar(FcCharSet *fcs, FcChar32 ucs4);
	DelChar(usc4 rune) error
	// FcCharSet * FcCharSetCopy(FcCharSet *src);
	Copy() CharSet
	// FcBool FcCharSetEqual(const FcCharSet *a, const FcCharSet *b);
	Equal(cs CharSet) bool
	// FcCharSet * FcCharSetIntersect(const FcCharSet *a, const FcCharSet *b);
	Intersect(cs CharSet) CharSet
	// FcCharSet * FcCharSetUnion(const FcCharSet *a, const FcCharSet *b);
	Union(cs CharSet) CharSet
	// FcCharSet * FcCharSetSubtract(const FcCharSet *a, const FcCharSet *b);
	Subtract(cs CharSet) CharSet
	// FcBool FcCharSetMerge(FcCharSet *a, const FcCharSet *b, FcBool *changed);
	Merge(cs CharSet) (bool, error)
	// FcBool FcCharSetHasChar(const FcCharSet *fcs, FcChar32 ucs4);
	HasChar(ucs4 rune) bool
	// FcChar32 FcCharSetCount(const FcCharSet *a);
	Count() uint32
	// FcChar32 FcCharSetIntersectCount(const FcCharSet *a, const FcCharSet *b);
	IntersectCount(cs CharSet) uint32
	// FcChar32 FcCharSetSubtractCount(const FcCharSet *a, const FcCharSet *b);
	SubtractCount(cs CharSet) uint32
	// FcBool FcCharSetIsSubset(const FcCharSet *a, const FcCharSet *b);
	IsSubset(cs CharSet) bool
	// FcChar32 FcCharSetFirstPage(const FcCharSet *a, FcChar32[FC_CHARSET_MAP_SIZE] map, FcChar32 *next);
	FirstPage() ([]rune, uint32, uint32, bool)
	// FcChar32 FcCharSetNextPage(const FcCharSet *a, FcChar32[FC_CHARSET_MAP_SIZE] map, FcChar32 *next);
	NextPage(next uint32) ([]rune, uint32, uint32, bool)
}

type stdCharSet struct {
	hnd *C.FcCharSet
}

func destroyCharSet(cs CharSet) {
	if f, ok := cs.(cairo.Finalizable); ok {
		f.Finalize(cs)
	}
	if scs, ok := cs.(*stdCharSet); ok {
		C.FcCharSetDestroy(scs.hnd)
	}
}

func blessCharSet(hnd *C.FcCharSet) CharSet {
	cs := &stdCharSet{hnd: hnd}
	runtime.SetFinalizer(cs, destroyCharSet)
	return cs
}

func CharSetCreate() CharSet {
	cset := C.FcCharSetCreate()
	return blessCharSet(cset)
}

func (cs *stdCharSet) AddChar(ucs4 rune) error {
	r := C.FcCharSetAddChar(cs.hnd, C.FcChar32(ucs4))
	if int(r) == 0 {
		return errors.New("AddChar() failed")
	}
	return nil
}

func (cs *stdCharSet) DelChar(usc4 rune) error {
	r := C.FcCharSetDelChar(cs.hnd, C.FcChar32(usc4))
	if int(r) == 0 {
		return errors.New("DelChar() failed")
	}
	return nil
}

func (cs *stdCharSet) Copy() CharSet {
	cref := C.FcCharSetCopy(cs.hnd)
	return blessCharSet(cref)
}

func (cs *stdCharSet) Equal(c CharSet) bool {
	if scs, ok := c.(*stdCharSet); ok {
		return C.FcCharSetEqual(cs.hnd, scs.hnd) > 0
	} else {
		panic("stdCharSet.Equal(cs) unsupported for non-standard character set arguments")
	}
}

// FcCharSet * FcCharSetIntersect(const FcCharSet *a, const FcCharSet *b);
func (cs *stdCharSet) Intersect(c CharSet) CharSet {
	if scs, ok := c.(*stdCharSet); ok {
		cref := C.FcCharSetIntersect(cs.hnd, scs.hnd)
		return blessCharSet(cref)
	} else {
		panic("stdCharSet.Intersect(cs) unsupported for non-standard character set arguments")
	}
}

func (cs *stdCharSet) Union(c CharSet) CharSet {
	if scs, ok := c.(*stdCharSet); ok {
		cset := C.FcCharSetUnion(cs.hnd, scs.hnd)
		return blessCharSet(cset)
	} else {
		panic("stdCharSet.Union(cs) unsupported for non-standard character set arguments")
	}
}

func (cs *stdCharSet) Subtract(c CharSet) CharSet {
	if scs, ok := c.(*stdCharSet); ok {
		cset := C.FcCharSetSubtract(cs.hnd, scs.hnd)
		return blessCharSet(cset)
	} else {
		panic("stdCharSet.Subtract(cs) unsupported for non-standard character set arguments")
	}
}

func (cs *stdCharSet) Merge(c CharSet) (bool, error) {
	if scs, ok := c.(*stdCharSet); ok {
		var changed C.FcBool
		if C.FcCharSetMerge(cs.hnd, scs.hnd, &changed) > 0 {
			return changed > 0, nil
		} else {
			return false, errors.New("merge failed")
		}
	}
	panic("stdCharSet.Merge(cs) unsupported for non-standard character set arguments")
}

func (cs *stdCharSet) HasChar(ucs4 rune) bool {
	return C.FcCharSetHasChar(cs.hnd, C.FcChar32(ucs4)) > 0
}

func (cs *stdCharSet) Count() uint32 {
	return uint32(C.FcCharSetCount(cs.hnd))
}

func (cs *stdCharSet) IntersectCount(c CharSet) uint32 {
	if scs, ok := c.(*stdCharSet); ok {
		return uint32(C.FcCharSetIntersectCount(cs.hnd, scs.hnd))
	} else {
		panic("stdCharSet.IntersectCount(cs) unsupported for non-standard character set arguments")
	}
}

func (cs *stdCharSet) SubtractCount(c CharSet) uint32 {
	if scs, ok := c.(*stdCharSet); ok {
		return uint32(C.FcCharSetSubtractCount(cs.hnd, scs.hnd))
	} else {
		panic("stdCharSet.SubtractCount(cs) unsupported for non-standard character set arguments")
	}
}

func (cs *stdCharSet) IsSubset(c CharSet) bool {
	if scs, ok := c.(*stdCharSet); ok {
		return C.FcCharSetIsSubset(cs.hnd, scs.hnd) > 0
	} else {
		panic("stdCharSet.IsSubset(cs) unsupported for non-standard character set arguments")
	}
}

func (cs *stdCharSet) FirstPage() ([]rune, uint32, uint32, bool) {
	page := make([]rune, int(C.FC_CHARSET_MAP_SIZE))
	var next C.FcChar32
	hdr := *(*reflect.SliceHeader)(unsafe.Pointer(&page))
	first := uint32(C.FcCharSetFirstPage(cs.hnd, (*C.FcChar32)(unsafe.Pointer(hdr.Data)), &next))
	if first == C.FC_CHARSET_DONE {
		return nil, 0, 0, false
	}
	return page, first, uint32(next), true
}

func (cs *stdCharSet) NextPage(pnext uint32) ([]rune, uint32, uint32, bool) {
	page := make([]rune, int(C.FC_CHARSET_MAP_SIZE))
	var next C.FcChar32
	next = C.FcChar32(pnext)
	hdr := *(*reflect.SliceHeader)(unsafe.Pointer(&page))
	first := uint32(C.FcCharSetNextPage(cs.hnd, (*C.FcChar32)(unsafe.Pointer(hdr.Data)), &next))
	if first == C.FC_CHARSET_DONE {
		return nil, 0, 0, false
	}
	return page, first, uint32(next), true
}
