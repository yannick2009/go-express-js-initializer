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
	fmt.Println("=========== EXPRESS JS INITIALIZER ==========")
	fmt.Println("😏 What is your new Express Js 📁 project name ?: ")
	reader := bufio.NewReader(os.Stdin)
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
}
