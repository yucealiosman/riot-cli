package util

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

const (
	defaultCfgFileName = "config"
	localCfgFileName   = "local_config"
	cfgType            = "json"
)

type ConfigHandler interface {
	InitConfig() (err error)
	WriteLocalConfig(string, string)
	GetString(string) string
}

type ConfigHandle struct {
}

func getCurrentDir() string {
	_, fileName, _, _ := runtime.Caller(0)
	dir := filepath.Dir(fileName)

	return dir
}

// LoadConfig reads configuration from file or environment variables.
func (c *ConfigHandle) InitConfig() (err error) {
	var v = viper.New()
	dir := getCurrentDir()
	viper.AddConfigPath(dir)
	viper.SetConfigName(defaultCfgFileName)
	viper.SetConfigType(cfgType)

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.SetConfigName(localCfgFileName)
	viper.SetConfigType(cfgType)
	wd, _ := os.Getwd()
	viper.AddConfigPath(wd)

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

func (c *ConfigHandle) WriteLocalConfig(key string, value string) {
	viper.Set(key, value)
	viper.WriteConfig()
}

func (c *ConfigHandle) GetString(key string) (value string) {
	return viper.GetString(key)
}
