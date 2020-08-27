package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // inject pq
	"github.com/williamchang80/sea-apd/infrastructure/config"
)

const driver = "postgres"

// Postgres conn config
func Postgres() *gorm.DB {
	conf := config.New()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.PgHost, conf.PgPort, conf.PgUser, conf.PgPassword, conf.PgName)
	db, err := gorm.Open(driver, psqlInfo)
	if err != nil {
		panic(err.Error())
	}

	return db
}
