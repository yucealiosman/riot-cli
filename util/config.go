package util

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	defaultCfgFileName = "config"
	localCfgFileName   = "local_config"
	cfgType            = "json"
)

type Config struct {
	RiotUrl   string `mapstructure:"riotUrl"`
	RiotToken string `mapstructure:"riotToken"`
	Region    string `mapstructure:"region"`
}

// LoadConfig reads configuration from file or environment variables.
func InitConfig() (config Config, err error) {
	wd, _ := os.Getwd()

	viper.AddConfigPath(wd)
	viper.SetConfigName(localCfgFileName)
	viper.SetConfigType(cfgType)

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.SetConfigName(defaultCfgFileName)
	viper.SetConfigType(cfgType)

	err = createLocalConfigFile(filepath.Join(wd, localCfgFileName+"."+cfgType))
	if err != nil {
		return
	}

	viper.MergeInConfig()

	err = viper.Unmarshal(&config)
	return
}

func createLocalConfigFile(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			err := viper.SafeWriteConfigAs(path)
			return err
		} else {
			return err
		}
	}
	return nil
}

func WriteLocalConfigFile(key string, value string) {
	wd, _ := os.Getwd()

	viper.AddConfigPath(wd)
	viper.SetConfigName(localCfgFileName)
	viper.SetConfigType(cfgType)

	viper.Set(key, value)
	viper.WriteConfigAs(localCfgFileName + "." + cfgType)
}
