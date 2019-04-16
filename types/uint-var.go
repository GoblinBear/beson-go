package types

import (
    "encoding/binary"
    "errors"

    "github.com/GoblinBear/beson/helper"
)

type UIntVar struct {
    bs []byte
}

func NewUIntVar(s string, base int, size int) *UIntVar {
    return newUIntVar(s, base, size).(*UIntVar)
}

func newUIntVar(s string, base int, size int) interface{} {
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
    
    newValue := &UIntVar {
        bs: bs,
    }
    return newValue
}

func ToUIntVar(value interface{}, size int) *UIntVar {
    return toUIntVar(value, size).(*UIntVar)
}

func toUIntVar(value interface{}, size int) interface{} {
    bs := make([]byte, size)
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
        bs = helper.Resize(value.(*UInt256).Get(), size, 0)
    case *UInt512:
        bs = helper.Resize(value.(*UInt512).Get(), size, 0)
    case *UIntVar:
        bs = helper.Resize(value.(*UIntVar).Get(), size, 0)
    default:
        return nil
    }
    newValue := &UIntVar {
        bs: bs,
    }
    return newValue
}

func (value *UIntVar) Get() []byte {
    bs := make([]byte, len(value.bs))
    copy(bs, value.bs)
    return bs
}

func (value *UIntVar) Set(bs []byte) {
    copy(value.bs, bs)
}

func (value *UIntVar) LShift(bits uint) *UIntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UIntVar {
        bs: newBytes,
    }
    helper.LeftShift(newValue.bs, bits, 0)
    return newValue
}

func (value *UIntVar) RShift(bits uint) *UIntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UIntVar {
        bs: newBytes,
    }
    helper.RightShift(newValue.bs, bits, 0)
    return newValue
}

func (value *UIntVar) Not() *UIntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UIntVar {
        bs: newBytes,
    }
    helper.Not(newValue.bs)
    return newValue
}

func (value *UIntVar) Or(val *UIntVar) *UIntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UIntVar {
        bs: newBytes,
    }
    helper.Or(newValue.bs, val.bs)
    return newValue
}

func (value *UIntVar) And(val *UIntVar) *UIntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UIntVar {
        bs: newBytes,
    }
    helper.And(newValue.bs, val.bs)
    return newValue
}

func (value *UIntVar) Xor(val *UIntVar) *UIntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UIntVar {
        bs: newBytes,
    }
    helper.Xor(newValue.bs, val.bs)
    return newValue
}

func (value *UIntVar) Add(val *UIntVar) *UIntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UIntVar {
        bs: newBytes,
    }
    helper.Add(newValue.bs, val.bs)
    return newValue
}

func (value *UIntVar) Sub(val *UIntVar) *UIntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UIntVar {
        bs: newBytes,
    }
    helper.Sub(newValue.bs, val.bs)
    return newValue
}

func (value *UIntVar) Multiply(val *UIntVar) *UIntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UIntVar {
        bs: newBytes,
    }
    helper.Multiply(newValue.bs, val.bs)
    return newValue
}

func (value *UIntVar) Divide(val *UIntVar) *UIntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UIntVar {
        bs: newBytes,
    }
    helper.Divide(newValue.bs, val.bs, false)
    return newValue
}

func (value *UIntVar) Modulo(val *UIntVar) *UIntVar {
    newBytes := make([]byte, len(value.bs))
    copy(newBytes, value.bs)
    newValue := &UIntVar {
        bs: newBytes,
    }
    ans := helper.Divide(newValue.bs, val.bs, false)
    remainder := &UIntVar {
        bs: ans,
    }
    return remainder
}

func (value *UIntVar) Compare(val *UIntVar) int {
    return helper.Compare(value.bs, val.bs)
}

func (value *UIntVar) IsZero() bool {
    return helper.IsZero(value.bs)
}

func (value *UIntVar) ToString(base int) (string, error) {
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

func (value *UIntVar) ToBytes() []byte {
    bs := make([]byte, len(value.bs))
    copy(bs, value.bs)

    return bs
}

func (value *UIntVar) IsSigned() bool {
    return false
}

func (value *UIntVar) ZERO() *UIntVar {
    bs := make([]byte, len(value.bs))
    newValue := &UIntVar {
        bs: bs,
    }
    return newValue;
}

func (value *UIntVar) MAX(size int) *UIntVar {
    bs := make([]byte, size)
    for i := 0; i < size; i++ {
        bs[i] = 0xFF
    }
    
    newValue := &UIntVar {
        bs: bs,
    }
    return newValue;
}

func uintToVarBytes(value uint64, byteNum int, size int) []byte {
    var mask uint64 = 1 << 8 - 1
    bs := make([]byte, size)
    
    for i := 0; i < byteNum; i++ {
        bs[i] = byte((value & mask) >> uint(i * 8))
        mask = mask << 8
    }
    return bs
}
