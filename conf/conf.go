package conf

import (
	"os"
)

var Conf *Config

type Config struct {
	DB    *DB
	Redis *Redis
}

type DB struct {
	Host     string
	User     string
	Password string
	Name     string
	Idles    int
	Opens    int
	LifeTime int
}

type Redis struct {
	Host     string
	Password string
}

func Init() {
	initEnv()
}

func initEnv() (c *Config) {
	Conf = &Config{}
	Conf.DB = &DB{
		Host:     os.Getenv("MIAOSHA_DB_HOST"),
		User:     os.Getenv("MIAOSHA_DB_USER"),
		Password: os.Getenv("MIAOSHA_DB_PASSWORD"),
		Name:     os.Getenv("MIAOSHA_DB_NAME"),
		Idles:    2,
		Opens:    5,
		LifeTime: 2,
	}
	Conf.Redis = &Redis{
		Host:     os.Getenv("MIAOSHA_REDIS_HOST"),
		Password: os.Getenv("MIAOSHA_REDIS_PASSWORD"),
	}
	return
}
