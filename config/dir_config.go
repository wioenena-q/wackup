package config

import (
	"path"
	"strings"
)

type WackupDirConfig struct {
	tag         string   // Tag for logging
	Path        string   `json:"path"`        // The absolute path of the directory. Example: "/home/<username>/<directory>"
	OutputInZip string   `json:"outputInZip"` // The relative output path of the directory in the zip file. If not specified, it will act on the result of path.Base(<path>)
	Ignores     []string `json:"ignores"`     // The list of files to be ignored. Glob supported. Paths must be specified knowing that the absolute path of this folder is prepended.
	// Files of the specially specified folder.
	// Why is this necessary?
	// Let's say you have a json file in a folder and if you are going to take some values from this json file and print a different json as output, you should specify it here.
	// Because you can print this file in a different way with the writeFn you specify in it.
	// Unspecified files use a default function.
	WriteHandlers map[string] /* path */ string/* handler function name */ `json:"writeHandlers"`
}

func (d *WackupDirConfig) handlePath(homedir string) {
	d.Path = strings.Replace(d.Path, "~", homedir, 1)
}

func (d *WackupDirConfig) handleOutputInZip() {
	if d.OutputInZip == "" {
		d.OutputInZip = path.Base(d.Path)
	}
}
