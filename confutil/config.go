package confutil

import (
	"github.com/spf13/viper"
)

// TryConfigParsing tries to parse a `conffile` file using the viper package.
func TryConfigParsing(out interface{}, conffile string) error {
	conf := viper.New()
	conf.SetConfigFile(conffile)
	if err := conf.ReadInConfig(); err != nil {
		return err
	}
	return conf.Unmarshal(&out)
}
