// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/shanbay/gobay/testdata/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	nickname *string
	username *string
}

// SetNickname sets the nickname field.
func (uc *UserCreate) SetNickname(s string) *UserCreate {
	uc.nickname = &s
	return uc
}

// SetNillableNickname sets the nickname field if the given value is not nil.
func (uc *UserCreate) SetNillableNickname(s *string) *UserCreate {
	if s != nil {
		uc.SetNickname(*s)
	}
	return uc
}

// SetUsername sets the username field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.username = &s
	return uc
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if uc.nickname == nil {
		v := user.DefaultNickname
		uc.nickname = &v
	}
	if uc.username == nil {
		return nil, errors.New("ent: missing required field \"username\"")
	}
	return uc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	var (
		builder = sql.Dialect(uc.driver.Dialect())
		u       = &User{config: uc.config}
	)
	tx, err := uc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(user.Table).Default()
	if value := uc.nickname; value != nil {
		insert.Set(user.FieldNickname, *value)
		u.Nickname = *value
	}
	if value := uc.username; value != nil {
		insert.Set(user.FieldUsername, *value)
		u.Username = *value
	}

	id, err := insertLastID(ctx, tx, insert.Returning(user.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	u.ID = int(id)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return u, nil
}
