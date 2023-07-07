package repository

import (
	"log"

	"github.com/AbdulrahmanDaud10/PaymentSystemBackEnd/pkg/api"
	// Dialect for mysql database
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	db.AutoMigrate(&api.User{}, &api.BankAccount{}, &api.Product{})
	return db

}
