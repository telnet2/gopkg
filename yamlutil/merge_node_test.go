package yamlutil

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeNode(t *testing.T) {
	yaml1 := `# yaml1
name: yaml1
mapping:
  foo: bar
  baz: baz
  submapping:
    sup: sub
sequence1: [1,2,3]
sequence: [1,2,3]
sequence_object:
  - key: 1
    value: a
  - key: 2
    value: b
`
	yaml2 := `# yaml1
name: yaml2
mapping:
  foo: BAR
  baz: BAZ
  submapping:
    sup: SUB
    sap: SAP
sequence1: [a,b]
sequence: [4,5,6,7]
sequence_object:
  - key: 1
    value: A
  - 2
  - key: 2
    value: B
yaml2: yaml2`
	n1, _ := ParseString(yaml1)
	n2, _ := ParseString(yaml2)

	base := NewDocument()
	require.NoError(t, Merge(&base, &n1))
	require.NoError(t, Merge(&base, &n2))

	expected := `# yaml1
name: yaml2
mapping:
    foo: BAR
    baz: BAZ
    submapping:
        sup: SUB
        sap: SAP
sequence1: [a, b, 3]
sequence: [4, 5, 6, 7]
sequence_object:
    - key: 1
      value: A
    - key: 2
      value: b
    - key: 2
      value: B
yaml2: yaml2
`
	w := strings.Builder{}
	Write(base, &w)
	require.Equal(t, expected, w.String())
}

func TestMergeFiles(t *testing.T) {
	y1, err := ParseFile("./testdata/1.yaml")
	require.NoError(t, err)
	y2, err := ParseFile("./testdata/2.yaml")
	require.NoError(t, err)

	require.NoError(t, Merge(&y1, &y2))

	y3 := strings.Builder{}
	Write(y1, &y3)
	require.Equal(t, mergedYaml, y3.String())
}

const mergedYaml = `template:
    properties: &properties
        gateway: "CG"
config:
    server: "production:80"
    properties:
        # use the properties defined above
        !!merge <<: *properties
        gateway: "CG"
        owner: "Michelle"
        flags: [4, 5, 6]
        region: "${REGION:-EU}"
__override__:
    - match_env:
        key: MY_POD_NAME
        value: "pod-1"
      yaml:
        config:
            server: "pod-1:9999"
features:
    enable_foo: true
`
