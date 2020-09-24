package user

import (
	"reddit/pkg/database"
)

type Database database.Database

func (d *Database) GetIdByLogin(login string) (*string, error) {
	row := d.QueryRow("SELECT HEX(id) FROM users WHERE login = ? LIMIT 1", login)
	var id string
	dbError := row.Scan(&id)
	if dbError == nil {
		//idStr := strconv.Itoa(id)
		return &id, nil
	}
	return nil, dbError
}

func (d *Database) GetUserByLogin(login string) (*User, error) {
	row := d.QueryRow("SELECT HEX(id), login, password FROM users WHERE login = ? LIMIT 1", login)
	usr := User{}
	dbError := row.Scan(&usr.Id, &usr.Login, &usr.Password)
	return &usr, dbError
}

func (d *Database) InsertNewUser(login, password string) error {
	_, dbError := d.Exec("INSERT INTO users (login, password) VALUES (?, ?)", login, password)
	return dbError
}
