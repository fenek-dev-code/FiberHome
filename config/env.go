package config

import "github.com/joho/godotenv"

func LoadEnv(path ...string) {
	// Implementation for loading environment variables
	if len(path) > 0 {
		godotenv.Load(path[0])
		return
	}
	godotenv.Load()
}

type DataBeseConfig struct {
	Url string
}

type LoggerConfig struct {
	Level  string
	Format string
}

type ServerConfig struct {
	Port int
}

type ConfigEnv struct {
	*DataBeseConfig
	*LoggerConfig
	*ServerConfig
}

func LoadConfigEnv() *ConfigEnv {
	return &ConfigEnv{
		DataBeseConfig: &DataBeseConfig{
			Url: getStringEnv("DATABASE_URL", "postgres://user:password@localhost:5432/dbname"),
		},
		LoggerConfig: &LoggerConfig{
			Level:  getStringEnv("LOGGER_LEVEL", "info"),
			Format: getStringEnv("LOGGER_FORMAT", "json"),
		},
		ServerConfig: &ServerConfig{
			Port: getIntEnv("SERVER_PORT", 8080),
		},
	}
}
