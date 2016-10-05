package fontconfig

import (
//cairo "github.com/dtromb/gocairo"
)

/*
	#cgo LDFLAGS: -lcairo -lfontconfig
    #include <cairo/cairo.h>
	#include <inttypes.h>
	#include <stdlib.h>

	#ifdef CAIRO_HAS_FC_FONT
		#include <fontconfig/fontconfig.h>
		int cgo_fontconfig_supported() { return 1; }
	#else
		int cgo_fontconfig_supported() { return 0; }
		typedef int FcBool;
		typedef uint32_t FcChar32;
	#endif

*/
import "C"

func assertSupport() {
	if C.cgo_fontconfig_supported() != 1 {
		unsupportedPanic()
	}
}

func unsupportedPanic() {
	panic("Fontconfig library functions not supported in this Cairo build")
}
