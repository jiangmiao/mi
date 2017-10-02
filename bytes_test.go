package mi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Foo struct {
	A int64
	B int64
}

func Test(tt *testing.T) {
	att := assert.New(tt)
	ok := att.NoError
	eq := att.Equal
	_ = ok

	eq(Start([2]int64{1, 2}), Start([]int64{1, 2}))
	eq(Start([2]int64{1, 2}).Hex().String(), "01000000000000000200000000000000")
	var foo = Foo{1, 2}
	eq(Start(foo).Hex().String(), "01000000000000000200000000000000")
	eq(Start(&foo).Hex().String(), "01000000000000000200000000000000")
}
