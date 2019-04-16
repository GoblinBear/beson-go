package types

import (
    "encoding/binary"
    "errors"

    "github.com/GoblinBear/beson/helper"
)

const BYTE_LENGTH_512 int = 64

type UInt512 struct {
    bs []byte
}

func NewUInt512(s string, base int) *UInt512 {
    return newUInt512(s, base).(*UInt512)
}

func newUInt512(s string, base int) interface{} {
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
    
    newValue := &UInt512 {
        bs: bs,
    }
    return newValue
}

func ToUInt512(value interface{}) *UInt512 {
    return toUInt512(value).(*UInt512)
}

func toUInt512(value interface{}) interface{} {
    bs := make([]byte, 64)
    switch value.(type) {
    case uint8:
        bs[0] = value.(uint8)
    case uint16:
        binary.LittleEndian.PutUint16(bs, value.(uint16))
    case uint32:
        binary.LittleEndian.PutUint32(bs, value.(uint32))
    case uint64:
        binary.LittleEndian.PutUint64(bs, value.(uint64))
    case *UInt128:
        binary.LittleEndian.PutUint64(bs[:8], value.(*UInt128).Low())
        binary.LittleEndian.PutUint64(bs[8:16], value.(*UInt128).High())
    case *UInt256:
        bs = helper.Resize(value.(*UInt256).Get(), 64, 0)
    case *UIntVar:
        bs = helper.Resize(value.(*UIntVar).Get(), 64, 0)
    default:
        return nil
    }
    newValue := &UInt512 {
        bs: bs,
    }
    return newValue
}

func (value *UInt512) Get() []byte {
    bs := make([]byte, len(value.bs))
    copy(bs, value.bs)
    return bs
}

func (value *UInt512) Set(bs []byte) {
    copy(value.bs, bs)
}

func (value *UInt512) LShift(bits uint) *UInt512 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt512 {
        bs: newBytes,
    }
    helper.LeftShift(newValue.bs, bits, 0)
    return newValue
}

func (value *UInt512) RShift(bits uint) *UInt512 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt512 {
        bs: newBytes,
    }
    helper.RightShift(newValue.bs, bits, 0)
    return newValue
}

func (value *UInt512) Not() *UInt512 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt512 {
        bs: newBytes,
    }
    helper.Not(newValue.bs)
    return newValue
}

func (value *UInt512) Or(val *UInt512) *UInt512 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt512 {
        bs: newBytes,
    }
    helper.Or(newValue.bs, val.bs)
    return newValue
}

func (value *UInt512) And(val *UInt512) *UInt512 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt512 {
        bs: newBytes,
    }
    helper.And(newValue.bs, val.bs)
    return newValue
}

func (value *UInt512) Xor(val *UInt512) *UInt512 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt512 {
        bs: newBytes,
    }
    helper.Xor(newValue.bs, val.bs)
    return newValue
}

func (value *UInt512) Add(val *UInt512) *UInt512 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt512 {
        bs: newBytes,
    }
    helper.Add(newValue.bs, val.bs)
    return newValue
}

func (value *UInt512) Sub(val *UInt512) *UInt512 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt512 {
        bs: newBytes,
    }
    helper.Sub(newValue.bs, val.bs)
    return newValue
}

func (value *UInt512) Multiply(val *UInt512) *UInt512 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt512 {
        bs: newBytes,
    }
    helper.Multiply(newValue.bs, val.bs)
    return newValue
}

func (value *UInt512) Divide(val *UInt512) *UInt512 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt512 {
        bs: newBytes,
    }
    helper.Divide(newValue.bs, val.bs, false)
    return newValue
}

func (value *UInt512) Modulo(val *UInt512) *UInt512 {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UInt512 {
        bs: newBytes,
    }
    ans := helper.Divide(newValue.bs, val.bs, false)
    remainder := &UInt512 {
        bs: ans,
    }
    return remainder
}

func (value *UInt512) Compare(val *UInt512) int {
    return helper.Compare(value.bs, val.bs)
}

func (value *UInt512) IsZero() bool {
    return helper.IsZero(value.bs)
}

func (value *UInt512) ToString(base int) (string, error) {
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

func (value *UInt512) ToBytes() []byte {
    bs := make([]byte, len(value.bs))
    copy(bs, value.bs)

    return bs
}

func (value *UInt512) IsSigned() bool {
    return false
}

func (value *UInt512) ZERO() *UInt512 {
    bs := make([]byte, len(value.bs))
    newValue := &UInt512 {
        bs: bs,
    }
    return newValue;
}

func (value *UInt512) MAX() *UInt512 {
    bs := []byte{
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
        0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
    }
    newValue := &UInt512 {
        bs: bs,
    }
    return newValue;
}

func uintTo64Bytes(value uint64, byteNum int) []byte {
    var mask uint64 = 1 << 8 - 1
    bs := make([]byte, BYTE_LENGTH_512)
    
    for i := 0; i < byteNum; i++ {
        bs[i] = byte((value & mask) >> uint(i * 8))
        mask = mask << 8
    }
    return bs
}
