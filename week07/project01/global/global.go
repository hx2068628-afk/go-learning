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

func Global() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	Rdb = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis_Addr"),
		Password: "",
		DB:       0,
	})
	Db, err = sql.Open("mysql", viper.Get("mysql").(string))
	if err != nil {
		panic(err.Error())
	}
}
