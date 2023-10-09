package domain

type UserStruct struct {
	Username string
	Email    string
	Password string
}

type UserRepository interface {
	CheckUser(email string) error
	CreateUser(user UserStruct) error
	GetPassword(email string) (string, error)
	UpdatePassword(email string, password string) error
	GetUserDetails(email string) (UserStruct, error)
}
