package types

import (
    "errors"
)

const HEX_FORMAT_CHECKER string = "^0x[0-9a-fA-F]+$";

type Binary struct {
    bs []byte
}

func NewBinary(length int) *Binary {
    return newBinary(length).(*Binary)
}

func newBinary(length int) interface{} {
    if length < 0 {
        return nil
    }

    newBytes := make([]byte, length)
    return &Binary { bs: newBytes }
}

func (bin *Binary) Size() int {
    return len(bin.bs)
}

func (bin *Binary) Clone() *Binary {
    newBytes := make([]byte, len(bin.bs))
    copy(newBytes, bin.bs)
    return &Binary { bs: newBytes }
}

func (bin *Binary) Append(segments ...*Binary) *Binary {
    segments = append([]*Binary{ bin }, segments...)
    bytesBuffer := bin.bufferConcat(segments...)

    return &Binary { bs: bytesBuffer }
}

func (bin *Binary) Resize(length int) *Binary {
    if length < 0 {
        return nil
    }
    
    if length == len(bin.bs) {
        return bin
    } else if length < len(bin.bs) {
        newBytes := make([]byte, length)
        copy(newBytes, bin.bs[:length])        
        return &Binary { bs: newBytes }
    }

    newBytes := make([]byte, length)
    copy(newBytes, bin.bs)
    return bin
}

func (bin *Binary) LeftShift(bits uint, padding uint8) *Binary {
    bin.leftShift(bin.bs, bits, padding);
    return bin
}

func (bin *Binary) RightShift(bits uint, padding uint8) *Binary {
    bin.rightShift(bin.bs, bits, padding);
    return bin
}

func (bin *Binary) Not() *Binary {
    bin.not(bin.bs)
    return bin
}

func (bin *Binary) Compare(value *Binary, align bool) int {
    return bin.compare(bin.bs, value.bs, align)
}

func (bin *Binary) ToBytes() []byte {
    return bin.bs
}

func (bin *Binary) ToString(base int) (string, error) {
    switch base {
    case 2:
        return bin.toBinaryString(bin.bs), nil
    case 16:
        return bin.toHexString(bin.bs), nil
    default:
        return "", errors.New("Only accepts 2 and 16 representations.")
    }
    return "", nil
}

func (bin *Binary) From(segments ...*Binary) *Binary {
    bs := bin.bufferConcat(segments...)
    return &Binary { bs: bs }
}

func (bin *Binary) FromHex(hexString string) *Binary {
    if len(hexString) < 2 || hexString[0:2] != "0x" {
        hexString = "0x" + hexString
    }
    bs := bin.bufferFromHex(hexString)
    return &Binary { bs: bs }
}

func (bin *Binary) FromBytes(b []byte) *Binary {
    return &Binary { bs: b }
}
