package user

import (
	"errors"
	"fmt"
	"time"
)

// Only caps lettered functions and structs are exported and fields
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type admin struct {
	email    string
	password string
	User     // anonymous embedding
}

// create methods for struct
func (user *User) OutputUserDetails() {
	// no need for dereferencing pointer to access fields of struct
	fmt.Println("User Details:")
	fmt.Println("First Name:", user.firstName)
	fmt.Println("Last Name:", user.lastName)
	fmt.Println("Birth Date:", user.birthDate)
	fmt.Println("Account Created At:", user.createdAt)
}

// pointer is needed to mutate struct fields
func (user *User) ClearUserDetails() {
	user.firstName = ""
	user.lastName = ""
	user.birthDate = ""
}

// constructor
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("all fields are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password, firstName, lastName, birthDate string) (*admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	user, err := New(firstName, lastName, birthDate)
	if err != nil {
		return nil, err
	}

	return &admin{
		email:    email,
		password: password,
		User:     *user,
	}, nil
}
