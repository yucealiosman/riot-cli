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

type ConfigHandler interface {
	InitConfig() (err error)
	WriteLocalConfigFile(string, string)
	GetString(string) string
}

type ConfigHandle struct {
}

// LoadConfig reads configuration from file or environment variables.
func (c *ConfigHandle) InitConfig() (err error) {
	var v = viper.New()
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

	err = viper.Unmarshal(&v)
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

func (c *ConfigHandle) WriteLocalConfigFile(key string, value string) {
	wd, _ := os.Getwd()

	viper.AddConfigPath(wd)
	viper.SetConfigName(localCfgFileName)
	viper.SetConfigType(cfgType)

	viper.Set(key, value)
	viper.WriteConfigAs(localCfgFileName + "." + cfgType)
}

func (c *ConfigHandle) GetString(key string) (value string) {
	return viper.GetString(key)
}
