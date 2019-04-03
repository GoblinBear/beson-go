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
    str string
}

type Slice struct {
	slice []RootType
}

type Map struct {
	m map[string]RootType
}


/* Type initializer */

func NewUInt8(value uint8) *UInt8 {
    return newUInt8(value).(*UInt8)
}

func NewUInt16(value uint16) *UInt16 {
    return newUInt16(value).(*UInt16)
}

func NewUInt32(value uint32) *UInt32 {
    return newUInt32(value).(*UInt32)
}

func NewUInt64(value uint64) *UInt64 {
    return newUInt64(value).(*UInt64)
}

func NewInt8(value int8) *Int8 {
    return newInt8(value).(*Int8)
}

func NewInt16(value int16) *Int16 {
    return newInt16(value).(*Int16)
}

func NewInt32(value int32) *Int32 {
    return newInt32(value).(*Int32)
}

func NewInt64(value int64) *Int64 {
    return newInt64(value).(*Int64)
}

func NewFloat32(value float32) *Float32 {
    return newFloat32(value).(*Float32)
}

func NewFloat64(value float64) *Float64 {
    return newFloat64(value).(*Float64)
}

func NewBool(value bool) *Bool {
    return newBool(value).(*Bool)
}

func NewString(value string) *String {
    return newString(value).(*String)
}

func NewSlice(value []RootType) *Slice {
    return newSlice(value).(*Slice)
}

func NewMap(value map[string]RootType) *Map {
    return newMap(value).(*Map)
}


/* Initialize type to RootType */

func newUInt8(value uint8) RootType {
    return &UInt8 { value }
}

func newUInt16(value uint16) RootType {
    return &UInt16 { value }
}

func newUInt32(value uint32) RootType {
    return &UInt32 { value }
}

func newUInt64(value uint64) RootType {
    return &UInt64 { value }
}

func newInt8(value int8) RootType {
    return &Int8 { value }
}

func newInt16(value int16) RootType {
    return &Int16 { value }
}

func newInt32(value int32) RootType {
    return &Int32 { value }
}

func newInt64(value int64) RootType {
    return &Int64 { value }
}

func newFloat32(value float32) RootType {
    return &Float32 { value }
}

func newFloat64(value float64) RootType {
    return &Float64 { value }
}

func newBool(value bool) RootType {
    return &Bool { value }
}

func newString(value string) RootType {
    return &String { value }
}

func newSlice(value []RootType) RootType {
    return &Slice { value }
}

func newMap(value map[string]RootType) RootType {
    return &Map { value }
}


/* Get value */

func (value *UInt8) Get() uint8 {
    return value.value
}

func (value *UInt16) Get() uint16 {
    return value.value
}

func (value *UInt32) Get() uint32 {
    return value.value
}

func (value *UInt64) Get() uint64 {
    return value.value
}

func (value *Int8) Get() int8 {
    return value.value
}

func (value *Int16) Get() int16 {
    return value.value
}

func (value *Int32) Get() int32 {
    return value.value
}

func (value *Int64) Get() int64 {
    return value.value
}

func (value *Float32) Get() float32 {
    return value.value
}

func (value *Float64) Get() float64 {
    return value.value
}

func (value *Bool) Get() bool {
    return value.value
}

func (value *String) Get() string {
    return value.str
}

func (value *Slice) Get() []RootType {
    return value.slice
}

func (value *Map) Get() map[string]RootType {
    return value.m
}


/* Set value */

func (value *UInt8) Set(newValue uint8) {
    value.value = newValue
}

func (value *UInt16) Set(newValue uint16) {
    value.value = newValue
}

func (value *UInt32) Set(newValue uint32) {
    value.value = newValue
}

func (value *UInt64) Set(newValue uint64) {
    value.value = newValue
}

func (value *Int8) Set(newValue int8) {
    value.value = newValue
}

func (value *Int16) Set(newValue int16) {
    value.value = newValue
}

func (value *Int32) Set(newValue int32) {
    value.value = newValue
}

func (value *Int64) Set(newValue int64) {
    value.value = newValue
}

func (value *Float32) Set(newValue float32) {
    value.value = newValue
}

func (value *Float64) Set(newValue float64) {
    value.value = newValue
}

func (value *Bool) Set(newValue bool) {
    value.value = newValue
}

func (value *String) Set(newValue string) {
    value.str = newValue
}

func (value *Slice) Set(newValue []RootType) {
    value.slice = newValue
}

func (value *Map) Set(newValue map[string]RootType) {
    value.m = newValue
}
