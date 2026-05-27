package config

import (
	"fmt"

	//"github.com/MatheusMikio/Nexus/internal/domain/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initPostgreSQL(cfg Config) (*gorm.DB, error) {
	dbConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode)

	db, err := gorm.Open(postgres.Open(dbConnection), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// err = db.AutoMigrate(
	// 	&schemas.Goal{},
	// 	&schemas.Task{},
	// 	&schemas.User{},
	// )

	// if err != nil {
	// 	return nil, err
	// }

	return db, nil
}
