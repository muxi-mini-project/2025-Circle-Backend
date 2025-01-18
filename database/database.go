package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB //必须大写表示公开
func InitDB() {
	var err error

	DB, err = sql.Open("mysql", "root:2388287244@tcp(112.126.68.22:3306)/circle?parseTime=true&charset=utf8mb4&loc=Local") //?...统一时间
	if err != nil {
		log.Fatalf("连接失败 %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("连接测试失败 %v", err)
	}
}