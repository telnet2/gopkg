package yamlutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type File struct {
	Input  string            `yaml:"input"`
	Output []string          `yaml:"output"`
	Maps   map[string]string `yaml:"maps"`
}

type FileConfig struct {
	File1 *File `yaml:"file1"`
	File2 *File `yaml:"file2"`
}

func TestReadYAMLFiles(t *testing.T) {
	f := FileConfig{}
	err := ReadYAMLFiles(&f, "./1.yaml", "./2.yaml")
	require.NoError(t, err)
	require.Equal(t, "file1", f.File1.Input)
	content := `file1:
    input: file1
    output:
        - array
        - replaced
    maps:
        hello: world
        key: value
        mapkey: merged
file2:
    input: hello
    output:
        - a
        - b
        - c
    maps:
        hello: world
        key: value
`
	require.Equal(t, content, AsYAML(&f))
}

func TestMergeYAML(t *testing.T) {
	f := FileConfig{}
	err := MergeYAML(&f, `#
file1:
  input: "file1"`,
		`#
file1:
  input: "file2"
`)
	require.NoError(t, err)
	require.Equal(t, "file2", f.File1.Input)
}
