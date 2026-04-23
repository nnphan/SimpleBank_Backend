package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Name string `mapstructure:"name"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"server"`
	Database []struct {
		Driver string `mapstructure:"driver"`
		Host string `mapstructure:"host"`
		Port int `mapstructure:"port"`
		Name string `mapstructure:"name"`
		User  string `mapstructure:"user"`
		Password  string `mapstructure:"password"`
	} `mapstructure:"database"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs") 
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	//Read config file
	err := viper.ReadInConfig()
	if err != nil {
		panic("cannot read config file: " + err.Error())
	}

	//read value from config file
	fmt.Println("App Name:", viper.GetString("server.name"))
	fmt.Println("App Port:", viper.GetInt("server.port"))

	//read configuration into struct
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic("cannot unmarshal config: " + err.Error())
	}

	fmt.Printf("Config Port:: %d\n", config.Server.Port)
	fmt.Printf("Config Name:: %s\n", config.Server.Name)

	for _,db := range config.Database {
		fmt.Printf("Database Driver: %s\n", db.Driver)
		fmt.Printf("Database Host: %s\n", db.Host)
		fmt.Printf("Database Port: %d\n", db.Port)
		fmt.Printf("Database Name: %s\n", db.Name)
		fmt.Printf("Database User: %s\n", db.User)
		fmt.Printf("Database Password: %s\n", db.Password)
	}

}
