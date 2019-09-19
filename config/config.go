package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Schema Config struct
type Schema struct {
	PostgresDB struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Debug    bool   `mapstructure:"debug"`
		Database string `mapstructure:"name"`
	} `mapstructure:"go_postgres_database"`

	MongoDB struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"go_mongo_database"`

	Paging struct {
		Limit string `mapstructure:"limit"`
	} `mapstructure:"paging"`

	Encryption struct {
		EncriptionKey    string `mapstructure:"oid_key"`
		EncriptionSecret string `mapstructure:"jwt_secret"`
	} `mapstructure:"encription"`
}

var (
	Config Schema
)

// Init config
func init() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		fmt.Printf("Unable to decode into config struct, %v", err)
	}
}
