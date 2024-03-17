package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/rkBekzat/films/internal/model"
)

type account struct {
	db *sqlx.DB
}

func NewAccount(db *sqlx.DB) Account {
	return &account{db: db}
}

func (a *account) CreateUser(user *model.User) error {
	fmt.Println("before add: ", user)
	query := fmt.Sprintf("INSERT INTO %s (username, email, gender, role, password) VALUES ($1, $2, $3, $4, $5)", userTable)
	_, err := a.db.Exec(query, user.Username, user.Email, user.Gender, user.Role, user.Password)
	return err
}

func (a *account) EmailExist(email string) (bool, error) {
	var cnt int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE email=$1", userTable)
	log.Println("QUERY: ", query)
	row := a.db.QueryRow(query, email)
	if err := row.Scan(&cnt); err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (a *account) GetUser(username, password string) (*model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id, role FROM %s WHERE email=$1 AND password=$2", userTable)
	err := a.db.Get(&user, query, username, password)
	return &user, err
}
