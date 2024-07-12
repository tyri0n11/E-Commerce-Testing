package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"mysql"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config: %w ", err))
	}
	//read server configuration
	fmt.Println("Server port::", viper.GetInt("server.port"))
	fmt.Println("Security key::", viper.GetString("security.jwt.key"))

	//configure structure

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("unable to decode configuration %v", err)

	}
	fmt.Println("Config port::", config.Server.Port)
	for _, db := range config.Databases {
		fmt.Printf("database User: %s, password: %s, host: %s", db.User, db.Password, db.Host)
		fmt.Print(1)
	}
}
