package beson

import (
    "bytes"
    "encoding/binary"
    "log"
    "math"
    "strings"

    "beson/types"
)

// Serialize convert common type data to binary sequence.
func Serialize(data interface{}) []byte {
    return serializeContent(data)
}

func serializeContent(data interface{}) []byte {
    t := getType(data)
    typeBuffer := serializeType(t)
    dataBuffers := serializeData(t, data)

    bytesBuffer := bytes.NewBuffer(make([]byte, 0))
    bytesBuffer.Write(typeBuffer)
    bytesBuffer.Write(dataBuffers)

    length := len(typeBuffer) + len(dataBuffers)
    serialContent := make([]byte, length)
    bytesBuffer.Read(serialContent)

    return serialContent
}

func getType(data interface{}) string {
    var t string

    if data == nil {
        t = DATA_TYPE["NULL"]
        return t
    }

    switch data.(type) {
    case bool:
        if data.(bool) {
            t = DATA_TYPE["TRUE"]
        } else {
            t = DATA_TYPE["FALSE"]
        }
    case float32:
        t = DATA_TYPE["FLOAT32"]
    case float64:
        t = DATA_TYPE["FLOAT64"]
    case int8:
        t = DATA_TYPE["INT8"]
    case int16:
        t = DATA_TYPE["INT16"]
    case int32:
        t = DATA_TYPE["INT32"]
    case int64:
        t = DATA_TYPE["INT64"]
    case *types.Int128:
        t = DATA_TYPE["INT128"]
    case *types.Int256:
        t = DATA_TYPE["INT256"]
    case *types.Int512:
        t = DATA_TYPE["INT512"]
    case *types.IntVar:
        t = DATA_TYPE["INTVAR"]
    // case *types.UInt8:
    case uint8:
        t = DATA_TYPE["UINT8"]
    case uint16:
        t = DATA_TYPE["UINT16"]
    case uint32:
        t = DATA_TYPE["UINT32"]
    case uint64:
        t = DATA_TYPE["UINT64"]
    case *types.UInt128:
        t = DATA_TYPE["UINT128"]
    case *types.UInt256:
        t = DATA_TYPE["UINT256"]
    case *types.UInt512:
        t = DATA_TYPE["UINT512"]
    case *types.UIntVar:
        t = DATA_TYPE["UINTVAR"]
    case *types.Binary:
        t = DATA_TYPE["BINARY"]
    case string:
        t = DATA_TYPE["STRING"]
    case []interface{}:
        t = DATA_TYPE["ARRAY"]
    case map[string]interface{}:
        t = DATA_TYPE["MAP"]
    default:
        t = ""
    }
    
    return t
}

func serializeType(t string) []byte {
    typeHeader := make([]byte, 0)
    if t != "" {
        t = strings.ToUpper(t)
        typeHeader = TYPE_HEADER[t]
    }
    return typeHeader
}

func serializeData(t string, data interface{}) []byte {
    var buffers []byte

    switch t {
    case DATA_TYPE["NULL"]:
        buffers = serializeNull()
    case DATA_TYPE["TRUE"], DATA_TYPE["FALSE"]:
        buffers = serializeBoolean()
    case DATA_TYPE["UINT8"]:
        buffers = make([]byte, 1)
        buffers[0] = data.(uint8)
    case DATA_TYPE["UINT16"]:
        buffers = make([]byte, 2)
        binary.LittleEndian.PutUint16(buffers, data.(uint16))
    case DATA_TYPE["UINT32"]:
        buffers = make([]byte, 4)
        binary.LittleEndian.PutUint32(buffers, data.(uint32))
    case DATA_TYPE["UINT64"]:
        buffers = make([]byte, 8)
        binary.LittleEndian.PutUint64(buffers, data.(uint64))
    case DATA_TYPE["UINT128"]:
        buffers = serializeUInt128(data.(*types.UInt128))
    case DATA_TYPE["UINT256"]:
        buffers = serializeUInt256(data.(*types.UInt256))
    case DATA_TYPE["UINT512"]:
        buffers = serializeUInt512(data.(*types.UInt512))
    case DATA_TYPE["UINTVAR"]:
        buffers = serializeUIntVar(data.(*types.UIntVar))
    case DATA_TYPE["INT8"]:
        buffers = make([]byte, 1)
        buffers[0] = uint8(data.(int8))
    case DATA_TYPE["INT16"]:
        buffers = make([]byte, 2)
        binary.LittleEndian.PutUint16(buffers, uint16(data.(int16)))
    case DATA_TYPE["INT32"]:
        buffers = make([]byte, 4)
        binary.LittleEndian.PutUint32(buffers, uint32(data.(int32)))
    case DATA_TYPE["INT64"]:
        buffers = make([]byte, 8)
        binary.LittleEndian.PutUint64(buffers, uint64(data.(int64)))
    case DATA_TYPE["INT128"]:
        buffers = serializeInt128(data.(*types.Int128))
    case DATA_TYPE["INT256"]:
        buffers = serializeInt256(data.(*types.Int256))
    case DATA_TYPE["INT512"]:
        buffers = serializeInt512(data.(*types.Int512))
    case DATA_TYPE["INTVAR"]:
        buffers = serializeIntVar(data.(*types.IntVar))
    case DATA_TYPE["FLOAT32"]:
        bits := math.Float32bits(data.(float32))
        buffers = make([]byte, 4)
        binary.LittleEndian.PutUint32(buffers, bits)
    case DATA_TYPE["FLOAT64"]:
        bits := math.Float64bits(data.(float64))
        buffers = make([]byte, 8)
        binary.LittleEndian.PutUint64(buffers, bits)
    case DATA_TYPE["STRING"]:
        s := data.(string)
        buffers = serializeString(s)
    case DATA_TYPE["ARRAY"]:
        slice := data.([]interface{})
        buffers = serializeSlice(slice)
    case DATA_TYPE["MAP"]:
        m := data.(map[string]interface{})
        buffers = serializeMap(m)
    case DATA_TYPE["BINARY"]:
        b := data.(*types.Binary)
        buffers = serializeBinary(b)
    }

    return buffers
}

