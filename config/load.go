package config

import (
	"archive/zip"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/wioenena-q/wackup/handlers"
	"github.com/wioenena-q/wackup/utils"
)

func (c *WackupConfig) Load() {
	if utils.Exist(c.Output) {
		panic("Output file already exists")
	}
	zipFile, err := os.Create(c.Output)
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, dir := range c.Dirs {
		dir.tag = strings.ToUpper(filepath.Base(dir.Path))
		dir.handlePath(c.HomeDir)
		dir.handleOutputInZip()
		loadDir(dir.tag, dir.Path, dir.Path, dir.OutputInZip, dir.OutputInZip, zipWriter, dir.Ignores, &dir.WriteHandlers)
	}

	panic("TODO: Implement files")
}

func loadDir(tag, rootAbsPath, absPath, rootParent, parent string, zw *zip.Writer, ignores []string, writeHandlers *map[string]string) {
	files, err := os.ReadDir(absPath)

	if err != nil {
		panic(err)
	}

MainLoop:
	for _, entry := range files {
		entryName := entry.Name()
		entryAbsPath := filepath.Join(absPath, entryName)
		relativePath := parent + "/" + entryName

		for _, ignore := range ignores {
			ignoreAbsPath := filepath.Join(rootAbsPath, ignore)
			matches, err := filepath.Glob(ignoreAbsPath)
			if err != nil {
				panic(err)
			}

			for _, match := range matches {
				if match == entryAbsPath {
					log.Printf("[%s]: Ignored %s", tag, match)
					continue MainLoop // Go to next entry
				}
			}
		}

		if entry.IsDir() {
			loadDir(tag, rootAbsPath, entryAbsPath, rootParent, filepath.Join(parent, entryName), zw, ignores, writeHandlers)
		} else {
			fileBytes := utils.ReadFileBytes(entryAbsPath)
			handlerName := (*writeHandlers)[strings.Replace(relativePath, rootParent+"/", "", 1)]
			var handlerFn handlers.WriteHandler = nil
			if handlerName != "" {
				handlerFn = handlers.Cache[handlerName]
			}

			if handlerFn == nil {
				handlerFn = handlers.Cache["Default"]
			}

			handlerFn(zw, fileBytes, relativePath)
		}
	}
}
