package config

import (
	"net/url"
	"strconv"

	"github.com/joho/godotenv"
)

type ConfigEnv struct {
	*DataBeseConfig
	*LoggerConfig
	*ServerConfig
	*RedisConfig
}

func LoadEnv(path ...string) {
	// Implementation for loading environment variables
	if len(path) > 0 {
		godotenv.Load(path[0])
		return
	}
	godotenv.Load()
}

func LoadConfigEnv() *ConfigEnv {
	return &ConfigEnv{
		DataBeseConfig: &DataBeseConfig{
			Host:     getStringEnv("DB_HOST", "localhost"),
			Port:     getIntEnv("DB_PORT", 5432),
			User:     getStringEnv("DB_USER", "postgres"),
			Password: getStringEnv("DB_PASSWORD", "postgres"),
			DBName:   getStringEnv("DB_NAME", "postgres"),
			SSLMode:  getStringEnv("DB_SSL_MODE", "disable"),
		},
		LoggerConfig: &LoggerConfig{
			Level:  getStringEnv("LOGGER_LEVEL", "info"),
			Format: getStringEnv("LOGGER_FORMAT", "json"),
		},
		ServerConfig: &ServerConfig{
			Port: getIntEnv("SERVER_PORT", 8080),
			Host: getStringEnv("SERVER_HOST", "localhost"),
		},
		RedisConfig: &RedisConfig{
			Host: getStringEnv("REDIS_HOST", "localhost"),
			Port: getIntEnv("REDIS_PORT", 6379),
		},
	}
}

func (c *ConfigEnv) GetPostgresURL() string {
	dbUrl := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(c.DataBeseConfig.User, c.DataBeseConfig.Password),
		Host:   c.DataBeseConfig.Host + ":" + strconv.Itoa(c.DataBeseConfig.Port),
		RawQuery: url.Values{
			"sslmode": []string{c.DataBeseConfig.SSLMode},
		}.Encode(),
	}
	return dbUrl.String()
}

func (c *ConfigEnv) GetRedisURL() string {
	redisUrl := &url.URL{
		Scheme: "redis",
		Host:   c.ServerConfig.Host + ":6379",
	}
	return redisUrl.String()
}

func (c *ConfigEnv) GetServerAddress() string {
	return c.ServerConfig.Host + ":" + strconv.Itoa(c.ServerConfig.Port)
}
