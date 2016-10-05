package fontconfig

import (
	//"unsafe"
	//"reflect"
	//"errors"
	//"runtime"
	cairo "github.com/dtromb/gocairo"
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
	#endif
*/
import "C"

type Pattern interface {
	//FcPattern * FcPatternDuplicate(const FcPattern *p);
	Duplicate() Pattern
	//FcBool FcPatternEqual(const FcPattern *pa, const FcPattern *pb);
	Equal(p Pattern) bool
	//FcBool FcPatternEqualSubset(const FcPattern *pa, const FcPattern *pb, const FcObjectSet *os);
	EqualSubset(p Pattern, os ObjectSet)
	//FcPattern * FcPatternFilter(FcPattern *p, const FcObjectSet *);
	Filter(os ObjectSet) Pattern
	//FcChar32 FcPatternHash(const FcPattern *p);
	Hash() uint32
	//FcBool FcPatternAdd(FcPattern *p, const char *object, FcValue value, FcBool append);
	Add(name string, value interface{}, addAppend bool)
	//FcBool FcPatternAddWeak(FcPattern *p, const char *object, FcValue value, FcBool append);
	AddWeak(name string, value interface{}, addAppend bool)
	//FcBool FcPatternAddInteger(FcPattern *p, const char *object, int i);
	//FcBool FcPatternAddDouble(FcPattern *p, const char *object, double d);
	//FcBool FcPatternAddString(FcPattern *p, const char *object, const FcChar8 *s);
	//FcBool FcPatternAddMatrix(FcPattern *p, const char *object, const FcMatrix *m);
	//FcBool FcPatternAddCharSet(FcPattern *p, const char *object, const FcCharSet *c);
	//FcBool FcPatternAddBool(FcPattern *p, const char *object, FcBool b);
	//FcBool FcPatternAddFTFace(FcPattern *p, const char *object, const FT_Facef);
	//FcBool FcPatternAddLangSet(FcPattern *p, const char *object, const FcLangSet *l);
	//FcBool FcPatternAddRange(FcPattern *p, const char *object, const FcRange *r);
	//FcResult FcPatternGet(FcPattern *p, const char *object, int id, FcValue *v);
	Get(name string, id int) (interface{}, bool)
	//FcResult FcPatternGetInteger(FcPattern *p, const char *object, int n, int *i);
	GetInteger(name string, n int) (int, bool)
	//FcResult FcPatternGetDouble(FcPattern *p, const char *object, int n, double *d);
	GetDouble(name string, n int) (float64, bool)
	//FcResult FcPatternGetString(FcPattern *p, const char *object, int n, FcChar8 **s);
	GetString(name string, n int) (string, bool)
	//FcResult FcPatternGetMatrix(FcPattern *p, const char *object, int n, FcMatrix **s);
	GetMatrix(name string, n int) (cairo.Matrix, bool)
	//FcResult FcPatternGetCharSet(FcPattern *p, const char *object, int n, FcCharSet **c);
	GetCharSet(name string, n int) (CharSet, bool)
	//FcResult FcPatternGetBool(FcPattern *p, const char *object, int n, FcBool *b);
	GetBool(name string, n int) (bool, bool)
	//FcResult FcPatternGetFTFace(FcPattern *p, const char *object, int n, FT_Face *f);
	GetFtFace(name string, n int) (cairo.FtFace, bool)
	//FcResult FcPatternGetLangSet(FcPattern *p, const char *object, int n, FcLangSet **l);
	GetLangSet(name string, n int) (LangSet, bool)
	//FcResult FcPatternGetRange(FcPattern *p, const char *object, int n, FcRange **r);
	GetRange(name string, n int) (Range, bool)
	//FcBool FcPatternDel(FcPattern *p, const char *object);
	Del(name string)
	//FcBool FcPatternRemove(FcPattern *p, const char *object, int id);
	Remove(name string, n int)
	//void FcDefaultSubstitute(FcPattern *pattern);
	DefaultSubstitute()
	//FcChar8 * FcNameUnparse(FcPattern *pat);
	NameUnparse() string
	//FcChar8 * FcPatternFormat(FcPattern *pat, const FcChar8 *format);
	Format(fmt string)
}

type stdPattern struct {
	hnd *C.FcPattern
}

/*
FcPattern * FcNameParse(const FcChar8 *name);
FcPattern * FcPatternBuild(FcPattern *pattern, ...);
FcPattern * FcPatternVaBuild(FcPattern *pattern, va_list va);
void FcPatternVapBuild(FcPattern *result, FcPattern *pattern, va_list va);
FcPattern * FcPatternCreate(void);
void FcPatternReference(FcPattern *p);
void FcPatternDestroy(FcPattern *p);
*/
