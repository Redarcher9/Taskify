package repository

import (
	"database/sql"
	"fmt"
	"taskify/domain"
)

type UserRepository struct {
	db      *sql.DB
	timeout int
}

func NewUserRepository(database *sql.DB, timeout int) domain.UserRepository {
	return &UserRepository{
		db:      database,
		timeout: timeout,
	}
}

func (ur *UserRepository) CheckUser(email string) bool {
	var useremail string
	rows, err := ur.db.Query("Select email from public.Users where email=($1)", email)
	if err != nil {
		return true
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&useremail)
		if err != nil {
			return true
		}
	}
	if useremail == "" {
		return false
	}
	return true
}

func (ur *UserRepository) CreateUser(user domain.UserStruct) error {
	fmt.Println("before inserting" + user.Email)
	_, err := ur.db.Exec("INSERT INTO public.Users(username,email,password) VALUES($1,$2,$3)", user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetPassword(email string) (string, error) {
	var password string
	rows, err := ur.db.Query("SELECT password from public.Users where email=($1)", email)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&password)
		if err != nil {
			fmt.Println(password)
			return "", err
		}
	}
	return password, err
}

func (ur *UserRepository) UpdatePassword(email string, password string) error {
	_, err := ur.db.Exec("UPDATE public.Users SET password=$2 where email=$1", email, password)
	if err != nil {
		return err
	}
	return err
}

func (ur *UserRepository) GetUserDetails(email string) (domain.UserStruct, error) {
	var User domain.UserStruct
	rows, err := ur.db.Query("SELECT password from public.Users where email=($1)", email)
	if err != nil {
		return User, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&User.Username, &User.Email)
		if err != nil {
			return User, err
		}
	}
	return User, err
}
