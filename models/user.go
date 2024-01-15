package models

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/flexconsulta")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) HashPassword() error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) SaveUser() error {
	u.ID = uuid.New().String()

	if isUsernameTaken(u.Username) {
		return errors.New("username already taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	result, err := db.Exec("INSERT INTO users (id, username, password) VALUES (?, ?, ?)", u.ID, u.Username, string(hashedPassword))
	if err != nil {
		log.Println("Error executing INSERT query:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting RowsAffected:", err)
		return err
	}

	log.Printf("User saved successfully. Rows affected: %d\n", rowsAffected)

	u.Password = ""

	return nil
}

func isUsernameTaken(username string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&count)
	if err != nil {
		log.Println("Error checking username availability:", err)
		return true
	}

	return count > 0
}

func UpdateUser(userID string, updatedUser User) error {
	_, err := db.Exec("UPDATE users SET username = ?, password = ? WHERE id = ?", updatedUser.Username, updatedUser.Password, userID)
	return err
}

func DeleteUser(userID string) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", userID)
	return err
}

func GetUserByID(userID string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, password FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUsers() ([]User, error) {
	rows, err := db.Query("SELECT id, username, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
