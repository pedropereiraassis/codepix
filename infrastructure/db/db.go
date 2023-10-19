package db

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/pedropereiraassis/codepix/domain/model"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "gorm.io/driver/sqlite"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

func ConnectDB(env string) *gorm.DB {
	var dsn string
	var db *gorm.DB
	var err error

	if env != "test" {
		dsn = os.Getenv("DSN")
		db, err = gorm.Open(os.Getenv("DB_TYPE"), dsn)
	} else {
		dsn = os.Getenv("DSN_TEST")
		db, err = gorm.Open(os.Getenv("DB_TYPE_TEST"), dsn)
	}

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv("DEBUG") == "true" {
		db.LogMode(true)
	}

	if os.Getenv("AUTO_MIGRATE_DB") == "true" {
		db.AutoMigrate(&model.Bank{}, &model.Account{}, &model.PixKey{}, &model.Transaction{})
	}

	return db
}