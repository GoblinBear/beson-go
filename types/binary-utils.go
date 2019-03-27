package types

import (
    "bytes"
    "regexp"
    "strconv"
)

var HEX_MAP_I = map[rune]uint8 {
    '0':0, '1':1, '2':2, '3':3, '4':4, '5':5, '6':6, '7':7, '8':8, '9':9,
    'a':10, 'b':11, 'c':12, 'd':13, 'e':14, 'f':15,
    'A':10, 'B':11, 'C':12, 'D':13, 'E':14, 'F':15,
};

func (bin *Binary) bufferFromHex(hexString string) []byte {
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
        
        newBytes := make([]byte, bytesBuffer.Len())
        bytesBuffer.Read(newBytes)

        return newBytes
    }

    return nil
}

func (bin *Binary) bufferConcat(segments ...*Binary) []byte {
    bytesBuffer := bytes.NewBuffer(make([]byte, 0))

    for _, seg := range segments {
        newBytes := make([]byte, len(seg.bs))
        bytesBuffer.Write(newBytes)
    }

    newBytes := make([]byte, bytesBuffer.Len())
    bytesBuffer.Read(newBytes)

    return newBytes
}

func (bin *Binary) compare(a []byte, b []byte, align bool) int {
    if len(a) == 0 && len(b) == 0 {
        return 0
    }

    newA := make([]byte, len(a))
    newB := make([]byte, len(b))
    copy(newA, a)
    copy(newB, b)

    if !align {
        var valA, valB, max int 
        if len(newA) > len(newB) {
            valA = len(newA)
            valB = len(newA)
            max = len(newA)
        } else {
            valA = len(newB)
            valB = len(newB)
            max = len(newB)
        }

        for i := 0; i < max; i++ {
            if i > len(newA) - 1 {
                valA = 0
            } else {
                valA = int(newA[i])
            }

            if i > len(newB) - 1 {
                valB = 0
            } else {
                valB = int(newB[i])
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
    } else {
        var shiftA, shiftB, valA, valB, max, offset int
        if len(newA) > len(newB) {
            max = len(newA)
            shiftA = 0
            shiftB = len(newB) - len(newA)
        } else if len(newA) < len(newB) {
            max = len(newB)
            shiftA = len(newA) - len(newB)
            shiftB = 0
        } else {
            max = len(newA)
            shiftA = 0
            shiftB = 0
        }

        for i := 0; i < max; i++ {
            offset = i + shiftA
            if offset >= 0 {
                valA = int(newA[offset])
            } else {
                valA = 0
            }

            offset = i + shiftB
            if offset >= 0 {
                valB = int(newB[offset])
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
        
        return 0;
    }

    return 0
}

func (bin *Binary) leftShift(value []byte, bits uint, padding uint8) {
    valueLength := uint(len(value))
    newBytes := make([]byte, valueLength)
    copy(newBytes, value)

    if bits > 0 {
        if padding == 0 {
            padding = 0x00
        } else {
            padding = 0xFF
        }

        if bits >= valueLength * 8 {
            var off uint
            for off = 0; off < valueLength; off++ {
                newBytes[off] = byte(padding)
            }
        } else {
            offset := (bits / 8) | 0
            lastOffset := valueLength - offset
            realShift := bits % 8
            realShiftI := 8 - realShift
            
            var lowMask uint8
            if realShift == 0 {
                lowMask = 0
            } else {
                lowMask = bin.genMask(realShift) << realShiftI
            }

            var off uint
            for off = 0; off < lastOffset; off++ {
                shift := off + offset
                
                if realShift == 0 {
                    newBytes[off] = newBytes[shift]
                } else {
                    shiftVal := newBytes[shift]
                    var next uint8
                    if shift >= (valueLength - 1) {
                        next = padding
                    } else {
                        next = uint8(newBytes[shift + 1])
                    }

                    newBytes[off] = byte(uint8((shiftVal << realShift) | ((next & lowMask) >> realShiftI)))
                }
            }

            for off = lastOffset; off < valueLength; off++ {
                newBytes[off] = byte(padding)
            }
        }
    }
}

func (bin *Binary) rightShift(value []byte, bits uint, padding uint8) {
    valueLength := uint(len(value))
    newBytes := make([]byte, valueLength)
    copy(newBytes, value)

    if bits > 0 {
        if padding == 0 {
            padding = 0x00
        } else {
            padding = 0xFF
        }

        if bits >= valueLength * 8 {
            var off uint
            for off = 0; off < valueLength; off++ {
                newBytes[off] = byte(padding)
            }
        } else {
            offset := (bits / 8) | 0
            realShift := bits % 8
            realShiftI := 8 - realShift
            
            var highMask uint8
            if realShift == 0 {
                highMask = 0
            } else {
                highMask = bin.genMask(realShift)
            }

            var off uint
            for off = valueLength; off > offset; off-- {
                shift := off - offset - 1
                
                if realShift == 0 {
                    newBytes[off - 1] = newBytes[shift]
                } else {
                    shiftVal := newBytes[shift]
                    var next uint8
                    if shift == 0 {
                        next = padding
                    } else {
                        next = uint8(newBytes[shift - 1])
                    }

                    newBytes[off - 1] = byte(uint8((shiftVal >> realShift) | ((next & highMask) << realShiftI)))
                }
            }
            
            for off = offset; off > 0; off-- {
                newBytes[off - 1] = byte(padding)
            }
        }
    }
}

func (bin *Binary) not(value []byte) {
    valueLength := len(value)
    newBytes := make([]byte, valueLength)
    copy(newBytes, value)

    for off := 0; off < valueLength; off++ {
        newBytes[off] = ^newBytes[off];
    }
}

func (bin *Binary) genMask(bits uint) uint8 {
    if bits > 8 {
        return 0xFF;
    }

    var val uint8 = 0
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

func (bin *Binary) toBinaryString(val []byte) string {
    str := ""
    newBytes := make([]byte, len(val))
    copy(newBytes, val)

    for _, byte := range newBytes {
        s := strconv.FormatInt(int64(byte), 2)
        str = str + bin.paddingZero(s, 8);
    }

    return str
}

func (bin *Binary) toHexString(val []byte) string {
    str := ""
    newBytes := make([]byte, len(val))
    copy(newBytes, val)

    for _, byte := range newBytes {
        s := strconv.FormatInt(int64(byte), 16)
        str = str + bin.paddingZero(s, 2);
    }

    return str
}
