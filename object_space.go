package ngt

// #cgo darwin LDFLAGS: -lngt -lm -lstdc++
// #cgo linux LDFLAGS: -Wl,-Bstatic -lngt -Wl,-Bdynamic -lm -lstdc++
// #include "NGT/Capi.h"
// #include <stdlib.h>
import "C"
import "unsafe"

type NGTObjectSpace struct {
	space C.NGTObjectSpace
	size  int32
}

func (s *NGTObjectSpace) GetObjectAsFloat(id int) ([]float32, error) {
	ngterr := newNGTError()
	defer ngterr.free()

	cObjectID := C.ObjectID(id)

	cObject := C.ngt_get_object_as_float(s.space, cObjectID, ngterr.err)
	// defer C.free(unsafe.Pointer(cObject)) // It is freed collectively on NGTIndex.Close()
	if err := newErrorFrom(ngterr); err != nil {
		return []float32{}, err
	}

	obj := (*[1 << 30]float32)(unsafe.Pointer(cObject))[:s.size]
	return obj, nil
}

func (s *NGTObjectSpace) GetObjectAsInteger(id int) ([]uint8, error) {
	ngterr := newNGTError()
	defer ngterr.free()

	cObjectID := C.ObjectID(id)

	cObject := C.ngt_get_object_as_integer(s.space, cObjectID, ngterr.err)
	defer C.free(unsafe.Pointer(cObject))
	if err := newErrorFrom(ngterr); err != nil {
		return []uint8{}, err
	}

	obj := (*[1 << 30]uint8)(unsafe.Pointer(cObject))[:s.size]
	return obj, nil
}
