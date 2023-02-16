package yamlutil

import "gopkg.in/yaml.v3"

// AsYAML marshals v into a string.
func AsYAML(v interface{}) string {
	yamlContent, _ := yaml.Marshal(v)
	return string(yamlContent)
}
