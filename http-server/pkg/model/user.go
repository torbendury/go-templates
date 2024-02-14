package model

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func New(id int, firstName string, lastName string) *User {
	return &User{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
	}
}
