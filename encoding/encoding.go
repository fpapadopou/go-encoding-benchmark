package encoding

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

type object struct {
	StringProp1 string     `json:"string_prop_1"`
	StringProp2 string     `json:"string_prop_2"`
	List        []embedded `json:"list"`
}

type embedded struct {
	StringProp string `json:"string_prop"`
	IntProp    int64  `json:"int_prop"`
}

func (o *object) stdJSONEncode() ([]byte, error) {
	return json.Marshal(o)
}

func (o *object) stdJSONDecode(e []byte) error {
	return json.Unmarshal(e, o)
}

func (o *object) stdGobEncode() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(o)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (o *object) stdGobDecode(e []byte) error {
	buf := bytes.NewBuffer(e)
	dec := gob.NewDecoder(buf)
	return dec.Decode(o)
}
