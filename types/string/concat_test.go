package main

// go test . -bench  BenchmarkLong -benchtime=3s -benchmem
import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var lBytes = make([]byte, 1<<10)
var longStr string = string(lBytes)

// The standard Go compiler makes optimizations for string concatenations by using the + operator.
// So generally, using + operator to concatenate strings is convenient and efficient
// if the number of the concatenated strings is known at compile time.
func BenchmarkLongPlus(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = "my_string" + longStr
	}
}

func BenchmarkLongSprint(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = fmt.Sprint("my_string", longStr)
	}
}
func BenchmarkLongSprintf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = fmt.Sprintf("my_string%s", longStr)
	}
}

func BenchmarkLongBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var b bytes.Buffer
		b.WriteString("my_string")
		b.WriteString(longStr)
	}
}
func BenchmarkLongBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var b strings.Builder
		b.WriteString("my_string")
		b.WriteString(longStr)
	}
}
