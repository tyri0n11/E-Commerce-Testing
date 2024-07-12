package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/tyri0n11/Muffin/global"
)

func LoadConfig() {
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
	fmt.Println("db host::", viper.GetString("mysql.host"))

	//configure structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("unable to decode configuration %v", err)
	}
}
