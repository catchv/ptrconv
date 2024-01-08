package ptrconv

// the set of all unsigned  8-bit integers (0 to 255)
func UInt8Ptr(i uint8) *uint8 {
	return &i
}

// the set of all unsigned 16-bit integers (0 to 65535)
func UInt16Ptr(i uint16) *uint16 {
	return &i
}

// the set of all unsigned 32-bit integers (0 to 4294967295)
func UInt32Ptr(i uint32) *uint32 {
	return &i
}

// the set of all unsigned 64-bit integers (0 to 18446744073709551615)
func UInt64Ptr(i uint64) *uint64 {
	return &i
}

// the set of all signed  8-bit integers (-128 to 127)
func Int8Ptr(i int8) *int8 {
	return &i

}

// the set of all signed 16-bit integers (-32768 to 32767)
func Int16Ptr(i int16) *int16 {
	return &i
}

// the set of all signed 32-bit integers (-2147483648 to 2147483647)
func Int32Ptr(i int32) *int32 {
	return &i
}

// the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
func Int64Ptr(i int64) *int64 {
	return &i
}

// the set of all IEEE-754 32-bit floating-point numbers
func Float32Ptr(f float32) *float32 {
	return &f
}

// the set of all IEEE-754 64-bit floating-point numbers
func Float64Ptr(f float64) *float64 {
	return &f
}

// the set of all complex numbers with float32 real and imaginary parts
func Complex64Ptr(c complex64) *complex64 {
	return &c
}

// the set of all complex numbers with float64 real and imaginary parts
func Complex128Ptr(c complex128) *complex128 {
	return &c
}

// the set of all complex numbers with float64 real and imaginary parts
func BytePtr(b byte) *byte {
	return &b
}

// either 32 or 64 bits
func IntPtr(i int) *int {
	return &i
}

// same size as uint
func UIntPtr(i uint) *uint {
	return &i
}

// string
func StringPtr(str string) *string {
	return &str
}

// Go generic Auto Pointer converter
func AutoPtr[T any](data T) *T {
	return &data
}
