package klogs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ymlConfig struct {
	// 日志配置
	Log struct {
		// 日志输出方式 console,file
		Out string `yaml:"out"`
		// 日志输出文件路径
		File string `yaml:"file"`
	} `yaml:"log"`
}

var config = new(ymlConfig)

func (cfg *ymlConfig) init(ymlPath string) error {
	file, err := ioutil.ReadFile(ymlPath)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(file, cfg); err != nil {
		return err
	}
	cfg.fillNull()
	return nil
}

func (cfg *ymlConfig) fillNull() {
	if cfg.Log.Out == "" {
		cfg.Log.Out = OutConsole + "," + OutFile
	}
}
