package config

import "github.com/spf13/viper"

func main() {
	viper.Set("redis_Addr", "localhost:6379")
	viper.Set("mysql", "root:123456@(127.0.0.1:3306)/test?charset=utf8mb4")
	viper.SetConfigType("yaml")
	viper.WriteConfigAs("config.yaml")
}
