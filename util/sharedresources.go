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

func (sr SharedResources) Generate() error {
	if db, err := getDb(); err != nil {
		return err
	} else {
		sr.Db = db
	}
	return nil
}

func getDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DryRun: false,
	})
	return conn, err
}
