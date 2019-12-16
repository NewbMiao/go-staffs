package main

import "fmt"

type T string

func main() {
	var a []T
	a = nil
	copySlice(a, "way 1 is not match")
	a = make([]T, 0)
	copySlice(a, "way 2 is not match")
	a = append(a, "test")
	copySlice(a, "all way match")

}

func copySlice(a []T, notes string) {
	fmt.Printf("Copy slice, Got a:(%p,%+v), notes:%v\n", a, a, notes)

	var b []T
	// way 1
	b = make([]T, len(a))
	copy(b, a)
	fmt.Printf("way 1: b(%p,%+v)\n", b, b)

	// way 2
	b = append([]T(nil), a...)
	fmt.Printf("way 2: b(%p,%+v)\n", b, b)

	// way 3
	b = append(a[:0:0], a...) // See https://github.com/go101/go101/wiki
	fmt.Printf("way 3: b(%p,%+v)\n", b, b)
}
