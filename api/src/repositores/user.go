package repositores

import (
	"api/src/models"
	"database/sql"
)

type user struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *user {
	return &user{db}
}

func (u user) CreateUser(user models.User) (models.User, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return models.User{}, err
	}

	userResult, err := u.addUserToTable(user)
	if err != nil {
		return models.User{}, err
	}

	user.UserAdress.UserID = userResult
	_, err = u.addUserAdrresstoTable(user.UserAdress)
	if err != nil {
		return models.User{}, err
	}

	user.UserConfig.UserID = userResult
	_, err = u.addUserConfigtoTable(user.UserConfig)
	if err != nil {
		return models.User{}, err
	}

	err = tx.Commit()
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Insert into users table
func (u user) addUserToTable(user models.User) (uint64, error) {
	statement, err := u.db.Prepare(
		"insert into users (name, nickname, email, password) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	userResult, err := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	userID, err := userResult.LastInsertId()
    if err != nil {
        return 0, err
    }

	return uint64(userID), nil
}

// Insert into userAdress table
func (u user) addUserAdrresstoTable(userAdress models.UserAdress) (sql.Result, error) {
	statement, err := u.db.Prepare(
		"insert into userAdress (user_id, zip_code, street, number, city, state, country) values (?, ?, ?, ?, ?, ?, ?)",
	)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	userAdressResult, err := statement.Exec(userAdress.UserID, userAdress.ZipCode, userAdress.Street, userAdress.Number, userAdress.City, userAdress.State, userAdress.Country)
	if err != nil {
		return nil, err
	}

	return userAdressResult, nil
}

// Insert into userConfig table
func (u user) addUserConfigtoTable(userConfig models.UserConfig) (sql.Result, error) {
	statement, err := u.db.Prepare(
		"insert into userConfig (user_id, theme, language) values (?, ?, ?)",
	)
	if err != nil {
		return nil, err
	}
	statement.Close()

	userConfigResult, err := statement.Exec(userConfig.UserID, userConfig.Theme, userConfig.Language)
	if err != nil {
		return nil, err
	}
	return userConfigResult, nil
}
