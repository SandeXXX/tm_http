package models

import "time"

type User struct {
	Id        string    `json:"id"`
	Email     string    `json:"email"`
	PWhash    string    `json:"pwhash"`
	Name      string    `json:"name"`
	Fullname  string    `json:"fullname"`
	Role      string    `json:"role"`
	Contacts  string    `json:"contacts"`
	Geo       string    `json:"geo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FormData struct {
	Email     string    `form:"email"`     //noreq
	Name      string    `form:"name"`      //noreq
	FullName  string    `form:"full_name"` //noreq
	Contacts  string    `form:"contacts"`  //req
}

type SignUpParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpUser struct {
	Email    string `json:"email"`
	HashedPW []byte `json:"hashedpw"`
}

type Geo struct {
}
