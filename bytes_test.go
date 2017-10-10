package mi

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Foo struct {
	A int64
	B int64
}

type Bar struct {
}

func (b Bar) Bytes() Bytes {
	return Bytes("hello")
}

func TestBytes(tt *testing.T) {
	att := assert.New(tt)
	ok := att.NoError
	eq := att.Equal
	_ = ok

	var foo = Foo{1, 2}

	var out = new(bytes.Buffer)
	binary.Write(out, binary.LittleEndian, []int64{1, 2})
	var expected = B(out.Bytes())

	eq(B([2]int64{1, 2}), B([]int64{1, 2}))
	eq(B([2]int64{1, 2}), expected)
	eq(B(foo), expected)
	eq(B(&foo), expected)
	eq(B(Bar{}), B("hello"))
	eq(B("abc"), B("a").Concat("b", "c"))
}
