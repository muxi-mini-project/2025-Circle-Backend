package database

import (
	"log"
	"time"
	"circle/models"
	"io/ioutil"
	"encoding/json"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB // 必须大写表示公开
type Config struct {
	DatabaseDSN string `json:"database_dsn"`
}

func InitDB() {
	var err error

	// 配置 MySQL 连接字符串
	data, _ := ioutil.ReadFile("data.json")
	var config Config
	_ = json.Unmarshal(data, &config)
	dsn := config.DatabaseDSN
	fmt.Println("数据库连接字符串: ", dsn)

	// 初始化 GORM 数据库实例
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 设置日志级别
	})
	if err != nil {
		
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 获取原生的 SQL DB 以进行额外配置
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("获取 SQL DB 实例失败: %v", err)
	}

	// 设置连接池配置
	sqlDB.SetMaxOpenConns(100)                // 最大连接数
	sqlDB.SetMaxIdleConns(10)                 // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(10 * time.Minute) // 连接最大生命周期

	// 测试连接
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("连接测试失败: %v", err)
	}

	if err := DB.AutoMigrate(
		&models.User{},
		&models.Practice{},
		&models.PracticeComment{},
		&models.PracticeOption{},
		&models.UserPractice{},
		&models.Practicehistory{},
		&models.Test{},
		&models.TestComment{},
		&models.TestOption{},
		&models.TestQuestion{},
		&models.Testhistory{},
		&models.Top{},
		&models.Circle{},
		&models.FollowCircle{},
		&models.SearchHistory{},
	); err != nil {
		log.Fatalf("自动迁移失败: %v", err)
	}
	

}
