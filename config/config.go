package config

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App             AppConfig
		Server          ServerConfig
		Database        DatabaseConfig
		Authentication  AuthenticationConfig
		Observalibility Observalibility
		JWT             JWTConfig
	}

	AppConfig struct {
		Name        string
		Version     string
		Schema      string
		Host        string
		Environment string
	}

	ServerConfig struct {
		Port     string
		Debug    bool
		TimeZone string
	}

	DatabaseConfig struct {
		Host     string
		Port     int
		Name     string
		User     string
		Password string
	}

	AuthenticationConfig struct {
		Key string
	}

	Observalibility struct {
		Enable bool
		Mode   string
	}

	JWTConfig struct {
		Key     string
		Expired int
		Label   string
	}
)

func LoadConfig(filename string) (Config, error) {
	v := viper.New()
	v.SetConfigName(fmt.Sprintf("config/%s", filename))
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %v\n", err)
		return Config{}, err
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshaling config: %v\n", err)
		return Config{}, err
	}

	return config, nil
}

func LoadConfigPath(path string) (Config, error) {
	v := viper.New()

	v.SetConfigName(path)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return Config{}, errors.New("config file not found")
		}
		return Config{}, err
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return Config{}, err
	}

	return c, nil
}
