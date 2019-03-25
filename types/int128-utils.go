package types

import (
    "strconv"
)

func (value *Int128) compare(a *Int128, b *Int128) int {
    ans := &Int128 {
        high: a.high,
        low: a.low,
    }

    value.sub(ans, b)
    if value.isZero(ans) {
        return 0
    } else if value.isNegative(ans) {
        return -1
    } else {
        return 1
    }
}

func (value *Int128) isZero(val *Int128) bool {
    return val.high == 0 && val.low == 0
}

func (value *Int128) isNegative(val *Int128) bool {
    return (val.high & 0x8000000000000000) != 0
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
    if bits >= 128 {
        val.high = 0
        val.low = 0
        return
    }

    if bits < 64 {
        mask := genMask(bits)
        shifted := (val.high & mask) >> 0
        val.high = val.high >> bits
        val.low = ((val.low >> bits) | (shifted << (64 - bits))) >> 0
        return
    }

    bits = bits - 64
    val.low = (val.high >> bits)
    val.high = 0
}

func (value *Int128) rightShiftSigned(val *Int128, bits uint) {
    if bits >= 128 {
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

    if bits < 64 {
        mask := genMask(bits)
        shifted := (val.high & mask) >> 0
        val.high = val.high >> bits
        val.low = ((val.low >> bits) | (shifted << (64 - bits))) >> 0
        return
    }

    bits = bits - 64
    val.low = (val.high >> bits)
    val.high = (val.high >> 32 >> 32)
}

func (value *Int128) leftShift(val *Int128, bits uint) {
    if bits >= 128 {
        val.high = 0
        val.low = 0
        return
    }
    
    if bits < 64 {
        mask := (^genMask(64 - bits)) >> 0
        shifted := (val.low & mask) >> (64 - bits)
        val.low = (val.low << bits) >> 0
        val.high = (val.high << bits | shifted) >> 0
        return
    }
    
    bits = bits - 64
    val.high = (val.low << bits) >> 0
    val.low = 0;
}

func (value *Int128) add(a *Int128, b *Int128) {
    var carry uint64 = 0
    low := a.low + b.low
    if a.low > UINT64_MAX - b.low {
        carry = 1
    }

    high := a.high + b.high + carry

    a.high = high
    a.low = low
}

func (value *Int128) sub(a *Int128, b *Int128) {
    newB := &Int128 {
        high: b.high,
        low: b.low,
    }
    
    value.twosComplement(newB)
    value.add(a, newB)
}

func (value *Int128) multiply(a *Int128, b *Int128) {
    multiplier := &Int128 {
        high: b.high,
        low: b.low,
    }
    ans := &Int128 {
        high: 0,
        low: 0,
    }

    bits := value.nbits(b)

    for i := 0; i < bits; i++ {
        if multiplier.low & 1 == 1 {
            value.add(ans, a)
        }
        value.leftShift(a, 1)
        value.rightShiftUnsigned(multiplier, 1)
    }
    
    a.high = ans.high
    a.low = ans.low
}

func (value *Int128) divide(a *Int128, b *Int128) *Int128 {
    quotient := &Int128 {
        high: 0,
        low: 0,
    }
    remainder := &Int128 {
        high: a.high,
        low: a.low,
    }
    divider := &Int128 {
        high: b.high,
        low: b.low,
    }

    if value.isZero(b) {
        return nil
    }
    if value.compare(a, b) < 0 {
        remainder.high = a.high
        remainder.low = a.low
        a.high = 0
        a.low = 0
        return remainder
    }

    var mask uint64 = 0x8000000000000000
    var dPadding uint = 0
    var rPadding uint = 0
    var count uint = 128

    for count > 0 {
        if (remainder.high & mask) != 0 {
            break
        }

        value.leftShift(remainder, 1)
        rPadding++
        count--
    }
    
    remainder.high = a.high
    remainder.low = a.low

    count = 128
    for count > 0 {
        if (divider.high & mask) != 0 {
            break
        }
        
        value.leftShift(divider, 1)
        dPadding++
        count--
    }
    value.rightShiftUnsigned(divider, rPadding)

    count = dPadding - rPadding + 1
    for count > 0 {
        count--

        if value.compare(remainder, divider) >= 0 {
            value.sub(remainder, divider)
            quotient.low = quotient.low | 1
        }
        if count > 0 {
            value.leftShift(quotient, 1)
            value.rightShiftUnsigned(divider, 1)
        }
    }

    a.high = quotient.high
    a.low = quotient.low

    return remainder
}

func (value *Int128) nbits(val *Int128) int {
    bits := 0
    high := val.high
    low := val.low

    if high == 0 {
        for low > 0 {
            low = low >> 1
            bits++
        }
        return bits
    }

    for high > 0 {
        high = high >> 1
        bits++
    }
    return bits + 64
}

func (value *Int128) twosComplement(val *Int128) {
    val.high = (^val.high) >> 0
    val.low = (^val.low) >> 0

    var carry uint64 = 0
    low := val.low + 1
    if val.low > UINT64_MAX - 1 {
        carry = 1
    }

    high := val.high + carry

    val.high = high
    val.low = low
}

func (value *Int128) toBinaryString(val *Int128) string {
    strLow := strconv.FormatUint(val.low, 2)
    if val.high == 0 {
        return strLow
    }

    strHigh := strconv.FormatUint(val.high, 2)
    str := strHigh + paddingZero(strLow, 64);

    return str
}

func (value *Int128) toHexString(val *Int128) string {
    strLow := strconv.FormatUint(val.low, 16)
    if val.high == 0 {
        return strLow
    }

    strHigh := strconv.FormatUint(val.high, 16)
    str := strHigh + paddingZero(strLow, 16);

    return str
}

func (value *Int128) toDecimalStringSigned(val *Int128) string {
    var output []string

    stepper := &Int128 {
        high: 0,
        low: DECIMAL_STEPPER,
    }
    
    quotient := &Int128 {
        high: val.high,
        low: val.low,
    }

    neg := quotient.isNegative(quotient)
    if neg {
        quotient.twosComplement(quotient)
    }

    for !value.isZero(quotient) {
        remain := value.divide(quotient, stepper)
        var slc []string
        slc = append(slc, strconv.FormatUint(remain.low, 10))
        output = append(slc, output...)
        
    }

    if len(output) == 0 {
        return "0"
    } else {
        var x string
        if neg {
            x = "-"
        } else {
            x = ""
        }

        x = x + output[0]
        for _, comp := range output [1:] {
            x = x + paddingZero(comp, DECIMAL_STEPPER_LEN)
        }
        
        return x
    }
    return "0"
}
