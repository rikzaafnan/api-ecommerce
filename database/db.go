package database

import (
	"api-ecommerce/config"
	"api-ecommerce/payment"
	"api-ecommerce/product"
	"api-ecommerce/transaction"
	"api-ecommerce/user"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeDB() *gorm.DB {

	configLoadENV := config.LoadENV()

	var username_db, password_db, database_db, host_db, port_db string

	username_db = configLoadENV.USERNAMEDB
	password_db = configLoadENV.PASSWORDDB
	database_db = configLoadENV.SCHEMADB
	host_db = configLoadENV.HOSTDB
	port_db = configLoadENV.PORTDB

	// Connection to database mysql
	// dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username_db, password_db, host_db, port_db, database_db)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	autoMigrate(db)

	return db
}

// func GetDB() *gorm.DB {
// 	return db
// }

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{}, &product.Product{}, &transaction.Transaction{}, &transaction.TransactionDetail{}, &payment.Payment{})
}
