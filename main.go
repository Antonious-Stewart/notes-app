package main

import (
	"fmt"
	"github.com/Antonious-Stewart/notes-app/config"
	"github.com/Antonious-Stewart/notes-app/modules/commands"
	"os"
	"path/filepath"
	"strings"
)

var dir string = filepath.Join("modules", "notes")

func init() {
	config.SetupFlags()
}

func main() {
	command := map[string]func(filePath string){
		"add":    commands.AddPost,
		"edit":   commands.AddPost,
		"remove": commands.DeletePost,
		"read":   commands.ReadPost,
	}

	err := os.MkdirAll(dir, 0755)

	if err != nil {
		fmt.Println("error creating directory")
		os.Exit(1)
	}

	filePath := filepath.Join(dir, config.Title+".txt")

	switch {
	case config.Add:
		command["add"](filePath)
	case config.Edit:
		command["edit"](filePath)
	case config.Remove:
		if strings.Trim(config.Title, config.CutSet) == "" {
			fmt.Println("Title is required")
			os.Exit(1)
		}
		command["remove"](filePath)
	case config.Read:
		if strings.Trim(config.Title, config.CutSet) == "" {
			fmt.Println("Title is required")
			os.Exit(1)
		}
		command["read"](filePath)
	default:
		fmt.Println("invalid command")
	}

}
