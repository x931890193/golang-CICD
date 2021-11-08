package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var Conf *config

type config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Model struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"model"`
	Mysql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Db       string `yaml:"db"`
	} `yaml:"mysql"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		Db       int `yaml:"db"`
	} `yaml:"redis"`
}

type envConf struct {
	Local config
	Prod  config
}

func ReadEnv(path string) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		panic("load conf error!")
	}
	conf := envConf{}
	err = yaml.Unmarshal(f, &conf)
	if err != nil {
		fmt.Println(err.Error())
		panic("load conf error!")
	}
	Conf = &conf.Local
	if os.Getenv("PROGRAM_ENV") == "pro" {
		Conf = &conf.Prod
	}
	fmt.Println("config file load success!")
}
