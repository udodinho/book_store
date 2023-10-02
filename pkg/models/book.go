package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/udodinho/bookstore/pkg/config"
	"gorm.io/gorm"
)

type Books struct {
	ID        uint    `gorm:"primarykey;autoIncrement" json:"id"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
}

var DB *gorm.DB

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})

	return err
}

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	configFile := &config.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	db, err := config.Connect(configFile)

	if err != nil {
		log.Fatal("Could not connect to db", err)
	}

	err = MigrateBooks(db)

	if err != nil {
		log.Fatal("Could not migrate db", err)
	}

	DB = db
	fmt.Println("Database connected successfully")

}

func (b *Books) CreateBook() (*Books, error) {
	DB.Create(&b)
	return b, nil
}

func GetAllBooks() ([]Books, error) {
	var Books []Books
	DB.Find(&Books)
	return Books, nil
}

func GetBookbyID(id int64) (*Books, *gorm.DB, error) {
	var book Books
	db := DB.Where("id=?", id).First(&book)
	return &book, db, nil
}

func DeleteBook(id int64) (Books, error) {
	var book Books
	DB.Where("id", id).Delete(book)
	return book, nil
}
