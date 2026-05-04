package global

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	Mysql struct {
		Dsn string `mapstructure:"dsn"`
	} `mapstructure:"mysql"`
	Redis struct {
		Addr     string `mapstructure:"addr"`
		Password string `mapstructure:"password"`
		Db       int    `mapstructure:"db"`
	} `mapstructure:"redis"`
	RabbitMq struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"rabbitmq"`
	Jwt struct {
		Secret string `mapstructure:"secret"`
	} `mapstructure:"jwt"`
}

var (
	Db     *sql.DB
	Rdb    *redis.Client
	Cnn    *amqp091.Connection
	Secret []byte
)

func Global() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Println("zap.logger初始化失败")
	}
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config/config.yaml")
	var config Config
	err = viper.ReadInConfig()
	if err != nil {
		logger.Info("viper读取文件失败")
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		logger.Info("viper解析失败")
	}
	Db, err = sql.Open("mysql", config.Mysql.Dsn)
	if err != nil {
		logger.Info("数据库打开失败")
	}
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.Db,
	})
	Cnn, err = amqp091.Dial(config.RabbitMq.Url)
	if err != nil {
		logger.Info("rabbitmq连接失败")
	}
	Secret = []byte(config.Jwt.Secret)
}
