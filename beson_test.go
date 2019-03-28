package beson

import (
    "reflect"
    "testing"

    "beson/types"
)

var originData = map[string]types.RootType {
    "NULL":     nil,
    "TRUE":     types.NewBool(true).(*types.Bool),
    "FALSE":    types.NewBool(false).(*types.Bool),
    "UINT8":    types.NewUInt8(2).(*types.UInt8),
    "UINT16":   types.NewUInt16(2).(*types.UInt16),
    "UINT32":   types.NewUInt32(2).(*types.UInt32),
    "UINT64":   types.NewUInt64(2).(*types.UInt64),
    "UINT128":  types.NewUInt128("2", 10).(*types.UInt128),
    "INT8":     types.NewInt8(-3).(*types.Int8),
    "INT16":    types.NewInt16(-3).(*types.Int16),
    "INT32":    types.NewInt32(-3).(*types.Int32),
    "INT64":    types.NewInt64(-3).(*types.Int64),
    "INT128":   types.NewInt128("-3", 10).(*types.Int128),
    "FLOAT32":  types.NewFloat32(0.456).(*types.Float32),
    "FLOAT64":  types.NewFloat64(0.456).(*types.Float64),
    "STRING":   types.NewString("Hello world").(*types.String),
    "ARRAY":    types.NewSlice([]types.RootType { 
        types.NewFloat32(0.456).(*types.Float32),
        types.NewInt32(-3).(*types.Int32),
    }).(*types.Slice),
    "MAP":      types.NewMap(map[string]types.RootType { 
        "apple":    types.NewUInt8(2).(*types.UInt8),
        "banana":   types.NewBool(false).(*types.Bool),
    }).(*types.Map),
    "BINARY":   types.NewBinary(0).(*types.Binary).FromHex("0x2564877"),
}

