package units

import (
	"encoding/json"
	"fmt"

	"github.com/dustin/go-humanize"
	"gopkg.in/yaml.v3"
)

// ByteSize is an uint64 with unmarshaling from human readable string.
// It doesn't support the default & validate tag 😰.
type ByteSize uint64

var (
	_i = ByteSize(0)
	_  = json.Unmarshaler(&_i)
	_  = yaml.Unmarshaler(&_i)
)

func (b *ByteSize) UnmarshalJSON(jz []byte) (err error) {
	var v interface{}
	if err = json.Unmarshal(jz, &v); err != nil {
		return err
	}
	return b.fromInterface(v)
}

func (b *ByteSize) UnmarshalYAML(n *yaml.Node) (err error) {
	var v interface{}
	if err = n.Decode(&v); err != nil {
		return err
	}
	return b.fromInterface(v)
}

func (b *ByteSize) fromInterface(v interface{}) (err error) {
	switch vv := v.(type) {
	case string:
		bb, err := humanize.ParseBytes(vv)
		if err != nil {
			return err
		}
		*b = ByteSize(bb)
	case uint64:
		*b = ByteSize(vv)
	case uint32:
		*b = ByteSize(vv)
	case int32:
		*b = ByteSize(vv)
	case int64:
		*b = ByteSize(vv)
	case int:
		*b = ByteSize(vv)
	case float64:
		*b = ByteSize(vv)
	case float32:
		*b = ByteSize(vv)
	default:
		return fmt.Errorf("invalid byte format: %v", v)
	}
	return nil
}
