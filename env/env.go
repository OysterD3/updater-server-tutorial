package env

import (
	"reflect"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

var (
	Config = struct {
		Env    string `env:"ENV,required"`
		Port   string `env:"PORT,required" envDefault:"1000"`
		GitHub struct {
			AccessToken string `env:"GITHUB_ACCESS_TOKEN,required"`
			Account     string `env:"GITHUB_ACCOUNT,required"`
			Repository  string `env:"GITHUB_REPOSITORY,required"`
		}
		ServiceURL string `env:"SERVICE_URL,required"`
	}{}
)

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		panic(err)
	}
	err = env.ParseWithFuncs(&Config, map[reflect.Type]env.ParserFunc{})
	if err != nil {
		panic(err)
	}
}
