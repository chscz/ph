package config

type Config struct {
	MySQL MySQL
}

func LoadConfig() (Config, error) {
	return Config{}, nil
}
