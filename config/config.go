package config

type Config struct {
	SERVER_ADDRESS string
	DB_USERNAME    string
	DB_PASSWORD    string
	DB_PORT        string
	DB_HOST        string
	DB_NAME        string
	JWT_KEY 	   string
}

func InitConfiguration() Config {

	return Config{
		SERVER_ADDRESS: "0.0.0.0:8888",
		DB_USERNAME:    "root",
		DB_PASSWORD:    "root",
		DB_NAME:        "vaccination",
		DB_PORT:        "3306",
		DB_HOST:        "127.0.0.1",
		JWT_KEY:		"123",
	}
}
