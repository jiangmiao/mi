package mi

import (
	"reflect"
	"unsafe"
)

func Bytes(v interface{}) []byte {
	src := BytesRef(v)
	dest := make([]byte, len(src))
	copy(dest, src)
	return dest
}

func BytesRef(v interface{}) []byte {
	var addr uintptr
	var size, n int
	var base reflect.Value
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Slice {
		n = rv.Len()
		if n == 0 {
			return []byte{}
		}
		base = rv.Index(0)
	} else {
		n = 1
		base = rv.Elem()
	}
	addr = base.UnsafeAddr()
	size = n * int(base.Type().Size())
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: addr,
		Len:  size,
		Cap:  size,
	}))
}

func BytesCopy(v interface{}, src []byte) {
	copy(BytesRef(v), src)
}
