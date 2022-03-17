package database

import (
	"os"
	"practiceTwo/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenConnection() (db *gorm.DB, err error) {
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: os.Getenv("DSN"),
	}))

	return DB, err
}

func MigrateEntities(entities ...interface{}) error {
	println("Migrate started")
	err := DB.AutoMigrate(entities...)

	return err
}

func Seed(amount uint, entities ...entities.BaseEntity) {
	for _, entity := range entities {
		data := entity.CreateFakeData(amount)

		for _,row := range data  {
			DB.Create(row)
		}
	}
}

