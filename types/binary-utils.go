package types

import (
    //"fmt"
    "bytes"
    "regexp"
    "strconv"
)

var HEX_MAP_I = map[rune]uint8 {
    '0':0, '1':1, '2':2, '3':3, '4':4, '5':5, '6':6, '7':7, '8':8, '9':9,
    'a':10, 'b':11, 'c':12, 'd':13, 'e':14, 'f':15,
    'A':10, 'B':11, 'C':12, 'D':13, 'E':14, 'F':15,
};

func (bin *Binary) bufferFromHex(hexString string) *bytes.Buffer {
    isMatch, _ := regexp.MatchString(HEX_FORMAT_CHECKER, hexString)
    if isMatch {
        hexString = hexString[2:]
        if len(hexString) % 2 == 1 {
            hexString = "0" + hexString
        }
        bytesBuffer := bytes.NewBuffer(make([]byte, 0))
        
        for pointer := 0; pointer < len(hexString) / 2; pointer++ {
            offset := pointer * 2
            buf := HEX_MAP_I[rune(hexString[offset+1])] | (HEX_MAP_I[rune(hexString[offset])] << 4)
            bytesBuffer.WriteByte(buf)
        }
        return bytesBuffer
    }

    return nil
}

func (bin *Binary) bufferConcat(segments ...*Binary) *bytes.Buffer {
    bytesBuffer := bytes.NewBuffer(make([]byte, 0))
    
    for _, seg := range segments {
        newBytes := make([]byte, seg.buf.Len())
        seg.buf.Read(newBytes)
        bytesBuffer.Write(newBytes)
    }

    return bytesBuffer
}

func (bin *Binary) leftShift(value *bytes.Buffer, bits int, padding int) {

}

func (bin *Binary) not(value *bytes.Buffer){
    valueLength := value.Len()
    newBytes := make([]byte, value.Len())
    value.Read(newBytes)

    for off := 0; off < valueLength; off++ {
        newBytes[off] = ^newBytes[off];
    }
    value.Write(newBytes)
}

func (bin *Binary) genMask(bits uint) int {
    if bits > 8 {
        return 0xFF;
    }
    if bits < 0 {
        return 0;
    }

    val := 0
    for bits > 0 {
        val = ((val << 1) | 1) >> 0
        bits--
    }
    return val;
}

func (bin *Binary) paddingZero(data string, length int) string {
    zeros := length - len(data)
    padded := ""
    for zeros > 0 {
        padded = padded + "0";
        zeros--
    }

    return padded + data;
}

func (bin *Binary) toBinaryString(val *bytes.Buffer) string {
    str := ""
    newBytes := make([]byte, val.Len())
    val.Read(newBytes)

    for _, byte := range newBytes {
        s := strconv.FormatInt(int64(byte), 2)
        str = str + bin.paddingZero(s, 8);
    }

    return str
}

func (bin *Binary) toHexString(val *bytes.Buffer) string {
    str := ""
    newBytes := make([]byte, val.Len())
    val.Read(newBytes)

    for _, byte := range newBytes {
        s := strconv.FormatInt(int64(byte), 16)
        str = str + bin.paddingZero(s, 2);
    }

    return str
}
