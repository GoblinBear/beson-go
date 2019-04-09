package types

import (
    "errors"

    "beson/helper"
)

const BYTE_LENGTH_256 int = 32

type UInt256 struct {
    bs []byte
}

func NewUInt256(s string, base int) *UInt256 {
    return newUInt256(s, base).(*UInt256)
}

func newUInt256(s string, base int) RootType {
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
    
    newValue := &UInt256 {
        bs: bs,
    }
    return newValue
}

func ToUInt256(value interface{}) *UInt256 {
    return toUInt256(value).(*UInt256)
}

// TODO: UInt128 to UInt256
func toUInt256(value interface{}) RootType {
    var bs []byte
    switch value.(type) {
    case *UInt8:
        v := uint64(value.(*UInt8).Get())
        bs = uintToBytes(v, 1)
    case *UInt16:
        v := uint64(value.(*UInt16).Get())
        bs = uintToBytes(v, 2)
    case *UInt32:
        v := uint64(value.(*UInt32).Get())
        bs = uintToBytes(v, 4)
    case *UInt64:
        v := value.(*UInt64).Get()
        bs = uintToBytes(v, 8)
    default:
        return nil
    }
    newValue := &UInt256 {
        bs: bs,
    }
    return newValue
}

func (value *UInt256) Get() []byte {
    bs := make([]byte, len(value.bs))
    copy(bs, value.bs)
    return bs
}

func (value *UInt256) Set(bs []byte) {
    copy(value.bs, bs)
}

func (value *UInt256) LShift(bits uint) *UInt256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt256 {
        bs: newBytes,
    }
    helper.LeftShift(newValue.bs, bits, 0)
    return newValue
}

func (value *UInt256) RShift(bits uint) *UInt256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt256 {
        bs: newBytes,
    }
    helper.RightShift(newValue.bs, bits, 0)
    return newValue
}

func (value *UInt256) Not() *UInt256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt256 {
        bs: newBytes,
    }
    helper.Not(newValue.bs)
    return newValue
}

func (value *UInt256) Or(val *UInt256) *UInt256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt256 {
        bs: newBytes,
    }
    helper.Or(newValue.bs, val.bs)
    return newValue
}

func (value *UInt256) And(val *UInt256) *UInt256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt256 {
        bs: newBytes,
    }
    helper.And(newValue.bs, val.bs)
    return newValue
}

func (value *UInt256) Xor(val *UInt256) *UInt256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt256 {
        bs: newBytes,
    }
    helper.Xor(newValue.bs, val.bs)
    return newValue
}

func (value *UInt256) Add(val *UInt256) *UInt256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt256 {
        bs: newBytes,
    }
    helper.Add(newValue.bs, val.bs)
    return newValue
}

func (value *UInt256) Sub(val *UInt256) *UInt256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt256 {
        bs: newBytes,
    }
    helper.Sub(newValue.bs, val.bs)
    return newValue
}

func (value *UInt256) Multiply(val *UInt256) *UInt256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt256 {
        bs: newBytes,
    }
    helper.Multiply(newValue.bs, val.bs)
    return newValue
}

func (value *UInt256) Divide(val *UInt256) *UInt256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt256 {
        bs: newBytes,
    }
    helper.Divide(newValue.bs, val.bs, false)
    return newValue
}

func (value *UInt256) Modulo(val *UInt256) *UInt256 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt256 {
        bs: newBytes,
    }
    ans := helper.Divide(newValue.bs, val.bs, false)
    remainder := &UInt256 {
        bs: ans,
    }
    return remainder
}

func (value *UInt256) Compare(val *UInt256) int {
    return helper.Compare(value.bs, val.bs)
}

func (value *UInt256) IsZero() bool {
    return helper.IsZero(value.bs)
}

func (value *UInt256) ToString(base int) (string, error) {
    switch base {
    case 2:
        return helper.ToBinaryString(value.bs), nil
    case 10:
        return helper.ToDecimalString(value.bs, false), nil
    case 16:
        return helper.ToHexString(value.bs), nil
    default:
        return "", errors.New("Only accepts 2 and 16 representations.")
    }
    return "", nil
}

func (value *UInt256) ToBytes() []byte {
    bs := make([]byte, len(value.bs))
    copy(bs, value.bs)

    return bs
}

func (value *UInt256) IsSigned() bool {
    return false
}

func (value *UInt256) ZERO() *UInt256 {
    bs := make([]byte, len(value.bs))
    newValue := &UInt256 {
        bs: bs,
    }
    return newValue;
}

func (value *UInt256) MAX() *UInt256 {
    bs := []byte{
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
    }
    newValue := &UInt256 {
        bs: bs,
    }
    return newValue;
}

func uintToBytes(value uint64, byteNum int) []byte {
    var mask uint64 = 1 << 8 - 1
    bs := make([]byte, BYTE_LENGTH_256)
    
    for i := 0; i < byteNum; i++ {
        bs[i] = byte((value & mask) >> uint(i * 8))
        mask = mask << 8
    }
    return bs
}
