package util

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type SharedResources struct {
	Db *gorm.DB
}

func (sr SharedResources) Generate() (*SharedResources, error) {
	if db, err := getDb(); err != nil {
		return &sr, err
	} else {
		sr.Db = db
	}
	return &sr, nil
}

func getDb() (*gorm.DB, error) {
	dialector := fmt.Sprintf("host=%s user=%s password=%s dbname=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PWD"),
		os.Getenv("POSTGRES_DB"))

	conn, err := gorm.Open(postgres.Open(dialector), &gorm.Config{
		DryRun: false,
	})
	return conn, err
}
