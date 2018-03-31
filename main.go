package main

import (
        "C"
        "unsafe"
)

func main() {
	// required so CGO can compile for c
}

//export writeData
func writeData(c_filename *C.char, c_data unsafe.Pointer, c_length C.int) {
        
}

//export readData
func readData(c_filename *C.char, c_data unsafe.Pointer, c_length C.int) {
        
}