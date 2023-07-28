package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("examples/packages/viper/")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(viper.GetString("app.name"))
	fmt.Println(viper.GetString("database.host"))
	fmt.Println(viper.GetString("database.port"))
}
