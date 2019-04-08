package types

import (
    "errors"

    "beson/helper"
)

type UIntVar struct {
    bs []byte
}

func NewUIntVar(s string, base int, byteNum int) *UIntVar {
    return newUIntVar(s, base, byteNum).(*UIntVar)
}

func newUIntVar(s string, base int, byteNum int) RootType {
    var bs []byte
    switch base {
    case 2:
        bs = helper.BinaryStringToBytes(s, byteNum)
    case 10:
        bs = helper.DecimalStringToBytes(s, byteNum)
    case 16:
        bs = helper.HexStringToBytes(s, byteNum)
    default:
        bs = helper.DecimalStringToBytes(s, byteNum)
    }
    
    newValue := &UIntVar {
        bs: bs,
    }
    return newValue
}

func ToUIntVar(value interface{}, newByteNum int) *UIntVar {
    return toUIntVar(value, newByteNum).(*UIntVar)
}

// TODO: UInt128 to UIntVar
func toUIntVar(value interface{}, newByteNum int) RootType {
    var bs []byte
    switch value.(type) {
    case *UInt8:
        v := uint64(value.(*UInt8).Get())
        bs = uintToVarBytes(v, 1, newByteNum)
    case *UInt16:
        v := uint64(value.(*UInt16).Get())
        bs = uintToVarBytes(v, 2, newByteNum)
    case *UInt32:
        v := uint64(value.(*UInt32).Get())
        bs = uintToVarBytes(v, 4, newByteNum)
    case *UInt64:
        v := value.(*UInt64).Get()
        bs = uintToVarBytes(v, 8, newByteNum)
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

func (value *UIntVar) MAX(byteNum int) *UIntVar {
    bs := make([]byte, byteNum)
    for i := 0; i < byteNum; i++ {
        bs[i] = 0xFF
    }
    
    newValue := &UIntVar {
        bs: bs,
    }
    return newValue;
}

func uintToVarBytes(value uint64, byteNum int, newByteNum int) []byte {
    var mask uint64 = 1 << 8 - 1
    bs := make([]byte, newByteNum)
    
    for i := 0; i < byteNum; i++ {
        bs[i] = byte((value & mask) >> uint(i * 8))
        mask = mask << 8
    }
    return bs
}
