package user

import (
	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
)

type app struct {
	DBConfig dbConfig
	DB       *gorm.DB
}

type dbConfig struct {
	Driver string
	DSN    string
}

func newDBConfig() dbConfig {
	var dbConfig dbConfig

	path := "./config/db.toml"
	if _, err := toml.DecodeFile(path, &dbConfig); err != nil {
		panic(err)
	}

	return dbConfig
}

func newDB(dbConfig dbConfig) *gorm.DB {
	db, err := gorm.Open(dbConfig.Driver, dbConfig.DSN)
	if err != nil {
		panic(err)
	}

	return db
}

func NewApp(dbConfig dbConfig, db *gorm.DB) app {
	return app{dbConfig, db}
}

var App app

func InitApp() {
	App = initApp()
}
