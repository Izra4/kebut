package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"kebut/entity"
	"log"
	"os"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	var err error

	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	DB, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal("init db failed", err)
	}
	return DB
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.Penyakit{},
		&entity.Link{})
}
