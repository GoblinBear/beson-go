package types

import (
    "reflect"
    "testing"
)

func TestNewUInt256(t *testing.T) {
    expect := &UInt256 {
        bs: []byte{ 57, 116, 79, 149, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    t.Run("base2", testNewUInt256Func("10010101010011110111010000111001", 2, expect))
    t.Run("base10", testNewUInt256Func("2505012281", 10, expect))
    t.Run("base16", testNewUInt256Func("0x954f7439", 16, expect))
}

func testNewUInt256Func(s string, base int, expect *UInt256) func(*testing.T) {  
    return func(t *testing.T) {
        actual := NewUInt256(s, base)
        if reflect.DeepEqual(actual.Get(), expect.Get()) {
            t.Log("NewUInt256 test passed.")
        } else {
            t.Error("NewUInt256 test failed.")
        }
    }
}

func TestToUInt256(t *testing.T) {
    u8 := NewUInt8(23)
    u16 := NewUInt16(258)
    u32 := NewUInt32(2505012281)
    u64 := NewUInt64(2505012281)

    expect8 := &UInt256 {
        bs: []byte{ 23, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    expect16 := &UInt256 {
        bs: []byte{ 2, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    expect32 := &UInt256 {
        bs: []byte{ 57, 116, 79, 149, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    expect64 := &UInt256 {
        bs: []byte{ 57, 116, 79, 149, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }

    t.Run("uint8", testToUInt256Func(u8, expect8))
    t.Run("uint16", testToUInt256Func(u16, expect16))
    t.Run("uint32", testToUInt256Func(u32, expect32))
    t.Run("uint64", testToUInt256Func(u64, expect64))
}

func testToUInt256Func(value interface{}, expect *UInt256) func(*testing.T) {  
    return func(t *testing.T) {
        actual := ToUInt256(value)
        if reflect.DeepEqual(actual.Get(), expect.Get()) {
            t.Log("ToUInt256 test passed.")
        } else {
            t.Error("ToUInt256 test failed.")
        }
    }
}

func TestLShift_UInt256(t *testing.T) {
    value := &UInt256 {
        bs: []byte{ 214, 48, 213, 181, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }

    expect := &UInt256 {
        bs: []byte{ 0, 192, 26, 166, 186, 22, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }

    t.Run("", testLShiftFunc_UInt256(value, 13, expect))
}

func testLShiftFunc_UInt256(value *UInt256, bits uint, expect *UInt256) func(*testing.T) {  
    return func(t *testing.T) {
        actual := value.LShift(bits)
        if reflect.DeepEqual(actual.Get(), expect.Get()) {
            t.Log("LShift test passed.")
        } else {
            t.Error("LShift test failed.")
        }
    }
}

func TestRShift_UInt256(t *testing.T) {
    value := &UInt256 {
        bs: []byte{ 214, 48, 213, 181, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    expect := &UInt256 {
        bs: []byte{ 169, 174, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }

    t.Run("", testRShiftFunc_UInt256(value, 13, expect))
}

func testRShiftFunc_UInt256(value *UInt256, bits uint, expect *UInt256) func(*testing.T) {  
    return func(t *testing.T) {
        actual := value.RShift(bits)
        if reflect.DeepEqual(actual.Get(), expect.Get()) {
            t.Log("RShift test passed.")
        } else {
            t.Error("RShift test failed.")
        }
    }
}

func TestNot_UInt256(t *testing.T) {
    value := &UInt256 {
        bs: []byte{ 92, 54, 161, 228, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    expect := &UInt256 {
        bs: []byte{ 163, 201, 94, 27, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255 },
    }

    t.Run("", testNotFunc_UInt256(value, expect))
}

func testNotFunc_UInt256(value *UInt256, expect *UInt256) func(*testing.T) {  
    return func(t *testing.T) {
        actual := value.Not()
        if reflect.DeepEqual(actual.Get(), expect.Get()) {
            t.Log("Not test passed.")
        } else {
            t.Error("Not test failed.")
        }
    }
}

func TestAnd_UInt256(t *testing.T) {
    a := &UInt256 {
        bs: []byte{ 92, 54, 161, 228, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    b := &UInt256 {
        bs: []byte{ 49, 11, 186, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    expect := &UInt256 {
        bs: []byte{ 16, 2, 160, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }

    t.Run("", testAndFunc_UInt256(a, b, expect))
}

func testAndFunc_UInt256(a *UInt256, b *UInt256, expect *UInt256) func(*testing.T) {  
    return func(t *testing.T) {
        actual := a.And(b)
        if reflect.DeepEqual(actual.Get(), expect.Get()) {
            t.Log("And test passed.")
        } else {
            t.Error("And test failed.")
        }
    }
}

func TestOr_UInt256(t *testing.T) {
    a := &UInt256 {
        bs: []byte{ 92, 54, 161, 228, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    b := &UInt256 {
        bs: []byte{ 49, 11, 186, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    expect := &UInt256 {
        bs: []byte{ 125, 63, 187, 228, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }

    t.Run("", testOrFunc_UInt256(a, b, expect))
}

func testOrFunc_UInt256(a *UInt256, b *UInt256, expect *UInt256) func(*testing.T) {  
    return func(t *testing.T) {
        actual := a.Or(b)
        if reflect.DeepEqual(actual.Get(), expect.Get()) {
            t.Log("Or test passed.")
        } else {
            t.Error("Or test failed.")
        }
    }
}

func TestXor_UInt256(t *testing.T) {
    a := &UInt256 {
        bs: []byte{ 92, 54, 161, 228, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    b := &UInt256 {
        bs: []byte{ 49, 11, 186, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    expect := &UInt256 {
        bs: []byte{ 109, 61, 27, 160, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }

    t.Run("", testXorFunc_UInt256(a, b, expect))
}

func testXorFunc_UInt256(a *UInt256, b *UInt256, expect *UInt256) func(*testing.T) {  
    return func(t *testing.T) {
        actual := a.Xor(b)
        if reflect.DeepEqual(actual.Get(), expect.Get()) {
            t.Log("Xor test passed.")
        } else {
            t.Error("Xor test failed.")
        }
    }
}

func TestAdd_UInt256(t *testing.T) {
    a := &UInt256 {
        bs: []byte{ 204, 19, 240, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    b := &UInt256 {
        bs: []byte{ 110, 10, 252, 24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    expect := &UInt256 {
        bs: []byte{ 58, 30, 236, 24, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }

    t.Run("", testAddFunc_UInt256(a, b, expect))
}

func testAddFunc_UInt256(a *UInt256, b *UInt256, expect *UInt256) func(*testing.T) {  
    return func(t *testing.T) {
        actual := a.Add(b)
        if reflect.DeepEqual(actual.Get(), expect.Get()) {
            t.Log("Add test passed.")
        } else {
            t.Error("Add test failed.")
        }
    }
}

func TestSub_UInt256(t *testing.T) {
    a := &UInt256 {
        bs: []byte{ 204, 19, 240, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    b := &UInt256 {
        bs: []byte{ 110, 10, 252, 24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    expect := &UInt256 {
        bs: []byte{ 94, 9, 244, 230, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }

    t.Run("", testSubFunc_UInt256(a, b, expect))
}

func testSubFunc_UInt256(a *UInt256, b *UInt256, expect *UInt256) func(*testing.T) {  
    return func(t *testing.T) {
        actual := a.Sub(b)
        if reflect.DeepEqual(actual.Get(), expect.Get()) {
            t.Log("Sub test passed.")
        } else {
            t.Error("Sub test failed.")
        }
    }
}

func TestMultiply_UInt256(t *testing.T) {
    a := &UInt256 {
        bs: []byte{ 204, 19, 46, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    b := &UInt256 {
        bs: []byte{ 117, 10, 68, 47, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    expect := &UInt256 {
        bs: []byte{ 60, 4, 5, 35, 76, 72, 29, 47, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }

    t.Run("", testMultiplyFunc_UInt256(a, b, expect))
}

func testMultiplyFunc_UInt256(a *UInt256, b *UInt256, expect *UInt256) func(*testing.T) {  
    return func(t *testing.T) {
        actual := a.Multiply(b)
        if reflect.DeepEqual(actual.Get(), expect.Get()) {
            t.Log("Multiply test passed.")
        } else {
            t.Error("Multiply test failed.")
        }
    }
}

func TestDivide_UInt256(t *testing.T) {
    a := &UInt256 {
        bs: []byte{ 60, 4, 5, 35, 76, 72, 29, 47, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    b := &UInt256 {
        bs: []byte{ 204, 19, 46, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    expect := &UInt256 {
        bs: []byte{ 117, 10, 68, 47, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }

    t.Run("", testDivideFunc_UInt256(a, b, expect))
}

func testDivideFunc_UInt256(a *UInt256, b *UInt256, expect *UInt256) func(*testing.T) {  
    return func(t *testing.T) {
        actual := a.Divide(b)
        if reflect.DeepEqual(actual.Get(), expect.Get()) {
            t.Log("Divide test passed.")
        } else {
            t.Error("Divide test failed.")
        }
    }
}

func TestModulo_UInt256(t *testing.T) {
    a := &UInt256 {
        bs: []byte{ 60, 4, 5, 35, 76, 72, 29, 47, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    b := &UInt256 {
        bs: []byte{ 204, 19, 46, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    expect := &UInt256 {
        bs: []byte{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }

    t.Run("", testModuloFunc_UInt256(a, b, expect))
}

func testModuloFunc_UInt256(a *UInt256, b *UInt256, expect *UInt256) func(*testing.T) {  
    return func(t *testing.T) {
        actual := a.Modulo(b)
        if reflect.DeepEqual(actual.Get(), expect.Get()) {
            t.Log("Modulo test passed.")
        } else {
            t.Error("Modulo test failed.")
        }
    }
}

func TestCompare_UInt256(t *testing.T) {
    a := &UInt256 {
        bs: []byte{ 60, 4, 5, 35, 76, 72, 29, 47, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    b := &UInt256 {
        bs: []byte{ 204, 19, 46, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    expect := 1

    t.Run("", testCompareFunc_UInt256(a, b, expect))
}

func testCompareFunc_UInt256(a *UInt256, b *UInt256, expect int) func(*testing.T) {  
    return func(t *testing.T) {
        actual := a.Compare(b)
        if actual == expect {
            t.Log("Compare test passed.")
        } else {
            t.Error("Compare test failed.")
        }
    }
}

func TestIsZero_UInt256(t *testing.T) {
    value := &UInt256 {
        bs: []byte{ 60, 4, 5, 35, 76, 72, 29, 47, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    expect := false

    t.Run("", testIsZeroFunc_UInt256(value, expect))
}

func testIsZeroFunc_UInt256(value *UInt256, expect bool) func(*testing.T) {  
    return func(t *testing.T) {
        actual := value.IsZero()
        if actual == expect {
            t.Log("IsZero test passed.")
        } else {
            t.Error("IsZero test failed.")
        }
    }
}

func TestToString_UInt256(t *testing.T) {
    value := &UInt256 {
        bs: []byte{ 57, 116, 79, 149, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }

    t.Run("base2", testToStringFunc_UInt256(value, 2, "0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010010101010011110111010000111001"))
    t.Run("base10", testToStringFunc_UInt256(value, 10, "2505012281"))
    t.Run("base16", testToStringFunc_UInt256(value, 16, "00000000000000000000000000000000000000000000000000000000954f7439"))
}

func testToStringFunc_UInt256(value *UInt256, base int, expect string) func(*testing.T) {  
    return func(t *testing.T) {
        actual, _ := value.ToString(base)
        if actual == expect {
            t.Log("ToString test passed.")
        } else {
            t.Error("ToString test failed.")
        }
    }
}

func TestToBytes_UInt256(t *testing.T) {
    value := &UInt256 {
        bs: []byte{ 57, 116, 79, 149, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }

    t.Run("", testToBytesFunc_UInt256(value, value.bs))
}

func testToBytesFunc_UInt256(value *UInt256, expect []byte) func(*testing.T) {  
    return func(t *testing.T) {
        actual := value.ToBytes()
        if reflect.DeepEqual(actual, expect) {
            t.Log("ToBytes test passed.")
        } else {
            t.Error("ToBytes test failed.")
        }
    }
}
