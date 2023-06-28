package config

import (
	"path/filepath"
	"strings"
)

type WackupFileConfig struct {
	Path         string `json:"path"`         // The absolute path of the file.
	OutputInZip  string `json:"outputInZip"`  // The relative output path of the file in the zip file. If not specified, it will act on the result of path.Base(<path>)
	WriteHandler string `json:"writeHandler"` // Handler to write the file.
}

func (d *WackupFileConfig) handlePath(homedir string) {
	d.Path = strings.Replace(d.Path, "~", homedir, 1)
}

func (d *WackupFileConfig) handleOutputInZip() {
	if d.OutputInZip == "" {
		d.OutputInZip = filepath.Base(d.Path)
	}
}
