package types

import (
    "encoding/binary"
    "errors"
)

const UINT64_MAX uint64 = 1 << 64 - 1
const DECIMAL_STEPPER uint64 = 10000000000000000000
const DECIMAL_STEPPER_LEN int = 19

type UInt128 struct {
    high uint64
    low uint64
}

func NewUInt128(s string, base int) RootType {
    switch base {
    case 2:
        return parseBinaryToUint(s)
    case 10:
        return parseDecimalToUint(s)
    case 16:
        return parseHexToUint(s)
    default:
        return parseDecimalToUint(s)
    }

    return parseDecimalToUint(s)
}

func ToUInt128(value interface{}) RootType {
    var low uint64
    switch value.(type) {
    case uint8:
        low = uint64(value.(uint8))
    case uint16:
        low = uint64(value.(uint16))
    case uint32:
        low = uint64(value.(uint32))
    case uint64:
        low = value.(uint64)
    default:
        return nil
    }

    newValue := &UInt128 {
        high: 0,
        low: low,
    }
    return newValue
}

func (value *UInt128) Rshift(bits uint) *UInt128 {
    newValue := &UInt128 {
        high: value.high,
        low: value.low,
    }
    value.rightShiftUnsigned(newValue, bits)
    return newValue
}

func (value *UInt128) Lshift(bits uint) *UInt128 {
    newValue := &UInt128 {
        high: value.high,
        low: value.low,
    }
    value.leftShift(newValue, bits)
    return newValue
}

func (value *UInt128) Not() *UInt128 {
    newValue := &UInt128 {
        high: value.high,
        low: value.low,
    }
    value.not(newValue)
    return newValue
}

func (value *UInt128) Or(val *UInt128) *UInt128 {
    newValue := &UInt128 {
        high: value.high,
        low: value.low,
    }
    newValue.or(newValue, val)
    return newValue
}

func (value *UInt128) And(val *UInt128) *UInt128 {
    newValue := &UInt128 {
        high: value.high,
        low: value.low,
    }
    newValue.and(newValue, val)
    return newValue
}

func (value *UInt128) Xor(val *UInt128) *UInt128 {
    newValue := &UInt128 {
        high: value.high,
        low: value.low,
    }
    newValue.xor(newValue, val)
    return newValue
}

func (value *UInt128) Add(val *UInt128) *UInt128 {
    newValue := &UInt128 {
        high: value.high,
        low: value.low,
    }
    newValue.add(newValue, val)
    return newValue
}

func (value *UInt128) Sub(val *UInt128) *UInt128 {
    newValue := &UInt128 {
        high: value.high,
        low: value.low,
    }
    newValue.sub(newValue, val)
    return newValue
}

func (value *UInt128) Multiply(val *UInt128) *UInt128 {
    newValue := &UInt128 {
        high: value.high,
        low: value.low,
    }
    newValue.multiply(newValue, val)
    return newValue
}

func (value *UInt128) Divide(val *UInt128) *UInt128 {
    newValue := &UInt128 {
        high: value.high,
        low: value.low,
    }
    newValue.divide(newValue, val)
    return newValue
}

func (value *UInt128) Modulo(val *UInt128) *UInt128 {
    newValue := &UInt128 {
        high: value.high,
        low: value.low,
    }
    ans := newValue.divide(newValue, val)
    return ans
}

func (value *UInt128) Compare(val *UInt128) int {
    return value.compare(value, val)
}

func (value *UInt128) IsZero() bool {
    return value.isZero(value)
}

func (value *UInt128) ToString(base int) (string, error) {
    switch base {
    case 2:
        return value.toBinaryString(value), nil
    case 10:
        return value.toDecimalString(value), nil
    case 16:
        return value.toHexString(value), nil
    default:
        return "", errors.New("Only accepts 2 and 16 representations.")
    }
    return "", nil
}

func (value *UInt128) ToBytes() []byte {
    b := make([]byte, 16)
    binary.LittleEndian.PutUint64(b[:8], value.low)
    binary.LittleEndian.PutUint64(b[8:], value.high)

    return b
}

func (value *UInt128) IsSigned() bool {
    return false
}

func (value *UInt128) SetValue(str string, base int) {
    newValue := NewUInt128(str, base).(*UInt128)
    value.high = newValue.high
    value.low = newValue.low
}

func (value *UInt128) High() uint64 {
    return value.high
}

func (value *UInt128) SetHigh(high uint64) {
    value.high = high
}

func (value *UInt128) Low() uint64 {
    return value.low
}

func (value *UInt128) SetLow(low uint64) {
    value.low = low
}

func (value *UInt128) ZERO() *UInt128 {
    newValue := &UInt128 {
        high: 0,
        low: 0,
    }
    return newValue;
}

func (value *UInt128) MAX() *UInt128 {
    newValue := &UInt128 {
        high: 0xFFFFFFFFFFFFFFFF,
        low: 0xFFFFFFFFFFFFFFFF,
    }
    return newValue;
}
