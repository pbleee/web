package config

import (
	"fmt"

	"github.com/pbleee/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:2224@tcp(127.0.0.1:3306)/learn_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("gagal koneksi")
	}
	fmt.Println("berhasil terhubung")

	db.AutoMigrate(
		&models.User{},
		&models.Enrollment{},
		&models.Course{},
		&models.Module{},
		&models.Grade{},
	)

	DB = db
}
