package main

import (
	cfg "app/cmd/config"
	"app/db"
	"app/model"
	"fmt"
	"os"
)

type modelConfig struct {
	Model  string `yaml:"model"`
	Fields []struct {
		Name       string `yaml:"name"`
		Type       string `yaml:"type"`
		Binding    string `yaml:"binding"`
		Constraint string `yaml:"constraint"`
		Input      bool   `yaml:"input"`
		Output     bool   `yaml:"output"`
	} `yaml:"fields"`
	Routers []struct {
		Method string `yaml:"method"`
		Path   string `yaml:"path"`
	} `yaml:"routers"`
}

func (m *modelConfig) generateDbReplacements() map[string]string {
	replacements := map[string]string{
		"{.model}": m.Model,
	}
	return replacements
}

func (m *modelConfig) generateHandlerReplacements() map[string]string {
	replacements := map[string]string{
		"{.model}": m.Model,
	}
	return replacements
}

func (m *modelConfig) generateServiceReplacements() map[string]string {
	replacements := map[string]string{
		"{.model}": m.Model,
	}
	return replacements
}

func (m *modelConfig) generateModelReplacements() map[string]string {
	var model_definition string
	var model_input string
	var PopulateFromDTOInput string
	var model_output string
	var PopulateDTOOutput string

	for _, field := range m.Fields {
		tag := "`gorm:\"" + field.Constraint + "\"`"
		model_definition += fmt.Sprintf("\t%s %s %s\n", field.Name, field.Type, tag)
	}

	for _, field := range m.Fields {
		if !field.Input {
			continue
		}
		var tag string
		if field.Binding == "none" {
			tag = fmt.Sprintf("`json:\"%s\"`", field.Name)
		} else {
			tag = fmt.Sprintf("`json:\"%s\" binding:\"%s\"`", field.Name, field.Binding)
		}
		model_input += fmt.Sprintf("\t%s %s %s\n", field.Name, field.Type, tag)

		PopulateFromDTOInput += fmt.Sprintf("\tm.%s = input.%s\n", field.Name, field.Name)
	}

	for _, field := range m.Fields {
		if !field.Output {
			continue
		}
		tag := fmt.Sprintf("`json:\"%s\"`", field.Name)
		model_output += fmt.Sprintf("\t%s %s %s\n", field.Name, field.Type, tag)

		PopulateDTOOutput += fmt.Sprintf("\toutput.%s = m.%s\n", field.Name, field.Name)
	}

	replacements := map[string]string{
		"{.model}":                 m.Model,
		"{.model_definition}":      model_definition,
		"{.model_input}":           model_input,
		"{.PopulateFromDTOInput}":  PopulateFromDTOInput,
		"{.model_output}":          model_output,
		"{.PopulateFromDTOOutput}": PopulateDTOOutput,
	}
	return replacements
}

func (m *modelConfig) generateRouteReplacements() map[string]string {
	var handler string
	var route string
	handler += fmt.Sprintf("\t%shandler:= handler.New%sHandler(database)", m.Model, m.Model)

	for _, r := range m.Routers {
		switch r.Method {
		case "get":
			route += fmt.Sprintf("\tr.GET(\"%s\", %shandler.GetList)\n", r.Path, m.Model)
		case "post":
			route += fmt.Sprintf("\tr.POST(\"%s\", %shandler.Insert)\n", r.Path, m.Model)
		}
	}
	replacements := map[string]string{
		"//{.NewHandler}": handler + "\n //{.NewHandler}",
		"{.model}":        m.Model,
		"//{.NewRoute}":   route + "	// ---- {.model} API ---\n //{.NewRoute}",
	}
	return replacements
}

func (m *modelConfig) generateMySQLQuery() string {
	mode := "debug"
	configfolder := os.Getenv("config")
	configuration := cfg.Load(mode, configfolder)
	// setup DB connection
	database := db.NewDatabase(configuration)

	migrator := database.Db.Migrator()
	// Generate table DDL for UserProfile
	createTableSQL := migrator.CreateTable(&model.User{})
	return createTableSQL.Error()
}
