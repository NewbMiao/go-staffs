package main

import (
	"log"
	"reflect"
	"unsafe"
)

type Slice []int

func logSliceHeader(A Slice, symbol string) {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&A))
	log.Printf("%s Data:%d,Len:%d,Cap:%d %v", symbol, sh.Data, sh.Len, sh.Cap, A)
}

func (A Slice) Append(value int) {
	log.Printf("\ntry append %v to A(len:%d), get A1", value, len(A))
	A1 := append(A, value)
	logSliceHeader(A, "A")
	logSliceHeader(A1, "A1")
}
func appendSliceDiff() {
	log.Println("Append slice: when cap is not enough, use rescale a new slice, otherwise use a new slice which len is different")
	mSlice20 := make(Slice, 2, 4)
	mSlice20.Append(5)

	mSlice10 := make(Slice, 2, 2)
	mSlice10.Append(5)
}
