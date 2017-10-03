package mi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(tt *testing.T) {
	att := assert.New(tt)
	ok := att.NoError
	eq := att.Equal
	_ = ok

	eq(String("f_o_o"), String("FOO").Underscorize())
	eq(String("fOO"), String("f_o_o").Camelize())
	eq(String("fOO_"), String("f_o_o_").Camelize())
	eq(String("FOO"), String("F_o_o").Camelize())
	eq(String("FOO"), String("f_o_o").Camelize(true))
}
