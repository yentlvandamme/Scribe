package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/yentlvandamme/Scribe/parse"
	"github.com/yentlvandamme/Scribe/snippets"
	"github.com/yentlvandamme/Scribe/storage"
)

var currentVersion string = "0.0.1"
var BaseStoragePath string
var FileStoragePath string

type Command struct {
	CommandName string
	Description string
	Execute     func(*Command)

	// Flags
	SnippetName        string
	SnippetContent     string
	SnippetDescription string
}

var commands = map[string]*Command{
	"add": {
		CommandName: "add",
		Description: "Add a snippet",
		Execute:     AddSnippet,
	},
	"delete": {
		CommandName: "delete",
		Description: "Delete a snippet",
		Execute:     RemoveSnippet,
	},
	"setup": {
		CommandName: "setup",
		Description: "Runs the setup script",
		Execute:     RunSetup,
	},
}

func main() {
	setup()

	commandName := os.Args[1]
	command, ok := commands[commandName]
	if !ok {
		fmt.Printf("Could not find command: %s", commandName)
		os.Exit(1)
	}

	switch commandName {
	case "add":
		flag.StringVar(&command.SnippetName, "name", "", "The snippet you wish to add")
		flag.StringVar(&command.SnippetContent, "snippet", "", "The snippet you wish to add")
		flag.StringVar(&command.SnippetDescription, "description", "", "The description of the snippet you wish to add")

	case "delete":
		flag.StringVar(&command.SnippetName, "name", "", "The snippet you wish to add")
	}
	flag.CommandLine.Parse(os.Args[2:])

	command.Execute(command)
}

func AddSnippet(cmd *Command) {
	file, err := storage.ReadFromFile(FileStoragePath)
	if err != nil {
		fmt.Println("Could not open snippets file: %w", err)
	}
	defer file.Close()

	parsedSnippets, err := parse.ParseJson(file)
	if err != nil {
		fmt.Println("Could not parse stored snippets: %w", err)
	}

	parsedSnippets.SnippetsMap.AddSnippet(snippets.Snippet{
		Name:        (*cmd).SnippetName,
		Value:       (*cmd).SnippetContent,
		Description: (*cmd).SnippetDescription,
		ModifiedOn:  time.Now(),
	})

	snippetsBytes, err := parse.ParseToBytes(parsedSnippets)
	if err != nil {
		fmt.Println("Could not parse the JSON structure to bytes: %w", err)
	}

	err = storage.WriteToFile(FileStoragePath, snippetsBytes)
	if err != nil {
		fmt.Println("Failed to write to file: %w", err)
	}
}

func RemoveSnippet(cmd *Command) {
	file, err := storage.ReadFromFile(FileStoragePath)
	if err != nil {
		fmt.Println("Could not open snippets file: %w", err)
	}
	defer file.Close()

	parsedSnippets, err := parse.ParseJson(file)
	if err != nil {
		fmt.Println("Could not parse stored snippets: %w", err)
	}

	parsedSnippets.SnippetsMap.DeleteSnippet(cmd.SnippetName)
	snippetsBytes, err := parse.ParseToBytes(parsedSnippets)
	if err != nil {
		fmt.Println("Could not parse the JSON structure to bytes: %w", err)
	}

	err = storage.WriteToFile(FileStoragePath, snippetsBytes)
	if err != nil {
		fmt.Println("Failed to write to file: %w", err)
	}
}

func RunSetup(cmd *Command) {
	err := setup()
	if err != nil {
		fmt.Println("Could not run setup: %w", err)
	}
}

func setup() error {
	var homeDir, err = os.UserHomeDir()
	if err != nil {
		return err
	}
	BaseStoragePath = filepath.Join(homeDir, "Documents/Scribe")
	FileStoragePath = filepath.Join(BaseStoragePath, "snippets.json")

	createStorageFolder()
	fileInfo, err := os.Stat(FileStoragePath)
	if err != nil {
		return err
	}
	if fileInfo.Size() == 0 {
		createEmptyStructure(FileStoragePath)
	}

	return nil
}

func createStorageFolder() {
	if _, err := os.Stat(FileStoragePath); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			err = os.MkdirAll(BaseStoragePath, os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}

			var file, err = os.Create(FileStoragePath)
			defer file.Close()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func createEmptyStructure(path string) error {
	var emptyStructure snippets.Snippets = snippets.Snippets{
		Version:     currentVersion,
		SnippetsMap: make(map[string]snippets.Snippet),
	}

	parsedEmptyStructure, err := parse.ParseToBytes(emptyStructure)
	if err != nil {
		return err
	}

	storage.WriteToFile(path, parsedEmptyStructure)
	return nil
}
