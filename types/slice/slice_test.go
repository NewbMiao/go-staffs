package main

import "testing"

type bigStruct struct {
	id   int
	data [1024]int
}

const loopCnt = 1000

func init() {
	discardLog()
}
func BenchmarkCutSlicePointer(b *testing.B) {
	var a = make([]T, loopCnt)
	for i := 0; i < loopCnt; i++ {
		a[i] = &bigStruct{id: i}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cutSlicePointer(a, 1, 800)
	}
}

func BenchmarkCopySlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// println(1)
	}
}

func BenchmarkCutSlice(b *testing.B) {
	var a = make([]T, loopCnt)
	for i := 0; i < loopCnt; i++ {
		a[i] = &bigStruct{id: i}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cutSlice(a, 1, 800)

	}
}
