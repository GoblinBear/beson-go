package helper

import (
    "fmt"
    "reflect"
    "testing"
)

func TestHexStringToBytes(t *testing.T) {
    t.Run("size4", testHexStringToBytesFunc("0xde4f9879", 4, []byte{ 121, 152, 79, 222 }))
    t.Run("size8", testHexStringToBytesFunc("0xde4f9879", 8, []byte{ 121, 152, 79, 222, 0, 0, 0, 0 }))
}

func testHexStringToBytesFunc(s string, size int, expect []byte) func(*testing.T) {  
    return func(t *testing.T) {
        actual := HexStringToBytes(s, size)
        if reflect.DeepEqual(actual, expect) {
            t.Log("HexStringToBytes test passed.")
        } else {
            t.Error("HexStringToBytes test failed.")
        }
    }
}

func TestBinaryStringToBytes(t *testing.T) {
    t.Run("size4", testBinaryStringToBytesFunc("11011110010011111001100001111001", 4, []byte{ 121, 152, 79, 222 }))
    t.Run("size8", testBinaryStringToBytesFunc("11011110010011111001100001111001", 8, []byte{ 121, 152, 79, 222, 0, 0, 0, 0 }))
}

func testBinaryStringToBytesFunc(s string, size int, expect []byte) func(*testing.T) {  
    return func(t *testing.T) {
        actual := BinaryStringToBytes(s, size)
        if reflect.DeepEqual(actual, expect) {
            t.Log("BinaryStringToBytes test passed.")
        } else {
            t.Error("BinaryStringToBytes test failed.")
        }
    }
}

func TestDecimalStringToBytes(t *testing.T) {
    t.Run("size4", testDecimalStringToBytesFunc("258487312", 4, []byte{ 16, 52, 104, 15 }))
    t.Run("size8", testDecimalStringToBytesFunc("258487312", 8, []byte{ 16, 52, 104, 15, 0, 0, 0, 0 }))
    t.Run("size4_neg", testDecimalStringToBytesFunc("-258487312", 4, []byte{ 240, 203, 151, 240 }))
    t.Run("size8_neg", testDecimalStringToBytesFunc("-258487312", 8, []byte{ 240, 203, 151, 240, 255, 255, 255, 255 }))
}

func testDecimalStringToBytesFunc(s string, size int, expect []byte) func(*testing.T) {  
    return func(t *testing.T) {
        actual := DecimalStringToBytes(s, size)
        if reflect.DeepEqual(actual, expect) {
            t.Log("DecimalStringToBytes test passed.")
        } else {
            t.Error("DecimalStringToBytes test failed.")
        }
    }
}

func TestConcat(t *testing.T) {
    b1 := []byte{ 27, 149, 207, 253 }
    b2 := []byte{ 59, 40, 16, 3 }
    b3 := []byte{ 119, 255, 37, 198 }
    expect := []byte{ 27, 149, 207, 253, 59, 40, 16, 3, 119, 255, 37, 198 }
    t.Run("3_bytes", testConcatFunc(expect, b1, b2, b3))
}

func testConcatFunc(expect []byte, segments ...[]byte) func(*testing.T) {  
    return func(t *testing.T) {
        actual := Concat(segments[0], segments[1], segments[2])
        if reflect.DeepEqual(actual, expect) {
            t.Log("Concat test passed.")
        } else {
            t.Error("Concat test failed.")
        }
    }
}

func TestLeftShift(t *testing.T) {
    t.Run("size4_pad0", testLeftShiftFunc([]byte{ 214, 48, 213, 181 }, 13, 0, []byte{ 0, 192, 26, 166 }))
    t.Run("size8_pad0", testLeftShiftFunc([]byte{ 214, 48, 213, 181, 0, 0, 0, 0 }, 13, 0, []byte{ 0, 192, 26, 166, 186, 22, 0, 0 }))
    t.Run("size4_pad1", testLeftShiftFunc([]byte{ 214, 48, 213, 181 }, 13, 1, []byte{ 255, 223, 26, 166 }))
    t.Run("size8_pad1", testLeftShiftFunc([]byte{ 214, 48, 213, 181, 0, 0, 0, 0 }, 13, 1, []byte{ 255, 223, 26, 166, 186, 22, 0, 0 }))
}

