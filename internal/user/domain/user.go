package domain

type User struct {
	Name  string
	Email string
}

func NewUser(name, email string) (*User, error) {
	return &User{
		Name:  name,
		Email: email,
	}, nil
}
