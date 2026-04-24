package initialize

import (
	"fmt"
	"simplebank/global"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper := viper.New()

	// for run debug
	//viper.AddConfigPath("../../configs")
	
	// for run from cmd/server
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
	err = viper.Unmarshal(&global.Config) // global.Config is the struct defined in global/global.go
	if err != nil {
		panic("cannot unmarshal config: " + err.Error())
	}
}