var serializedData = map[string][]byte {
    "NULL":     []byte{ 0, 0 },
    "TRUE":     []byte{ 1, 1 },
    "FALSE":    []byte{ 1, 0 },
    "UINT8":    []byte{ 3, 4, 2 },
    "UINT16":   []byte{ 3, 5, 2, 0 },
    "UINT32":   []byte{ 3, 0, 2, 0, 0, 0 },
    "UINT64":   []byte{ 3, 1, 2, 0, 0, 0, 0, 0, 0, 0 },
    "UINT128":  []byte{ 3, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
    "INT8":     []byte{ 2, 4, 253 },
    "INT16":    []byte{ 2, 5, 253, 255 },
    "INT32":    []byte{ 2, 0, 253, 255, 255, 255 },
    "INT64":    []byte{ 2, 1, 253,  255,  255,  255,  255,  255,  255, 255 },
    "INT128":   []byte{ 2, 2, 253,  255,  255,  255,  255,  255,  255,  255,  255,  255,  255,  255,  255,  255,  255, 255 },
    "FLOAT32":  []byte{ 4, 1, 213, 120, 233, 62 },
    "FLOAT64":  []byte{ 4, 0, 201, 118, 190, 159, 26, 47, 221, 63 },
    "STRING":   []byte{ 5, 0, 11, 0, 0, 0, 72, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100 },
    "ARRAY":    []byte{ 6, 0, 12, 0, 0, 0, 4, 1, 213, 120, 233, 62, 2, 0, 253, 255, 255, 255 },
    "MAP":      []byte{ 9, 0, 20, 0, 0, 0, 3, 4, 5, 0, 97, 112, 112, 108, 101, 2, 1, 0, 6, 0, 98, 97, 110, 97, 110, 97 },
    "BINARY":   []byte{ 14, 0, 4, 0, 0, 0, 2, 86, 72, 119 },
}

func TestSerialize(t *testing.T) {
    t.Run("NULL", testSerializeFunc(originData["NULL"], serializedData["NULL"]))
    t.Run("FALSE", testSerializeFunc(originData["FALSE"], serializedData["FALSE"]))
    t.Run("TRUE", testSerializeFunc(originData["TRUE"], serializedData["TRUE"]))
    t.Run("UINT8", testSerializeFunc(originData["UINT8"], serializedData["UINT8"]))
    t.Run("UINT16", testSerializeFunc(originData["UINT16"], serializedData["UINT16"]))
    t.Run("UINT32", testSerializeFunc(originData["UINT32"], serializedData["UINT32"]))
    t.Run("UINT64", testSerializeFunc(originData["UINT64"], serializedData["UINT64"]))
    t.Run("UINT128", testSerializeFunc(originData["UINT128"], serializedData["UINT128"]))
    t.Run("INT8", testSerializeFunc(originData["INT8"], serializedData["INT8"]))
    t.Run("INT16", testSerializeFunc(originData["INT16"], serializedData["INT16"]))
    t.Run("INT32", testSerializeFunc(originData["INT32"], serializedData["INT32"]))
    t.Run("INT64", testSerializeFunc(originData["INT64"], serializedData["INT64"]))
    t.Run("INT128", testSerializeFunc(originData["INT128"], serializedData["INT128"]))
    t.Run("FLOAT32", testSerializeFunc(originData["FLOAT32"], serializedData["FLOAT32"]))
    t.Run("FLOAT64", testSerializeFunc(originData["FLOAT64"], serializedData["FLOAT64"]))
    t.Run("STRING", testSerializeFunc(originData["STRING"], serializedData["STRING"]))
    t.Run("ARRAY", testSerializeFunc(originData["ARRAY"], serializedData["ARRAY"]))
    t.Run("MAP", testSerializeFunc(originData["MAP"], serializedData["MAP"]))
    t.Run("BINARY", testSerializeFunc(originData["BINARY"], serializedData["BINARY"]))
}

func testSerializeFunc(data interface{}, expect []byte) func(*testing.T) {  
    return func(t *testing.T) {
        ser := Serialize(data)
        if reflect.DeepEqual(ser, expect) {
            t.Log("Serialize test passed.")
        } else {
            t.Error("Serialize test failed.")
        }
    }
}

func TestDeserialize(t *testing.T) {
    t.Run("NULL", testDeserializeFunc(serializedData["NULL"], originData["NULL"]))
    t.Run("FALSE", testDeserializeFunc(serializedData["FALSE"], originData["FALSE"]))
    t.Run("TRUE", testDeserializeFunc(serializedData["TRUE"], originData["TRUE"]))
    t.Run("UINT8", testDeserializeFunc(serializedData["UINT8"], originData["UINT8"]))
    t.Run("UINT16", testDeserializeFunc(serializedData["UINT16"], originData["UINT16"]))
    t.Run("UINT32", testDeserializeFunc(serializedData["UINT32"], originData["UINT32"]))
    t.Run("UINT64", testDeserializeFunc(serializedData["UINT64"], originData["UINT64"]))
    t.Run("UINT128", testDeserializeFunc(serializedData["UINT128"], originData["UINT128"]))
    t.Run("INT8", testDeserializeFunc(serializedData["INT8"], originData["INT8"]))
    t.Run("INT16", testDeserializeFunc(serializedData["INT16"], originData["INT16"]))
    t.Run("INT32", testDeserializeFunc(serializedData["INT32"], originData["INT32"]))
    t.Run("INT64", testDeserializeFunc(serializedData["INT64"], originData["INT64"]))
    t.Run("INT128", testDeserializeFunc(serializedData["INT128"], originData["INT128"]))
    t.Run("FLOAT32", testDeserializeFunc(serializedData["FLOAT32"], originData["FLOAT32"]))
    t.Run("FLOAT64", testDeserializeFunc(serializedData["FLOAT64"], originData["FLOAT64"]))
    t.Run("STRING", testDeserializeFunc(serializedData["STRING"], originData["STRING"]))
    t.Run("ARRAY", testDeserializeFunc(serializedData["ARRAY"], originData["ARRAY"]))
    t.Run("MAP", testDeserializeFunc(serializedData["MAP"], originData["MAP"]))
    t.Run("BINARY", testDeserializeFunc(serializedData["BINARY"], originData["BINARY"]))
}

func testDeserializeFunc(ser []byte, expect interface{}) func(*testing.T) { 
    return func(t *testing.T) {
        _, data := Deserialize(ser, 0)
        if reflect.DeepEqual(data, expect) {
            t.Log("Deserialize test passed.")
        } else {
            t.Error("Deserialize test failed.")
        }
    }
}
