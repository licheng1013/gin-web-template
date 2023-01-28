package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

var Db *gorm.DB

func init() {
	MysqlInit()
}

func MysqlInit() {
	// 日志打印
	newLogger := logger.Default
	log.Println("Mysql:初始化！")
	dsn := "root:root@tcp(192.168.101.8:3306)/t_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	v, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 定义表前缀
			SingularTable: true, // true不在表后面+ s，
		},
	})
	if err != nil {
		log.Panic(err)
	}
	Db = v
}
