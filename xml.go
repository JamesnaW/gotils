package gotils

import (
  "encoding/xml"
)
// XMLDecode - send v as address `&v`
func XMLDecode(data []byte, v interface{}) (interface{}, error) {
	err := xml.Unmarshal(data, v)
	if err != nil {
		return v, err
	}
	return v, nil
}
