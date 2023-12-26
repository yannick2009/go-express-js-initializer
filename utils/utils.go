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
// Function to ask between Javascript (Js) and Typescript (Ts)
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

// CHOICE FOR JAVASCRIPT

// Function to create sub-folders
func CreateJsSubFolders(projectName string, subfolders []string) {
	// create sub-folders in root
	for _, sub := range subfolders {
		sub = fmt.Sprintf("%s/%s", projectName, sub)
		err := os.MkdirAll(sub, os.ModePerm)
		if err != nil && !os.IsExist(err) {
			log.Fatalf("Erreur lors de la création du sous dossier %s : %T - %v", sub, err, err)
		}
	}
}

// Functions to create differents main files 
func CreateJsFiles(projectName string, jsonfilesList []string, envFile string) {
	// create differents json files
	var filesNames []string = []string{"package"}
	for index, file := range filesNames {
		err := os.WriteFile(fmt.Sprintf("%s/%s.json", projectName, file), []byte(jsonfilesList[index]), os.ModePerm)
		if err != nil {
			log.Fatalf("Erreur lors de la création du fichier %s: %T - %v", file, err, err)
		}
	}
	// create the env example file
	err := os.WriteFile(fmt.Sprintf("%s/.env.example", projectName), []byte(envFile), os.ModePerm)
	if err != nil {
		log.Fatalf("Erreur lors de la création du fichier exemple env")
	}
	// create server.js and app.js files
	var jsMainFiles []string = []string{"server", "app"}
	for _, file := range jsMainFiles {
		err := os.WriteFile(fmt.Sprintf("%s/%s.js", projectName, file), []byte(fmt.Sprintf("//%s.js", file)), os.ModePerm)
		if err != nil {
			log.Fatalf("Erreur lors de la création du fichier %s.js: %T - %v", file, err, err)
		}
	}
}

// CHOICE FOR TYPESCRIPT
const sourceDir string = "src"
// Function to create src sub-folders
func CreateTsSubFolders(projectName string, subfolders []string) {
	// create the source folder
	err := os.MkdirAll(fmt.Sprintf("%s/%s", projectName, sourceDir), os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatalf("Erreur lors de la création du sous dossier source %s : %T - %v", sourceDir, err, err)
	}
	// create sub-folders in src
	for _, sub := range subfolders {
		sub = fmt.Sprintf("%s/%s/%s", projectName, sourceDir, sub)
		err := os.MkdirAll(sub, os.ModePerm)
		if err != nil && !os.IsExist(err) {
			log.Fatalf("Erreur lors de la création du sous dossier %s : %T - %v", sub, err, err)
		}
	}
}

// Functions to create differents main files 
func CreateTsFiles(projectName string, jsonfilesList []string, envFile string) {
	// create differents json files
	var filesNames []string = []string{"package", "tsconfig"}
	for index, file := range filesNames {
		err := os.WriteFile(fmt.Sprintf("%s/%s.json", projectName, file), []byte(jsonfilesList[index]), os.ModePerm)
		if err != nil {
			log.Fatalf("Erreur lors de la création du fichier %s: %T - %v", file, err, err)
		}
	}
	// create the env example file
	err := os.WriteFile(fmt.Sprintf("%s/.env.example", projectName), []byte(envFile), os.ModePerm)
	if err != nil {
		log.Fatalf("Erreur lors de la création du fichier exemple env")
	}
	// create server.ts and app.ts files
	var tsMainFiles []string = []string{"server", "app"}
	for _, file := range tsMainFiles {
		err := os.WriteFile(fmt.Sprintf("%s/%s/%s.ts", projectName, sourceDir, file), []byte(fmt.Sprintf("//%s.ts", file)), os.ModePerm)
		if err != nil {
			log.Fatalf("Erreur lors de la création du fichier %s.ts: %T - %v", file, err, err)
		}
	}
}

// END WHITH PRINT
func PrintAtheEnd(projectName string) {
	fmt.Println("|--------------------------------------------|")
	fmt.Println(" Congratulations your project is ready ✅🎊")
	fmt.Println("|--------------------------------------------|")
	fmt.Println("")
	fmt.Println("execute these commands below and enjoy your project 💫😁")
	fmt.Println("")
	fmt.Printf("cd %s\n", projectName)
	fmt.Println("npm install")
}
