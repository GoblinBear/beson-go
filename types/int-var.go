package types

import (
    "errors"

    "beson/helper"
)

type IntVar struct {
    bs []byte
}

func NewIntVar(s string, base int, size int) *IntVar {
    return newIntVar(s, base, size).(*IntVar)
}

func newIntVar(s string, base int, size int) RootType {
    var bs []byte
    switch base {
    case 2:
        bs = helper.BinaryStringToBytes(s, size)
    case 10:
        bs = helper.DecimalStringToBytes(s, size)
    case 16:
        bs = helper.HexStringToBytes(s, size)
    default:
        bs = helper.DecimalStringToBytes(s, size)
    }
    
    newValue := &IntVar {
        bs: bs,
    }
    return newValue
}

func ToIntVar(value interface{}, size int) *IntVar {
    return toIntVar(value, size).(*IntVar)
}

// TODO: Int128 to IntVar
func toIntVar(value interface{}, size int) RootType {
    var bs []byte
    switch value.(type) {
    case *Int8:
        v := int64(value.(*Int8).Get())
        bs = intToVarBytes(v, 1, size)
    case *Int16:
        v := int64(value.(*Int16).Get())
        bs = intToVarBytes(v, 2, size)
    case *Int32:
        v := int64(value.(*Int32).Get())
        bs = intToVarBytes(v, 4, size)
    case *Int64:
        v := value.(*Int64).Get()
        bs = intToVarBytes(v, 8, size)
    default:
        return nil
    }
    newValue := &IntVar {
        bs: bs,
    }
    return newValue
}

func (value *IntVar) Get() []byte {
    bs := make([]byte, len(value.bs))
    copy(bs, value.bs)
    return bs
}

func (value *IntVar) LShift(bits uint) *IntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &IntVar {
        bs: newBytes,
    }
    helper.LeftShift(newValue.bs, bits, 0)
    return newValue
}

func (value *IntVar) RShift(bits uint) *IntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &IntVar {
        bs: newBytes,
    }

    var padding uint8 = 0
    if helper.IsNegative(newValue.bs) {
        padding = 1
    }

    helper.RightShift(newValue.bs, bits, padding)
    return newValue
}

func (value *IntVar) Not() *IntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &IntVar {
        bs: newBytes,
    }
    helper.Not(newValue.bs)
    return newValue
}

func (value *IntVar) Or(val *IntVar) *IntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &IntVar {
        bs: newBytes,
    }
    helper.Or(newValue.bs, val.bs)
    return newValue
}

func (value *IntVar) And(val *IntVar) *IntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &IntVar {
        bs: newBytes,
    }
    helper.And(newValue.bs, val.bs)
    return newValue
}

func (value *IntVar) Xor(val *IntVar) *IntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &IntVar {
        bs: newBytes,
    }
    helper.Xor(newValue.bs, val.bs)
    return newValue
}

func (value *IntVar) Add(val *IntVar) *IntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &IntVar {
        bs: newBytes,
    }
    helper.Add(newValue.bs, val.bs)
    return newValue
}

func (value *IntVar) Sub(val *IntVar) *IntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &IntVar {
        bs: newBytes,
    }
    helper.Sub(newValue.bs, val.bs)
    return newValue
}

func (value *IntVar) Multiply(val *IntVar) *IntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &IntVar {
        bs: newBytes,
    }
    helper.Multiply(newValue.bs, val.bs)
    return newValue
}

func (value *IntVar) Divide(val *IntVar) *IntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &IntVar {
        bs: newBytes,
    }
    helper.Divide(newValue.bs, val.bs, true)
    return newValue
}

func (value *IntVar) Modulo(val *IntVar) *IntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &IntVar {
        bs: newBytes,
    }
    ans := helper.Divide(newValue.bs, val.bs, true)
    remainder := &IntVar {
        bs: ans,
    }
    return remainder
}

func (value *IntVar) Compare(val *IntVar) int {
    negA := helper.IsNegative(value.bs)
    negB := helper.IsNegative(val.bs)

    ans := helper.Compare(value.bs, val.bs)
    if negA && negB {
        ans = ans * -1
    } else if negA && ans != 0 {
        ans = -1
    } else if negB && ans != 0 {
        ans = 1
    }

    return ans
}

func (value *IntVar) IsZero() bool {
    return helper.IsZero(value.bs)
}

func (value *IntVar) ToString(base int) (string, error) {
    switch base {
    case 2:
        return helper.ToBinaryString(value.bs), nil
    case 10:
        return helper.ToDecimalString(value.bs, true), nil
    case 16:
        return helper.ToHexString(value.bs), nil
    default:
        return "", errors.New("Only accepts 2 and 16 representations.")
    }
    return "", nil
}

func (value *IntVar) ToBytes() []byte {
    bs := make([]byte, len(value.bs))
    copy(bs, value.bs)

    return bs
}

func (value *IntVar) IsSigned() bool {
    return true
}

func (value *IntVar) ZERO() *IntVar {
    bs := make([]byte, len(value.bs))
    newValue := &IntVar {
        bs: bs,
    }
    return newValue;
}

func (value *IntVar) MAX(size int) *IntVar {
    bs := make([]byte, size)
    for i := 0; i < size - 1; i++ {
        bs[i] = 0xFF
    }
    bs[size - 1] = 0x7F

    newValue := &IntVar {
        bs: bs,
    }
    return newValue;
}

func (value *IntVar) MIN(size int) *IntVar {
    bs := make([]byte, size)
    for i := 0; i < size - 1; i++ {
        bs[i] = 0xFF
    }
    bs[size - 1] = 0x80

    newValue := &IntVar {
        bs: bs,
    }
    return newValue;
}

func intToVarBytes(value int64, byteNum int, size int) []byte {
    var mask int64 = 1 << 8 - 1
    bs := make([]byte, size)
    
    for i := 0; i < byteNum; i++ {
        bs[i] = byte((value & mask) >> uint(i * 8))
        mask = mask << 8
    }
    return bs
}