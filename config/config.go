package config

import (
	"encoding/json"
	"github.com/dandelabs/ghostbuster-backend-libs/dandelog"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	EnvDev  = "DEV"
	EnvStg  = "STG"
	EnvProd = "PROD"
)

type Config struct {
	DB     *DB    `json:"db,omitempty"`
	Main   *Main  `json:"main,omitempty"`
	Crypto *Crypt `json:"crypto,omitempty"`
}

type DB struct {
	URL string `json:"url,omitempty"`
}

type Main struct {
	Version  string `json:"version,omitempty"`
	Hostname string `json:"hostname,omitempty"`
}

type Crypt struct {
	SaltSize   int `json:"salt_size,omitempty"`
	Iterations int `json:"iterations,omitempty"`
	KeySize    int `json:"key_size,omitempty"`
}

//Cfg exported
var Cfg Config

var (
	configFileDEV, _ = filepath.Abs("config_files/config_dev.json")
)

// SetEnvironment according to buildenv ENV variable
func SetEnvironment(environment string) {

	if environment == EnvDev {
		dandelog.InitLog(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
		dandelog.Trace.Println("Setting Dev environment")
		LoadConfig(configFileDEV)
	} else {
		dandelog.InitLog(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
		dandelog.Trace.Println("Setting Dev environment")
		LoadConfig(configFileDEV)
	}
}

func LoadConfig(configFile string) {

	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		dandelog.Error.Fatal("Config File Missing. ", err)

	}
	err = json.Unmarshal(file, &Cfg)
	dandelog.Trace.Printf("%+v\n%+v\n%+v\n", Cfg.DB, Cfg.Crypto, Cfg.Main)
	if err != nil {
		dandelog.Error.Fatal("Reiew config file format. ", err)

	}

}
