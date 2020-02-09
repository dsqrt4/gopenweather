package current

import (
	"encoding/json"
	"io"
)

type Decoder func(data io.Reader, v interface{}) error

var DefaultDecoder Decoder = func(data io.Reader, v interface{}) error {
	return json.NewDecoder(data).Decode(v)
}
