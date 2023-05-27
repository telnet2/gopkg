package units

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestByteSize(t *testing.T) {
	var bs ByteSize
	require.NoError(t, bs.UnmarshalJSON([]byte(`"1M"`)))
	require.Equal(t, ByteSize(1000000), bs)
	require.NoError(t, bs.UnmarshalJSON([]byte(`"1m"`)))
	require.Equal(t, ByteSize(1000000), bs)
	require.NoError(t, bs.UnmarshalJSON([]byte(`"1mb"`)))
	require.Equal(t, ByteSize(1000000), bs)
	require.NoError(t, bs.UnmarshalJSON([]byte(`"1mib"`)))
	require.Equal(t, ByteSize(1048576), bs)
}

func TestJSONByteSize(t *testing.T) {
	type conf struct {
		MaxSize ByteSize `json:"max_size" default:"5m"`
	}

	cc := conf{}
	require.NoError(t, json.Unmarshal([]byte(`{"max_size":"1m"}`), &cc))
	require.Equal(t, ByteSize(1_000_000), cc.MaxSize)
}

func TestYAMLByteSize(t *testing.T) {
	type conf struct {
		MaxSize ByteSize `yaml:"max_size"`
	}

	cc := conf{}
	require.NoError(t, yaml.Unmarshal([]byte(`{"max_size":"1m"}`), &cc))
	require.Equal(t, ByteSize(1_000_000), cc.MaxSize)
}
