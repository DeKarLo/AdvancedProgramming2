package models

import (
	"database/sql"
	"errors"
)

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

type UserRepository struct {
	DB *sql.DB
}

func (ur *UserRepository) InsertUser(user *User) error {
	_, err := ur.DB.Exec("INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)",
		user.Username, user.Email, user.PasswordHash)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetUserByID(userID int) (*User, error) {
	user := &User{}
	err := ur.DB.QueryRow("SELECT id, username, email FROM users WHERE id = $1", userID).
		Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserByUsername(username string) (*User, error) {
	user := &User{}
	err := ur.DB.QueryRow("SELECT id, username, email, password_hash FROM users WHERE username = $1", username).
		Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}
