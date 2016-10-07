package freetype

import (
	"unsafe"
	cairo "github.com/dtromb/gocairo"
	"github.com/dtromb/gocairo/fontconfig"
)

/*
	#cgo CFLAGS: -I/usr/include/freetype2
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
	#include <inttypes.h>
	#include <stdlib.h>
	#include <freetype2/ft2build.h>
	#include FT_FREETYPE_H

	// #define CAIRO_HAS_FT_FONT 1
	// #define CAIRO_HAS_FC_FONT 1
*/
import "C"

type stdFtFace struct {
	hnd C.FT_Face
}

type FtFont interface {
	cairo.FontFace
	FtFace() cairo.FtFace
}

type stdFtFont struct {
	cairo.StdFontFace
	ftFace *stdFtFace
}

func (sff *stdFtFace) Hnd() uintptr {
	return uintptr(unsafe.Pointer(sff.hnd))
}

func blessFtFace(hnd C.FT_Face, addRef bool) *stdFtFace {
	panic("blessFtFace() unimplemented")
}

// XXX -  Yucky.   Maybe someday, the Go authors will see the wisdom of relaxing
//  	  the oppressive no-package-cycles rule.  =/
func UnsafeBlessFtFace(hnd uintptr, addRef bool) cairo.FtFace {
	return blessFtFace(C.FT_Face(unsafe.Pointer(hnd)), addRef)
}
func init() {
	cairo.RegisterSubordinate("unsafe-bless-ft-face", UnsafeBlessFtFace)
}

// cairo_font_face_t *cairo_ft_font_face_create_for_ft_face (FT_Face face, int load_flags);
func FtFontCreateForFtFace(face cairo.FtFace, loadFlags cairo.FtLoadFlags) FtFont {
	panic("unimplemented")
}

// cairo_font_face_t *cairo_ft_font_face_create_for_pattern (FcPattern *pattern);
func FtFontCreateForPattern(pattern fontconfig.Pattern) FtFont {
	panic("unimplemented")
}

//void cairo_ft_font_options_substitute (const cairo_font_options_t *options, FcPattern *pattern);
func FtFontOptionsSubstitute(opts cairo.FontOptions, pattern fontconfig.Pattern) {
	panic("unimplemented")
} // XXX - Move to FontOptions.

// FT_Face cairo_ft_scaled_font_lock_face (cairo_scaled_font_t *scaled_font);
func FtScaledFontLockFace(font cairo.ScaledFont) {
	panic("unimplemented")
} // XXX - Move to ScaledFont