func serializeNull() []byte {
    buf := make([]byte, 0)
    return buf
}

func serializeBoolean() []byte {
    buf := make([]byte, 0)
    return buf
}

func serializeUInt128(value *types.UInt128) []byte {
    buf := value.ToBytes()
    return buf
}

func serializeUInt256(value *types.UInt256) []byte {
    buf := value.ToBytes()
    return buf
}

func serializeUInt512(value *types.UInt512) []byte {
    buf := value.ToBytes()
    return buf
}

func serializeUIntVar(value *types.UIntVar) []byte {
    dataBytes := value.ToBytes()
    length := len(dataBytes)
    if length > 127 {
        log.Fatal("Cannot support IntVar whose size is greater than 127 bytes.")
    }

    lengthBytes := []byte{ byte(length) }
    buf := concatBytesArray(lengthBytes, dataBytes)

    return buf
}

func serializeInt128(value *types.Int128) []byte {
    buf := value.ToBytes()
    return buf
}

func serializeInt256(value *types.Int256) []byte {
    buf := value.ToBytes()
    return buf
}

func serializeInt512(value *types.Int512) []byte {
    buf := value.ToBytes()
    return buf
}

func serializeIntVar(value *types.IntVar) []byte {
    dataBytes := value.ToBytes()
    length := len(dataBytes)
    if length > 127 {
        log.Fatal("Cannot support IntVar whose size is greater than 127 bytes.")
    }

    lengthBytes := []byte{ byte(length) }
    buf := concatBytesArray(lengthBytes, dataBytes)

    return buf
}

func serializeString(value string) []byte {
    str := value
    length := len(str)
    lengthBytes := make([]byte, 4)
    binary.LittleEndian.PutUint32(lengthBytes, uint32(length))
    
    dataBytes := []byte(str)
    buf := concatBytesArray(lengthBytes, dataBytes)
    return buf
}

func serializeShortString(value string) []byte {
    str := value
    length := len(str)
    lengthBytes := make([]byte, 2)
    binary.LittleEndian.PutUint16(lengthBytes, uint16(length))
    
    dataBytes := []byte(str)
    buf := concatBytesArray(lengthBytes, dataBytes)
    return buf
}

func serializeSlice(value []interface{}) []byte {
    slice := value
    subBytesBuffer := bytes.NewBuffer(make([]byte, 0))
    for _, element := range slice {
        subType := getType(element)
        subTypeBytes := serializeType(subType)
        subDataBytes := serializeData(subType, element)
        subBytesBuffer.Write(subTypeBytes)
        subBytesBuffer.Write(subDataBytes)
    }

    length := subBytesBuffer.Len()
    lengthBytes := make([]byte, 4)
    binary.LittleEndian.PutUint32(lengthBytes, uint32(length))

    dataBytes := make([]byte, length)
    subBytesBuffer.Read(dataBytes)

    buf := concatBytesArray(lengthBytes, dataBytes)
    return buf
}

func serializeMap(value map[string]interface{}) []byte {
    subBytesBuffer := bytes.NewBuffer(make([]byte, 0))
    m := value
    for key, value := range m {
        // serialize key
        keyBytes := serializeShortString(key)

        // serialize value
        subType := getType(value)
        subTypeBytes := serializeType(subType)
        subDataBytes := serializeData(subType, value)

        subBytesBuffer.Write(subTypeBytes)
        subBytesBuffer.Write(keyBytes)
        subBytesBuffer.Write(subDataBytes)
    }

    length := subBytesBuffer.Len()
    lengthBytes := make([]byte, 4)
    binary.LittleEndian.PutUint32(lengthBytes, uint32(length))

    dataBytes := make([]byte, length)
    subBytesBuffer.Read(dataBytes)

    buf := concatBytesArray(lengthBytes, dataBytes)
    return buf
}

func serializeBinary(value *types.Binary) []byte {
    dataBytes := value.ToBytes()
    length := len(dataBytes)
    lengthBytes := make([]byte, 4)
    binary.LittleEndian.PutUint32(lengthBytes, uint32(length))

    buf := concatBytesArray(lengthBytes, dataBytes)
    return buf
}

func concatBytesArray(b1 []byte, b2 ...[]byte) []byte {
    buf := bytes.NewBuffer(make([]byte, 0))
    
    buf.Write(b1)
    for _, element := range b2 {
        buf.Write(element)
    }

    newBytes := make([]byte, buf.Len())
    buf.Read(newBytes)

    return newBytes
}
