package utils

import "os"

func ReadFileBytes(path string) []byte {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return bytes
}

func Exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			panic(err)
		}
	}

	return true
}
