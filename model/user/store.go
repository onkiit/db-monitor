package user

import "context"

type Store interface {
	AddUser(context.Context, *User) error
	GetById(context.Context, string) (*User, error)
	GetByUsername(context.Context, string) (*User, error)
	GetByEmail(context.Context, string) (*User, error)
}
