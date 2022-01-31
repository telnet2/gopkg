package confutil

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v3"
)

const (
	UNKNOWN ConfigType = iota
	JSON
	YAML
)

type ConfigType int

var ErrUnknownConfigType = errors.New("Unknown config type")

// GuessConfigType guesses the config file type with its extension
func GuessConfigType(confname string) ConfigType {
	ext := strings.ToLower(filepath.Ext(confname))

	switch ext {
	case ".yaml":
		return YAML
	case ".yml":
		return YAML
	case ".json":
		return JSON
	}
	return UNKNOWN
}

// parseYamlFile parses the conffile with yaml.v3 package.
func parseYamlFile(out interface{}, conffile string) error {
	r, err := os.Open(conffile)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(data, out); err != nil {
		return err
	}
	return nil
}

// parseYamlFile parses the conffile with yaml.v3 package.
func parseJsonFile(out interface{}, conffile string) error {
	r, err := os.Open(conffile)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if err := jsoniter.Unmarshal(data, &out); err != nil {
		return err
	}
	return nil
}

// TryConfigParsing tries to parse a `conffile` file with the config type guessed from its name.
// If the config file cannot be guessed, returns ErrUnknownConfigType
func TryConfigParsing(out interface{}, conffile string) error {
	switch GuessConfigType(conffile) {
	case YAML:
		return parseYamlFile(out, conffile)
	case JSON:
		return parseJsonFile(out, conffile)
	}
	return ErrUnknownConfigType
}
