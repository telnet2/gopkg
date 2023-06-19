package yamlutil

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestPathParser(t *testing.T) {
	e, err := parsePathExpr(`hello["x.y.z"]. a. b.c`)
	require.NoError(t, err)
	require.Equal(t, `hello[x.y.z].a.b.c`, e.String())
}

func TestFindNode(t *testing.T) {
	y1 := `# test YAML
mapping:
  child:
    name: "child"
    "x.y.z": "value"
    integer: 1`

	root, err := ParseString(y1)
	require.NoError(t, err)
	require.NotNil(t, root)

	mapping, err := FindNode(&root, "mapping")
	require.NoError(t, err)
	require.NotNil(t, mapping)
	require.Equal(t, yaml.MappingNode, mapping.Kind)

	child, err := FindNode(&root, "mapping.child")
	require.NoError(t, err)
	require.NotNil(t, child)
	require.Equal(t, yaml.MappingNode, child.Kind)

	child, err = FindNode(&root, "mapping.not_exist")
	require.Error(t, err)
	require.ErrorIs(t, err, ErrNotFound)
	require.Nil(t, child)

	child, err = FindNode(&root, `mapping["child"]`)
	require.NoError(t, err)
	require.NotNil(t, child)

	child, err = FindNode(&root, `mapping["child"]["x.y.z"]`)
	require.NoError(t, err)
	require.NotNil(t, child)

	var value string
	require.NoError(t, child.Decode(&value))
	require.Equal(t, "value", value)

	child, err = FindNode(&root, `mapping["child"].integer`)
	require.NoError(t, err)
	require.NotNil(t, child)

	var i int
	require.NoError(t, child.Decode(&i))
	require.Equal(t, 1, i)
}
