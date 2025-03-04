package models

import (
	"errors"
	"strings"
)

type UserAdress struct {
	ID        uint64 `json:"id,omitempty"`
	UserID    uint64 `json:"user_id,omitempty"`
	ZipCode   string `json:"zip_code,omitempty"`
	Street    string `json:"street,omitempty"`
	//TODO: Change number to int
	Number    string `json:"number,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Country   string `json:"country,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type UserConfig struct {
	ID        uint64 `json:"id,omitempty"`
	UserID    uint64 `json:"user_id,omitempty"`
	Theme     string `json:"theme,omitempty"`
	Language  string `json:"language,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type User struct {
	ID         uint64     `json:"id,omitempty"`
	Name       string     `json:"name,omitempty"`
	Nickname   string     `json:"nickname,omitempty"`
	Email      string     `json:"email,omitempty"`
	Password   string     `json:"password,omitempty"`
	CreatedAt  string     `json:"created_at,omitempty"`
	UserAdress UserAdress `json:"user_adress,omitempty"`
	UserConfig UserConfig `json:"user_config,omitempty"`
}

func (u *User) validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Nickname == "" {
		return errors.New("nickname is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	// TODO: Change the zip code to a number and add a validation
	if u.UserAdress.ZipCode == "" {
		return errors.New("zip code is required")
	}
	if u.UserAdress.Street == "" {
		return errors.New("street is required")
	}
	//TODO: Change number to int
	if u.UserAdress.Number == "" {
		return errors.New("number is required")
	}
	if u.UserAdress.City == "" {
		return errors.New("city is required")
	}
	if u.UserAdress.State == "" {
		return errors.New("state is required")
	}
	if u.UserAdress.Country == "" {
		return errors.New("country is required")
	}
	if u.UserConfig.Theme == "" {
		u.UserConfig.Theme = "01"
	}
	if u.UserConfig.Language == "" {
		u.UserConfig.Language = "PT-BR"
	}	
	return nil
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nickname = strings.TrimSpace(u.Nickname)
	u.Email = strings.TrimSpace(u.Email)
	u.UserAdress.ZipCode = strings.TrimSpace(u.UserAdress.ZipCode)
	u.UserAdress.Number = strings.TrimSpace(u.UserAdress.Number)
}

// Prepare calls the validation and formatting functions
func (u *User)Prepare() error {
	if err := u.validate(); err != nil {
		return err
	}
	u.format()
	return nil
}
