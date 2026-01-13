package config

type HTTP struct {
	Addr string `env:"HTTP_ADDR,default=:8080"`
}
