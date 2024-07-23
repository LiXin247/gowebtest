package data

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database 定义数据库连接类型
type Database struct {
	DB  *gorm.DB
	err error
}

func (db *Database) DatabaseConnect() {
	db.DB, db.err = gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/gowebtest"), &gorm.Config{})
	if db.err != nil {
		panic("failed to connect database")
	}
	fmt.Printf("The database connection is successful")
}
