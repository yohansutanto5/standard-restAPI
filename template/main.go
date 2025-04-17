package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var config modelConfig

func init() {
	yamlFile, err := os.ReadFile("./manifest/" + os.Args[1])
	if err != nil {
		log.Fatalf("error load yaml")
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("error unmarshalling yaml")
	}
}

func main() {
	generateFile(config, "model", config.generateModelReplacements())
	generateFile(config, "db", config.generateDbReplacements())
	generateFile(config, "service", config.generateServiceReplacements())
	generateFile(config, "handler", config.generateHandlerReplacements())
	
	if len(os.Args) > 2 {
		generateRouteFile(config, config.generateRouteReplacements())
		os.WriteFile("../db/sql_migrations/generated.sql", []byte(config.generateMySQLQuery()), 0644)
	}
}

func generateFile(config modelConfig, file string, Modelreplacement map[string]string) {

	templateFile := fmt.Sprintf("./blueprint/%s.txt", file)
	content, err := os.ReadFile(templateFile)
	if err != nil {
		fmt.Println("Error reading template file:", err)
		return
	}

	// Convert content to string
	contentStr := string(content)
	// Perform replacements
	for placeholder, replacement := range Modelreplacement {
		contentStr = strings.ReplaceAll(contentStr, placeholder, replacement)
	}
	// Define the target folder
	targetFolder := "../" + file + "/"
	modifiedContent := []byte(contentStr)
	// Write the modified content to a new file in the target folder
	targetFile := targetFolder + Modelreplacement["{.model}"] + ".go"
	err = os.WriteFile(targetFile, modifiedContent, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File successfully created:", targetFile)
}

func generateRouteFile(config modelConfig, Modelreplacement map[string]string) {
	content, err := os.ReadFile("../cmd/route.go")
	if err != nil {
		fmt.Println("Error reading template file:", err)
		return
	}

	// Convert content to string
	contentStr := string(content)
	// Perform replacements
	for placeholder, replacement := range Modelreplacement {
		contentStr = strings.ReplaceAll(contentStr, placeholder, replacement)
	}
	// Define the target folder
	targetFolder := "../cmd/"
	modifiedContent := []byte(contentStr)
	// Write the modified content to a new file in the target folder
	targetFile := targetFolder + "route.go"
	err = os.WriteFile(targetFile, modifiedContent, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File successfully created:", targetFile)
}
