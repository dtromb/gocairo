package fontconfig

// XXX - Range support is untested.   Need to compile a recent FC and do so!

// XXX - This bit of fontconfig appears incomplete and poorly documented -
//       this functionality in gocairo should be considered provisional / upcoming.

import (
	cairo "github.com/dtromb/gocairo"
	"runtime"
)

/*
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
	#include <inttypes.h>
	#include <stdlib.h>

	extern int cgo_fontconfig_supported();

	#ifdef CAIRO_HAS_FC_FONT
		#include <fontconfig/fontconfig.h>
		int cgo_fontconfig_range_supported() {
			return FC_MAJOR >= 2 &&
				   FC_MINOR >= 11 &&
				   FC_REVISION >= 91;
		}
		#ifndef FcRangeCreateDouble
			typedef char *FcRange;
			FcRange *FcRangeCreateDouble(double a, double b) { return NULL; }
			void FcRangeDestroy(FcRange *range) {}
			FcRange *FcRangeCreateInteger(int begin, int end) { return NULL; }
			FcBool FcRangeGetDouble(const FcRange *range, double *begin, double *end) { return 0; }
		#endif
	#else
		int cgo_fontconfig_range_supported() { return 0; }
		typedef char *FcRange;
		FcRange *FcRangeCreateDouble(double a, double b) { return NULL; }
		void FcRangeDestroy(FcRange *range) {}
		FcRange *FcRangeCreateInteger(int begin, int end) { return NULL; }
		FcBool FcRangeGetDouble(const FcRange *range, double *begin, double *end) { return 0; }
	#endif
*/
import "C"

func assertRangeSupport() {
	if C.cgo_fontconfig_range_supported() != 1 {
		unsupportedRangePanic()
	}
}

func unsupportedRangePanic() {
	panic("Fontconfig library range functions not supported in this Cairo build")
}

type Range interface {
	GetDouble() (float64, float64)
}

type stdRange struct {
	hnd *C.FcRange
}

func destroyRange(x Range) {
	if xf, ok := x.(cairo.Finalizable); ok {
		xf.Finalize(x)
	}
	if sr, ok := x.(*stdRange); ok {
		C.FcRangeDestroy(sr.hnd)
	}
}

func blessRange(hnd *C.FcRange) Range {
	sr := &stdRange{
		hnd: hnd,
	}
	runtime.SetFinalizer(sr, destroyRange)
	return sr
}

func RangeCreateDouble(a, b float64) Range {
	assertRangeSupport()
	crng := C.FcRangeCreateDouble(C.double(a), C.double(b))
	return blessRange(crng)
}

func RangeCreateInteger(a, b int) Range {
	assertRangeSupport()
	crng := C.FcRangeCreateInteger(C.int(a), C.int(b))
	return blessRange(crng)
}

func (sr *stdRange) GetDouble() (float64, float64) {
	var a, b C.double
	assertRangeSupport()
	C.FcRangeGetDouble(sr.hnd, &a, &b)
	return float64(a), float64(b)
}
