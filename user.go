package user

// User : represents a user with a name and login

type User struct {
	Name  string
	Login string
}

// New : create a new user
func New(name, login string) *User {
	return &User{Name: name, Login: login}
}

// UpdateName : update the user's name
func (u *User) UpdateName(name string) {
	u.Name = name
}

// GetName : get the user's name
func (u *User) GetName() string {
	return u.Name
}
