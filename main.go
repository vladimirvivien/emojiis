package main

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed all:data
var mbedFS embed.FS

//var emojis []search.Emoji

func main() {
	files, err := fs.ReadDir(mbedFS, ".")
	if err != nil {
		panic(err.Error())
	}

	for _, file := range files {
		fmt.Printf("Found embedded file: %s\n", file.Name())
	}

}
