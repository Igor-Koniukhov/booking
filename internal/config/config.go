package config

import (
	"html/template"
	"log"

	"github.com/Igor-Koniukhov/bookings/internal/models"
	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	Str           string
	MailChan      chan models.MailData
}
