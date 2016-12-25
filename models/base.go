package models

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Config struct {
	DB_HOST     string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
}

func ReadConfig() Config {
	var configfile = "config"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	log.Print(config.DB_NAME)
	return config
}

// InitDB creates and migrates the database
func InitDB() (*gorm.DB, error) {
	var err error

	var dbConfig = ReadConfig()

	dbUser := dbConfig.DB_USER
	dbPassword := dbConfig.DB_PASSWORD
	dbName := dbConfig.DB_NAME
	connectionString := dbUser + ":" + dbPassword + "@/" + dbName + "?charset=utf8&parseTime=True"
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)
	db.AutoMigrate(&User{}, &Task{}, &Status{})

	return db, nil
}
