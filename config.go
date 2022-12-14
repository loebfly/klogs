package klogs

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

type ymlConfig struct {
	// 日志配置
	Logger struct {
		// 日志输出方式 console,file
		Out string `yml:"out"`
		// 日志输出文件路径
		File string `yml:"file"`
	} `yml:"logger"`
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
	if strings.Contains(cfg.Logger.Out, OutFile) {
		if cfg.Logger.File == "" {
			return errors.New("if out contains file, log.file not null")
		}
	}
	cfg.fillNull()
	return nil
}

func (cfg *ymlConfig) fillNull() {
	if cfg.Logger.Out == "" {
		cfg.Logger.Out = OutConsole
	}
}
