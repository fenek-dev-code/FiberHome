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

type ConfigEnv struct {
	DatabaseURL string
	Port        int
	DebugMode   bool
	ServiceName string
}

func LoadConfigEnv() *ConfigEnv {
	return &ConfigEnv{
		DatabaseURL: getStringEnv("DATABASE_URL", "postgres://localhost:5432/mydb"),
		Port:        getIntEnv("PORT", 8080),
		DebugMode:   getBoolEnv("DEBUG_MODE", false),
		ServiceName: getStringEnv("SERVICE_NAME", "MyService"),
	}
}
