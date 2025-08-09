package config

type ServerConfig struct {
	Url  string `env:"SERVER_URL" envDefault:"http://localhost:3000"`
	Port int    `env:"PORT" envDefault:"3000"`
}
