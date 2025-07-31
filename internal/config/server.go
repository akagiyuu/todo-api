package config

type ServerConfig struct {
	Port int `env:"PORT" envDefault:"3000"`
}
