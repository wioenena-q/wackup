package handlers

import "archive/zip"

type WriteHandler func(*zip.Writer, []byte, string) int

var Cache = make(map[string]WriteHandler, 1)

func init() {
	Cache["Default"] = DefaultWriteHandler
}
