package mi

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
	"reflect"
	"unsafe"
)

type Byteser interface {
	Bytes() Bytes
}

type Bytes []byte

func BLE(v interface{}) Bytes {
	return B(v).Reverse()
}
func B(v interface{}) Bytes {
	switch r := v.(type) {
	case Byteser:
		return r.Bytes()
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

	switch rv.Kind() {
	case reflect.Slice:
		n = rv.Len()
		if n == 0 {
			return []byte{}
		}
		base = rv.Index(0)
		addr = base.UnsafeAddr()
	case reflect.Ptr:
		n = 1
		base = rv.Elem()
		addr = base.UnsafeAddr()
	default:
		n = 1
		base = rv
		ptrs := *(*[2]uintptr)(unsafe.Pointer(&v))
		addr = ptrs[1]
	}
	size = n * int(base.Type().Size())
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: addr,
		Len:  size,
		Cap:  size,
	}))
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
func (b Bytes) Clone() Bytes {
	r := make(Bytes, len(b))
	copy(r, b)
	return r
}
func (b Bytes) Save(r *Bytes) Bytes {
	*r = b.Clone()
	return b
}
func (b Bytes) WriteTo(w io.Writer) (int, error) {
	return w.Write(b)
}
func (b Bytes) MustWriteTo(w io.Writer) Bytes {
	_, err := b.WriteTo(w)
	if err != nil {
		panic(err)
	}
	return b
}
func (b Bytes) Len() int {
	return len(b)
}
func (b Bytes) Reader() io.Reader {
	return bytes.NewBuffer(b)
}
func (b Bytes) Concat(vs ...interface{}) Bytes {
	var r Bytes
	r = b
	for _, v := range vs {
		r = append(r, B(v)...)
	}
	return r
}

// Crypto
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

// Encoding
func (b Bytes) Hex() Bytes {
	r := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(r, b)
	return r
}
