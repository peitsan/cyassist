package main

import (
	"checkinWebsite/app/login"
	"checkinWebsite/app/utils"
	"checkinWebsite/database"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type HttpConfig struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

type SecurityConfig struct {
	Key string `yaml:"key"`
}

type wxPusher struct {
	AppToken string `yaml:"appToken"`
}

type YamlConfig struct {
	Mysql    MysqlConfig    `yaml:"mysql"`
	Http     HttpConfig     `yaml:"http"`
	Security SecurityConfig `yaml:"security"`
	wxPusher wxPusher       `yaml:"wxPusher"`
}

func main() {
	config := readFromConfig("config.yaml")
	// 无法连接则panic
	database.InitDatabase(config.Mysql.Username, config.Mysql.Password, config.Mysql.Host, config.Mysql.Port, config.Mysql.Database)

	logFile, err := os.OpenFile("server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	key := keyTo16Key(config.Security.Key)
	r := InitialRouter(key)
	utils.APPTOKEN = config.wxPusher.AppToken
	login.Key = []byte(key)
	err = r.Run(fmt.Sprintf("%v:%v", config.Http.Address, config.Http.Port))
	if err != nil {
		panic("服务启动失败，发生以下错误:" + err.Error())
	}
}

func keyTo16Key(key string) string {
	if len(key) >= 16 {
		return key[:17]
	} else {
		for i, n := 0, 16-len(key); i < n; i++ {
			key += "a"
		}
		return key
	}
}

func readFromConfig(path string) *YamlConfig {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("读取配置文件失败，在打开文件%v时,发生了以下错误:\n%v", path, err)
		os.Exit(1)
	}
	config := &YamlConfig{}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		fmt.Printf("读取配置文件失败,发生了以下错误:\n%v", err)
		os.Exit(1)
	}
	return config
}
