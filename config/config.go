package config

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

// 自动加载配置项
type configInfo struct {
	Server   *serverModel   `yaml:"server"`
	Wechat   *wechatModel   `yaml:"wechat"`
	DataBase *dataBaseModel `yaml:"database"`
	Redis    *redisModel    `yaml:"redis"`
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
	Dburl    string `yaml:"dburl"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

// redis配置项
type redisModel struct {
	redisUrl string `yaml:"redisUrl"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

// 消息队列配置项
type mqModel struct {
	MqUrl    string `yaml:"mqurl"`
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
var Config configInfo

// 加载配置
func loadConfigInformation(fPath string) error {
	filePath := path.Join(fPath, "config.yaml")
	configData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf(" config file read failed: %s", err)
		return err
	}
	err = yaml.Unmarshal(configData, &Config)
	if err != nil {
		log.Printf(" config parse failed: %s", err)
		return err
	}
	return nil
}

func init() {
	LoadConfig()
}

func LoadConfig() {
	var err error
	fPath, _ := os.Getwd()
	fPath = path.Join(fPath, "config")
	configPath := flag.String("c", fPath, "config file path")
	flag.Parse()
	err = loadConfigInformation(*configPath)
	// log.Printf("%+v\n%+v",Config.Server, Config.Wechat)
	if err != nil {
		panic(err)
		return
	}
}
