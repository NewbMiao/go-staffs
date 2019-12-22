package main

import (
	"log"
	"reflect"
	"time"
)

func main() {
	log.SetFlags(0)
	log.Printf("Type %#v:", reflect.StringHeader{})

	// compare is about underlying bytes sequence
	longStringCmp()

	// concat string avoid copy ( + is recomend way for concat, see concat_test's benchmark)
	concatStringWithCopyUnderlyingBytes()

	// 3way for-range
	forRangeString()
}

func longStringCmp() {
	bs := make([]byte, 1<<26) //64MB
	s0 := string(bs)
	s1 := string(bs)
	s2 := s1

	// s0, s1 and s2 are three equal strings.
	// The underlying bytes of s0 is a copy of bs.
	// The underlying bytes of s1 is also a copy of bs.
	// The underlying bytes of s0 and s1 are two
	// different copies of bs.
	// s2 shares the same underlying bytes with s1.

	startTime := time.Now()
	_ = s0 == s1
	duration := time.Since(startTime)
	log.Println("duration for (s0 == s1):", duration)

	startTime = time.Now()
	_ = s1 == s2
	duration = time.Since(startTime)
	log.Println("duration for (s1 == s2):", duration)
	log.Println("1ms is 1000000ns! So please try to avoid comparing two long strings if they don't share the same underlying byte sequence.")

}
