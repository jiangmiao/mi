package mi

import (
	"bytes"
	"encoding/json"
)

type JSONBytes []byte

func (v JSONBytes) Beauty() []byte {
	var out = new(bytes.Buffer)
	json.Indent(out, v, "", "\t")
	return out.Bytes()
}

func JSON(v interface{}) JSONBytes {
	r, _ := json.Marshal(v)
	return r
}

func JSONBeauty(v interface{}) []byte {
	return JSON(v).Beauty()
}
