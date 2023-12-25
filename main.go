package main

import (
	"fmt"
	"go-express-initializer/utils"
)

func main() {
	// ========== ASK FOR THE PROJECT NAME ========== //
	var projectName string = fmt.Sprintf("%v", utils.FormatProjectName())
	// ========== CREATE THE PROJECT FOLDER ========== //
	utils.CreateFolder(projectName)
}
