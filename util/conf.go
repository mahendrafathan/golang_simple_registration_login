package util

import (
	"encoding/json"
	"io/ioutil"
)

var c Conf

type Conf struct {
	DbHost string `json:"db_host"`
	DbPort int    `json:"db_port"`
	DbUser string `json:"db_user"`
	DbPass string `json:"db_pass"`
	DbName string `json:"db_name"`
}

func (c *Conf) getConf() *Conf {

	yamlFile, err := ioutil.ReadFile("conf.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(yamlFile, c)
	if err != nil {
		panic(err)
	}

	return c
}

func InitConfig() {
	c.getConf()
}
