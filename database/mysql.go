package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"memoryDataBase/config"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB(cfg config.MySQLConfig) (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return DB, nil
}
