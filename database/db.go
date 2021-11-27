package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"github.com/felipeas314/go/models"
)

/*DB is connected database object*/
var DB *gorm.DB

func Setup() {
	host := "localhost"
	port := "5432"
	dbname := "teste_go"
	user := "postgres"
	password := "postgres"

	db, err := gorm.Open("postgres",
			"host="+host+
				" port="+port+
				" user="+user+
				" dbname="+dbname+
				" sslmode=disable password="+password)

	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)
	db.AutoMigrate([]models.Book{})
	DB = db
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}