package main

import (
	"github.com/wioenena-q/wackup/config"
	"github.com/wioenena-q/wackup/utils"
)

func main() {
	json_contens := utils.ReadFileBytes("wackup.json")
	config := config.NewWackupConfig(json_contens)
	config.Load()
}
