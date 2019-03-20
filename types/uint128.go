package types

import (
	"bytes"
	"encoding/binary"
	"errors"
	"strconv"
)

const UINT64_MAX uint64 = 1 << 64 - 1
const DECIMAL_STEPPER uint64 = 10000000000000000000
const DECIMAL_STEPPER_LEN int = 19

type UInt128 struct {
	high uint64
	low uint64
}

// TODO: base = 2 / 10 / 16
func NewUInt128(s string) *UInt128 {
	if len(s) == 0 {
		return nil
	}

	remain := s
	newValue := UInt128 {
		high: 0,
		low: 0,
	}
	stepper := UInt128 {
		high: 0,
		low: DECIMAL_STEPPER,
	}
	pow := UInt128 {
		high: 0,
		low: 1,
	}

	cutoff := 0
	for remain != "" {
		if len(remain) < DECIMAL_STEPPER_LEN {
			cutoff = len(remain)
		} else {
			cutoff = DECIMAL_STEPPER_LEN
		}

		low, _ := strconv.ParseUint(remain[len(remain) - cutoff:], 10, 64)
		add := UInt128 {
			high: 0,
			low: low,
		}
		newValue.multiply(&add, &pow)
		newValue.add(&newValue, &add)

		remain = remain[:len(remain) - cutoff]
		newValue.multiply(&pow, &stepper)
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
