package conf

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wire/internal/config"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Conf.DB.DSN()), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	return db
}
