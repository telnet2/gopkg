package yamlutil

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsEmpty(t *testing.T) {
	assert.True(t, IsEmpty(""))
	assert.True(t, IsEmpty(0))
	assert.True(t, IsEmpty(int64(0)))
	assert.True(t, IsEmpty(0.0))
	assert.True(t, IsEmpty(nil))
	assert.True(t, IsEmpty(interface{}(nil)))
	assert.True(t, IsEmpty(map[string]string{}))
	assert.True(t, IsEmpty([]string{}))
}

func TestRemoveEmpty(t *testing.T) {
	x := map[string]interface{}{}
	err := json.Unmarshal([]byte(`{
		"string": "",
		"integer": 0,
		"number": 0.0,
		"null": null,
		"array": [],
		"object": {
			"nested": {
				"object": ""
			}
		},
		"object_array": [
			{ "type": "", "value": "" }
		]
	}`), &x)
	RemoveEmptyFromMap(x)
	assert.NoError(t, err)
	assert.Empty(t, x)

	xx, _ := json.Marshal(x)
	require.JSONEq(t, `{}`, string(xx))

	var (
		nilptr *jsonField
		ptr    = &jsonField{}
	)

	y := map[string]interface{}{
		"string":                map[string]string{},
		"int":                   map[string]int{},
		"mapOfstringArray":      map[string][]string{"arr": {}},
		"arrayOfStrings":        []string{},
		"nil":                   nil,
		"object":                jsonField{},
		"nilpointer":            nilptr,
		"pointer":               ptr,
		"pointerToMapToPointer": &map[string]*map[string]string{"map": {"empty": ""}},
	}
	RemoveEmptyFromMap(y)
	assert.NoError(t, err)
	assert.Empty(t, y)

	z := map[string]map[string]string{
		"x": {"y": ""},
	}
	RemoveEmptyFromMap(z)
	assert.NoError(t, err)
	assert.Empty(t, z)

}

func TestMarshalJSONOmitEmpty(t *testing.T) {
	x := jsonField{
		Description: "root",
		Type:        "object",
		Properties: map[string]*jsonField{
			"emptyObjcet": {},
		},
	}

	jzon, _ := MarshalJSONOmitEmptyIndented(x)
	assert.JSONEq(t, `{"description":"root", "type":"object"}`, string(jzon))
}

func TestRemoveEmpty2(t *testing.T) {
	x := map[string]interface{}{}
	err := json.Unmarshal([]byte(`{
		"string": "",
		"integer": 0,
		"number": 0.0,
		"null": null,
		"array": [],
		"object": {
			"nested": {
				"object": ""
			}
		},
		"object_array": [
			{ "type": "", "value": "" },
			{ "type": "string" }
		],
		"non-empty": {
			"type": "string"
		}
	}`), &x)
	RemoveEmptyFromMap(x)
	assert.NoError(t, err)
	assert.NotEmpty(t, x)

	xx, _ := json.Marshal(x)
	require.JSONEq(t, `{"object_array":[{"type":"string"}],"non-empty":{"type":"string"}}`, string(xx))
}

type jsonField struct {
	Type        string                `json:"type,omitempty" yaml:"type,omitempty"`
	Description string                `json:"description,omitempty" yaml:"description,omitempty"`
	Properties  map[string]*jsonField `json:"properties,omitempty" yaml:"properties,omitempty"`
}
