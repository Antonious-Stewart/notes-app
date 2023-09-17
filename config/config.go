package config

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var Body string
var Title string
var Author string

var Add bool
var Edit bool
var Remove bool
var Read bool
var CutSet = " \t\n\r"

func SetupFlags() {
	flag.StringVar(&Title, "Title", "", "Title of the post")
	flag.StringVar(&Body, "Body", "", "Body of the post")
	flag.StringVar(&Author, "Author", "", "Author of the post")
	flag.BoolVar(&Add, "add", false, "add a new post")
	flag.BoolVar(&Edit, "edit", false, "edit an existing post")
	flag.BoolVar(&Remove, "remove", false, "delete an existing post")
	flag.BoolVar(&Read, "read", false, "read an existing post")

	flag.Parse()
}

func ValidateFlags() {
	if strings.Trim(Title, CutSet) == "" || strings.Trim(Body, CutSet) == "" || strings.Trim(Author, CutSet) == "" {
		fmt.Println("Title, Body and Author are required")
		os.Exit(1)
	}
}
