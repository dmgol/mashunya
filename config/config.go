package config

import (
	"html/template"
	"os"

	"github.com/jinzhu/configor"
	"github.com/microcosm-cc/bluemonday"
	"github.com/qor/render"
)

type SMTPConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Site     string
}

var Config = struct {
	AdminsPort  uint `default:"7000" env:"ADMINS_PORT"`
	ClientsPort uint `default:"80" env:"CLIENTS_PORT"`
	DB          struct {
		Name     string `default:"testdb"`
		Adapter  string `default:"mysql"`
		User     string
		Password string
	}
	SMTP SMTPConfig
}{}

var (
	Root = os.Getenv("GOPATH") + "/src/github.com/dmgol/mashunya"
	View *render.Render
)

func init() {
	if err := configor.Load(&Config, "config/database.yml", "config/smtp.yml"); err != nil {
		panic(err)
	}

	View = render.New()

	htmlSanitizer := bluemonday.UGCPolicy()
	View.RegisterFuncMap("raw", func(str string) template.HTML {
		return template.HTML(htmlSanitizer.Sanitize(str))
	})
}

func (s SMTPConfig) HostWithPort() string {
	return s.Host + ":" + s.Port
}
