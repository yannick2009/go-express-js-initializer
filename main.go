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
		utils.CreateJsFiles(projectName, []string{assets.F_js_packages}, assets.F_env_example)
	} else if choice == "ts" {
		utils.CreateTsSubFolders(projectName, assets.Folders_list_array)
		utils.CreateTsFiles(projectName, []string{assets.F_ts_packages, assets.F_tsconfig}, assets.F_env_example)
	}
	// ========== END OF THE EXECUTIONS ========== //
	utils.PrintAtheEnd(projectName)
}
