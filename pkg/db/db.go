package db

import (
	"go/adv-demo/configs"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.DB.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Db{db}
}
