package ngt

// #cgo LDFLAGS: -lm -lstdc++ -lngt
// #include "NGT/Capi.h"
// #include <stdlib.h>
import "C"
import "fmt"

func NewNGTProperty(dim int) (NGTProperty, error) {
	property, err := newNGTProperty()
	if err != nil {
		return property, err
	}
	property.SetDimension(dim)
	return property, err
}

func (p *NGTProperty) SetObjectType(type_ ObjectType) error {
	switch type_ {
	case ObjectTypeUint8:
		return p.setObjectTypeInteger()
	case ObjectTypeFloat:
		return p.setObjectTypeFloat()
	default:
		return fmt.Errorf("Invalid object type: %s", type_)
	}
}

func (p *NGTProperty) SetDistanceType(type_ DistanceType) error {
	switch type_ {
	case DistanceTypeL1:
		return p.setDistanceTypeL1()
	case DistanceTypeL2:
		return p.setDistanceTypeL2()
	case DistanceTypeAngle:
		return p.setDistanceTypeAngle()
	case DistanceTypeHamming:
		return p.setDistanceTypeHamming()
	default:
		return fmt.Errorf("Invalid distance type: %s", type_)
	}
}

func newNGTProperty() (NGTProperty, error) {
	ngterr := newNGTError()
	defer ngterr.free()

	cProperty := C.ngt_create_property(ngterr.err)
	return NGTProperty{property: cProperty}, newErrorFrom(ngterr)
}

type NGTProperty struct {
	property C.NGTProperty
}

func (p *NGTProperty) getDimension() (int, error) {
	ngterr := newNGTError()
	defer ngterr.free()

	cDim := C.ngt_get_property_dimension(p.property, ngterr.err)
	return int(cDim), newErrorFrom(ngterr)
}

func (p *NGTProperty) SetDimension(dim int) error {
	ngterr := newNGTError()
	defer ngterr.free()

	C.ngt_set_property_dimension(p.property, C.int(dim), ngterr.err)
	return newErrorFrom(ngterr)
}

func (p *NGTProperty) getEdgeSizeForCreation() (int, error) {
	ngterr := newNGTError()
	defer ngterr.free()

	cSize := C.ngt_get_property_edge_size_for_creation(p.property, ngterr.err)
	return int(cSize), newErrorFrom(ngterr)
}

func (p *NGTProperty) SetEdgeSizeForCreation(size int) error {
	ngterr := newNGTError()
	defer ngterr.free()

	C.ngt_set_property_edge_size_for_creation(p.property, C.int(size), ngterr.err)
	return newErrorFrom(ngterr)
}

func (p *NGTProperty) getEdgeSizeForSearch() (int, error) {
	ngterr := newNGTError()
	defer ngterr.free()

	cSize := C.ngt_get_property_edge_size_for_search(p.property, ngterr.err)
	return int(cSize), newErrorFrom(ngterr)
}

func (p *NGTProperty) SetEdgeSizeForSearch(size int) error {
	ngterr := newNGTError()
	defer ngterr.free()

	C.ngt_set_property_edge_size_for_search(p.property, C.int(size), ngterr.err)
	return newErrorFrom(ngterr)
}

func (p *NGTProperty) getObjectType() (ObjectType, error) {
	ngterr := newNGTError()
	defer ngterr.free()

	cType := C.ngt_get_property_object_type(p.property, ngterr.err)
	return ObjectType(C.int(cType)), newErrorFrom(ngterr)
}

func (p *NGTProperty) setObjectTypeFloat() error {
	ngterr := newNGTError()
	defer ngterr.free()

	C.ngt_set_property_object_type_float(p.property, ngterr.err)
	return newErrorFrom(ngterr)
}

func (p *NGTProperty) setObjectTypeInteger() error {
	ngterr := newNGTError()
	defer ngterr.free()

	C.ngt_set_property_object_type_integer(p.property, ngterr.err)
	return newErrorFrom(ngterr)
}

func (p *NGTProperty) getDistanceType() (DistanceType, error) {
	ngterr := newNGTError()
	defer ngterr.free()

	cType := C.ngt_get_property_distance_type(p.property, ngterr.err)
	return DistanceType(C.int(cType)), newErrorFrom(ngterr)
}

func (p *NGTProperty) setDistanceTypeL1() error {
	ngterr := newNGTError()
	defer ngterr.free()

	C.ngt_set_property_distance_type_l1(p.property, ngterr.err)
	return newErrorFrom(ngterr)
}

func (p *NGTProperty) setDistanceTypeL2() error {
	ngterr := newNGTError()
	defer ngterr.free()

	C.ngt_set_property_distance_type_l2(p.property, ngterr.err)
	return newErrorFrom(ngterr)
}

func (p *NGTProperty) setDistanceTypeAngle() error {
	ngterr := newNGTError()
	defer ngterr.free()

	C.ngt_set_property_distance_type_angle(p.property, ngterr.err)
	return newErrorFrom(ngterr)
}

func (p *NGTProperty) setDistanceTypeHamming() error {
	ngterr := newNGTError()
	defer ngterr.free()

	C.ngt_set_property_distance_type_hamming(p.property, ngterr.err)
	return newErrorFrom(ngterr)
}

func (p *NGTProperty) Free() {
	C.ngt_destroy_property(p.property)
}
