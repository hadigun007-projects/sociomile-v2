package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort         string `mapstructure:"SERVER_PORT"`
	DBHost             string `mapstructure:"DB_HOST"`
	DBPort             string `mapstructure:"DB_PORT"`
	DBUser             string `mapstructure:"DB_USER"`
	DBPassword         string `mapstructure:"DB_PASSWORD"`
	DBName             string `mapstructure:"DB_NAME"`
	JWTSecret          string `mapstructure:"JWT_SECRET"`
	OwnerEmail         string `mapstructure:"OWNER_EMAIL"`
	OwnerPassword      string `mapstructure:"OWNER_PASSWORD"`
	AdminAlphaEmail    string `mapstructure:"ADMIN_ALPHA_EMAIL"`
	AdminAlphaPassword string `mapstructure:"ADMIN_ALPHA_PASSWORD"`
	AgentAlphaEmail    string `mapstructure:"AGENT_ALPHA_EMAIL"`
	AgentAlphaPassword string `mapstructure:"AGENT_ALPHA_PASSWORD"`
	AdminBetaEmail     string `mapstructure:"ADMIN_BETA_EMAIL"`
	AdminBetaPassword  string `mapstructure:"ADMIN_BETA_PASSWORD"`
	AgentBetaEmail     string `mapstructure:"AGENT_BETA_EMAIL"`
	AgentBetaPassword  string `mapstructure:"AGENT_BETA_PASSWORD"`
	AllowedOrigins     string `mapstructure:"ALLOWED_ORIGINS"`
	AllowCredentials   bool   `mapstructure:"ALLOW_CREDENTIALS"`
	AllowedMethods     string `mapstructure:"ALLOWED_METHODS"`
	AllowedHeaders     string `mapstructure:"ALLOWED_HEADERS"`
	ExposeHeaders      string `mapstructure:"EXPOSE_HEADERS"`
	MaxAge             int    `mapstructure:"MAX_AGE"`
	APIKey             string `mapstructure:"API_KEY"`
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
