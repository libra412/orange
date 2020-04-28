package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

// 自动加载配置项
type ConfigInfo struct {
	Server   *serverModel   `yaml:"server"`
	Wechat   *wechatModel   `yaml:"wechat"`
	DataBase *dataBaseModel `yaml:"database"`
	MQ       *mqModel       `yaml:"mq"`
	Email    *emailModel    `yaml:"email"`
}

// 服务器配置项
type serverModel struct {
	Mode string `yaml:"mode"` // run mode
	Host string `yaml:"host"` // server host
	Port string `yaml:"port"` // server port
}

// 微信配置项
type wechatModel struct {
	AppID     string `yaml:"appID"`
	AppSecret string `yaml:"appSecret"`
}

// 数据库配置项
type dataBaseModel struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Dburl    string `yaml:"dburl"`
}

// 消息队列配置项
type mqModel struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// 邮箱服务配置
type emailModel struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

//
var Config ConfigInfo

// 加载配置
func loadConfigInformation(fPath string) error {
	filePath := path.Join(fPath, "config.yaml")
	configData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf(" config file read failed: %s", err)
		os.Exit(-1)
		return err
	}
	err = yaml.Unmarshal(configData, &Config)
	if err != nil {
		fmt.Printf(" config parse failed: %s", err)
		return err
	}
	return nil
}

func LoadConfig() {
	var err error
	fPath, _ := os.Getwd()
	fPath = path.Join(fPath, "config")
	configPath := flag.String("c", fPath, "config file path")
	flag.Parse()
	err = loadConfigInformation(*configPath)
	// fmt.Printf("%+v\n%+v",Config.Server, Config.Wechat)
	if err != nil {
		return
	}
}
