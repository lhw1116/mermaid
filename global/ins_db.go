package global

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(info *db) (err error) {
	if DB, err = gorm.Open(mysql.Open(info.Dsn), &gorm.Config{}); err != nil {
		Logger.Error(err)
		return err
	}
	return
}
