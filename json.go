package mi

import (
	"bytes"
	"encoding/json"

	jsonitor "github.com/json-iterator/go"
)

type JSONBytes []byte

func (v JSONBytes) Beauty() Bytes {
	var out = new(bytes.Buffer)
	json.Indent(out, v, "", "\t")
	return out.Bytes()
}

func JSON(v interface{}) JSONBytes {
	r, err := jsonitor.Marshal(v)
	if err != nil {
		panic(err)
	}
	return r
}

func JSONBeauty(v interface{}) Bytes {
	return JSON(v).Beauty()
}
