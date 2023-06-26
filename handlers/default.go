package handlers

import "archive/zip"

func DefaultWriteHandler(zw *zip.Writer, bytes []byte, outputInZip string) int {
	zfw, err := zw.Create(outputInZip)
	if err != nil {
		panic(err)
	}

	n, err := zfw.Write(bytes)
	if err != nil {
		panic(err)
	}

	return n
}
