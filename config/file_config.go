package config

type WackupFileConfig struct {
	Path         string `json:"path"`         // The absolute path of the file.
	OutputInZip  string `json:"outputInZip"`  // The relative output path of the file in the zip file. If not specified, it will act on the result of path.Base(<path>)
	WriteHandler string `json:"writeHandler"` // Handler to write the file.
}
