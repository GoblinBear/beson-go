package beson

import (
	//"bytes"
	//"encoding/binary"
	"fmt"
	"strings"

	"beson/types"
)

func GetType(data interface{}) string {
	fmt.Println(".")
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
	case types.Int128:
		t = DATA_TYPE["INT128"]
	case uint8:
		t = DATA_TYPE["UINT8"]
	case uint16:
		t = DATA_TYPE["UINT16"]
	case uint32:
		t = DATA_TYPE["UINT32"]
	case uint64:
		t = DATA_TYPE["UINT64"]
	case types.UInt128:
		t = DATA_TYPE["UINT128"]
	case types.Binary:
		t = DATA_TYPE["BINARY"]
	default:
		t = ""
	}

	return t
}

func SerializeType(t string) []byte {
	typeHeader := make([]byte, 0)
	if t != "" {
		t = strings.ToUpper(t)
		typeHeader = TYPE_HEADER[t]
	}
	return typeHeader
}

func SerializeData(t string, data interface{}) []byte {
	var buffers []byte

	switch t {
	case DATA_TYPE["NULL"]:
		buffers = serializeNull()
	case DATA_TYPE["TRUE"], DATA_TYPE["FALSE"]:
		buffers = serializeBoolean()
	case DATA_TYPE["UINT8"]:
		buffers = make([]byte, 1)
		buffers[0] = data.(uint8)
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
