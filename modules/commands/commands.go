package commands

import (
	"fmt"
	"github.com/Antonious-Stewart/notes-app/config"
	"os"
)

func AddPost(filePath string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error creating file")
		os.Exit(1)
	}

	// before the function returns, close the file
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing file")
		}
	}()

	_, err = file.WriteString("title: " + config.Title + "\n" + "body: " + config.Body + "\n" + "author: " + config.Author + "\n")
	if err != nil {
		fmt.Println("error writing to file")
		os.Exit(1)
	}

	fmt.Println("post added successfully")
}

func DeletePost(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Error trying to remove file")
		os.Exit(1)
	}
	fmt.Println("File", filePath, "has been deleted.")
}

func ReadPost(filePath string) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error trying to read file")
		os.Exit(1)
	}
	for _, b := range bytes {
		contents := string(b)
		fmt.Print(contents)
	}
}
