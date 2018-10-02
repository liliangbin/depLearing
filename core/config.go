package core

import (
	"os"
	"encoding/json"
)

var Conf *Config
/*全局变量，初始化后可以被全局调用*/
type Config struct {
	Db Database `json:"db"`
}

type Database struct {
	Address  string `json:"address" form:"address"`
	Port     string `json:"port" form:"port"`
	Dbname   string `json:"dbname" form:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ParseConf(configFile string) error {

	var config Config

	conf, err := os.Open(configFile)
	if err != nil {
		panic(err.Error())
	}
	err = json.NewDecoder(conf).Decode(&config)

	Conf = &config
	return nil
}
