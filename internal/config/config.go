package config

import (
	"fmt"
)

var Conf = new(Config)

type Config struct {
	App App `mapstructure:"app"`
	DB  DB  `mapstructure:"db"`
}
type App struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	Env  string `mapstructure:"env"`
	Log  string `mapstructure:"log"`
}

func (app *App) Link() string {
	return fmt.Sprintf("%s:%s", app.Host, app.Port)
}

type DB struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

func (db *DB) DSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", db.User, db.Password, db.Host, db.Port, db.DBName, "5s")
	return dsn
}
