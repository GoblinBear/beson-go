package types

import (
	//"fmt"
	"strconv"
)

const UINT64_MAX uint64 = 18446744073709551615
const UINT64_MAX_LENGTH int = 20

type UInt128 struct {
	high uint64
	low uint64
}

func NewUInt128(value string) *UInt128 {
	newValue := UInt128 {
		high: 0,
		low: 0,
	}

	binaryString := newValue.decimalStringToBinaryString(value)
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

func (value *UInt128) decimalStringToBinaryString(str string) string {
	var newString string
	if (str == "0") {
		newString = "0"
	} else {
		newString = ""
		for (str != "0") {
			lastChar := str[len(str) - 1]
			remainder := (int(lastChar) - int('0')) % 2
			newString = strconv.Itoa(remainder) + newString
			str = value.divideByTwo(str)
		}
	}
	return newString
}

func (value *UInt128) divideByTwo(str string) string {
    newString := ""
	newDigit := 0
	add := 0

	for _, ch := range str {
        newDigit = (int(ch) - int('0')) / 2 + add
        newString = newString + strconv.Itoa(newDigit)
		if ((int(ch) - int('0')) % 2 == 1) {
			add = 5
		} else {
			add = 0
		}
	}

    if (string(newString) != "0" && newString[0:1] == "0") {
		newString = newString[1:]
	}

	return newString
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

func (value *UInt128) rightShiftUnsigned(val *UInt128, bits uint) {
	if (bits >= 128) {
		val.high = 0
		val.low = 0
		return
	}

	if (bits < 64) {
		mask := value.genMask(bits);
		shifted := (val.high & mask) >> 0
		val.high = val.high >> bits
		val.low = ((val.low >> bits) | (shifted << (64 - bits))) >> 0
		return
	}

	bits = bits - 64;
	val.low = (val.high >> bits)
	val.high = 0;
}

func (value *UInt128) leftShift(val *UInt128, bits uint) {
	if (bits >= 128) {
		val.high = 0
		val.low = 0
		return
	}
	
	if ( bits < 64 ) {
		mask := (^value.genMask(64 - bits)) >> 0
		shifted := (val.low & mask) >> (64 - bits)
		val.low = (val.low << bits) >> 0
		val.high = (val.high << bits | shifted) >> 0
		return
	}
	
	bits = bits - 64
	val.high = (val.low << bits) >> 0
	val.low = 0;
}

func (value *UInt128) genMask(bits uint) uint64 {
	if (bits > 64) {
		bits = 64
	}
	if (bits < 0) {
		bits = 0
	}

	var val uint64 = 0
	for (bits > 0) {
		val = ((val << 1) | 1) >> 0
		bits--
	}
	return val;
}
