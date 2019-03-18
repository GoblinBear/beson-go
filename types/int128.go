package types

import (
	"fmt"
	"strconv"
)

type Int128 struct {
	high int64
	low int64
}

func NewInt128(value string) *Int128 {
	fmt.Println("....")
	newValue := Int128 {
		high: 0,
		low: 0,
	}

	binaryString := decimalStringToBinaryString(value)
	length := len(binaryString)
	
	if (length > 128) {
		return nil
	}

	if (length > 64) {
		newValue.high, _ = strconv.ParseInt(binaryString[:length - 64], 2, 64)
		newValue.low, _ = strconv.ParseInt(binaryString[length - 64:], 2, 64)
	} else {
		newValue.high = 0
		newValue.low, _ = strconv.ParseInt(binaryString, 2, 64)
	}
	
	return &newValue
}


func (value *Int128) IsNegative() bool {
	return isNegative(value)
}