func testLeftShiftFunc(value []byte, bits uint, padding uint8, expect []byte) func(*testing.T) {  
    return func(t *testing.T) {
        LeftShift(value, bits, padding)
        if reflect.DeepEqual(value, expect) {
            t.Log("LeftShift test passed.")
        } else {
            t.Error("LeftShift test failed.")
        }
    }
}

func TestRightShift(t *testing.T) {
    t.Run("size4_pad0", testRightShiftFunc([]byte{ 214, 48, 213, 181 }, 13, 0, []byte{ 169, 174, 5, 0 }))
    t.Run("size8_pad0", testRightShiftFunc([]byte{ 214, 48, 213, 181, 0, 0, 0, 0 }, 13, 0, []byte{ 169, 174, 5, 0, 0, 0, 0, 0 }))
    t.Run("size4_pad1", testRightShiftFunc([]byte{ 214, 48, 213, 181 }, 13, 1, []byte{ 169, 174, 253, 255 }))
    t.Run("size8_pad1", testRightShiftFunc([]byte{ 214, 48, 213, 181, 0, 0, 0, 0 }, 13, 1, []byte{ 169, 174, 5, 0, 0, 0, 248, 255 }))
}

func testRightShiftFunc(value []byte, bits uint, padding uint8, expect []byte) func(*testing.T) {  
    return func(t *testing.T) {
        RightShift(value, bits, padding)
        if reflect.DeepEqual(value, expect) {
            t.Log("RightShift test passed.")
        } else {
            t.Error("RightShift test failed.")
        }
    }
}

func TestNot(t *testing.T) {
    t.Run("size4", testNotFunc([]byte{ 92, 54, 161, 228 }, []byte{ 163, 201, 94, 27 }))
}

func testNotFunc(value []byte, expect []byte) func(*testing.T) {  
    return func(t *testing.T) {
        Not(value)
        if reflect.DeepEqual(value, expect) {
            t.Log("Not test passed.")
        } else {
            t.Error("Not test failed.")
        }
    }
}

func TestAnd(t *testing.T) {
    t.Run("size4", testAndFunc([]byte{ 92, 54, 161, 228 }, []byte{ 49, 11, 186, 68 }, []byte{ 16, 2, 160, 68 }))
}

func testAndFunc(a []byte, b []byte, expect []byte) func(*testing.T) {  
    return func(t *testing.T) {
        And(a, b)
        if reflect.DeepEqual(a, expect) {
            t.Log("And test passed.")
        } else {
            t.Error("And test failed.")
        }
    }
}

func TestOr(t *testing.T) {
    t.Run("size4", testOrFunc([]byte{ 92, 54, 161, 228 }, []byte{ 49, 11, 186, 68 }, []byte{ 125, 63, 187, 228 }))
}

func testOrFunc(a []byte, b []byte, expect []byte) func(*testing.T) {  
    return func(t *testing.T) {
        Or(a, b)
        if reflect.DeepEqual(a, expect) {
            t.Log("Or test passed.")
        } else {
            t.Error("Or test failed.")
        }
    }
}

func TestXor(t *testing.T) {
    t.Run("size4", testXorFunc([]byte{ 92, 54, 161, 228 }, []byte{ 49, 11, 186, 68 }, []byte{ 109, 61, 27, 160 }))
}

func testXorFunc(a []byte, b []byte, expect []byte) func(*testing.T) {  
    return func(t *testing.T) {
        Xor(a, b)
        fmt.Println(a, expect)
        if reflect.DeepEqual(a, expect) {
            t.Log("Xor test passed.")
        } else {
            t.Error("Xor test failed.")
        }
    }
}
