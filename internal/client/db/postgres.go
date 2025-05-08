package client

import (
	"fleet_management/internal/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var DB *sqlx.DB

func InitPostgreSQL() *sqlx.DB {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Cfg.DBHost, config.Cfg.DBPort, config.Cfg.DBUser, config.Cfg.DBPass, config.Cfg.DBName,
	)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		logrus.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(0)

	logrus.Info("PostgreSQL connected successfully")

	DB = db
	return db
}
