package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// CREATE AND FORMAT THE PROJECT NAME
func FormatProjectName() string {
	fmt.Println("================ EXPRESS JS INITIALIZER ================")
	fmt.Println("|----------- Developed by Yannick KOUAKOU ----------|")
	fmt.Println()
	fmt.Println("😏 What is your new Express Js 📁 project name ?: ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("your project name: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("an error has occurred while reading the name of your project:", err)
		os.Exit(1)
	}
	input = strings.TrimSpace(strings.ToLower(input))
	input = fmt.Sprint(strings.Replace(input, " ", "-", -1))
	return input
}

// FUNCTION TO CREATE THE PROJECT ROOT FOLDER
func CreateFolder(projectName string) {
	err := os.Mkdir(projectName, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(fmt.Sprintf("Erreur lors de la création du dossier %s : %T - %v", projectName, err, err))
	}
	fmt.Println(" ")
}

// FUNCTION TO ASK BETWEEN JS & TS
var choice int

func AskJsorTs() string {
	fmt.Println("What's your choice between JavaScript🟨 and TypeScript🟦: ")
	fmt.Println("1 => JavaScript🟨 ")
	fmt.Println("2 => TypeScript🟦")
	fmt.Print("your choice: ")
	fmt.Scan(&choice)
	if int(choice) == 1 {
		fmt.Println("you choose Javascript🟨")
		fmt.Println(" ")
		return "js"
	} else {
		fmt.Println("you choose TypeScript🟦")
		fmt.Println(" ")
		return "ts"
	}
}

// CHOICE FOR TYPESCRIPT
func CreateJsSubFolders(projectName string, subfolders []string) {
	// create sub-folders in root
	for _, sub := range subfolders {
		sub = fmt.Sprintf("%s/%s", projectName, sub)
		err := os.MkdirAll(sub, os.ModePerm)
		if err != nil && !os.IsExist(err) {
			log.Fatal(fmt.Sprintf("Erreur lors de la création du sous dossier %s : %T - %v", sub, err, err))
		}
	}
}

func CreateTsSubFolders(projectName string, subfolders []string) {
	// create the source folder
	const sourceDir string = "src"
	err := os.MkdirAll(fmt.Sprintf("%s/%s", projectName, sourceDir), os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(fmt.Sprintf("Erreur lors de la création du sous dossier source %s : %T - %v", sourceDir, err, err))
	}
	// create sub-folders in src
	for _, sub := range subfolders {
		sub = fmt.Sprintf("%s/%s/%s", projectName, sourceDir, sub)
		err := os.MkdirAll(sub, os.ModePerm)
		if err != nil && !os.IsExist(err) {
			log.Fatal(fmt.Sprintf("Erreur lors de la création du sous dossier %s : %T - %v", sub, err, err))
		}
	}
}
