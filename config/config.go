package config

import (
	"io/ioutil"
	"github.com/double-chosen-will-server/logutil"
	"github.com/juju/errors"
	"encoding/json"
)

//Config contain configuration options
type Config struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`
	Log         `json:"log"`
}

//Log configuration options
type Log struct {
	//Log level
	Level            string                `json:"levle"`
	Format           string                `json:"format"`
	DisableTimestamp bool                  `json:"disableTimeStamp"`
	LogFile          logutil.FileLogConfig `json:"file"`
}

var defaultConfig = Config{
	Host: "0.0.0.0",
	Port: 8080,
	Log: Log{
		Level:            "info",
		Format:           "text",
		DisableTimestamp: false,
		LogFile: logutil.FileLogConfig{
			FileName:   "",
			MaxSize:    0,
			MaxDays:    0,
			MaxBackups: 0,
		},
	},
}

//NewConfig return a new config with default value
func NewConfig() *Config {
	conf := defaultConfig
	return &conf
}

//load configuration from configuration file
func (conf *Config) LoadConfig(configPath string) error {
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return errors.Trace(err)
	}
	c := Config{}
	if err := json.Unmarshal(configFile, &c); err != nil {
		return errors.Trace(err)
	}
	*conf = c
	return nil
}

func (conf *Config) ToLogConfig() *logutil.LogConfig {
	return &logutil.LogConfig{
		Level:            conf.Level,
		Format:           conf.Format,
		DisableTimestamp: conf.DisableTimestamp,
		File:             conf.LogFile,
	}
}
