package db

import (
	"log"
    "os"
    "path/filepath"
    "io/ioutil"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"events-service/pkg/entities"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:postgres@localhost:5432/postgres"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&entities.Event{})
	db.AutoMigrate(&entities.Notification{})

	// Read the files in the migrations folder
	filepath.Walk("migrations", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalln(err)
			return err
		}

		// Read the contents of each file and execute the SQL commands
		if !info.IsDir() {
			migration, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatalln(err)
				return err
			}
			db.Exec(string(migration))
		}
		return nil
	})

	return db
}
