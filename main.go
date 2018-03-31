package main

//#include <string.h>
import "C"
import (
	"unsafe"
	"os"
	"image/jpeg"
)


func main() {
	// required so CGO can compile for c
}

//export writeData
func writeData(c_filename *C.char, c_data unsafe.Pointer, c_length C.int) C.int {
	filename := C.GoString(c_filename)
	data := C.GoBytes(c_data, c_length)

	f, err := os.Open(filename)
	if err != nil {
		return C.int(-1)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		return C.int(-1)
	}

	// hide data in img
	err = Hide(f, img, data, nil)
	if err != nil {
		return C.int(-1)
	}

	return 0
}

//export readData
func readData(c_filename *C.char, c_data unsafe.Pointer, c_length C.size_t) C.int {
	filename := C.GoString(c_filename)

	f, err := os.Open(filename)
	if err != nil {
		return C.int(-1)
	}
	defer f.Close()

	// reveal data
	revealed, err := Reveal(f)
	if err != nil {
		return C.int(-1)
	}

	if C.size_t(len(revealed)) < c_length {
		c_length = C.size_t(len(revealed))
	}

	C.memcpy(c_data, unsafe.Pointer(&revealed[0]), c_length)

	return 0
}
