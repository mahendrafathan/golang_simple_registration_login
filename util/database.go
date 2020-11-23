package util

import (
	"database/sql"
	"fmt"
)

type User struct {
	PhoneNumber int64  `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DoB         string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
}

var (
	db *sql.DB

	queryUserFunc  = queryUser
	insertUserFunc = insertUser
)

func ConnectDB() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=false",
		c.DbHost, c.DbPort, c.DbUser, c.DbPass, c.DbUser)
	fmt.Println(psqlInfo)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database Successfully connected!")
}

func queryUser(email string) (user User, err error) {

	query := fmt.Sprintf(`
		SELECT email, 
		first_name, 
		last_name, 
		phone_number,
		date_of_birth,
		gender 
		FROM users WHERE email=$1
		`)

	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}

	row := stmt.QueryRow(email)
	if row == nil {
		err = fmt.Errorf("row is nil")
		return
	}

	row.Scan(
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.PhoneNumber,
		&user.DoB,
		&user.Gender,
	)
	return
}

func insertUser(user User) (err error) {
	stmt, err := db.Prepare(`INSERT INTO users (email, first_name, last_name, 
		phone_number, date_of_birth, gender) values ($1, $2, $3, $4, $5, $6)`)
	if err != nil {
		return
	}

	_, err = stmt.Exec(user.Email, user.FirstName, user.LastName,
		user.PhoneNumber, user.DoB, user.Gender)
	return
}
