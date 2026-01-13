package config

type Logging struct {
	Level string `env:"LOG_LEVEL,default=info"`
}
