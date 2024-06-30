package models

import (
	"fmt"
	"go-rest-api/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
    var err error
    config.LoadConfig()

    dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
        config.AppConfig.DBHost,
        config.AppConfig.DBPort,
        config.AppConfig.DBUser,
        config.AppConfig.DBName,
        config.AppConfig.DBPassword,
    )

    DB, err = gorm.Open("postgres", dsn)
    if err != nil {
        panic("Failed to connect to database!")
    }

    DB.AutoMigrate(&Product{})
}
