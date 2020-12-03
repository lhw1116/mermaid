package global

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mermaid/model"
)

func InitDB(info *db) (err error) {
	if model.DB, err = gorm.Open(mysql.Open(info.Dsn), &gorm.Config{}); err != nil {
		model.Logger.Error(err)
		return err
	}
	return
}
