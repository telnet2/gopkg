package yamlutil

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// ReadYAMLFiles reads files in the given order and merge them.
// A later file will overwrite the previous yaml data.
func ReadYAMLFiles(v interface{}, files ...string) ([]*yaml.Node, error) {
	var nodes []*yaml.Node
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}

		root := &yaml.Node{}
		if err = yaml.Unmarshal(data, root); err != nil {
			return nodes, fmt.Errorf("fail to unmarshal %s: %v", f, err)
		}
		if err = root.Decode(v); err != nil {
			return nodes, fmt.Errorf("fail to decode %s: %v", f, err)
		}
		nodes = append(nodes, root)
	}
	return nodes, nil
}

// MergeYAML merges a list of given YAML contents.
func MergeYAML(v interface{}, contents ...string) (nodes []*yaml.Node, err error) {
	for i, data := range contents {
		root := &yaml.Node{}
		if err = yaml.Unmarshal([]byte(data), root); err != nil {
			return nodes, fmt.Errorf("fail to unmarshal %d: %v", i, err)
		}
		if err = root.Decode(v); err != nil {
			return nodes, fmt.Errorf("fail to decode %d: %v", i, err)
		}
		nodes = append(nodes, root)
	}
	return nodes, nil
}
