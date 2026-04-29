package global

import (
	"context"
	"database/sql"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var MySigningKey = []byte("nihao")
var Ctx = context.Background()
var Rdb *redis.Client

var Db *sql.DB

type Config struct {
	// Mysql struct {
	// 	Dsn string `mapstructure:"dsn"`
	// } `mapstructure:"mysql"`
	Redis struct {
		Addr     string `mapstructure:"addr"`
		Password string `mapstructure:"password"`
		Db       int    `mapstructure:"db"`
	} `mapstructure:"redis"`
}
type Mysql struct {
	Dsn string `mapstructure:"dsn"`
}

func Global() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config/config.yaml")
	var config Config
	var m Mysql
	err := viper.ReadInConfig()

	if err != nil {
		panic(err.Error())
	}
	viper.Unmarshal(&config)
	viper.Sub("mysql").Unmarshal(&m)
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: "",
		DB:       0,
	})
	Db, err = sql.Open("mysql", m.Dsn)
	if err != nil {
		panic(err.Error())
	}
}
