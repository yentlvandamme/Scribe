package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var BaseStoragePath string
var FileStoragePath string

func main() {
	setup()
}

func setup() {
	var homeDir, err = os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	BaseStoragePath = filepath.Join(homeDir, "Documents/Scribe")
	FileStoragePath = filepath.Join(BaseStoragePath, "snippets.json")

	createStorageFolder()
}

func createStorageFolder() {
	if _, err := os.Stat(FileStoragePath); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			err = os.MkdirAll(BaseStoragePath, os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}

			var file, err = os.Create(FileStoragePath)
			if err != nil {
				fmt.Println(err)
			}
			file.Close()
		}
	}
}
