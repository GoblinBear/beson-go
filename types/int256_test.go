package types

import (
    "reflect"
    "testing"
)

func TestNewInt256(t *testing.T) {
    expect := &Int256 {
        bs: []byte{ 199, 139, 176, 106, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255 },
    }
    t.Run("base10", testNewInt256Func("-2505012281", 10, expect))
}

func testNewInt256Func(s string, base int, expect *Int256) func(*testing.T) {  
    return func(t *testing.T) {
        actual := NewInt256(s, base)
        if reflect.DeepEqual(actual.Get(), expect.Get()) {
            t.Log("NewInt256 test passed.")
        } else {
            t.Error("NewInt256 test failed.")
        }
    }
}

func TestToInt256(t *testing.T) {
    u8 := NewInt8(23)
    u16 := NewInt16(258)
    u32 := NewInt32(258)
    u64 := NewInt64(2505012281)

    expect8 := &Int256 {
        bs: []byte{ 23, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    expect16 := &Int256 {
        bs: []byte{ 2, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    expect32 := &Int256 {
        bs: []byte{ 2, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }
    expect64 := &Int256 {
        bs: []byte{ 57, 116, 79, 149, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }

    t.Run("uint8", testToInt256Func(u8, expect8))
    t.Run("uint16", testToInt256Func(u16, expect16))
    t.Run("uint32", testToInt256Func(u32, expect32))
    t.Run("uint64", testToInt256Func(u64, expect64))
}

func testToInt256Func(value interface{}, expect *Int256) func(*testing.T) {  
    return func(t *testing.T) {
        actual := ToInt256(value)
        if reflect.DeepEqual(actual.Get(), expect.Get()) {
            t.Log("ToInt256 test passed.")
        } else {
            t.Error("ToInt256 test failed.")
        }
    }
}

func TestLShift_Int256(t *testing.T) {
    value := &Int256 {
        bs: []byte{ 214, 48, 213, 181, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }

    expect := &Int256 {
        bs: []byte{ 0, 192, 26, 166, 186, 22, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    }

    t.Run("", testLShiftFunc_Int256(value, 13, expect))
}

func testLShiftFunc_Int256(value *Int256, bits uint, expect *Int256) func(*testing.T) {  
    return func(t *testing.T) {
        actual := value.LShift(bits)
        if reflect.DeepEqual(actual.Get(), expect.Get()) {
            t.Log("LShift test passed.")
        } else {
            t.Error("LShift test failed.")
        }
    }
}

func TestRShift_Int256(t *testing.T) {
    value := &Int256 {
        bs: []byte{ 214, 48, 213, 181, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128 },
    }
    expect := &Int256 {
        bs: []byte{ 169, 174, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 252, 255 },
    }

    t.Run("", testRShiftFunc_Int256(value, 13, expect))
}

func testRShiftFunc_Int256(value *Int256, bits uint, expect *Int256) func(*testing.T) {  
    return func(t *testing.T) {
        actual := value.RShift(bits)
        if reflect.DeepEqual(actual.Get(), expect.Get()) {
            t.Log("RShift test passed.")
        } else {
            t.Error("RShift test failed.")
        }
    }
}
