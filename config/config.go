package config

import (
	"encoding/json"
	"os"
	"strings"
)

type WackupConfig struct {
	HomeDir string              // The home directory of the user. This is automatically set.
	Output  string              `json:"output"` // The zip file to be created after the backup is made. Default "output.zip"
	Dirs    []*WackupDirConfig  `json:"dirs"`   // The configs of the directories to be backed up.
	Files   []*WackupFileConfig `json:"files"`  // Configurations of files to be backed up. This is for individual files only.
}

func NewWackupConfig(contents []byte) *WackupConfig {
	var config WackupConfig
	err := json.Unmarshal(contents, &config)
	if err != nil {
		panic(err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	config.HomeDir = homeDir

	if config.Output == "" {
		config.Output = "output.zip"
	} else if !strings.HasSuffix(config.Output, ".zip") {
		config.Output += ".zip"
	}

	return &config
}
