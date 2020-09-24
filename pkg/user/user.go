package user

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reddit/pkg/errors"
)

type User struct {
	Id       string `json:"id"`
	Login    string `json:"username"`
	Password string `json:"password,omitempty"`
}

func NewUser(id, login, password string) *User {
	return &User{
		Id:       id,
		Login:    login,
		Password: password,
	}
}

func (u *User) Parse(body *io.ReadCloser, funcName string) *errors.Error {
	defer (*body).Close()
	bodyData, _ := ioutil.ReadAll(*body)
	// it will never happen
	//if bodyError != nil {
	//	log.Printf("%s error: %s\n", funcName, bodyError.Error())
	//	return errors.New(http.StatusInternalServerError, errors.ReadErr)
	//}

	jsonError := json.Unmarshal(bodyData, u)
	if jsonError != nil {
		log.Printf("%s error: %s\n", funcName, jsonError.Error())
		return errors.New(http.StatusBadRequest, errors.InvalidBody)
	}

	return nil
}

func (u *User) GetFromContext(ctx context.Context, key string) {
	data := ctx.Value(key).(*User)
	u.Id = data.Id
	u.Login = data.Login
}
