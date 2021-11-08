package lib

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang-CICD/config"
)

var (
	db *gorm.DB
)

func InitMysql() {
	env := "test"
	if os.Getenv("PROGRAM_ENV") == "pro" {
		env = "pro"
	}
	fmt.Println(fmt.Sprintf("当前处于%s环境", env))

	db = OpenDb(fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.Mysql.User,
		config.Conf.Mysql.Password,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Port,
		config.Conf.Mysql.Db,
	))
}

// 返回数据库连接对象
func DBConn() *gorm.DB {
	return db
}

func OpenDb(source string) *gorm.DB {
	db, err := gorm.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)
	db.LogMode(true)
	return db
}
