package db

import (
	"errors"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type BookDB struct {
	ID             string `gorm:"primaryKey"`
	Title          string
	CreatedAt      time.Time
	CurrentChapter int
	CurrentPage    int
}

type DB struct {
	db   *gorm.DB
	Book *BookDB
}

var GlobalDB *DB

func New(path string) {
	book := BookDB{}
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger: logger.Discard,
	})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(book)
	GlobalDB = &DB{db: db, Book: &book}
}

func (d *DB) CreateBook(book *BookDB) {
	d.db.Create(book)
}

func (d *DB) GetBook(id string) (BookDB, error) {
	var book BookDB
	err := d.db.First(&book, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return book, err
	}
	return book, nil
}

func (d *DB) UpdateBook(book BookDB) {
	d.db.Save(&book)
}

func (d *DB) DeleteBook(id string) {
	d.db.Delete(&BookDB{}, id)
}

func (d *DB) GetAllBooks() []BookDB {
	var books []BookDB
	d.db.Find(&books)
	return books
}
