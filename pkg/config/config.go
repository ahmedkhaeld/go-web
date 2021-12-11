package config

import (
	"html/template"
	"log"
)

// AppConfig holds the application wide config, to be available to all app's packages
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
