package main

import (
	"database/sql"
)

var db *sql.DB

// Por ahora es debug la función, cambiar luego
func InitializeDB() error {
	var err error
	db, err = sql.Open("sqlite3", "test.db")
	if err != nil {
		return err
	}

	_, err = db.Exec(`
	DROP TABLE IF EXISTS User;
	DROP TABLE IF EXISTS Chat;
	DROP TABLE IF EXISTS User_Chat;
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS User (
		id INTEGER PRIMARY KEY,
		email TEXT NOT NULL,
		username TEXT NOT NULL,
		description TEXT DEFAULT '',
		password TEXT NOT NULL,
		profile_picture BLOB DEFAULT NULL,
		session_id TEXT DEFAULT NULL
	);
	
	CREATE TABLE IF NOT EXISTS Chat (
		id INTEGER PRIMARY KEY
	);

	CREATE TABLE IF NOT EXISTS User_Chat (
		user_id INTEGER,
		chat_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES User(id)
		FOREIGN KEY(chat_id) REFERENCES Chat(id)
	);
	`)
	if err != nil {
		return err
	}

	return nil
}

func CreateUser(newUser UserInput) error {
	_, err := db.Exec(`
	INSERT INTO User (email, username, password)
	VALUES (?, ?, ?)
	`, newUser.Email, newUser.Username, newUser.Password)
	if err != nil {
		return err
	}

	return nil
}

type User struct {
	Email string
	Username string
	ProfilePic []byte
	Description string
}

func ReadUser(id int) (User, error) {
	var email string
	var username string
	var profilePic []byte
	var description string

	query := `
	SELECT email, username, profile_picture, description
	FROM User
	WHERE id = ?;`
	err := db.QueryRow(query, id).Scan(&email, &username, &profilePic, &description)
	if err != nil {
		return User{}, err
	}

	return User{email, username, profilePic, description}, nil
}
