package models


import "gorm.io/driver/mysql"
import "gorm.io/gorm"

var DB *gorm.DB

func ConnectToDb() error {
	dns := "root:@tcp(localhost:3306)/golang?multiStatements=true"
	db , err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db

	return nil
}