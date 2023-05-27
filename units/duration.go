package units

import (
	"encoding/json"
	"time"

	"gopkg.in/yaml.v3"
)

type Duration time.Duration

// UnmarshalJSON handles a string representation of time.Duration.
// e.g., 3s, 2m
func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	// depending on the type of `v`, parse differently.
	switch value := v.(type) {
	case float64:
		*d = Duration(value)
	case string:
		if _d, err := time.ParseDuration(value); err != nil {
			return err
		} else {
			*d = Duration(_d)
		}
	}
	return nil
}

// UnmarshalYAML unmarshals a string node as time.Duration
func (d *Duration) UnmarshalYAML(value *yaml.Node) error {
	// handle it as if it was time.Duration
	return value.Decode((*time.Duration)(d))
}

// D returns time.Duration conversion of this.
func (d *Duration) D() time.Duration {
	return time.Duration(*d)
}

// ParseDuration parses a string format of the duration
func ParseDuration(s string) (Duration, error) {
	d, err := time.ParseDuration(s)
	return Duration(d), err
}
