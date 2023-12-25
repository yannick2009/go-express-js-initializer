package main

import (
	"fmt"
	"go-express-initializer/assets"
	"go-express-initializer/utils"
)

func main() {
	// VARIABLES
	// ========== ASK FOR THE PROJECT NAME ========== //
	var projectName string = fmt.Sprintf("%v", utils.FormatProjectName())
	// ========== CREATE THE PROJECT FOLDER ========== //
	utils.CreateFolder(projectName)
	// ========== ASK FOR JAVASCRIPT🟨 OR TYPESCRIPT🟦 ========== //
	var choice string = utils.AskJsorTs()
	// ========== MAKE ACTIONS DEPENDING OF THE CHOICE ========== //
	if choice == "js" {
		utils.CreateJsSubFolders(projectName, assets.Folders_list_array)
	} else if choice == "ts" {
		utils.CreateTsSubFolders(projectName, assets.Folders_list_array)
	}
}
