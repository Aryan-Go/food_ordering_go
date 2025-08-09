package backend

import (
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Secret_key  string `mapstructure:"secret_key"`
	Db_host     string `mapstructure:"db_host"`
	Db_user     string `mapstructure:"db_user"`
	Db_password string `mapstructure:"db_password"`
	Db_database string `mapstructure:"db_database"`
	Db_port     int    `mapstructure:"db_port"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(filepath.Join(path, ".env"))

	viper.AutomaticEnv()

    err = viper.ReadInConfig()
    if err != nil {
        return
    }

    err = viper.Unmarshal(&config)
    return
}
