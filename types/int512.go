package types

import (
    "errors"

    "beson/helper"
)

type Int512 struct {
    bs []byte
}

func NewInt512(s string, base int) *Int512 {
    return newInt512(s, base).(*Int512)
}

func newInt512(s string, base int) RootType {
    var bs []byte
    switch base {
    case 2:
        bs = helper.BinaryStringToBytes(s, BYTE_LENGTH_512)
    case 10:
        bs = helper.DecimalStringToBytes(s, BYTE_LENGTH_512)
    case 16:
        bs = helper.HexStringToBytes(s, BYTE_LENGTH_512)
    default:
        bs = helper.DecimalStringToBytes(s, BYTE_LENGTH_512)
    }
    
    newValue := &Int512 {
        bs: bs,
    }
    return newValue
}

func ToInt512(value interface{}) *Int512 {
    return toInt512(value).(*Int512)
}

// TODO: Int128/Int256 to Int512
func toInt512(value interface{}) RootType {
    var bs []byte
    switch value.(type) {
    case *Int8:
        v := int64(value.(*Int8).Get())
        bs = intTo64Bytes(v, 1)
    case *Int16:
        v := int64(value.(*Int16).Get())
        bs = intTo64Bytes(v, 2)
    case *Int32:
        v := int64(value.(*Int32).Get())
        bs = intTo64Bytes(v, 4)
    case *Int64:
        v := value.(*Int64).Get()
        bs = intTo64Bytes(v, 8)
    default:
        return nil
    }
    newValue := &Int512 {
        bs: bs,
    }
    return newValue
}

func (value *Int512) Get() []byte {
    bs := make([]byte, BYTE_LENGTH_512)
    copy(bs, value.bs)
    return bs
}

func (value *Int512) LShift(bits uint) *Int512 {
    newBytes := make([]byte, BYTE_LENGTH_512)
    copy(newBytes, value.bs)
    newValue := &Int512 {
        bs: newBytes,
    }
    helper.LeftShift(newValue.bs, bits, 0)
    return newValue
}

func (value *Int512) RShift(bits uint) *Int512 {
    newBytes := make([]byte, BYTE_LENGTH_512)
    copy(newBytes, value.bs)
    newValue := &Int512 {
        bs: newBytes,
    }

    var padding uint8 = 0
    if helper.IsNegative(newValue.bs) {
        padding = 1
    }

    helper.RightShift(newValue.bs, bits, padding)
    return newValue
}

func (value *Int512) Not() *Int512 {
    newBytes := make([]byte, BYTE_LENGTH_512)
    copy(newBytes, value.bs)
    newValue := &Int512 {
        bs: newBytes,
    }
    helper.Not(newValue.bs)
    return newValue
}

func (value *Int512) Or(val *Int512) *Int512 {
    newBytes := make([]byte, BYTE_LENGTH_512)
    copy(newBytes, value.bs)
    newValue := &Int512 {
        bs: newBytes,
    }
    helper.Or(newValue.bs, val.bs)
    return newValue
}

func (value *Int512) And(val *Int512) *Int512 {
    newBytes := make([]byte, BYTE_LENGTH_512)
    copy(newBytes, value.bs)
    newValue := &Int512 {
        bs: newBytes,
    }
    helper.And(newValue.bs, val.bs)
    return newValue
}

func (value *Int512) Xor(val *Int512) *Int512 {
    newBytes := make([]byte, BYTE_LENGTH_512)
    copy(newBytes, value.bs)
    newValue := &Int512 {
        bs: newBytes,
    }
    helper.Xor(newValue.bs, val.bs)
    return newValue
}

func (value *Int512) Add(val *Int512) *Int512 {
    newBytes := make([]byte, BYTE_LENGTH_512)
    copy(newBytes, value.bs)
    newValue := &Int512 {
        bs: newBytes,
    }
    helper.Add(newValue.bs, val.bs)
    return newValue
}

func (value *Int512) Sub(val *Int512) *Int512 {
    newBytes := make([]byte, BYTE_LENGTH_512)
    copy(newBytes, value.bs)
    newValue := &Int512 {
        bs: newBytes,
    }
    helper.Sub(newValue.bs, val.bs)
    return newValue
}

func (value *Int512) Multiply(val *Int512) *Int512 {
    newBytes := make([]byte, BYTE_LENGTH_512)
    copy(newBytes, value.bs)
    newValue := &Int512 {
        bs: newBytes,
    }
    helper.Multiply(newValue.bs, val.bs)
    return newValue
}

func (value *Int512) Divide(val *Int512) *Int512 {
    newBytes := make([]byte, BYTE_LENGTH_512)
    copy(newBytes, value.bs)
    newValue := &Int512 {
        bs: newBytes,
    }
    helper.Divide(newValue.bs, val.bs, true)
    return newValue
}

func (value *Int512) Modulo(val *Int512) *Int512 {
    newBytes := make([]byte, BYTE_LENGTH_512)
    copy(newBytes, value.bs)
    newValue := &Int512 {
        bs: newBytes,
    }
    ans := helper.Divide(newValue.bs, val.bs, true)
    remainder := &Int512 {
        bs: ans,
    }
    return remainder
}

func (value *Int512) Compare(val *Int512) int {
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

func (value *Int512) IsZero() bool {
    return helper.IsZero(value.bs)
}

func (value *Int512) ToString(base int) (string, error) {
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

func (value *Int512) ToBytes() []byte {
    bs := make([]byte, BYTE_LENGTH_512)
    copy(bs, value.bs)

    return bs
}

func (value *Int512) IsSigned() bool {
    return true
}

func (value *Int512) ZERO() *Int512 {
    bs := make([]byte, BYTE_LENGTH_512)
    newValue := &Int512 {
        bs: bs,
    }
    return newValue;
}

func (value *Int512) MAX() *Int512 {
    bs := []byte{
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F,
    }
    newValue := &Int512 {
        bs: bs,
    }
    return newValue;
}

func (value *Int512) MIN() *Int512 {
    bs := []byte{
        0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
        0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
        0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
        0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
        0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
        0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
        0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
        0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80,
    }
    newValue := &Int512 {
        bs: bs,
    }
    return newValue;
}

func intTo64Bytes(value int64, byteNum int) []byte {
    var mask int64 = 1 << 8 - 1
    bs := make([]byte, BYTE_LENGTH_512)
    
    for i := 0; i < byteNum; i++ {
        bs[i] = byte((value & mask) >> uint(i * 8))
        mask = mask << 8
    }
    return bs
}