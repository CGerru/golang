package userRepository

import (
	"database/sql"
	"log"

	"local/jwt-api/models"
)

type UserRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Signup inserts one record in database
func (u UserRepository) Signup(db *sql.DB, user models.User) models.User {
	stmt := "insert into users (email, password) values($1, $2) RETURNING id;"
	err := db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)

	logFatal(err)

	user.Password = ""

	return user
}

// Login retrieves one record from database
func (u UserRepository) Login(db *sql.DB, user models.User) (models.User, error) {
	row := db.QueryRow("select * from users where email=$1", user.Email)

	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return user, err
	}

	return user, nil
}
