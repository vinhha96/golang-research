package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetDBConnection(dialect, url string) (*gorm.DB, error) {
	fmt.Println("[MySQL] Open connection to database")
	db, err := gorm.Open(dialect, url)

	return db, err
}
