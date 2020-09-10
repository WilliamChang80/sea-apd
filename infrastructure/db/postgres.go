package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // inject postgres
	"github.com/joho/godotenv"
)

const driver = "postgres"

var db *gorm.DB

// Postgres ...
func Postgres() *gorm.DB {
	if db == nil {
		godotenv.Load()
		host := os.Getenv("PG_HOST")
		port := os.Getenv("PG_PORT")
		dbname := os.Getenv("PG_NAME")
		user := os.Getenv("PG_USER")
		password := os.Getenv("PG_PASSWORD")
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		d, err := gorm.Open(driver, psqlInfo)
		db = d
		if err != nil {
			panic(err)
		}
	}
	return db
}
