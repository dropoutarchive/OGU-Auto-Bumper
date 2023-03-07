package config

import (
	"github.com/BurntSushi/toml"
)

var Conf Config

func init() {
	_, err := toml.DecodeFile("resources/config.toml", &Conf)
	if err != nil {
		panic(err)
	}
}
