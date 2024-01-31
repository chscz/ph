package config

type Config struct {
	MySQL MySQL
}

func LoadConfig() (Config, error) {
	return Config{
		MySQL: MySQL{
			UserName: "root",
			Password: "1111",
			Host:     "localhost",
			Port:     "3306",
			Schema:   "ph",
		}}, nil
}
