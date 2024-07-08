package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read config %w \n", err))
	}
	//read server configuration
	fmt.Println("Server port::", viper.GetInt("server.port"))
	fmt.Println("Server port::", viper.GetString("security.jwt.key"))
}
