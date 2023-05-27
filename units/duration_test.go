package units

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/creasty/defaults"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestSetDuration(t *testing.T) {
	var v = Duration(time.Duration(1 * time.Second))
	require.Equal(t, time.Duration(1*time.Second), v.D())
}

func TestJsonDuration(t *testing.T) {
	var timeConf struct {
		Interval Duration
	}

	require.NoError(t, json.Unmarshal([]byte(`{
		"Interval": "3m"
	}`), &timeConf))

	require.Equal(t, time.Minute*3, timeConf.Interval.D())
}

func TestYamlDuration(t *testing.T) {
	var timeConf struct {
		Interval Duration `yaml:"interval"`
	}

	require.NoError(t, yaml.Unmarshal([]byte(`interval: "3m"`), &timeConf))
	require.Equal(t, time.Minute*3, timeConf.Interval.D())
}

func TestDefaultDuration(t *testing.T) {
	var timeConf struct {
		Interval Duration `default:"3m"`
	}

	require.EqualValues(t, 0, timeConf.Interval)
	defaults.MustSet(&timeConf)
	require.EqualValues(t, time.Minute*3, timeConf.Interval)
}
