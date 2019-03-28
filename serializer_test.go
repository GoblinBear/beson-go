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

var expect = map[string][]byte {
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
    t.Run("NULL", testSerializeFunc(originData["NULL"], expect["NULL"]))
    t.Run("FALSE", testSerializeFunc(originData["FALSE"], expect["FALSE"]))
    t.Run("TRUE", testSerializeFunc(originData["TRUE"], expect["TRUE"]))
    t.Run("UINT8", testSerializeFunc(originData["UINT8"], expect["UINT8"]))
    t.Run("UINT16", testSerializeFunc(originData["UINT16"], expect["UINT16"]))
    t.Run("UINT32", testSerializeFunc(originData["UINT32"], expect["UINT32"]))
    t.Run("UINT64", testSerializeFunc(originData["UINT64"], expect["UINT64"]))
    t.Run("UINT128", testSerializeFunc(originData["UINT128"], expect["UINT128"]))
    t.Run("INT8", testSerializeFunc(originData["INT8"], expect["INT8"]))
    t.Run("INT16", testSerializeFunc(originData["INT16"], expect["INT16"]))
    t.Run("INT32", testSerializeFunc(originData["INT32"], expect["INT32"]))
    t.Run("INT64", testSerializeFunc(originData["INT64"], expect["INT64"]))
    t.Run("INT128", testSerializeFunc(originData["INT128"], expect["INT128"]))
    t.Run("FLOAT32", testSerializeFunc(originData["FLOAT32"], expect["FLOAT32"]))
    t.Run("FLOAT64", testSerializeFunc(originData["FLOAT64"], expect["FLOAT64"]))
    t.Run("STRING", testSerializeFunc(originData["STRING"], expect["STRING"]))
    t.Run("ARRAY", testSerializeFunc(originData["ARRAY"], expect["ARRAY"]))
    t.Run("MAP", testSerializeFunc(originData["MAP"], expect["MAP"]))
    t.Run("BINARY", testSerializeFunc(originData["BINARY"], expect["BINARY"]))
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