package user

import (
	"database/sql"
	"log"
	"net/http"
	"reddit/pkg/errors"
	"sync"
)

type Repo struct {
	Storage *Database
	Mutex   *sync.RWMutex
}

func NewRepo(db *Database) *Repo {
	return &Repo{
		Storage: db,
		Mutex:   &sync.RWMutex{},
	}
}

func (r *Repo) Add(login, password string) (*User, *errors.Error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	log.Printf("Try to get user %s from database\n", login)
	_, dbError := r.Storage.GetIdByLogin(login)
	switch dbError {
	case nil:
		log.Printf("User %s already exists\n", login)
		return nil, errors.New(http.StatusBadRequest, errors.UserExists)
	case sql.ErrNoRows:
		log.Printf("Try to insert user %s to database\n", login)
		dbError = r.Storage.InsertNewUser(login, password)
		if dbError != nil {
			log.Println(dbError.Error())
			return nil, errors.New(http.StatusInternalServerError, errors.InternalError)
		}
		log.Printf("Try to get id for user %s\n", login)
		id, dbError := r.Storage.GetIdByLogin(login)
		if dbError != nil {
			log.Println(dbError.Error())
			return nil, errors.New(http.StatusInternalServerError, errors.InternalError)
		}
		log.Printf("Added new user %s with id %s\n", login, *id)
		return NewUser(*id, login, password), nil
	default:
		log.Printf("Error from database: %s", dbError.Error())
		return nil, errors.New(http.StatusInternalServerError, errors.InternalError)
	}
}

func (r *Repo) Authorize(login, password string) (*User, *errors.Error) {
	r.Mutex.Lock()
	usr, dbError := r.Storage.GetUserByLogin(login)
	r.Mutex.Unlock()
	switch dbError {
	case sql.ErrNoRows:
		log.Printf("User %s not found\n", login)
		return nil, errors.New(http.StatusUnauthorized, errors.NoUser)
	case nil:
		log.Printf("User id: %s, login: %s", usr.Id, login)
		if usr.Password != password {
			return nil, errors.New(http.StatusUnauthorized, errors.InvalidPassword)
		}
		return usr, nil
	default:
		log.Println(dbError.Error())
		return nil, errors.New(http.StatusInternalServerError, errors.InternalError)
	}
}
