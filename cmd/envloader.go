package cmd

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// deprecated function name in favor of EnvLoader
type OptionsEnv struct {
	Prefix string

	DotEnv bool
}

// deprecated function name in favor of EnvLoader
func EnvLoader(v any, opt OptionsEnv) (err error) {
	if opt.DotEnv {
		if err = godotenv.Load(); err != nil {
			return
		}
	}

	err = env.ParseWithOptions(v, env.Options{
		Prefix: opt.Prefix,
	})
	return
}
