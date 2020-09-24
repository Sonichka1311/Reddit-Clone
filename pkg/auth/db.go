package auth

import (
	"log"
	"reddit/pkg/database"
	"reddit/pkg/user"
	"strconv"
	"time"
)

type Database database.Database

var GenerateTime = GenerateExpireTime

func GenerateExpireTime(hours int) string {
	return strconv.Itoa(int(time.Now().Add(time.Hour * time.Duration(hours)).Unix()))
}

func (d *Database) GetByToken(token string) (*user.User, int64, error) {
	row := d.QueryRow("SELECT id, login, expire FROM sessions WHERE token = ? LIMIT 1", token)
	var usr user.User
	var expire int64
	dbError := row.Scan(&usr.Id, &usr.Login, &expire)
	return &usr, expire, dbError
}

func (d *Database) GetByUser(login string) (string, int64, error) {
	row := d.QueryRow("SELECT token, expire FROM sessions WHERE login = ? LIMIT 1", login)
	var expire int64
	var token string
	dbError := row.Scan(&token, &expire)
	return token, expire, dbError
}

func (d *Database) AddNewSession(token string, usr *user.User) error {
	log.Printf("Trying to add session for user id %s\n", usr.Id)
	_, dbError := d.Exec(
		"INSERT INTO sessions (token, id, login, expire) VALUES (?, ?, ?, ?)",
		token, usr.Id, usr.Login, GenerateTime(12))
	return dbError
}

func (d *Database) UpdateSession(token string, usr *user.User) error {
	_, dbError := d.Exec(
		"UPDATE sessions SET token = ?, expire = ? WHERE login = ?",
		token, GenerateTime(12), usr.Login)
	return dbError
}
