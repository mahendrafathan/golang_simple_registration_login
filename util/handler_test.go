package util

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_RegisterHandler(t *testing.T) {
	var oriQueryUser = queryUserFunc
	var oriInsertuserFunc = insertUserFunc

	defer func() {
		queryUserFunc = oriQueryUser
		insertUserFunc = oriInsertuserFunc
	}()

	type args struct {
		w          http.ResponseWriter
		r          *http.Request
		queryUser  func(string) (User, error)
		insertUser func(User) error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodPost,
					"http://test.com",
					strings.NewReader(`{
						"phone_number": 0,
						"first_name": "Fathan",
						"last_name": "Mahendra",
						"date_of_birth": "1995-06-13",
						"gender": "male",
						"email": "mahendraf@mail.com"
					}`),
				),

				queryUser: func(string) (User, error) {
					return User{}, nil
				},
				insertUser: func(User) error {
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "error insert",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodPost,
					"http://test.com",
					strings.NewReader(`{
						"phone_number": 0,
						"first_name": "Fathan",
						"last_name": "Mahendra",
						"date_of_birth": "1995-06-13",
						"gender": "male",
						"email": "mahendraf@mail.com"
					}`),
				),

				queryUser: func(string) (User, error) {
					return User{}, nil
				},
				insertUser: func(User) error {
					return fmt.Errorf("error")
				},
			},
			wantErr: true,
		},
		{
			name: "query user found",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodPost,
					"http://test.com",
					strings.NewReader(`{
						"phone_number": 0,
						"first_name": "Fathan",
						"last_name": "Mahendra",
						"date_of_birth": "1995-06-13",
						"gender": "male",
						"email": "mahendraf@mail.com"
					}`),
				),
				queryUser: func(string) (User, error) {
					return User{
						Email: "mail@mail",
					}, nil
				},
			},
			wantErr: true,
		},
		{
			name: "error query user",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodPost,
					"http://test.com",
					strings.NewReader(`{
						"phone_number": 0,
						"first_name": "Fathan",
						"last_name": "Mahendra",
						"date_of_birth": "1995-06-13",
						"gender": "male",
						"email": "mahendraf@mail.com"
					}`),
				),
				queryUser: func(string) (User, error) {
					return User{}, fmt.Errorf("error")
				},
			},
			wantErr: true,
		},
		{
			name: "error query user",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodPost,
					"http://test.com",
					strings.NewReader(`{
						"phone_number": 0,
						"first_name": "Fathan",
						"last_name": "Mahendra",
						"date_of_birth": "1995-06-13",
						"gender": "male",
						"email": "mahendraf@mail.com"
					}`),
				),

				queryUser: func(string) (User, error) {
					return User{}, fmt.Errorf("error")
				},
			},
			wantErr: true,
		},
		{
			name: "error unmarshal",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodPost,
					"http://test.com",
					strings.NewReader(`{ashj}`),
				),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryUserFunc = tt.args.queryUser
			insertUserFunc = tt.args.insertUser
			RegisterHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_LoginHandler(t *testing.T) {
	var oriQueryUser = queryUserFunc

	defer func() {
		queryUserFunc = oriQueryUser
	}()

	type args struct {
		w         http.ResponseWriter
		r         *http.Request
		queryUser func(string) (User, error)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodPost,
					"http://test.com",
					strings.NewReader(`{
						"phone_number": 0,
						"first_name": "Fathan",
						"last_name": "Mahendra",
						"date_of_birth": "1995-06-13",
						"gender": "male",
						"email": "mahendraf@mail.com"
					}`),
				),

				queryUser: func(string) (User, error) {
					return User{
						Email: "mahendraf@mail.com",
					}, nil
				},
			},
			wantErr: false,
		},
		{
			name: "query user not found",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodPost,
					"http://test.com",
					strings.NewReader(`{
						"phone_number": 0,
						"first_name": "Fathan",
						"last_name": "Mahendra",
						"date_of_birth": "1995-06-13",
						"gender": "male",
						"email": "mahendraf@mail.com"
					}`),
				),
				queryUser: func(string) (User, error) {
					return User{}, nil
				},
			},
			wantErr: true,
		},
		{
			name: "error query user",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodPost,
					"http://test.com",
					strings.NewReader(`{
						"phone_number": 0,
						"first_name": "Fathan",
						"last_name": "Mahendra",
						"date_of_birth": "1995-06-13",
						"gender": "male",
						"email": "mahendraf@mail.com"
					}`),
				),
				queryUser: func(string) (User, error) {
					return User{}, fmt.Errorf("error")
				},
			},
			wantErr: true,
		},
		{
			name: "error query user",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodPost,
					"http://test.com",
					strings.NewReader(`{
						"phone_number": 0,
						"first_name": "Fathan",
						"last_name": "Mahendra",
						"date_of_birth": "1995-06-13",
						"gender": "male",
						"email": "mahendraf@mail.com"
					}`),
				),

				queryUser: func(string) (User, error) {
					return User{}, fmt.Errorf("error")
				},
			},
			wantErr: true,
		},
		{
			name: "error unmarshal",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodPost,
					"http://test.com",
					strings.NewReader(`{ashj}`),
				),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryUserFunc = tt.args.queryUser
			LoginHandler(tt.args.w, tt.args.r)
		})
	}
}
