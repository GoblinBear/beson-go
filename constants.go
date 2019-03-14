package beson

import (
	//"fmt"
   	"bytes"
)

var dataType = map[string]string {
	"NULL":           	"null",
	"FALSE":          	"false",
	"TRUE":           	"true",
	
	"INT32":          	"int32",
	"INT64":          	"int64",
	"INT128":         	"int128",
	"INT8":				"int8",
	"INT16":			"int16",
	
	"UINT32":         	"uint32",
	"UINT64":         	"uint64",
	"UINT128":        	"uint128",
	"UINT8":			"uint8",
	"UINT16":			"uint16",
	
	"FLOAT64":        	"float64",
	"FLOAT32":			"float32",
	
	"STRING":         	"string",
	"ARRAY":          	"array",
	"ARRAY_START":    	"array_start",
	"ARRAY_END":      	"array_end",
	"OBJECT":         	"object",
	"OBJECT_START":   	"object_start",
	"OBJECT_END":     	"object_end",
	"DATE":           	"date",
	"OBJECTID":       	"objectid",
	"BINARY":         	"binary",
	
	"ARRAY_BUFFER":		"array_buffer",
	"DATA_VIEW":		"data_view",
	"UINT8_ARRAY":		"uint8_array",
	"INT8_ARRAY":		"int8_array",
	"UINT16_ARRAY":		"uint16_array",
	"INT16_ARRAY":		"int16_array",
	"UINT32_ARRAY":		"uint32_array",
	"INT32_ARRAY":		"int32_array",
	"FLOAT32_ARRAY":	"float32_array",
	"FLOAT64_ARRAY":	"float64_array",
	
	"SPECIAL_BUFFER": 	"special_buffer",
}

var typeHeader = map[string][2]uint8 {
	"NULL":           	{ 0x00, 0x00 },
	"FALSE":          	{ 0x01, 0x00 },
	"TRUE":           	{ 0x01, 0x01 },
	
	"INT32":          	{ 0x02, 0x00 },
	"INT64":          	{ 0x02, 0x01 },
	"INT128":         	{ 0x02, 0x02 },
	"INT8":           	{ 0x02, 0x04 },
	"INT16":          	{ 0x02, 0x05 },
	
	"UINT32":         	{ 0x03, 0x00 },
	"UINT64":        	{ 0x03, 0x01 },
	"UINT128":        	{ 0x03, 0x02 },
	"UINT8":          	{ 0x03, 0x04 },
	"UINT16":         	{ 0x03, 0x05 },
	
	"FLOAT64":       	{ 0x04, 0x00 },
	"FLOAT32":			{ 0x04, 0x01 },
	
	"STRING":         	{ 0x05, 0x00 },
	"ARRAY":          	{ 0x06, 0x00 },
	"ARRAY_START":    	{ 0x07, 0x00 },
	"ARRAY_END":      	{ 0x08, 0x00 },
	"OBJECT":         	{ 0x09, 0x00 },
	"OBJECT_START":   	{ 0x0a, 0x00 },
	"OBJECT_END":     	{ 0x0b, 0x00 },
	"DATE":           	{ 0x0c, 0x00 },
	"OBJECTID":       	{ 0x0d, 0x00 },
	"BINARY":         	{ 0x0e, 0x00 },
	
	"ARRAY_BUFFER":		{ 0x0f, 0x00 },
	"DATA_VIEW":		{ 0x0f, 0x01 },
	"UINT8_ARRAY":		{ 0x0f, 0x02 },
	"INT8_ARRAY":		{ 0x0f, 0x03 },
	"UINT16_ARRAY":		{ 0x0f, 0x04 },
	"INT16_ARRAY":		{ 0x0f, 0x05 },
	"UINT32_ARRAY":		{ 0x0f, 0x06 },
	"INT32_ARRAY":		{ 0x0f, 0x07 },
	"FLOAT32_ARRAY":	{ 0x0f, 0x08 },
	"FLOAT64_ARRAY":	{ 0x0f, 0x09 },
	
	"SPECIAL_BUFFER":	{ 0x0f, 0xff },
}

func TypeHeader(str string) []bytes.Buffer {
	var header []bytes.Buffer

	for _, t := range typeHeader[str] {
		var buf bytes.Buffer
		buf.WriteByte(byte(t))
		header = append(header, buf)
	}
	return header
}
