package types

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
)

const UINT64_MAX uint64 = 1 << 64 - 1
const UINT64_MAX_LENGTH int = 20

type UInt128 struct {
	high uint64
	low uint64
}

func NewUInt128(s string) *UInt128 {
	fmt.Println("....")
	if len(s) == 0 {
		return nil
	}

	newValue := UInt128 {
		high: 0,
		low: 0,
	}

	binaryString := decimalStringToBinaryString(s)
	length := len(binaryString)
	
	if (length > 128) {
		return nil
	}

	if (length > 64) {
		newValue.high, _ = strconv.ParseUint(binaryString[:length - 64], 2, 64)
		newValue.low, _ = strconv.ParseUint(binaryString[length - 64:], 2, 64)
	} else {
		newValue.high = 0
		newValue.low, _ = strconv.ParseUint(binaryString, 2, 64)
	}
	
	return &newValue
}

func (value *UInt128) Rshift(bits uint) *UInt128 {
	newValue := UInt128 {
		high: value.high,
		low: value.low,
	}
	value.rightShiftUnsigned(&newValue, bits)
	return &newValue
}

func (value *UInt128) Lshift(bits uint) *UInt128 {
	newValue := UInt128 {
		high: value.high,
		low: value.low,
	}
	value.leftShift(&newValue, bits)
	return &newValue
}

func (value *UInt128) Not() *UInt128 {
	newValue := UInt128 {
		high: value.high,
		low: value.low,
	}
	value.not(&newValue)
	return &newValue
}

func (value *UInt128) Or(val *UInt128) *UInt128 {
	newValue := UInt128 {
		high: value.high,
		low: value.low,
	}
	newValue.or(&newValue, val)
	return &newValue
}

func (value *UInt128) And(val *UInt128) *UInt128 {
	newValue := UInt128 {
		high: value.high,
		low: value.low,
	}
	newValue.and(&newValue, val)
	return &newValue
}

func (value *UInt128) Xor(val *UInt128) *UInt128 {
	newValue := UInt128 {
		high: value.high,
		low: value.low,
	}
	newValue.xor(&newValue, val)
	return &newValue
}

func (value *UInt128) Add(val *UInt128) *UInt128 {
	newValue := UInt128 {
		high: value.high,
		low: value.low,
	}
	newValue.add(&newValue, val)
	return &newValue
}

func (value *UInt128) Sub(val *UInt128) *UInt128 {
	newValue := UInt128 {
		high: value.high,
		low: value.low,
	}
	newValue.sub(&newValue, val)
	return &newValue
}

func (value *UInt128) Multiply(val *UInt128) *UInt128 {
	newValue := UInt128 {
		high: value.high,
		low: value.low,
	}
	newValue.multiply(&newValue, val)
	return &newValue
}

func (value *UInt128) Divide(val *UInt128) *UInt128 {
	newValue := UInt128 {
		high: value.high,
		low: value.low,
	}
	newValue.divide(&newValue, val)
	return &newValue
}

func (value *UInt128) Modulo(val *UInt128) *UInt128 {
	newValue := UInt128 {
		high: value.high,
		low: value.low,
	}
	ans := newValue.divide(&newValue, val)
	return ans
}

func (value *UInt128) Compare(val *UInt128) int {
	return compare(value, val)
}

func (value *UInt128) IsZero() bool {
	return isZero(value)
}

func (value *UInt128) ToString(base int) (string, error) {
	switch base {
	case 2:
		return value.toBinaryString(value), nil
	case 10:
		return value.ToDecimalString(value), nil // TODO
	case 16:
		return value.toHexString(value), nil
	default:
		return "", errors.New("Only accepts 2 and 16 representations.")
	}
	return "", nil
}

func (value *UInt128) ToBytes() *bytes.Buffer {
	bytesBuffer := bytes.NewBuffer(make([]byte, 0))
	binary.Write(bytesBuffer, binary.BigEndian, value)

	return bytesBuffer
}

func (value *UInt128) IsSigned() bool {
	return false
}

func (value *UInt128) SetValue(str string) {
	newValue := NewUInt128(str)
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
	newValue := UInt128 {
		high: 0,
		low: 0,
	}

	return &newValue;
}

func (value *UInt128) MAX() *UInt128 {
	newValue := UInt128 {
		high: 0xFFFFFFFFFFFFFFFF,
		low: 0xFFFFFFFFFFFFFFFF,
	}

	return &newValue;
}
