package config

type CorsConfig struct {
	AllowOrigin string `env:"CORS_ALLOW_ORIGIN" envDefault:"*"`
}
