package helper

import (
    "bytes"
    "encoding/hex"
    "log"
    "strconv"
    "regexp"
)

const BYTE_MAX byte = 255
const DECIMAL_STEPPER byte = 100
const DECIMAL_STEPPER_LEN int = 2
const HEX_FORMAT_CHECKER string = "^0x[0-9a-fA-F]+$";

func HexStringToBytes(s string, size int) []byte {
    if len(s) == 0 {
        return nil
    }

    result := make([]byte, size)
    isMatch, _ := regexp.MatchString(HEX_FORMAT_CHECKER, s)
    if isMatch {
        s = s[2:]
        decoded, err := hex.DecodeString(s)
        if err != nil {
            log.Fatal(err)
            return nil
        }

        // convert to little endian
        reverse(decoded)
        copy(result, decoded)
        return result
    }

    log.Fatal("Input string must beginning with '0x'.")
    return nil
}

func BinaryStringToBytes(s string, size int) []byte {
    if len(s) == 0 {
        return nil
    }

    str := s
    if len(s) & 7 > 0 {
        str = paddingZero(s, len(s) + 8 - (len(s) & 7))
    }

    byteNum := len(str) >> 3
    result := make([]byte, size)
    
    for i := byteNum; i > 0; i-- {
        b, err := strconv.ParseUint(str[i*8-8:i*8], 2, 8)
        if err != nil {
            return nil
        }
        result[byteNum-i] = byte(b)
    }
    return result
}

func DecimalStringToBytes(s string, size int) []byte {
    if len(s) == 0 {
        return nil
    }

    str := ""
    neg := false
    if s[0] == '+' {
        str = s[1:]
    } else if s[0] == '-' {
        neg = true
        str = s[1:]
    } else {
        str = s
    }

    if len(str) & 1 > 0 {
        str = "0" + str
    }

    result := make([]byte, size)

    pow := make([]byte, size)
    pow[0] = 1

    stepper := make([]byte, size)
    stepper[0] = byte(DECIMAL_STEPPER)

    anchor := len(str)
    for anchor > 0 {
        b, err := strconv.ParseUint(str[anchor - DECIMAL_STEPPER_LEN:anchor], 10, 8)
        if err != nil {
            return nil
        }

        add := make([]byte, size)
        add[0] = byte(b)
        
        Multiply(add, pow)
        Add(result, add)
        Multiply(pow, stepper)
        anchor = anchor - DECIMAL_STEPPER_LEN
    }

    if neg {
        TwosComplement(result)
    }

    return result
}

func Concat(segments ...[]byte) []byte {
    buffer := bytes.NewBuffer(make([]byte, 0))

    for _, seg := range segments {
        buffer.Write(seg)
    }

    newBytes := make([]byte, buffer.Len())
    buffer.Read(newBytes)

    return newBytes
}

func LeftShift(value []byte, bits uint, padding uint8)  {
    if bits == 0 {
        return
    }
    if padding == 0 {
        padding = 0x00
    } else {
        padding = 0xFF
    }

    valueLength := uint(len(value))
    
    if bits >= valueLength * 8 {
        var off uint
        for off = 0; off < valueLength; off++ {
            value[off] = byte(padding)
        }
        return
    }

    byteNum := bits >> 3
    bitNum := bits & 7
    
    // copy bits
    var i uint
    for i = valueLength - 1; int(i) >= int(byteNum); i-- {
        high := (value[i-byteNum] & (1 << (8 - bitNum) - 1)) << bitNum
        var low byte = 0
        if int(i-byteNum-1) >= 0 {
            low = (value[i-byteNum-1] & ^(1 << (8 - bitNum) - 1)) >> (8 - bitNum)
        }
        value[i] = high | low
    }
    
    // padding
    for i = 0; i < byteNum; i++ {
        value[i] = padding
    }
    if padding > 0 {
        value[byteNum] = value[byteNum] | (1 << bitNum - 1)
    } else {
        value[byteNum] = value[byteNum] & ^(1 << bitNum - 1)
    }
}

