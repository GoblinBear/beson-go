package types


/* Root type */
type RootType interface {
}


/* Type definition */

type UInt8 struct {
    value uint8
}

type UInt16 struct {
    value uint16
}

type UInt32 struct {
    value uint32
}

type UInt64 struct {
    value uint64
}

type Int8 struct {
    value int8
}

type Int16 struct {
    value int16
}

type Int32 struct {
    value int32
}

type Int64 struct {
    value int64
}

type Float32 struct {
    value float32
}

type Float64 struct {
    value float64
}

type Bool struct {
    value bool
}

type String struct {
    value string
}

type Slice struct {
	slice []RootType
}

type Map struct {
	m map[RootType]RootType
}


/* Type initializer */

func NewUInt8(value uint8) RootType {
    return &UInt8 { value }
}

func NewUInt16(value uint16) RootType {
    return &UInt16 { value }
}

func NewUInt32(value uint32) RootType {
    return &UInt32 { value }
}

func NewUInt64(value uint64) RootType {
    return &UInt64 { value }
}

func NewInt8(value int8) RootType {
    return &Int8 { value }
}

func NewInt16(value int16) RootType {
    return &Int16 { value }
}

func NewInt32(value int32) RootType {
    return &Int32 { value }
}

func NewInt64(value int64) RootType {
    return &Int64 { value }
}

func NewFloat32(value float32) RootType {
    return &Float32 { value }
}

func NewFloat64(value float64) RootType {
    return &Float64 { value }
}

func NewBool(value bool) RootType {
    return &Bool { value }
}

func NewString(value string) RootType {
    return &String { value }
}

func NewSlice(value []RootType) RootType {
    return &Slice { value }
}

func NewMap(value map[RootType]RootType) RootType {
    return &Map { value }
}
