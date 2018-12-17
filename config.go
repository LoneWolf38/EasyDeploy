package main

import (
		"fmt"
		"github.com/spf13/viper"
)

func ReadConfig() {
	fmt.Println("Searching for config files")
	viper.SetConfigName("deploy")
}