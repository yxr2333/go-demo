package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type OSSConfigYaml struct {
	OSSConfig *OSSConfig `yaml:"oss"`
}

type OSSConfig struct {
	SecretKey string `yaml:"secretKey"`
	SecretID  string `yaml:"secretID"`
}

var OSSConf *OSSConfigYaml

func InitOSSConfig() (err error) {
	file, err := ioutil.ReadFile("config/oss.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(file, &OSSConf)
	fmt.Printf("id:%s\n", OSSConf.OSSConfig.SecretID)
	fmt.Printf("key:%s\n", OSSConf.OSSConfig.SecretKey)
	if err != nil {
		return err
	}
	return nil
}

func GetOSSConfig() (conf *OSSConfig) {
	return OSSConf.OSSConfig
}
