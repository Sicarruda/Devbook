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
	defer u.db.Close()

	userResult, err := u.addUserToTable(user)
	if err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	user.UserAdress.UserID = userResult
	_, err = u.addUserAdrresstoTable(user.UserAdress)
	if err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	user.UserConfig.UserID = userResult
	_, err = u.addUserConfigtoTable(user.UserConfig)
	if err != nil {
		tx.Rollback()
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

	userConfigResult, err := statement.Exec(userConfig.UserID, userConfig.Theme, userConfig.Language)
	if err != nil {
		return nil, err
	}
	return userConfigResult, nil
}

// Find user by name or nickname or email
func (u user) FindUsersByNameOrNick(nameOrNick string) ([]models.User, error) {
	nameOrNick = "%" + nameOrNick + "%"

	lines, err := u.db.Query(
		"select id, name, nickname, email, crated_at from users where name like ? or nickname like ? or email like ?",
		nameOrNick, nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User
		if err := lines.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt, &user.UserAdress, &user.UserConfig); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
