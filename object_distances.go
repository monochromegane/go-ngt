package ngt

// #cgo darwin LDFLAGS: -lngt -lm -lstdc++
// #cgo linux LDFLAGS: -Wl,-Bstatic -lngt -Wl,-Bdynamic -lm -lstdc++
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
	Id       uint
	Distance float64
}

func (o *NGTObjectDistances) getResults() ([]NGTObjectDistance, error) {
	ngterr := newNGTError()
	defer ngterr.free()

	cSize := C.ngt_get_result_size(o.obj, ngterr.err)
	if err := newErrorFrom(ngterr); err != nil {
		return nil, err
	}

	size := uint32(cSize)
	results := make([]NGTObjectDistance, size)
	for i := 0; uint32(i) < size; i++ {
		ngterr.clear()
		cResult := C.ngt_get_result(o.obj, C.uint32_t(i), ngterr.err)
		if err := newErrorFrom(ngterr); err != nil {
			return nil, err
		}
		results[i] = NGTObjectDistance{
			Id:       uint(cResult.id),
			Distance: float64(cResult.distance),
		}
	}
	return results, nil
}

func (o *NGTObjectDistances) free() {
	C.ngt_destroy_results(o.obj)
}
