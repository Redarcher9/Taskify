package domain

type SignUp interface {
	SignUp(user UserStruct) error
}
