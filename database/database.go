package database

import (
	"fmt"
	"github.com/evsamsonov/trading-history-server/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

func Init() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(fmt.Sprintf("database: failed to connect: %v", err))
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Bar{})
	db.AutoMigrate(&model.Symbol{})
	db.AutoMigrate(&model.Resolution{})
}
