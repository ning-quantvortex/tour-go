package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	s = s[2:4]
	printSlice(s)
	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("len=%d cap=%d %v array ptr: %v, header: %v \n", len(s), cap(s), s, (*unsafe.Pointer)(unsafe.Pointer(&s)), hdr)
}
