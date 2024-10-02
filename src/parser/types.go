package parser

import (
	"unsafe"
)

// DataType represents the possible data types in Kvexium
type DataType int

const (
	U8 DataType = iota
	U16
	U32
	U64
	I8
	I16
	I32
	I64
	F32
	F64
	F80
	Auto
	NAuto
	Str
	Char
	C64
	C128
	Bool
)

// TypeInfo holds information about a data type
type TypeInfo struct {
	Name  string
	Size  uintptr // Size in bytes
	Alloc func() unsafe.Pointer
}

// allocMemory is a generic memory allocation function
func allocMemory(size uintptr) unsafe.Pointer {
	return unsafe.Pointer(&make([]byte, size)[0])
}

// allocString is a special case for string allocation
func allocString() unsafe.Pointer {
	s := new(string)
	return unsafe.Pointer(s)
}

var typeInfoMap = map[DataType]TypeInfo{
	U8:    {"u8", 1, func() unsafe.Pointer { return allocMemory(1) }},
	U16:   {"u16", 2, func() unsafe.Pointer { return allocMemory(2) }},
	U32:   {"u32", 4, func() unsafe.Pointer { return allocMemory(4) }},
	U64:   {"u64", 8, func() unsafe.Pointer { return allocMemory(8) }},
	I8:    {"i8", 1, func() unsafe.Pointer { return allocMemory(1) }},
	I16:   {"i16", 2, func() unsafe.Pointer { return allocMemory(2) }},
	I32:   {"i32", 4, func() unsafe.Pointer { return allocMemory(4) }},
	I64:   {"i64", 8, func() unsafe.Pointer { return allocMemory(8) }},
	F32:   {"f32", 4, func() unsafe.Pointer { return allocMemory(4) }},
	F64:   {"f64", 8, func() unsafe.Pointer { return allocMemory(8) }},
	F80:   {"f80", 10, func() unsafe.Pointer { return allocMemory(10) }},
	Auto:  {"auto", 0, nil},   // Size determined at runtime
	NAuto: {"n_auto", 0, nil}, // Size determined at runtime
	Str:   {"str", unsafe.Sizeof(""), allocString},
	Char:  {"char", 1, func() unsafe.Pointer { return allocMemory(1) }},
	C64:   {"c64", 8, func() unsafe.Pointer { return allocMemory(8) }},
	C128:  {"c128", 16, func() unsafe.Pointer { return allocMemory(16) }},
	Bool:  {"bool", 1, func() unsafe.Pointer { return allocMemory(1) }},
}

// GetTypeInfo returns the TypeInfo for a given DataType
func GetTypeInfo(dt DataType) TypeInfo {
	return typeInfoMap[dt]
}

// IsUnsigned checks if the type is an unsigned integer
func IsUnsigned(dt DataType) bool {
	return dt >= U8 && dt <= U64
}

// IsFloat checks if the type is a floating-point number
func IsFloat(dt DataType) bool {
	return dt >= F32 && dt <= F80
}

// IsComplex checks if the type is a complex number
func IsComplex(dt DataType) bool {
	return dt == C64 || dt == C128
}

// FreeMemory is a generic memory deallocation function
func FreeMemory(p unsafe.Pointer) {
	// In Go, the garbage collector handles deallocation
	// This function is a placeholder for explicit deallocation if needed
}
