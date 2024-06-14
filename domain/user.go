package domain

import "time"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(id *int, firstName, LastName, Email string) *User {
	user := &User{
		FirstName: firstName,
		LastName:  LastName,
		Email:     Email,
	}

	if id != nil {
		user.ID = *id
	}

	return user
}

func (u *User) Update(firstName, lastName, email string) {
	u.FirstName = firstName
	u.LastName = lastName
	u.Email = email
}
