package config

type MySQL struct {
	UserName string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	DB       string `env:"DB"`
}
