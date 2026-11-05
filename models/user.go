package models

import (
	"errors"
	"lukasrar/go-rest-api-with-db/db"
	"lukasrar/go-rest-api-with-db/utils"
)

type User struct {
	ID       int64
	Email    string `binding: "required"`
	Password string `binding: "required"`
}

func (u User) Save() error {

	query := `
		INSERT INTO users(email, password) VALUES(?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPass, errorPass := utils.HashPassword(u.Password)

	if errorPass != nil {
		return errorPass
	}

	result, errExec := stmt.Exec(u.Email, hashedPass)

	if errExec != nil {
		return errExec
	}

	userId, errId := result.LastInsertId()

	u.ID = userId

	return errId

}

func GetUsers() ([]User, error) {

	query := `
		SELECT * FROM users
	`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u User) ValidateCredentials() error {

	query := `
		SELECT password FROM users WHERE email = ?
	`

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPass string

	err := row.Scan(&retrievedPass)
	if err != nil {
		return errors.New("Credentials invalid")
	}

	passIsValid := utils.CheckPassword(u.Password, retrievedPass)

	if !passIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
