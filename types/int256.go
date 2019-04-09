package types

import (
    "errors"

    "beson/helper"
)

type Int256 struct {
    bs []byte
}

func NewInt256(s string, base int) *Int256 {
    return newInt256(s, base).(*Int256)
}

func newInt256(s string, base int) RootType {
    var bs []byte
    switch base {
    case 2:
        bs = helper.BinaryStringToBytes(s, BYTE_LENGTH_256)
    case 10:
        bs = helper.DecimalStringToBytes(s, BYTE_LENGTH_256)
    case 16:
        bs = helper.HexStringToBytes(s, BYTE_LENGTH_256)
    default:
        bs = helper.DecimalStringToBytes(s, BYTE_LENGTH_256)
    }
    
    newValue := &Int256 {
        bs: bs,
    }
    return newValue
}

func ToInt256(value interface{}) *Int256 {
    return toInt256(value).(*Int256)
}

// TODO: Int128 to Int256
func toInt256(value interface{}) RootType {
    var bs []byte
    switch value.(type) {
    case *Int8:
        v := int64(value.(*Int8).Get())
        bs = intToBytes(v, 1)
    case *Int16:
        v := int64(value.(*Int16).Get())
        bs = intToBytes(v, 2)
    case *Int32:
        v := int64(value.(*Int32).Get())
        bs = intToBytes(v, 4)
    case *Int64:
        v := value.(*Int64).Get()
        bs = intToBytes(v, 8)
    default:
        return nil
    }
    newValue := &Int256 {
        bs: bs,
    }
    return newValue
}

func (value *Int256) Get() []byte {
    bs := make([]byte, len(value.bs))
    copy(bs, value.bs)
    return bs
}

func (value *Int256) Set(bs []byte) {
    copy(value.bs, bs)
}

func (value *Int256) LShift(bits uint) *Int256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &Int256 {
        bs: newBytes,
    }
    helper.LeftShift(newValue.bs, bits, 0)
    return newValue
}

func (value *Int256) RShift(bits uint) *Int256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &Int256 {
        bs: newBytes,
    }

    var padding uint8 = 0
    if helper.IsNegative(newValue.bs) {
        padding = 1
    }

    helper.RightShift(newValue.bs, bits, padding)
    return newValue
}

func (value *Int256) Not() *Int256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &Int256 {
        bs: newBytes,
    }
    helper.Not(newValue.bs)
    return newValue
}

func (value *Int256) Or(val *Int256) *Int256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &Int256 {
        bs: newBytes,
    }
    helper.Or(newValue.bs, val.bs)
    return newValue
}

func (value *Int256) And(val *Int256) *Int256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &Int256 {
        bs: newBytes,
    }
    helper.And(newValue.bs, val.bs)
    return newValue
}

func (value *Int256) Xor(val *Int256) *Int256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &Int256 {
        bs: newBytes,
    }
    helper.Xor(newValue.bs, val.bs)
    return newValue
}

func (value *Int256) Add(val *Int256) *Int256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &Int256 {
        bs: newBytes,
    }
    helper.Add(newValue.bs, val.bs)
    return newValue
}

func (value *Int256) Sub(val *Int256) *Int256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &Int256 {
        bs: newBytes,
    }
    helper.Sub(newValue.bs, val.bs)
    return newValue
}

func (value *Int256) Multiply(val *Int256) *Int256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &Int256 {
        bs: newBytes,
    }
    helper.Multiply(newValue.bs, val.bs)
    return newValue
}

func (value *Int256) Divide(val *Int256) *Int256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &Int256 {
        bs: newBytes,
    }
    helper.Divide(newValue.bs, val.bs, true)
    return newValue
}

func (value *Int256) Modulo(val *Int256) *Int256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &Int256 {
        bs: newBytes,
    }
    ans := helper.Divide(newValue.bs, val.bs, true)
    remainder := &Int256 {
        bs: ans,
    }
    return remainder
}

func (value *Int256) Compare(val *Int256) int {
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

func (value *Int256) IsZero() bool {
    return helper.IsZero(value.bs)
}

func (value *Int256) ToString(base int) (string, error) {
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

func (value *Int256) ToBytes() []byte {
    bs := make([]byte, len(value.bs))
    copy(bs, value.bs)

    return bs
}

func (value *Int256) IsSigned() bool {
    return true
}

func (value *Int256) ZERO() *Int256 {
    bs := make([]byte, len(value.bs))
    newValue := &Int256 {
        bs: bs,
    }
    return newValue;
}

func (value *Int256) MAX() *Int256 {
    bs := []byte{
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F,
    }
    newValue := &Int256 {
        bs: bs,
    }
    return newValue;
}

func (value *Int256) MIN() *Int256 {
    bs := []byte{
        0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
        0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
        0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
        0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80,
    }
    newValue := &Int256 {
        bs: bs,
    }
    return newValue;
}

func intToBytes(value int64, byteNum int) []byte {
    var mask int64 = 1 << 8 - 1
    bs := make([]byte, BYTE_LENGTH_256)
    
    for i := 0; i < byteNum; i++ {
        bs[i] = byte((value & mask) >> uint(i * 8))
        mask = mask << 8
    }
    return bs
}