package yamlutil

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// ParseFile parses a YAML file and returns the root node.
func ParseFile(f string) (yaml.Node, error) {
	r, err := os.ReadFile(f)
	if err != nil {
		return yaml.Node{}, err
	}
	return Parse(bytes.NewReader(r))
}

// ParseString parses a YAML document from a string.
func ParseString(s string) (yaml.Node, error) {
	return Parse(strings.NewReader(s))
}

// Parse is helper function that decodes yaml from io.Reader
func Parse(reader io.Reader) (yaml.Node, error) {
	d := yaml.NewDecoder(reader)
	var n yaml.Node
	err := d.Decode(&n)
	if err != nil {
		return n, err
	}
	return n, nil
}

// Write is helper function that writes yaml tree to io.Writer
func Write(doc yaml.Node, writer io.Writer) error {
	e := yaml.NewEncoder(writer)
	err := e.Encode(&doc)
	if err != nil {
		return err
	}
	return e.Close()
}

// NewDocument creates a new document node based on mapping.
func NewDocument() yaml.Node {
	return yaml.Node{Kind: yaml.DocumentNode, Content: []*yaml.Node{{Kind: yaml.MappingNode}}}
}

// Merge tries to merge two yaml nodes a and b into a.
// It merges the nodes with the same kind.
func Merge(a, b *yaml.Node) error {
	// Document nodes are simple, just merge first children
	if a.Kind == yaml.DocumentNode && b.Kind == yaml.DocumentNode {
		return Merge(a.Content[0], b.Content[0])
	}

	// Scalar node is overwritten.
	if a.Kind == yaml.ScalarNode && b.Kind == yaml.ScalarNode {
		*a = *b
		return nil
	}

	// Sequences are trickier. We need to check whether the index is in the array.
	// If it is not or if it is negative, we append new element, otherwise merge recursively
	// Little hack, we use spare variable in yaml.Node to signalize on which position in sequence we shall merge it
	if a.Kind == yaml.SequenceNode && b.Kind == yaml.SequenceNode {
		return MergeSequenceNode(a, b)
	}

	if a.Kind == yaml.MappingNode && b.Kind == yaml.MappingNode {
		return MergeMappingNode(a, b)
	}

	return fmt.Errorf("cannot merge incompatible nodes: [%s:(%s) %s:(%s)]", kindStr(a.Kind), a.Value, kindStr(b.Kind), b.Value)
}

// MergeSequenceNode attempts to merge two sequence nodes, a and b, into node a.
// Corresponding items are merged recursively. If the corresponding item is a scalar value,
// the item in 'a' is replaced with the corresponding item in 'b'.
// Any items in 'b' that come after the last item in 'a' are appended to 'a'.
func MergeSequenceNode(a, b *yaml.Node) error {
	if a.Kind == yaml.SequenceNode && b.Kind == yaml.SequenceNode {
		for i := 0; i < len(a.Content); i++ {
			if len(b.Content) > i {
				_ = Merge(a.Content[i], b.Content[i])
			}
		}
		if len(a.Content) < len(b.Content) {
			a.Content = append(a.Content, b.Content[len(a.Content):]...)
		}
		return nil
	}
	return fmt.Errorf("cannot merge non-sequence nodes [%s %s]", kindStr(a.Kind), kindStr(b.Kind))
}

// MergeMappingNode tries to merge two yaml nodes a and b into a.
// If it encounters incompatible nodes along the way, it returns error.
// We need to go through all content of a map and check whether the key is there.
// If yes, we merge it recursively, otherwise just append new key
func MergeMappingNode(a, b *yaml.Node) error {
	if a.Kind == yaml.MappingNode && b.Kind == yaml.MappingNode {
		for i := 0; i < len(b.Content); i += 2 {
			bval := b.Content[i].Value
			found := false
			for j := 0; j < len(a.Content); j += 2 {
				aval := a.Content[j].Value
				if bval == aval {
					err := Merge(a.Content[j+1], b.Content[i+1])
					if err != nil {
						return err
					}
					found = true
				}
			}
			if !found {
				a.Content = append(a.Content, b.Content[i], b.Content[i+1])
			}
		}
		return nil
	}
	return fmt.Errorf("cannot merge non-mapping nodes [%s %s]", kindStr(a.Kind), kindStr(b.Kind))
}

func kindStr(kind yaml.Kind) string {
	switch kind {
	case yaml.DocumentNode:
		return "document"
	case yaml.MappingNode:
		return "mapping"
	case yaml.SequenceNode:
		return "sequence"
	case yaml.AliasNode:
		return "alias"
	case yaml.ScalarNode:
		return "scalar"
	default:
		return fmt.Sprintf("unknown:%d", kind)
	}
}
