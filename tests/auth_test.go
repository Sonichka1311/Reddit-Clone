package tests

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"reddit/pkg/auth"
	"reddit/pkg/errors"
	"reddit/pkg/user"
	"reddit/tests/data"
	"testing"
)

type AuthTestCase = data.AuthTestCase

func TestRegister(t *testing.T) {
	defer func() {
		auth.GenerateTime = auth.GenerateExpireTime
	}()

	register := AuthTestCase{
		Data:    *data.GetTestUserReader(),
		Type:    "POST",
		Handler: "/api/register",
	}

	tests := []AuthTestCase{
		data.GetAuthTestCase(register, nil, "register", 0, http.StatusOK, data.GetTestAccessToken()),
		data.GetAuthTestCase(register, data.GetWrongTestUserReader(), "register", 0, http.StatusBadRequest, `Invalid body`),
		data.GetAuthTestCase(register, &bytes.Reader{}, "register", 0, http.StatusBadRequest, `Invalid body`),
		data.GetAuthTestCase(register, nil, "register", 1, http.StatusBadRequest, `User with this login already exists`),
		data.GetAuthTestCaseWithAction(
			data.GetAuthTestCase(register, nil, "register", 0, http.StatusInternalServerError, `Internal server error`),
			func() {
				auth.GenerateTokenFunc = func(userData user.User) (*auth.AccessToken, *errors.Error) {
					return nil, errors.New(http.StatusInternalServerError, errors.InternalError)
				}
			},
			func() {
				auth.GenerateTokenFunc = auth.Generate
			},
		),
		data.GetAuthTestCase(register, nil, "register", 2, http.StatusInternalServerError, `Internal server error`),
		data.GetAuthTestCase(register, nil, "register", 3, http.StatusInternalServerError, `Internal server error`),
		data.GetAuthTestCase(register, nil, "register", 4, http.StatusInternalServerError, `Internal server error`),
		data.GetAuthTestCase(register, nil, "register", 5, http.StatusInternalServerError, `Internal server error`),
		data.GetAuthTestCase(register, nil, "register", 6, http.StatusInternalServerError, `Internal server error`),
		data.GetAuthTestCase(register, nil, "register", 7, http.StatusInternalServerError, `Internal server error`),
	}

	for idx, test := range tests {
		log.Printf("TEST CASE %d\n", idx)
		req, err := http.NewRequest(test.Type, test.Handler, &test.Data)
		if err != nil {
			t.Fatalf("TEST %d of TestRegister failed create new request for %s with error %s", idx, test.Handler, err)
		}

		userHandler := data.GetTestUserHandler(test.DB)
		if test.Action != nil {
			test.Action()
		}

		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(userHandler.Register)
		handler.ServeHTTP(recorder, req)

		if ok, err := data.CheckStatus(recorder.Code, test.Status); !ok {
			t.Errorf("TEST %d of TestRegister failed with error: %s", idx, *err)
		}
		if ok, err := data.CheckResponse(recorder.Body.String(), test.Response); !ok {
			t.Errorf("TEST %d of TestRegister failed with error: %s", idx, *err)
		}

		test.DB.Close()
		if test.Unaction != nil {
			test.Unaction()
		}
	}
}

func TestLogin(t *testing.T) {
	defer func() {
		auth.GenerateTime = auth.GenerateExpireTime
	}()

	login := AuthTestCase{
		Data:    *data.GetTestUserReader(),
		Type:    "POST",
		Handler: "/api/login",
	}

	tests := []AuthTestCase{
		data.GetAuthTestCase(login, nil, "login", 0, http.StatusOK, data.GetTestAccessToken()),
		data.GetAuthTestCase(login, data.GetWrongTestUserReader(), "login", 0, http.StatusBadRequest, `Invalid body`),
		data.GetAuthTestCase(login, &bytes.Reader{}, "login", 0, http.StatusBadRequest, `Invalid body`),
		data.GetAuthTestCase(login, nil, "login", 1, http.StatusUnauthorized, `No user with this login`),
		data.GetAuthTestCaseWithAction(
			data.GetAuthTestCase(login, nil, "login", 0, http.StatusInternalServerError, `Internal server error`),
			func() {
				auth.GenerateTokenFunc = func(userData user.User) (*auth.AccessToken, *errors.Error) {
					return nil, errors.New(http.StatusInternalServerError, errors.InternalError)
				}
			},
			func() {
				auth.GenerateTokenFunc = auth.Generate
			},
		),
		data.GetAuthTestCase(login, nil, "login", 2, http.StatusInternalServerError, `Internal server error`),
		data.GetAuthTestCase(login, nil, "login", 3, http.StatusUnauthorized, `Invalid password`),
		data.GetAuthTestCase(login, nil, "login", 4, http.StatusInternalServerError, `Internal server error`),
		data.GetAuthTestCase(login, nil, "login", 5, http.StatusInternalServerError, `Internal server error`),
		data.GetAuthTestCase(login, nil, "login", 6, http.StatusInternalServerError, `Internal server error`),
		data.GetAuthTestCase(login, nil, "login", 7, http.StatusOK, data.GetTestAccessToken()),
		data.GetAuthTestCase(login, nil, "login", 8, http.StatusInternalServerError, `Internal server error`),
		data.GetAuthTestCase(login, nil, "login", 9, http.StatusInternalServerError, `Internal server error`),
	}

	for idx, test := range tests {
		log.Printf("TEST CASE %d\n", idx)
		req, err := http.NewRequest(test.Type, test.Handler, &test.Data)
		if err != nil {
			t.Fatalf("TEST %d of TestRegister failed create new request for %s with error %s", idx, test.Handler, err)
		}

		userHandler := data.GetTestUserHandler(test.DB)

		if test.Action != nil {
			test.Action()
		}

		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(userHandler.Login)
		handler.ServeHTTP(recorder, req)

		if ok, err := data.CheckStatus(recorder.Code, test.Status); !ok {
			t.Errorf("TEST %d of TestLogin failed with error: %s", idx, *err)
		}
		if ok, err := data.CheckResponse(recorder.Body.String(), test.Response); !ok {
			t.Errorf("TEST %d of TestLogin failed with error: %s", idx, *err)
		}

		test.DB.Close()
		if test.Unaction != nil {
			test.Unaction()
		}
	}
}

func TestNewUser(t *testing.T) {
	usr := user.User{
		Id:       "1",
		Login:    "test login",
		Password: "test password",
	}

	log.Printf("TEST CASE 0\n")
	newUser := user.NewUser(usr.Id, usr.Login, usr.Password)
	if usr != *newUser {
		t.Errorf("TEST 0 of TestLogin failed with error: Unexpected body: got: %v, expected: %v", *newUser, usr)

	}
}
