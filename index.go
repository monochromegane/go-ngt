package ngt

// #cgo darwin LDFLAGS: -lngt -lm -lstdc++
// #cgo linux LDFLAGS: -Wl,-Bstatic -lngt -Wl,-Bdynamic -lm -lstdc++
// #include "NGT/Capi.h"
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

func CreateGraphAndTree(database string, property NGTProperty) (NGTIndex, error) {
	cDatabase := C.CString(database)
	defer C.free(unsafe.Pointer(cDatabase))

	ngterr := newNGTError()
	defer ngterr.free()

	index := C.ngt_create_graph_and_tree(cDatabase, property.property, ngterr.err)
	return NGTIndex{index: index}, newErrorFrom(ngterr)
}

func OpenIndex(database string) (NGTIndex, error) {
	cDatabase := C.CString(database)
	defer C.free(unsafe.Pointer(cDatabase))

	ngterr := newNGTError()
	defer ngterr.free()

	index := C.ngt_open_index(cDatabase, ngterr.err)
	return NGTIndex{index: index}, newErrorFrom(ngterr)
}

type NGTIndex struct {
	index C.NGTIndex
}

func (i NGTIndex) getProperty() (NGTProperty, error) {
	property, err := newNGTProperty()
	if err != nil {
		return NGTProperty{}, err
	}

	ngterr := newNGTError()
	defer ngterr.free()

	C.ngt_get_property(i.index, property.property, ngterr.err)
	return property, newErrorFrom(ngterr)
}

func (i NGTIndex) SaveIndex(database string) error {
	cDatabase := C.CString(database)
	defer C.free(unsafe.Pointer(cDatabase))

	ngterr := newNGTError()
	defer ngterr.free()

	C.ngt_save_index(i.index, cDatabase, ngterr.err)
	return newErrorFrom(ngterr)
}

func (i *NGTIndex) InsertIndex(obj []float64) (uint, error) {
	cObj := (*C.double)(&obj[0])
	cObjDim := C.int(len(obj))

	ngterr := newNGTError()
	defer ngterr.free()

	cObjectID := C.ngt_insert_index(i.index, cObj, cObjDim, ngterr.err)
	return uint(cObjectID), newErrorFrom(ngterr)
}

func (i *NGTIndex) CreateIndex(thread int) error {
	cThread := C.int(thread)

	ngterr := newNGTError()
	defer ngterr.free()

	C.ngt_create_index(i.index, cThread, ngterr.err)
	return newErrorFrom(ngterr)
}

func (i *NGTIndex) RemoveIndex(id uint) error {
	cObjectID := C.ObjectID(id)

	ngterr := newNGTError()
	defer ngterr.free()

	C.ngt_remove_index(i.index, cObjectID, ngterr.err)
	return newErrorFrom(ngterr)
}

func (i *NGTIndex) GetObjectSpace() (NGTObjectSpace, error) {
	ngterr := newNGTError()
	defer ngterr.free()

	cObjectSpace := C.ngt_get_object_space(i.index, ngterr.err)
	if err := newErrorFrom(ngterr); err != nil {
		return NGTObjectSpace{}, err
	}

	property, err := i.getProperty()
	if err != nil {
		return NGTObjectSpace{}, err
	}
	defer property.Free()

	size, err := property.getDimension()
	if err != nil {
		return NGTObjectSpace{}, err
	}
	return NGTObjectSpace{space: cObjectSpace, size: size}, newErrorFrom(ngterr)
}

func (i *NGTIndex) Close() {
	C.ngt_close_index(i.index)
}

func (i NGTIndex) Search(query []float64, size int, epsilon float32) ([]NGTObjectDistance, error) {
	cQuery := (*C.double)(&query[0])
	cQueryDim := C.int(len(query))
	cSize := C.int(size)
	cEpsilon := C.float(epsilon)

	objectDistances, err := newNGTObjectDistances()
	if err != nil {
		return nil, err
	}
	defer objectDistances.free()

	ngterr := newNGTError()
	defer ngterr.free()

	C.ngt_search_index(i.index, cQuery, cQueryDim, cSize, cEpsilon, objectDistances.obj, ngterr.err)
	if err := newErrorFrom(ngterr); err != nil {
		return nil, err
	}

	return objectDistances.getResults()
}
