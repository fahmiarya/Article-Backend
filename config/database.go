package config

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"

    "github.com/fahmiarya/Article-Backend/internal/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
    dbUser := os.Getenv("DB_USER")
    if dbUser == "" {
        dbUser = "root"
    }
    dbPass := os.Getenv("DB_PASS")
    if dbPass == "" {
        dbPass = "root"
    }
    dbHost := os.Getenv("DB_HOST")
    if dbHost == "" {
        dbHost = "localhost"
    }
    dbPort := os.Getenv("DB_PORT")
    if dbPort == "" {
        dbPort = "3306"
    }
    dbName := os.Getenv("DB_NAME")
    if dbName == "" {
        dbName = "article"
    }

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
    gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to GORM database: %v", err)
    }
    // AutoMigrate the Post model
    if err := gormDB.AutoMigrate(&model.Post{}); err != nil {
        log.Fatalf("AutoMigrate failed: %v", err)
    }
    DB = gormDB
    return gormDB
}
