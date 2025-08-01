package config

type JwtConfig struct {
	Secret    string `env:"JWT_SECRET" envDefault:"secret"`
	ExpiredIn int    `env:"JWT_EXPIRED_IN" envDefault:"24"` // hour
}
