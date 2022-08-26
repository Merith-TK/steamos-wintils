package config

import (
	"os"

	"github.com/BurntSushi/toml"

	"steamos-wintils/util"
)

var (
	configFile = os.ExpandEnv("${USERPROFILE}/.config/steamos-utils/config.toml")
)
var Config ConfigType

type ConfigType struct {
	SelectBranch struct {
		Branch string `toml:"branch"`
	} `toml:"select-branch"`
}

func SetupConfig() error {
	// if ~/.config/steamos-utils/ doesn't exist, create it
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		util.DebugPrint("Creating ~/.config/steamos-utils/")
		os.MkdirAll(os.ExpandEnv("${USERPROFILE}/.config/steamos-utils"), 0755)
	}
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		util.DebugPrint("Creating ~/.config/steamos-utils/config.toml")
		f, err := os.Create(configFile)
		if err != nil {
			return err
		}

		// set default config
		Config.SelectBranch.Branch = "rel"
		// write config to file
		toml.NewEncoder(f).Encode(Config)
		f.Close()
	} else {
		util.DebugPrint("Reading ~/.config/steamos-utils/config.toml")
		_, err := toml.DecodeFile(configFile, &Config)
		if err != nil {
			return err
		}
		util.DebugPrint("Config:", Config)
	}
	return nil
}

func WriteConfig(conf ConfigType) error {
	util.DebugPrint("Writing Config to ~/.config/steamos-utils/config.toml")
	util.DebugPrint("Config:", conf)
	f, err := os.Create(configFile)
	if err != nil {
		return err
	}
	toml.NewEncoder(f).Encode(conf)
	f.Close()
	return nil
}
