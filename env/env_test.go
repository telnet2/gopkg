package env

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	tests := []struct {
		key string
		val string
		def string
	}{
		{"key1", "val1", "def1"},
		{"key2", "", "def2"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%+v", test), func(t *testing.T) {
			t.Setenv(test.key, test.val)
			got := Get(test.key, test.def)
			if test.val != "" {
				assert.Equal(t, test.val, got, "must match value")
			} else {
				assert.Equal(t, test.def, got, "must match default")
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	tests := []struct {
		key string
		val int
	}{
		{"key1", 1},
		{"key2", 50000},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%+v", test), func(t *testing.T) {
			t.Setenv(test.key, strconv.FormatInt(int64(test.val), 10))
			got := GetInt(test.key, 0)
			assert.Equal(t, test.val, got, "must match value")
		})
	}
}

func TestGetIntDefault(t *testing.T) {
	tests := []struct {
		key string
		def int
	}{
		{"key1", 1},
		{"key2", 50000},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%+v", test), func(t *testing.T) {
			got := GetInt(test.key, test.def)
			assert.Equal(t, test.def, got, "must match default")
		})
	}
}

func TestGetIntDefaultNotIntVal(t *testing.T) {
	tests := []struct {
		key string
		val string
		def int
	}{
		{"key1", "val%1", 1},
		{"key2", "v.+2_", 50000},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%+v", test), func(t *testing.T) {
			t.Setenv(test.key, test.val)
			got := GetInt(test.key, test.def)
			assert.Equal(t, test.def, got, "must match default")
		})
	}
}

func TestIsSetReturnsTrue(t *testing.T) {
	tests := []struct {
		key string
		val string
	}{
		{"key1", "val1"},
		{"k%$2", "2"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%+v", test), func(t *testing.T) {
			t.Setenv(test.key, test.val)
			got := IsSet(test.key)
			assert.True(t, got)
		})
	}
}

func TestIsSetEmptyKeys(t *testing.T) {
	tests := []string{"emptykey1", "emptykey%+2", "k3"}

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			got := IsSet(test)
			assert.False(t, got)
		})
	}
}

func TestIsSetUnSetKeys(t *testing.T) {
	tests := []string{"unset_key1", "unset_key%+2", "unSetk3"}

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			got := IsSet(test)
			assert.False(t, got)
		})
	}
}
