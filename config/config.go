package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	AiClient AIClientConfig
}

type ServerConfig struct {
	Host string `json:"host" validate:"required"`
	Port string `json:"port" validate:"required"`
}

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"DBName"`
	SSLMode  string `json:"sslMode"`
	PgDriver string `json:"pgDriver"`
}

type AIClientConfig struct {
	BaseUrlBrigAI string `json:"brig_ai"`
	VendorEUAI    string `json:"vandor_eu_ai"`
	VendorRUAI    string `json:"vendor_ru_ai"`
}

func LoadConfig() (*viper.Viper, error) {

	viperInstance := viper.New()

	if _, ok := os.LookupEnv("LOCAL"); ok {
		viperInstance.AddConfigPath("config")
	} else {
		viperInstance.AddConfigPath("./config")
	}
	viperInstance.SetConfigName("config")
	viperInstance.SetConfigType("yml")

	err := viperInstance.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return viperInstance, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {

	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode config into struct, %v", err)
		return nil, err
	}
	return &c, nil
}
