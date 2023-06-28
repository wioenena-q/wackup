package config

import (
	"archive/zip"
	"fmt"
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

		if !utils.Exist(dir.Path) {
			panic(fmt.Sprintf("Dir path does not exist: %s", dir.Path))
		}

		if !filepath.IsAbs(dir.Path) {
			panic(fmt.Sprintf("Dir path must be absolute: %s", dir.Path))
		}

		loadDir(dir.tag, dir.Path, dir.Path, dir.OutputInZip, dir.OutputInZip, zipWriter, dir.Ignores, &dir.WriteHandlers)
	}

	for _, file := range c.Files {
		file.handlePath(c.HomeDir)
		file.handleOutputInZip()

		if !utils.Exist(file.Path) {
			panic(fmt.Sprintf("File path does not exist: %s", file.Path))
		}

		if !filepath.IsAbs(file.Path) {
			panic(fmt.Sprintf("File path must be absolute: %s", file.Path))
		}

		loadFile("FILE", file.Path, file.WriteHandler, file.OutputInZip, zipWriter)
	}
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
			loadFile(tag, entryAbsPath, strings.Replace(relativePath, rootParent+"/", "", 1), relativePath, zw)
		}
	}
}

func loadFile(tag, sourcePath, handerName, outputInzip string, zw *zip.Writer) {
	fileBytes := utils.ReadFileBytes(sourcePath)
	handlerFn := getHandler(handerName)

	if strings.HasSuffix(outputInzip, "/") {
		outputInzip += filepath.Base(sourcePath)
	}

	handlerFn(zw, fileBytes, outputInzip)
	log.Printf("[%s]: Backed up %s", tag, sourcePath)
}

func getHandler(name string) handlers.WriteHandler {
	fn := handlers.Cache[name]
	if fn == nil {
		fn = handlers.Cache["Default"]
	}

	return fn
}
