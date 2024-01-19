package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // 记得要加入驱动
	"time"
)

var DB *gorm.DB

// 数据库连接
func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		panic(fmt.Sprintf("数据库连接错误: %s\nerror: %s", connString, err))
	}
	db.LogMode(true)

	if gin.Mode() == "release" {
		db.LogMode(false)
	}

	db.SingularTable(true)                       // 表名不加s
	db.DB().SetMaxIdleConns(20)                  // 设置连接池
	db.DB().SetMaxOpenConns(100)                 // 设置最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 30) // 设置连接等待时间

	DB = db

	fmt.Println("数据库连接成功！")

	migration()
}
