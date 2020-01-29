package config

//
type ConfigInfo struct {
	Server *serverModel `yaml:"server"`
	Wechat *wechatModel `yaml:"wechat"`
}

//
type serverModel struct {
	Mode string `yaml:"mode"` // run mode
	Host string `yaml:"host"` // server host
	Port string `yaml:"port"` // server port
}

//
type wechatModel struct {
	AppID     string `yaml:"appID"`
	AppSecret string `yaml:"appSecret"`
}

var Config ConfigInfo

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
