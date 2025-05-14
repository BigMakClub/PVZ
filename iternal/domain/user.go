package domain

import (
	"errors"
	"github.com/google/uuid"
)

type Role string

const (
	Employee  Role = "employee"
	Moderator Role = "moderator"
)

var ErrInvalidRole = errors.New("invalid role")

func (r Role) IsRoleValid() bool {
	if r == Employee || r == Moderator {
		return true
	}
	return false
}

type User struct {
	Id       uuid.UUID
	Email    string
	Role     Role
	Password string
}

func NewUser(email string, password string, role Role) (*User, error) {
	if !role.IsRoleValid() {
		return nil, ErrInvalidRole
	}

	return &User{
		Id:       uuid.New(),
		Email:    email,
		Role:     role,
		Password: password,
	}, nil
}
