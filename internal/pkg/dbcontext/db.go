package dbcontext

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DB struct {
	*gorm.DB
}

func NewDB(dsn string) *DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return &DB{
		db,
	}
}
