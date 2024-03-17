package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	userTable         = "users"
	filmTable         = "film"
	actorTable        = "actor"
	participantsTable = "participants"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = migrate(db)
	if err != nil {
		return nil, err
	}
	return db, err
}

func migrate(db *sqlx.DB) error {
	for _, v := range quries {
		err := migrateTable(db, v)
		if err != nil {
			log.Println("Error with query: ", v)
			return err
		}
	}
	return nil
}

func migrateTable(db *sqlx.DB, query string) error {
	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("error creating database table: %s", err.Error())
	}
	return nil
}
