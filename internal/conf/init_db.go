package conf

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"wire/internal/config"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Conf.DB.DSN()), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent), // 必须为Silent
	})
	if err != nil {
		panic(err)
	}
	return db
}
