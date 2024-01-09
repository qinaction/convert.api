package configs

import (
	"convert.api/libs/common"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var GConfig *Config

func init() {
	GConfig = &Config{}
}

func InitConfig() (err error) {
	var (
		env = common.CommandParameterGet("env").(string)
		dir = "./configs/" + env + ".yml"
	)

	yamlConf, err := ioutil.ReadFile(dir)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(yamlConf, GConfig); err != nil {
		return err
	}

	return nil

}
