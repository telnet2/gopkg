package yamlutil

import (
	"errors"
	"fmt"
	"strings"

	"github.com/alecthomas/participle/v2"
	"gopkg.in/yaml.v3"
)

var (
	ErrNotFound = errors.New("no such node")
)

// Get returns the node at the given path.
// The format of the path is similar to JSON path but simpler.
// It works only with mapping nodes.
// e.g. node1.node2 returns node2 and node1 must be a mapping node.
func FindNode(n *yaml.Node, path string) (*yaml.Node, error) {
	pe, err := parsePathExpr(path)
	if err != nil {
		return nil, err
	}
	for _, s := range pe.Segments {
		n = lookupMappingNode(n, s.Name)
		if n == nil {
			return nil, fmt.Errorf("%w: %s", ErrNotFound, s.Name)
		}
		for _, p := range s.Props {
			n = lookupMappingNode(n, p)
			if n == nil {
				return nil, fmt.Errorf("%w: %s", ErrNotFound, p)
			}
		}
	}
	return n, nil
}

func lookupMappingNode(n *yaml.Node, name string) *yaml.Node {
	if n.Kind == yaml.DocumentNode {
		// If the node is a document node, find from the first child.
		return lookupMappingNode(n.Content[0], name)
	}
	if n.Kind != yaml.MappingNode {
		return nil
	}
	for i := 0; i < len(n.Content); i += 2 {
		if n.Content[i].Value == name {
			return n.Content[i+1]
		}
	}
	return nil
}

func parsePathExpr(path string) (*pathExpr, error) {
	return parser.ParseString("", path)
}

// pathExpr represents the AST for yaml path.
type pathExpr struct {
	Segments []segment `parser:"@@ ( '.' @@ )*"`
}

func (p *pathExpr) String() string {
	sb := strings.Builder{}
	for i, pp := range p.Segments {
		if i > 0 {
			sb.WriteString(".")
		}
		sb.WriteString(pp.String())
	}
	return sb.String()
}

// segment represents a segment of the path delimited by a dot.
// It may have a list of property names. We don't support integer index.
type segment struct {
	Name  string   `parser:"@Ident"`
	Props []string `parser:"('[' @(String|Char|RawString) ']')*"`
}

func (p *segment) String() string {
	sb := strings.Builder{}
	sb.WriteString(p.Name)
	for _, prop := range p.Props {
		sb.WriteString(fmt.Sprintf("[%s]", prop))
	}
	return sb.String()
}

var (
	parser = participle.MustBuild[pathExpr](participle.Unquote())
)
