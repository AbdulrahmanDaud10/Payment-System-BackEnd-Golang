package repository

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//dialect for mysql database
)

// DatabaseConnection -> connection to Database
func DatabaseConnection() *gorm.DB {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to Database")
		return nil
	}

	// db.AutoMigrate(&model.User{}, &model.Accounts{}, &model.Order{})
	return db

}
