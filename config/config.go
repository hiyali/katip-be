package config

import (
  "fmt"
  "time"

  "github.com/jinzhu/configor"
)

type ConfigType struct {
	App struct {
    Name    string `default:"Katip"`
    Link    string `required:"true"`
    Secret  string `default:"katip_known_secret"`
  }

	DB struct {
    Name     string `default:"katip_v1"`
		User     string `default:"root"`
    Password string `default:""`
	}

	Email struct {
    Name        string  `default:"Katip team"`
		Address     string  `required:"true"`
		Password    string  `required:"true"`
		ServerHost  string  `required:"true"`
    ServerPort  uint    `default:"994"`
	}
}

func (cf *ConfigType) AppCopyright() string {
  return fmt.Sprintf("Copyright © %v %v. All rights reserved.", time.Now().Year(), cf.App.Name)
}

func (cf *ConfigType) EmailHostName() string {
  return fmt.Sprintf("%v:%v", cf.Email.ServerHost, cf.Email.ServerPort)
}

var configs ConfigType

func init() {
	configor.Load(&configs, "./config.yml")
}

func Get() ConfigType {
  return configs
}
