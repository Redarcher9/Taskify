package domain

type UserStruct struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type UserRepository interface {
	CheckUser(email string) bool
	CreateUser(user UserStruct) error
	GetPassword(email string) (string, error)
	UpdatePassword(email string, password string) error
	GetUserDetails(email string) (UserStruct, error)
}
