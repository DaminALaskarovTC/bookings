package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// Het is belangrijk dat voor dit soort packages(die je vanaf OVERAL op kan roepen) Niks importeert vanuit andere packages.

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
