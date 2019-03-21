package types

import (
    "errors"
    "fmt"
    "bytes"
)

const HEX_FORMAT_CHECKER string = "^0x[0-9a-fA-F]+$";

type Binary struct {
    buf *bytes.Buffer
}

func NewBinary(length int) *Binary {
    fmt.Println("....b")
    if length < 0 {
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
    if length < 0 {
        return nil
    }
    
    if length == bin.buf.Len() {
        return bin
    } else if length < bin.buf.Len() {
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
    return "", nil
}

func (bin *Binary) From(segments ...*Binary) *Binary {
    bytesBuffer := bin.bufferConcat(segments...)
    return &Binary { buf: bytesBuffer }
}

func (bin *Binary) FromHex(hexString string) *Binary {
    if hexString[0:2] != "0x" {
        hexString = "0x" + hexString
    }
    bytesBuffer := bin.bufferFromHex(hexString)
    return &Binary { buf: bytesBuffer }
}
