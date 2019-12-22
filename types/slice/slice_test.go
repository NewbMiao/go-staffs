package main

// go test . -bench  BenchmarkCutSlice -benchmem   -memprofile=mem.out
// go tool pprof -http=localhost:8080 mem.out

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
