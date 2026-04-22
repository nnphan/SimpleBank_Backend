package configs

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
    DB_URL      string `mapstructure:"DB_URL"`
    Port        string `mapstructure:"PORT"`
    DBDriver    string `mapstructure:"DBDriver"`
}

func LoadEnvConfig() *Config {
    return &Config{
        DB_URL: os.Getenv("DB_URL"),
        Port:     os.Getenv("PORT"),
        DBDriver:     os.Getenv("DB_DRIVER"),
    }
}

func LoadConfig(path string, env string) (config Config, err error) {
    
    viper.AddConfigPath(path)
    viper.SetConfigName(env)
    viper.SetConfigType("env")
    viper.AutomaticEnv()
    err = viper.ReadInConfig()

    if err != nil{
        return
    }

    err = viper.Unmarshal(&config)
    return config, nil
}
    
    
