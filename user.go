package user

// User : represents a user with a name and login

type User struct {
	Name  string
	Login string
	Email string
	age   int
}

// New : create a new user
func New(name, login, email string) *User {
	return &User{
		Name:  name,
		Login: login,
		Email: email,
	}
}

// UpdateName : update the user's name
func (u *User) UpdateName(name string) {
	u.Name = name
}

// UpdateAge : update age of the user
func (u *User) UpdateAge(age int) {
	u.age = age
}

// GetAge : get age of the user
func (u User) GetAge() int {
	return u.age
}

// GetName : get the user's name
func (u *User) GetName() string {
	return u.Name
}
