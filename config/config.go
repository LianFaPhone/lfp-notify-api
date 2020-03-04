package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var GConfig Config
var GPreConfig PreConfig

func LoadConfig(path string) *Config {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("Read yml config[%s] err[%v]", path, err).Error())
	}

	err = yaml.Unmarshal([]byte(data), &GConfig)
	if err != nil {
		panic(fmt.Errorf("yml.Unmarshal config[%s] err[%v]", path, err).Error())
	}
	PreProcess()
	return &GConfig
}

func PreProcess() error {
	//GPreConfig.DayMaxCoins = decimal.NewFromFloat(GConfig.Activity.DayMaxCoins)
	//GPreConfig.EachPlayerCoin = decimal.NewFromFloat(GConfig.Activity.EachPlayerCoin)
	return nil
}

type PreConfig struct {
}

type Config struct {
	Server System `yaml:"system"`
	//Redis   Redis       `yaml:"redis"`
	Db     Mysql  `yaml:"mysql"`
	Qcloud Qcloud `yaml:"qcloud"`
	ChuangLan ChuangLan `yaml:"chuanglan"`
	//User    User        `yaml:"user"`
	//Aws     Aws          `yaml:"aws"`
	//BussinessLimits BussinessLimits   `yaml:"bussiness_limits"`
	//BasNotify       BasNotify         `yaml:"bas_notify"`
	//TestPaper     TestPaper         `yaml:"test_paper"`
	//Activity      Activity           `yaml:"activity"`

}

type System struct {
	Port    string `yaml:"port"`
	Debug   bool   `yaml:"debug"`
	LogPath string `yaml:"log_path"`
	Monitor string `yaml:"monitor"`
	BkPort  string `yaml:"bk_port"`
}

type Redis struct {
	Network     string `yaml:"network"`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Password    string `yaml:"password"`
	Database    int    `yaml:"database"`
	MaxIdle     int    `yaml:"maxIdle"`
	MaxActive   int    `yaml:"maxActive"`
	IdleTimeout int    `yaml:"idleTimeout"`
	Prefix      string `yaml:"prefix"`
}

type Mysql struct {
	Host          string `yaml:"host"`
	Port          string `yaml:"port"`
	User          string `yaml:"user"`
	Pwd           string `yaml:"password"`
	Dbname        string `yaml:"dbname"`
	Charset       string `yaml:"charset"`
	Max_idle_conn int    `yaml:"maxIdle"`
	Max_open_conn int    `yaml:"maxOpen"`
	Debug         bool   `yaml:"debug"`
	ParseTime     bool   `yaml:"parseTime"`
}

type Qcloud struct {
	AppId  string `yaml:"appid"`
	AppKey string `yaml:"appkey"`
	Sign   string `yaml:"sign"`
}


type ChuangLan struct {
	YzmAccount string `yaml:"yzm_account"`
	YzmPwd     string `yaml:"yzm_pwd"`
	HyyxAccount string `yaml:"hyyx_account"`
	HyyxPwd     string `yaml:"hyyx_pwd"`
	Url     string `yaml:"url"`
}

