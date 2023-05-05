package db

import (
	"fmt"
	"log"

	"github.com/RiddhiNahata/shorten-url/app/common/config"
	"github.com/RiddhiNahata/shorten-url/app/common/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// initialize DB configs
func Init(c *config.Config) *gorm.DB {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// creates the tables as per the models
	db.AutoMigrate(&model.Urls{})
	db.AutoMigrate(&model.Users{})

	return db
}
