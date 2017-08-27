package ngt

// #cgo LDFLAGS: -lm -lstdc++ -lngt
// #include "NGT/Capi.h"
// #include <stdlib.h>
import "C"
import "fmt"

func newErrorFrom(err NGTError) error {
	if str := err.Error(); str == "" {
		return nil
	} else {
		return fmt.Errorf("%s", str)
	}
}

func newNGTError() NGTError {
	return NGTError{
		err: C.ngt_create_error_object(),
	}
}

type NGTError struct {
	err C.NGTError
}

func (e *NGTError) Error() string {
	str := C.ngt_get_error_string(e.err)
	return C.GoString(str)
}

func (e *NGTError) clear() {
	C.ngt_clear_error_string(e.err)
}

func (e *NGTError) free() {
	C.ngt_destroy_error_object(e.err)
}
