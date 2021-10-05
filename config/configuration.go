package config

import "os"

type databaseConfig struct {
	Username     string
	Password     string
	Address      string
	DatabaseName string
}

func LoadDatabaseConfiguration() (databaseConfig, error) {
	var result databaseConfig
	result.Username = os.Getenv("DB_USER")
	result.Password = os.Getenv("DB_PASS")
	result.Address = os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
	result.DatabaseName = "order-validator"
	return result, nil

}
func LoadPort() string {
	return os.Getenv("PORT")
}
