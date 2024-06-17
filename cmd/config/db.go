package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "/home/joao/microservicosGOlang/Atualiza-o-cadastral-Emserh/cmd/config/Users.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
}

func GetDatabase() *gorm.DB {
	return DB
}
