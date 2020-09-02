// This file is generated by gen_compare_funcs.sh.
// Don't edit manually.

package gozset

import "bytes"

type CompareFunc func(a, b interface{}) int

var (
	Byte CompareFunc = func(a, b interface{}) int {
		if a.(byte) == b.(byte) {
			return 0
		} else if a.(byte) < b.(byte) {
			return -1
		} else {
			return 1
		}
	}

	Float32 CompareFunc = func(a, b interface{}) int {
		if a.(float32) == b.(float32) {
			return 0
		} else if a.(float32) < b.(float32) {
			return -1
		} else {
			return 1
		}
	}

	Float64 CompareFunc = func(a, b interface{}) int {
		if a.(float64) == b.(float64) {
			return 0
		} else if a.(float64) < b.(float64) {
			return -1
		} else {
			return 1
		}
	}

	Int CompareFunc = func(a, b interface{}) int {
		if a.(int) == b.(int) {
			return 0
		} else if a.(int) < b.(int) {
			return -1
		} else {
			return 1
		}
	}

	Int16 CompareFunc = func(a, b interface{}) int {
		if a.(int16) == b.(int16) {
			return 0
		} else if a.(int16) < b.(int16) {
			return -1
		} else {
			return 1
		}
	}

	Int32 CompareFunc = func(a, b interface{}) int {
		if a.(int32) == b.(int32) {
			return 0
		} else if a.(int32) < b.(int32) {
			return -1
		} else {
			return 1
		}
	}

	Int64 CompareFunc = func(a, b interface{}) int {
		if a.(int64) == b.(int64) {
			return 0
		} else if a.(int64) < b.(int64) {
			return -1
		} else {
			return 1
		}
	}

	Int8 CompareFunc = func(a, b interface{}) int {
		if a.(int8) == b.(int8) {
			return 0
		} else if a.(int8) < b.(int8) {
			return -1
		} else {
			return 1
		}
	}

	Rune CompareFunc = func(a, b interface{}) int {
		if a.(rune) == b.(rune) {
			return 0
		} else if a.(rune) < b.(rune) {
			return -1
		} else {
			return 1
		}
	}

	String CompareFunc = func(a, b interface{}) int {
		if a.(string) == b.(string) {
			return 0
		} else if a.(string) < b.(string) {
			return -1
		} else {
			return 1
		}
	}

	Uint CompareFunc = func(a, b interface{}) int {
		if a.(uint) == b.(uint) {
			return 0
		} else if a.(uint) < b.(uint) {
			return -1
		} else {
			return 1
		}
	}

	Uint16 CompareFunc = func(a, b interface{}) int {
		if a.(uint16) == b.(uint16) {
			return 0
		} else if a.(uint16) < b.(uint16) {
			return -1
		} else {
			return 1
		}
	}

	Uint32 CompareFunc = func(a, b interface{}) int {
		if a.(uint32) == b.(uint32) {
			return 0
		} else if a.(uint32) < b.(uint32) {
			return -1
		} else {
			return 1
		}
	}

	Uint64 CompareFunc = func(a, b interface{}) int {
		if a.(uint64) == b.(uint64) {
			return 0
		} else if a.(uint64) < b.(uint64) {
			return -1
		} else {
			return 1
		}
	}

	Uint8 CompareFunc = func(a, b interface{}) int {
		if a.(uint8) == b.(uint8) {
			return 0
		} else if a.(uint8) < b.(uint8) {
			return -1
		} else {
			return 1
		}
	}

	Uintptr CompareFunc = func(a, b interface{}) int {
		if a.(uintptr) == b.(uintptr) {
			return 0
		} else if a.(uintptr) < b.(uintptr) {
			return -1
		} else {
			return 1
		}
	}

	// the type []byte.
	Bytes CompareFunc = func(a, b interface{}) int {
		return bytes.Compare(a.([]byte), b.([]byte))
	}
)