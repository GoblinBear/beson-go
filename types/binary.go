package types

import (
	"errors"
	"fmt"
	"bytes"
	"regexp"
	"strconv"
)

const HEX_FORMAT_CHECKER string = "^0x[0-9a-fA-F]+$";

type Binary struct {
	buf *bytes.Buffer
}

func NewBinary(length int) *Binary {
	if (length < 0) {
		return nil
	}

	bytesBuffer := bytes.NewBuffer(make([]byte, length))
	return &Binary { buf: bytesBuffer }
}

func (bin *Binary) Size() int {
	return bin.buf.Len()
}

func (bin *Binary) Clone() *Binary {
	newBytes := make([]byte, bin.buf.Len())
    bin.buf.Read(newBytes)

	bytesBuffer := bytes.NewBuffer(make([]byte, bin.buf.Len()))
	bytesBuffer.Write(newBytes)

	return &Binary { buf: bytesBuffer }
}

func (bin *Binary) Append(segments ...*Binary) *Binary {
	segments = append([]*Binary{ bin }, segments...)
	bytesBuffer := bin.bufferConcat(segments...)

	return &Binary { buf: bytesBuffer }
}

func (bin *Binary) Resize(length int) *Binary {
	if (length < 0) {
		return nil
	}
	
	if (length == bin.buf.Len()) {
		return bin
	} else if (length < bin.buf.Len()) {
		newBytes := make([]byte, length)
		bin.buf.Read(newBytes)
		
		bytesBuffer := bytes.NewBuffer(make([]byte, 0))
		bytesBuffer.Write(newBytes)
		
		return &Binary { buf: bytesBuffer }
	}

	newBytes := make([]byte, length - bin.buf.Len())
	bin.buf.Write(newBytes)

	return bin
}

func (bin *Binary) Not() *Binary {
	bin.not(bin.buf);
	return bin;
}

func (bin *Binary) ToBytes() *bytes.Buffer {
	return bin.buf
}

func (bin *Binary) ToString(base int) (string, error) {
	switch base {
	case 2:
		return bin.toBinaryString(bin.buf), nil
	case 16:
		return bin.toHexString(bin.buf), nil
	default:
		return "", errors.New("Only accepts 2 and 16 representations.")
	}
	fmt.Println("")
	return "", nil
}

func (bin *Binary) From(segments ...*Binary) *Binary {
	bytesBuffer := bin.bufferConcat(segments...)
	return &Binary { buf: bytesBuffer }
}

func (bin *Binary) FromHex(hexString string) *Binary {
	if (hexString[0:2] != "0x") {
		hexString = "0x" + hexString
	}
	bytesBuffer := bin.bufferFromHex(hexString)
	return &Binary { buf: bytesBuffer }
}

var HEX_MAP_I = map[rune]uint8 {
	'0':0, '1':1, '2':2, '3':3, '4':4, '5':5, '6':6, '7':7, '8':8, '9':9,
	'a':10, 'b':11, 'c':12, 'd':13, 'e':14, 'f':15,
	'A':10, 'B':11, 'C':12, 'D':13, 'E':14, 'F':15,
};

func (bin *Binary) bufferFromHex(hexString string) *bytes.Buffer {
	isMatch, _ := regexp.MatchString(HEX_FORMAT_CHECKER, hexString)
	if (isMatch) {
		hexString = hexString[2:]
		if (len(hexString) % 2 == 1) {
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

func (bin *Binary) genMask(bits int) int {
	if (bits > 8) {
		return 0xFF;
	}
	if (bits < 0) {
		return 0;
	}

	val := 0
	for (bits > 0) {
		val = ((val << 1) | 1) >> 0
		bits--
	}
	return val;
}

func (bin *Binary) paddingZero(data string, length int) string {
	zeros := length - len(data)
	padded := ""
	for (zeros > 0) {
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
