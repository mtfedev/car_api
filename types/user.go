package types

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost     = 12
	minUserName    = 2
	maxUserName    = 64
	minUserSurname = 2
	maxUserSurname = 64
	minPasswodLen  = 7
	minEmail       = 3
)

type User struct {
	ID      string `bson:"_id,omitempty" json:"id,omitempty"`
	Name    string `bson:"name" json:"name"`
	Surname string `bson:"surname" json:"surname"`
	Email   string `bson:"email" json:"email"`
	EncPwd  string `bson:"encpwd" json:"-"`
	Phone   string `bson:"phone" json:"phone"`
}

type CreateUserParams struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func (p *CreateUserParams) ToUser() (*User, error) {

	p.Email = strings.ToLower(p.Email)

	var (
		encpwd, err = bcrypt.GenerateFromPassword([]byte(p.Password), bcryptCost)
	)

	if err != nil {
		return nil, err
	}

	var (
		user = &User{
			Name:    p.Name,
			Surname: p.Surname,
			Email:   p.Email,
			EncPwd:  string(encpwd),
			Phone:   p.Phone,
		}
	)

	return user, nil
}
func (p *CreateUserParams) ValidateName() bool {
	if len(p.Name) < minUserName || len(p.Name) > maxUserName {
		return false
	}
	if len(p.Surname) < minUserSurname || len(p.Surname) > maxUserSurname {
		return false
	}

	return false
}
func (params CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.Password) < minPasswodLen {
		errors["password"] = fmt.Sprintf("password length should be at lesast %d characters", minPasswodLen)
	}
	if !isEmailValid(params.Email) {
		errors["email"] = fmt.Sprintf("email is invalid %d", minEmail)
	}
	return errors
}

func IsValidPassword(encpw, pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encpw), []byte(pw)) == nil
}
func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