func RightShift(value []byte, bits uint, padding uint8)  {
    if bits == 0 {
        return
    }
    if padding == 0 {
        padding = 0x00
    } else {
        padding = 0xFF
    }

    valueLength := uint(len(value))
    
    if bits >= valueLength * 8 {
        var off uint
        for off = 0; off < valueLength; off++ {
            value[off] = byte(padding)
        }
        return
    }

    byteNum := bits >> 3
    bitNum := bits & 7

    // copy bits
    var i uint
    for i = 0; i < valueLength - byteNum; i++ {
        var high byte = 0
        if i+byteNum+1 < valueLength {
            high = (value[i+byteNum+1] & (1 << bitNum - 1)) << (8 - bitNum)
        }
        low := ((value[i+byteNum] & ^(1 << bitNum - 1))) >> bitNum
        value[i] = high | low
    }
    
    // padding
    for i = valueLength - 1; i >= valueLength - byteNum; i-- {
        value[i] = padding
    }
    if padding > 0 {
        value[valueLength - byteNum - 1] = value[valueLength - byteNum - 1] | ^(1 << (8 - bitNum) - 1)
    } else {
        value[valueLength - byteNum - 1] = value[valueLength - byteNum - 1] & (1 << (8 - bitNum) - 1)
    }
}

func Not(value []byte) {
    for i := 0; i < len(value); i++ {
        value[i] = ^value[i];
    }
}

func And(a []byte, b []byte) {
    for i := 0; i < len(a); i++ {
        a[i] = a[i] & b[i];
    }
}

func Or(a []byte, b []byte) {
    for i := 0; i < len(a); i++ {
        a[i] = a[i] | b[i];
    }
}

func Xor(a []byte, b []byte) {
    for i := 0; i < len(a); i++ {
        a[i] = a[i] ^ b[i];
    }
}

func Add(a []byte, b []byte) {
    var carry, nextCarry byte = 0, 0
    for i := 0; i < len(a); i++ {
        if i < len(b) && ((a[i] > BYTE_MAX - b[i] - carry) || (b[i] > BYTE_MAX - a[i] - carry)) {
            nextCarry = 1
        } else {
            nextCarry = 0
        }
        if i < len(b) {
            a[i] = a[i] + b[i] + carry
        } else {
            a[i] = a[i] + carry
        }
        carry = nextCarry
    }
}

func Sub(a []byte, b []byte) {
    newB := make([]byte, len(b))
    copy(newB, b)
    TwosComplement(newB)
    Add(a, newB)
}

func Multiply(a []byte, b []byte) {
    ans := make([]byte, len(a) + len(b))
    bits := nbits(b)

    var i uint
    for i = bits - 1; int(i) >= 0; i-- {
        byteNum := i >> 3
        bitNum := i & 7

        LeftShift(ans, 1, 0)
        if (b[byteNum] & (1 << bitNum)) > 0 {
            Add(ans, a)
        }
    }
    copy(a, ans)
}

func Divide(a []byte, b []byte, signed bool) []byte {
    quotient := make([]byte, len(a))
    remainder := make([]byte, len(a))
    copy(remainder, a)
    divider := make([]byte, len(b))
    copy(divider, b)
    
    if IsZero(b) {
        log.Fatal("Divisor cannot be zero.")
    }
    if Compare(a, b) < 0 {
        for i := 0; i < len(a); i++ {
            a[i] = 0
        }
        return remainder
    }

    var negA, negB bool
    if signed {
        negA = IsNegative(a)
        if negA {
            TwosComplement(a)
        }
        negB = IsNegative(b)
        if negB {
            TwosComplement(b)
        }
    }

    var dPadding int = 0
    var rPadding int = 0
    var count int = len(remainder) * 8

    for count > 0 {
        count--
        if (remainder[len(remainder) - 1] & 0x80) != 0 {
            break
        }
        LeftShift(remainder, 1, 0)
        rPadding++
    }

    copy(remainder, a)
    count = len(divider) * 8

    for count > 0 {
        count--
        if (divider[len(divider) - 1] & 0x80) != 0 {
            break
        }
        LeftShift(divider, 1, 0)
        dPadding++
    }
    
    RightShift(divider, uint(rPadding), 0)
    count = dPadding - rPadding + 1

    for count > 0 {
        count--
        if Compare(remainder, divider) >= 0 {
            Sub(remainder, divider)
            quotient[0] = quotient[0] | 0x01
        }
        if count > 0 {
            LeftShift(quotient, 1, 0)
            RightShift(divider, 1, 0)
        }
    }

    if negA != negB {
       TwosComplement(quotient)
    }

    copy(a, quotient)
    return remainder
}

