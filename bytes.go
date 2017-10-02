package mi

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
	"reflect"
	"unsafe"
)

type Bytes []byte

func StartLE(v interface{}) Bytes {
	return Start(v).Reverse()
}
func Start(v interface{}) Bytes {
	switch r := v.(type) {
	case string:
		return Bytes(r)
	case Bytes:
		return r
	case []byte:
		return r
	}
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
		addr = base.UnsafeAddr()
	} else {
		n = 1
		if rv.Kind() == reflect.Ptr {
			base = rv.Elem()
			addr = base.UnsafeAddr()
		} else {
			base = rv
			ptrs := *(*[2]uintptr)(unsafe.Pointer(&v))
			addr = ptrs[1]
		}
	}
	size = n * int(base.Type().Size())
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: addr,
		Len:  size,
		Cap:  size,
	}))
}
func (b Bytes) Hex() Bytes {
	r := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(r, b)
	return r
}
func (b Bytes) MD5() Bytes {
	r := md5.Sum(b)
	return r[:]
}
func (b Bytes) SHA256() Bytes {
	r := sha256.Sum256(b)
	return r[:]
}
func (b Bytes) SHA512() Bytes {
	r := sha512.Sum512(b)
	return r[:]
}
func (b Bytes) String() String {
	return String(b)
}
func (b Bytes) Reverse() Bytes {
	var l = len(b)
	var r = make(Bytes, l)
	for i := 0; i < l; i++ {
		r[l-i-1] = b[i]
	}
	return r
}
func (b Bytes) WriteTo(w io.Writer) (int, error) {
	return w.Write(b)
}
func (b Bytes) Clone() Bytes {
	r := make(Bytes, len(b))
	copy(r, b)
	return r
}
func (b Bytes) Save(r *Bytes) Bytes {
	*r = b.Clone()
	return b
}
func (b Bytes) MustWriteTo(w io.Writer) Bytes {
	_, err := b.WriteTo(w)
	if err != nil {
		panic(err)
	}
	return b
}
