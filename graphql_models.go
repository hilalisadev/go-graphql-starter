// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package dexp

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Entity interface {
	IsEntity()
}

type Filter struct {
	Search  *string `json:"search"`
	OrderBy *string `json:"order_by"`
	Limit   int     `json:"limit"`
	Offset  int     `json:"offset"`
}

type NewUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

type SystemLog struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	Message   string    `json:"message"`
	UserID    int       `json:"user_id"`
	Read      bool      `json:"read"`
	CreatedAt time.Time `json:"created_at"`
}

type Token struct {
	Token     string    `json:"token"`
	User      *User     `json:"user"`
	ExpiredAt time.Time `json:"expired_at"`
}

type UserRole string

const (
	UserRoleAdministrator UserRole = "administrator"
	UserRoleAuthenticated UserRole = "authenticated"
)

var AllUserRole = []UserRole{
	UserRoleAdministrator,
	UserRoleAuthenticated,
}

func (e UserRole) IsValid() bool {
	switch e {
	case UserRoleAdministrator, UserRoleAuthenticated:
		return true
	}
	return false
}

func (e UserRole) String() string {
	return string(e)
}

func (e *UserRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserRole", str)
	}
	return nil
}

func (e UserRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
