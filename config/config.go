package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Db struct {
	User     string `json:"user"`
	Password string `json:"pass"`
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DbType   string `json:"dbType"`
	Location string `json:"location"`
	SslMode  string `json:"sslMode"`
	Alias    string `json:"alias"`
	Force    string `json:"force"`
	Verbose  string `json:"verbose"`
	Debug    string `json:"debug"`
}

type Configuration struct {
	Db Db `json:"db"`
}

var Config Configuration

func init() {
	env := "dev"

	if len(os.Args) > 1 {
		env = os.Args[1]
	}
	log.Println("INITIALIZE", "SERVER RUN ON "+env+" ENV")
	path := "/resource/" + env + ".jsonConfig.json"

	basePath, _ := filepath.Abs(filepath.Dir(""))

	byteConf, err := ioutil.ReadFile(basePath + path)
	if err != nil {
		log.Panic("env file not found")
	}
	json.Unmarshal(byteConf, &Config)

}
