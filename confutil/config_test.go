package confutil

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"github.com/telnet2/gopkg/ioutil"
)

type TestMe struct {
	suite.Suite
}

func (tm *TestMe) TestTryConfigParsingYAML() {
	remove := ioutil.MustWriteFile("yamlfile.yaml", `
config:
    field: value
    array:
        - hello
        - world
    object:
        inner:
            field: value
`)
	defer remove()

	data := map[string]interface{}{}
	tm.NoError(TryConfigParsing(&data, "yamlfile.yaml"))
	tm.Equal("hello", data["config"].(map[string]interface{})["array"].([]interface{})[0])
}

func (tm *TestMe) TestTryConfigParsingJSON() {
	remove := ioutil.MustWriteFile("jsonfile.JSON", `{
		"config": {
			"field": "value",
			"array": ["hello", "world"],
			"object": { 
				"inner": {
					"field": "value"
				}
			}
		}
	}`)
	defer remove()

	conf := viper.New()
	conf.SetConfigFile("jsonfile.json")
	err := conf.ReadInConfig()
	tm.NoError(err)

	data := map[string]interface{}{}
	tm.NoError(TryConfigParsing(&data, "jsonfile.json"))
	tm.Equal("hello", data["config"].(map[string]interface{})["array"].([]interface{})[0])
}

func TestConfUtil(t *testing.T) {
	tm := new(TestMe)
	suite.Run(t, tm)
}
