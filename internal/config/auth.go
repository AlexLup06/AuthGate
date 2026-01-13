package config

type Auth struct {
	JWTKey         string `env:"JWT_KEY,required"`
	GoogleClientID string `env:"GOOGLE_CLIENT_ID"`
}
