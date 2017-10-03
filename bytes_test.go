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

func Test(tt *testing.T) {
	att := assert.New(tt)
	ok := att.NoError
	eq := att.Equal
	_ = ok

	var foo = Foo{1, 2}

	var out = new(bytes.Buffer)
	binary.Write(out, binary.LittleEndian, []int64{1, 2})
	var expected = Bytes(out.Bytes())

	eq(Start([2]int64{1, 2}), Start([]int64{1, 2}))
	eq(Start([2]int64{1, 2}), expected)
	eq(Start(foo), expected)
	eq(Start(&foo), expected)
	eq(Start(Bar{}), Bytes("hello"))
}
