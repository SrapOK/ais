package repository

import (
	"fmt"
	"kis/internal/config"
	model "kis/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewPostgresDB(cfg config.Postgres) (*gorm.DB, error) {

	config := &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), config)

	if err != nil {
		return nil, fmt.Errorf("cannot open connection: %s", err.Error())
	}

	db.AutoMigrate(&model.Vacancy{}, &model.Candidate{}, &model.Customer{}, &model.Employee{}, &model.Grade{})

	return db, nil
}
