package cairo

import (
	"io"
	"runtime"
	"unsafe"
)

/*

	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
    #include <cairo/cairo-pdf.h>
	#include <inttypes.h>
	#include <stdlib.h>

	#define PDF_VERSION_LIST_MAX	2
	int cgo_cairo_has_pdf_version(cairo_pdf_version_t v) {
		const cairo_pdf_version_t *versions;
		int i, num_versions;
		cairo_pdf_get_versions(&versions, &num_versions);
		for (i = 0; i < num_versions; i++) {
			if (versions[i] == v) return 1;
		}
		return 0;
	}

	#ifdef CAIRO_HAS_PDF_SURFACE
		int cgo_cairo_has_pdf = 1;
	#else
		int cgo_cairo_has_pdf = 0;
	#endif

	cairo_status_t cgo_cairo_pdf_surface_write_c(void *closure, const unsigned char *data, unsigned int len) {
		WriteStdPdfSurface(closure,data,len);
	}
	cairo_write_func_t cgo_cairo_pdf_surface_write = cgo_cairo_pdf_surface_write_c;

*/
import "C"

type PdfVersion uint32

const (
	PdfVersion_1_4 PdfVersion = C.CAIRO_PDF_VERSION_1_4
	PdfVersion_1_5            = C.CAIRO_PDF_VERSION_1_5
)

func PdfGetVersions() []PdfVersion {
	var versions []PdfVersion
	for _, v := range []PdfVersion{PdfVersion_1_4, PdfVersion_1_5} {
		if C.cgo_cairo_has_pdf_version(C.cairo_pdf_version_t(v)) > 0 {
			versions = append(versions, v)
		}
	}
	return versions
}

func (pv PdfVersion) String() string {
	return C.GoString(C.cairo_pdf_version_to_string(C.cairo_pdf_version_t(pv)))
}

func HasPdfSurface() bool {
	return C.cgo_cairo_has_pdf > 0
}

type PdfSurface interface {
	Surface
	RestrictToVersion(version PdfVersion)
	SetSize(widthPts, heightPts float64)
}

func blessPdfSurface(hnd *C.cairo_surface_t, addRef bool) PdfSurface {
	ss := stdPdfSurface{}
	ss.stdSurface = blessSurface(hnd, addRef).(*stdSurface)
	return &ss
}

func finalizePdfSurface(ps PdfSurface) {
	if fn, ok := ps.(Finalizable); ok {
		fn.Finalize(ps)
	}
}

func PdfSurfaceCreate(filename string, width, height float64) PdfSurface {
	csurf := C.cairo_pdf_surface_create(C.CString(filename), C.double(width), C.double(height))
	return blessPdfSurface(csurf, false)
}

func PdfSurfaceCreateForStream(writer io.Writer, width, height float64) PdfSurface {
	ps := &stdPdfSurface{
		pdfOut: writer,
	}
	csurf := C.cairo_pdf_surface_create_for_stream(C.cgo_cairo_pdf_surface_write,
		unsafe.Pointer(ps),
		C.double(width), C.double(height))
	ps.stdSurface = blessSurface(csurf, false).(*stdSurface)
	runtime.SetFinalizer(ps, finalizePdfSurface)
	return ps
}

type stdPdfSurface struct {
	*stdSurface
	pdfOut io.Writer
}

func (ps *stdPdfSurface) GetStandardSurface() *stdSurface {
	return ps.stdSurface
}

func (ps *stdPdfSurface) Finalize(x interface{}) {
	ps.pdfOut = nil
}

func (ps *stdPdfSurface) RestrictToVersion(version PdfVersion) {
	C.cairo_pdf_surface_restrict_to_version(ps.hnd, C.cairo_pdf_version_t(version))
}

func (ps *stdPdfSurface) SetSize(widthPts, heightPts float64) {
	C.cairo_pdf_surface_set_size(ps.hnd, C.double(widthPts), C.double(heightPts))
}
