package util

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func Test_queryUser(t *testing.T) {
	var mock sqlmock.Sqlmock
	db, mock = NewMock()
	defer func() {
		db.Close()
	}()

	mockRow := sqlmock.NewRows([]string{
		"email",
		"first_name",
		"last_name",
		"phone_number",
		"date_of_birth",
		"gender",
	}).AddRow(
		"mail", "a", "b", 1, "2020-01-02", "male",
	)

	query := "SELECT (.+) FROM users WHERE (.+)"

	mock.ExpectPrepare(query)
	mock.ExpectQuery(query).WithArgs("mail").WillReturnRows(mockRow)

	wantUser := User{
		Email:       "mail",
		FirstName:   "a",
		LastName:    "b",
		PhoneNumber: 1,
		DoB:         "2020-01-02",
		Gender:      "male",
	}

	resp, err := queryUser("mail")
	assert.Equal(t, false, (err != nil))
	assert.Equal(t, wantUser, resp)

	// test error
	mock.ExpectQuery(query).WithArgs("mail").WillReturnError(fmt.Errorf("error"))
	resp, err = queryUser("mail")
	assert.Equal(t, true, (err != nil))
	assert.Equal(t, User{}, resp)
}

func Test_insertUser(t *testing.T) {
	var mock sqlmock.Sqlmock
	db, mock = NewMock()
	defer func() {
		db.Close()
	}()

	query := "INSERT INTO users (.+) values (.+)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WillReturnResult(sqlmock.NewResult(0, 1))

	user := User{
		Email:       "mail",
		FirstName:   "a",
		LastName:    "b",
		PhoneNumber: 1,
		DoB:         "2020-01-02",
		Gender:      "male",
	}

	err := insertUser(user)
	assert.Equal(t, false, (err != nil))

	// test error
	prep.ExpectExec().WillReturnError(fmt.Errorf("error"))
	err = insertUser(user)
	assert.Equal(t, true, (err != nil))
}
