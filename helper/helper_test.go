package helper

import (
    "testing"
)

var b1 []byte = []byte{ 255, 154, 45, 26, 85 }
// var b2 []byte = []byte{ 1 }
var bs []byte= b1

func BenchmarkLeftShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
        LeftShift(bs, 12, 0)
    }
}

func BenchmarkRightShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
        RightShift(bs, 12, 0)
    }
}

