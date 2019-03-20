package types

import (
	//"fmt"
	"strconv"
)

func (value *Int128) compare(a *Int128, b *Int128) int {
	if (a.high < b.high) {
		return -1
	} else if (a.high > b.high) {
		return 1
	} else if (a.low < b.low) {
		return -1
	} else if (a.low > b.low) {
		return 1
	} else {
		return 0
	}
}

func (value *Int128) isZero(val *Int128) bool {
	return val.high == 0 && val.low == 0
}

func (value *Int128) isNegative(val *Int128) bool {
	return (val.high & -0x8000000000000000) != 0
}


func (value *Int128) not(val *Int128) {
	val.high = (^val.high) >> 0
	val.low = (^val.high) >> 0
}

func (value *Int128) or(a *Int128, b *Int128) {
	a.high = (a.high | b.high) >> 0
	a.low = (a.low | b.low) >> 0
}

func (value *Int128) and(a *Int128, b *Int128) {
	a.high = (a.high & b.high) >> 0
	a.low = (a.low & b.low) >> 0
}

func (value *Int128) xor(a *Int128, b *Int128) {
	a.high = (a.high ^ b.high) >> 0
	a.low = (a.low ^ b.low) >> 0
}

func (value *Int128) rightShiftUnsigned(val *Int128, bits uint) {
	if (bits >= 128) {
		val.high = 0
		val.low = 0
		return
	}

	if (bits < 64) {
		mask := genMask(bits)
		shifted := (val.high & int64(mask)) >> 0
		val.high = val.high >> bits
		val.low = ((val.low >> bits) | (shifted << (64 - bits))) >> 0
		return
	}

	bits = bits - 64
	val.low = (val.high >> bits)
	val.high = 0
}

func (value *Int128) rightShiftSigned(val *Int128, bits uint) {
	if (bits >= 128) {
		neg := value.isNegative(val)
		if neg {
			val.high = 1 << 63 - 1
			val.low = 1 << 63 - 1
		} else {
			val.high = 0
			val.low = 0
		}
		
		return
	}

	if (bits < 64) {
		mask := genMask(bits)
		shifted := (val.high & int64(mask)) >> 0
		val.high = val.high >> bits
		val.low = ((val.low >> bits) | (shifted << (64 - bits))) >> 0
		return
	}

	bits = bits - 64
	val.low = (val.high >> bits)
	val.high = (val.high >> 32 >> 32)
}

func (value *Int128) leftShift(val *Int128, bits uint) {
	if (bits >= 128) {
		val.high = 0
		val.low = 0
		return
	}
	
	if ( bits < 64 ) {
		mask := (^genMask(64 - bits)) >> 0
		shifted := (val.low & int64(mask)) >> (64 - bits)
		val.low = (val.low << bits) >> 0
		val.high = (val.high << bits | shifted) >> 0
		return
	}
	
	bits = bits - 64
	val.high = (val.low << bits) >> 0
	val.low = 0;
}

func (value *Int128) add(a *Int128, b *Int128) {
	var carry int64 = 0
	low := a.low + b.low
	if (a.low > INT64_MAX - b.low) {
		carry = 1
	}

	high := a.high + b.high + carry

	a.high = high
	a.low = low
}

func (value *Int128) sub(a *Int128, b *Int128) {
	newB := Int128 {
		high: b.high,
		low: b.low,
	}
	
	value.twosComplement(&newB)
	value.add(a, &newB)
}


// TODO
func (value *Int128) multiply(a *Int128, b *Int128) {

}

// TODO
func (value *Int128) divide(a *Int128, b *Int128) *Int128 {
	return nil
}

func (value *Int128) nbits(val *Int128) int {
	bits := 0
	high := val.high
	low := val.low

	if (high == 0) {
		for (low > 0) {
			low = low >> 1
			bits++
		}
		return bits
	}

	for (high > 0) {
		high = high >> 1
		bits++
	}
	return bits + 64
}

//TODO
func (value *Int128) twosComplement(val *Int128) {
	val.high = (^val.high) >> 0
	val.low = (^val.low) >> 0

	var carry int64 = 0
	low := val.low + 1
	if (val.low > INT64_MAX - 1) {  //TODO: INT64_MAX ?
		carry = 1
	}

	high := val.high + carry

	val.high = high
	val.low = low
}

func (value *Int128) toBinaryString(val *Int128) string {
	strLow := strconv.FormatInt(val.low, 2)
	if (val.high == 0) {
		return strLow
	}

	strHigh := strconv.FormatInt(val.high, 2)
	str := strHigh + paddingZero(strLow, 64);

	return str
}

func (value *Int128) toHexString(val *Int128) string {
	strLow := strconv.FormatInt(val.low, 16)
	if (val.high == 0) {
		return strLow
	}

	strHigh := strconv.FormatInt(val.high, 16)
	str := strHigh + paddingZero(strLow, 16);

	return str
}


// TODO
func (value *Int128) toDecimalString(val *Int128) string {
	var output []string

	stepper := Int128 {
		high: 0,
		low: DECIMAL_STEPPER_INT64, // TODO: DECIMAL_STEPPER_INT64?
	}
	
	quotient := Int128 {
		high: val.high,
		low: val.low,
	}

	for (!value.isZero(&quotient)) {
		remain := value.divide(&quotient, &stepper)
		var slc []string
		slc = append(slc, strconv.FormatInt(remain.low, 10))
		output = append(slc, output...)
		
	}

	if (len(output) == 0) {
		return "0"
	} else {
		x := output[0]
		for _, comp := range output[1:] {
			x = x + paddingZero(comp, DECIMAL_STEPPER_LEN)
		}
		
		return x
	}
	return "0"
}

