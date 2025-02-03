package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var BaseStoragePath string
var FileStoragePath string

type Command struct {
	Name        string
	Description string
	Execute     func(*Command) error

	// Flags
	SnippetName    string
	SnippetContent string
}

var commands = map[string]*Command{
	"add": {
		Name:        "add",
		Description: "Add a snippet",
		Execute:     AddSnippet,
	},
}

func main() {
	setup()

	commandName := os.Args[1]
	command, ok := commands[commandName]
	if !ok {
		fmt.Printf("Could not find command: %s", commandName)
	}

	switch commandName {
	case "add":
		flag.StringVar(&command.SnippetName, "name", "", "The snippet you wish to add")
		flag.StringVar(&command.SnippetContent, "snippet", "", "The snippet you wish to add")
	}
	flag.Parse()

	if err := command.Execute(command); err != nil {
		fmt.Printf("Failed to execute command: %s", commandName)
	}
}

func AddSnippet(cmd *Command) error {
	// Add the snippet
	fmt.Println("Adding snippet")
	return nil
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