func Compare(a []byte, b []byte) int {
    if len(a) == 0 && len(b) == 0 {
        return 0
    }

    var valA, valB byte
    for i := max(len(a), len(b)) - 1; i >= 0; i-- {
        if i < len(a) {
            valA = a[i]
        } else {
            valA = 0
        }
        if i < len(b) {
            valB = b[i]
        } else {
            valB = 0
        }

        if valA == valB {
            continue
        }
        
        if valA > valB {
            return 1
        } else {
            return -1
        }
    }
    return 0
}

func IsZero(value []byte) bool {
    isZero := true
    for _, v := range value {
        isZero = isZero && (v == 0)
    }
    return isZero
}

func IsNegative(value []byte) bool {
    return (value[len(value) - 1] & 0x80) > 0
}

func TwosComplement(value []byte) {
    var carry, nextCarry byte = 1, 0
    for i := 0; i < len(value); i++ {
        if ^value[i] > BYTE_MAX - carry {
            nextCarry = 1
        } else {
            nextCarry = 0
        }
        value[i] = ^value[i] + carry
        carry = nextCarry
    }
}

func ToBinaryString(value []byte) string {
    str := ""
    if len(value) == 0 {
        return str
    }

    for i := len(value) - 1; i >= 0; i-- {
        subStr := strconv.FormatUint(uint64(value[i]), 2)
        str = str + paddingZero(subStr, 8)
    }
    return str
}

func ToHexString(value []byte) string {
    str := ""
    if len(value) == 0 {
        return str
    }
    
    for i := len(value) - 1; i >= 0; i-- {
        subStr := strconv.FormatUint(uint64(value[i]), 16)
        str = str + paddingZero(subStr, 2)
    }
    return str
}

func ToDecimalString(value []byte, signed bool) string {
    if len(value) == 0 {
        return ""
    }
    var output []string

    stepper := make([]byte, len(value))
    stepper[0] = byte(DECIMAL_STEPPER)

    quotient := make([]byte, len(value))
    copy(quotient, value)

    neg := IsNegative(quotient)
    if signed && neg {
        TwosComplement(quotient)
    }

    for !IsZero(quotient) {
        remain := Divide(quotient, stepper, false)
        var slc []string
        slc = append(slc, strconv.FormatUint(uint64(remain[0]), 10))
        output = append(slc, output...)
    }

    if len(output) == 0 {
        return "0"
    }

    var str string
    if signed && neg {
        str = "-"
    } else {
        str = ""
    }

    str = str + output[0]
    for _, comp := range output [1:] {
        str = str + paddingZero(comp, DECIMAL_STEPPER_LEN)
    }
    
    return str
}

func Resize(value []byte, size int, padding int) []byte {
    newValue := make([]byte, size)
    if padding > 0 {
        for i := 0; i < size; i++ {
            newValue[i] = 255
        }
    }
    copy(newValue, value)
    return newValue
}

func PaddingOne(value []byte) {
    for i := 0; i < len(value); i++ {
        value[i] = 255
    }
}

func paddingZero(data string, length int) string {
    zeros := length - len(data)
    padded := ""
    for zeros > 0 {
        padded = padded + "0";
        zeros--
    }

    return padded + data;
}

func nbits(value []byte) uint {
    var byteNum, bitNum int = 0, 0
    for i := len(value) - 1; i >= 0; i-- {
        if value[i] != 0 {
            byteNum = i
            break
        }
    }
    for i := 7; i >= 0; i-- {
        if value[byteNum] & (1 << byte(i)) > 0 {
            bitNum = i
            break
        } 
    }
    bits := byteNum * 8 + bitNum + 1
    return uint(bits)
}

func max(a int, b int) int {
    if a < b {
        return b
    }
    return a
}

func reverse(bs []byte) {
    for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
        bs[i], bs[j] = bs[j], bs[i]
    }
}
