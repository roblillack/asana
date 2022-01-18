package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/thash/asana/utils"
	"gopkg.in/yaml.v1"
)

type Conf struct {
	Personal_access_token string
	Workspace             string
}

func Load() Conf {
	var dat []byte
	var err error
	dat, err = ioutil.ReadFile(utils.Home() + "/.asana.yml")
	if err != nil {
		fmt.Println("Config file isn't set.\n  ==> $ asana config")
		os.Exit(1)
	}
	conf := Conf{}
	err = yaml.Unmarshal(dat, &conf)
	utils.Check(err)
	return conf
}
