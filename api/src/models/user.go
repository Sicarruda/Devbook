package models

type UserAdress struct {
	ID       uint64 `json:"id,omitempty"`
	UserID   uint64 `json:"user_id,omitempty"`
	ZipCode  string `json:"zip_code,omitempty"`
	Street   string `json:"street,omitempty"`
	Number   string `json:"number,omitempty"`
	City     string `json:"city,omitempty"`
	State    string `json:"state,omitempty"`
	Country  string `json:"country,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type UserConfig struct {
	ID       uint64 `json:"id,omitempty"`
	UserID   uint64 `json:"user_id,omitempty"`
	Theme    string `json:"theme,omitempty"`
	Language string `json:"language,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type User struct {
	ID       uint64 `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Nickname     string `json:"nick,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UserAdress UserAdress `json:"user_adress,omitempty"`
	UserConfig UserConfig `json:"user_config,omitempty"`	
}