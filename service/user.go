package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserName       string
	HashedPassword string
	Role           string
}

func NewUser(username string, password string, role string) (*User, error) {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}

	user := &User{
		UserName: username,
		HashedPassword: string(hasedPassword),
		Role: role,
	}

	return user, nil
}

func (user *User) IsCorrecPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return err == nil
}

func (user *User) Clone() *User {
	return &User {
		UserName: user.UserName,
		HashedPassword: user.HashedPassword,
		Role: user.Role,
	}
}
