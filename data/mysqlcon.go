package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Database 定义数据库连接类型
type Database struct {
	DB             *sql.DB
	dataSourceName string
	err            error
}

func NewDatabase(dataSourceName string) *Database {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Errorf("failed to open database: %v", err)
		return nil
	}

	if err = db.Ping(); err != nil {
		db.Close()
		fmt.Errorf("failed to connect to database: %v", err)
		return nil
	}
	return &Database{DB: db, err: nil}
}
func (db *Database) DatabaseConnect() {
	db.dataSourceName = "root:123456@tcp(127.0.0.1:3306)/gowebtest"
	db = NewDatabase(db.dataSourceName)
	if db.err != nil {
		fmt.Println("DatabaseError:", db.err)
		return
	}
	defer db.DB.Close()
	fmt.Println("Database connected successfully")
}
