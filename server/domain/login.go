package domain

type Loginstruct struct {
	Email    string
	Password string
}

type LoginUsecase interface {
	Login(ls Loginstruct) error
}
