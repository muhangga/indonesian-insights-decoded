package config

import (
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/muhangga/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg Config) *gorm.DB {

	dbHost := cfg.Get("POSTGRES_HOST")
	dbPort := cfg.Get("POSTGRES_PORT")
	dbUser := cfg.Get("POSTGRES_USER")
	dbPass := cfg.Get("POSTGRES_PASSWORD")
	dbName := cfg.Get("POSTGRES_DB")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&domain.Product{})

	pool, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	pool.SetMaxIdleConns(10)
	pool.SetMaxOpenConns(100)
	pool.SetConnMaxLifetime(time.Hour)

	return db
}

func CloseDatabase(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}

	dbSQL.Close()
}
