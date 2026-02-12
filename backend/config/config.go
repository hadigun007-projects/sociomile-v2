package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort       string `mapstructure:"SERVER_PORT"`
	DBHost           string `mapstructure:"DB_HOST"`
	DBPort           string `mapstructure:"DB_PORT"`
	DBUser           string `mapstructure:"DB_USER"`
	DBPassword       string `mapstructure:"DB_PASSWORD"`
	DBName           string `mapstructure:"DB_NAME"`
	JWTSecret        string `mapstructure:"JWT_SECRET"`
	AdminEmail       string `mapstructure:"ADMIN_EMAIL"`
	AdminPassword    string `mapstructure:"ADMIN_PASSWORD"`
	AgentEmail       string `mapstructure:"AGENT_EMAIL"`
	AgentPassword    string `mapstructure:"AGENT_PASSWORD"`
	AllowedOrigins   string `mapstructure:"ALLOWED_ORIGINS"`
	AllowCredentials bool   `mapstructure:"ALLOW_CREDENTIALS"`
	AllowedMethods   string `mapstructure:"ALLOWED_METHODS"`
	AllowedHeaders   string `mapstructure:"ALLOWED_HEADERS"`
	ExposeHeaders    string `mapstructure:"EXPOSE_HEADERS"`
	MaxAge           int    `mapstructure:"MAX_AGE"`
}

func LoadConfig() (config Config) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: .env file not found, using system environment variables")
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error mapping config: ", err)
	}
	return
}
