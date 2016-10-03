package cairo

/*
	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
	
*/
import "C"

type Operator uint32
const (
	OperatorClear			Operator	= C.CAIRO_OPERATOR_CLEAR
	OperatorSource						= C.CAIRO_OPERATOR_SOURCE
	OperatorOver						= C.CAIRO_OPERATOR_OVER
	OperatorIn							= C.CAIRO_OPERATOR_IN
	OperatorOut							= C.CAIRO_OPERATOR_OUT
	OperatorAtop						= C.CAIRO_OPERATOR_ATOP
	OperatorDest						= C.CAIRO_OPERATOR_DEST
	OperatorDestOver					= C.CAIRO_OPERATOR_DEST_OVER
	OperatorDestIn						= C.CAIRO_OPERATOR_DEST_IN
	OperatorDestOut						= C.CAIRO_OPERATOR_DEST_OUT
	OperatorDestAtop					= C.CAIRO_OPERATOR_DEST_ATOP
	OperatorXor							= C.CAIRO_OPERATOR_XOR
	OperatorAdd							= C.CAIRO_OPERATOR_ADD
	OperatorSaturate					= C.CAIRO_OPERATOR_SATURATE
	OperatorMultiply					= C.CAIRO_OPERATOR_MULTIPLY
	OperatorScreen						= C.CAIRO_OPERATOR_SCREEN
	OperatorOverlay						= C.CAIRO_OPERATOR_OVERLAY
	OperatorDarken						= C.CAIRO_OPERATOR_DARKEN
	OperatorLighten						= C.CAIRO_OPERATOR_LIGHTEN
	OperatorColorDodge					= C.CAIRO_OPERATOR_COLOR_DODGE
	OperatorColorBurn					= C.CAIRO_OPERATOR_COLOR_BURN
	OperatorHardLight					= C.CAIRO_OPERATOR_HARD_LIGHT
	OperatorSoftLight					= C.CAIRO_OPERATOR_SOFT_LIGHT
	OperatorDifference					= C.CAIRO_OPERATOR_DIFFERENCE
	OperatorExclusion					= C.CAIRO_OPERATOR_EXCLUSION
	OperatorHslHue						= C.CAIRO_OPERATOR_HSL_HUE
	OperatorHslSaturation				= C.CAIRO_OPERATOR_HSL_SATURATION
	OperatorHslColor					= C.CAIRO_OPERATOR_HSL_COLOR
	OperatorHslLuminosity				= C.CAIRO_OPERATOR_HSL_LUMINOSITY
)