package ngt

// #cgo LDFLAGS: -lm -lstdc++ -lngt
// #include "NGT/Capi.h"
// #include <stdlib.h>
import "C"

func newNGTObjectDistances() (NGTObjectDistances, error) {
	ngterr := newNGTError()
	defer ngterr.free()

	obj := C.ngt_create_empty_results(ngterr.err)
	return NGTObjectDistances{obj: obj}, newErrorFrom(ngterr)
}

type NGTObjectDistances struct {
	obj C.NGTObjectDistances
}

type NGTObjectDistance struct {
	id       uint
	distance float64
}

func (o NGTObjectDistances) getResults() ([]NGTObjectDistance, error) {
	ngterr := newNGTError()
	defer ngterr.free()

	cSize := C.ngt_get_size(o.obj, ngterr.err)
	if err := newErrorFrom(ngterr); err != nil {
		return nil, err
	}

	size := int(cSize)
	results := make([]NGTObjectDistance, size)
	for i := 0; i < size; i++ {
		ngterr.clear()
		cResult := C.ngt_get_result(o.obj, C.int(i), ngterr.err)
		if err := newErrorFrom(ngterr); err != nil {
			return nil, err
		}
		results[i] = NGTObjectDistance{
			id:       uint(cResult.id),
			distance: float64(cResult.distance),
		}
	}
	return results, nil
}

func (o *NGTObjectDistances) free() {
	C.ngt_destroy_results(o.obj)
}