package main

import (
	"github.com/BurntSushi/toml"
	"github.com/kardianos/osext"
)

var (
	FULL_EXE_PATH string
)

type config struct {
	LogPath string
}

func loadFullExePath() {
	if exe, err := osext.Executable(); err != nil {
		panic("Cannot find EXE path, error: " + err.Error())
	} else {
		FULL_EXE_PATH = exe
	}
}

func getDefaultConfig() *config {
	return &config{
		FULL_EXE_PATH + ".log",
	}
}

func loadConfig() *config {
	configFile := FULL_EXE_PATH + ".toml"

	conf := &config{}
	if _, err := toml.DecodeFile(configFile, conf); err != nil {
		return getDefaultConfig()
	} else {
		return conf
	}
}
