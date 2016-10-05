package cairo

/*
	#cgo CFLAGS: -I/usr/include/freetype2
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
    #include <cairo/cairo-ft.h>
	#include <inttypes.h>
	#include <stdlib.h>
	#include <freetype2/ft2build.h>
	#include FT_FREETYPE_H

	// #define CAIRO_HAS_FT_FONT 1
	// #define CAIRO_HAS_FC_FONT 1
*/
import "C"

type FtFace interface {
	FontFace
	//FT_Long           num_faces;
	NumFaces() int
	// FT_Long           face_index;
	Index() int
	// FT_Long           face_flags;
	FaceFlags() FtFaceFlags
	// FT_Long           style_flags;
	StyleFlags() FtStyleFlags
	// FT_Long           num_glyphs;
	NumGlyphs() int
	// FT_String*        family_name;
	Family() string
	// FT_String*        style_name;
	Style() string
	//FT_Int            num_fixed_sizes;
	//FT_Bitmap_Size*   available_sizes;
	Sizes() []BitmapSize
	// FT_Int            num_charmaps;
	// FT_CharMap*       charmaps;
	Charmaps() []CharMap
	//unsigned int cairo_ft_font_face_get_synthesize (cairo_font_face_t *font_face);
	GetSynthesize() FtSynthesize
	//void cairo_ft_font_face_set_synthesize (cairo_font_face_t *font_face,unsigned int synth_flags);
	SetSynthesize(flags FtSynthesize)
	//void cairo_ft_font_face_unset_synthesize (cairo_font_face_t *font_face,unsigned int synth_flags);
	UnsetSynthesize(flags FtSynthesize)
}
type FtFaceFlags uint32

const (
	FtFaceFlagScalable        FtFaceFlags = C.FT_FACE_FLAG_SCALABLE
	FtFaceFlagFixedSizes                  = C.FT_FACE_FLAG_FIXED_SIZES
	FtFaceFlagFixedWidth                  = C.FT_FACE_FLAG_FIXED_WIDTH
	FtFaceFlagSfnt                        = C.FT_FACE_FLAG_SFNT
	FtFaceFlagHorizontal                  = C.FT_FACE_FLAG_HORIZONTAL
	FtFaceFlagVertical                    = C.FT_FACE_FLAG_VERTICAL
	FtFaceFlagKerning                     = C.FT_FACE_FLAG_KERNING
	FtFaceFlagFastGlyphs                  = C.FT_FACE_FLAG_FAST_GLYPHS
	FtFaceFlagMultipleMasters             = C.FT_FACE_FLAG_MULTIPLE_MASTERS
	FtFaceFlagGlyphNames                  = C.FT_FACE_FLAG_GLYPH_NAMES
	FtFaceFlagExternalStream              = C.FT_FACE_FLAG_EXTERNAL_STREAM
	FtFaceFlagFlagHinter                  = C.FT_FACE_FLAG_HINTER
	FtFaceFlagCidKeyed                    = C.FT_FACE_FLAG_CID_KEYED
	FtFaceFlagTricky                      = C.FT_FACE_FLAG_TRICKY
	FtFaceFlagColor                       = C.FT_FACE_FLAG_COLOR
)

type FtEncoding uint32

const (
	FtEncodingNone          FtEncoding = C.FT_ENCODING_NONE
	FtEncodingMsSymbol                 = C.FT_ENCODING_MS_SYMBOL
	FtEncodingUnicode                  = C.FT_ENCODING_UNICODE
	FtEncodingSjis                     = C.FT_ENCODING_SJIS
	FtEncodingGb2312                   = C.FT_ENCODING_GB2312
	FtEncodingBig5                     = C.FT_ENCODING_BIG5
	FtEncodingWansung                  = C.FT_ENCODING_WANSUNG
	FtEncodingJohab                    = C.FT_ENCODING_JOHAB
	FtEncodingMsSjis                   = C.FT_ENCODING_MS_SJIS
	FtEncodingMsGb2312                 = C.FT_ENCODING_MS_GB2312
	FtEncodingMsBig5                   = C.FT_ENCODING_MS_BIG5
	FtEncodingMsWansung                = C.FT_ENCODING_MS_WANSUNG
	FtEncodingMsJohab                  = C.FT_ENCODING_MS_JOHAB
	FtEncodingAdobeStandard            = C.FT_ENCODING_ADOBE_STANDARD
	FtEncodingAdobeExpert              = C.FT_ENCODING_ADOBE_EXPERT
	FtEncodingAdobeCustom              = C.FT_ENCODING_ADOBE_CUSTOM
	FtEncodingAdobeLatin1              = C.FT_ENCODING_ADOBE_LATIN_1
	FtEncodingOldLatin2                = C.FT_ENCODING_OLD_LATIN_2
	FtEncodingAppleRoman               = C.FT_ENCODING_APPLE_ROMAN
)

type FtLoadFlags uint32

const (
	FtLoadFlagsDefault                  FtLoadFlags = C.FT_LOAD_DEFAULT
	FtLoadFlagsNoScale                              = C.FT_LOAD_NO_SCALE
	FtLoadFlagsNoHinting                            = C.FT_LOAD_NO_HINTING
	FtLoadFlagsRender                               = C.FT_LOAD_RENDER
	FtLoadFlagsNoBitmap                             = C.FT_LOAD_NO_BITMAP
	FtLoadFlagsVerticalLayout                       = C.FT_LOAD_VERTICAL_LAYOUT
	FtLoadFlagsAutohint                             = C.FT_LOAD_FORCE_AUTOHINT
	FtLoadFlagsCropBitmap                           = C.FT_LOAD_CROP_BITMAP
	FtLoadFlagsPedantic                             = C.FT_LOAD_PEDANTIC
	FtLoadFlagsIgnoreGlobalAdvanceWidth             = C.FT_LOAD_IGNORE_GLOBAL_ADVANCE_WIDTH
	FtLoadFlagsNoRecurse                            = C.FT_LOAD_NO_RECURSE
	FtLoadFlagsIgnoreTransform                      = C.FT_LOAD_IGNORE_TRANSFORM
	FtLoadFlagsMonochrome                           = C.FT_LOAD_MONOCHROME
	FtLoadFlagsLinearDesign                         = C.FT_LOAD_LINEAR_DESIGN
	FtLoadFlagsNoAutohint                           = C.FT_LOAD_NO_AUTOHINT
	FtLoadFlagsColor                                = C.FT_LOAD_COLOR
	FtLoadFlagsComputeMetrics                       = C.FT_LOAD_COMPUTE_METRICS
)

type FtStyleFlags uint32

const (
	FtStyleFlagItalic FtStyleFlags = C.FT_STYLE_FLAG_ITALIC
	FtStyleFlagBold                = C.FT_STYLE_FLAG_BOLD
)

type FtSynthesize uint32

const (
	FtSynthesizeBold    FtSynthesize = C.CAIRO_FT_SYNTHESIZE_BOLD
	FtSynthesizeOblique              = C.CAIRO_FT_SYNTHESIZE_OBLIQUE
)

type BitmapSize interface {
	Height() uint16
	Width() uint16
	Size() int
	XPpem() int
	YPpem() int
}

type CharMap interface {
	Face() FtFace
	Encoding() FtEncoding
	PlatformId() uint16
	EncodingId() uint16
}
