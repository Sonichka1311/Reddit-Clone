package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"reddit/pkg/auth"
	"reddit/pkg/database"
	"reddit/pkg/middleware"
	"reddit/pkg/user"
	"reddit/tests/data"
	"testing"
)

type MiddlewareTestCase = data.MiddlewareTestCase
type TokenTestCase = data.TokenTestCase

func TestGenerateExpireTime(t *testing.T) {
	timeResp := auth.GenerateTime(12)
	if len(timeResp) == 0 {
		t.Errorf("TEST %d of TestGenerateExpireTime failed with error: incorrect format of result", 0)
	}
}

func TestMiddlewareAuth(t *testing.T) {
	testUser := user.User{
		Login:    "Sonia",
		Password: "testpassword",
	}

	tests := []MiddlewareTestCase{
		data.GetMiddlewareTestCase("/", http.MethodGet),
		data.GetMiddlewareTestCase("/static/", http.MethodGet),
		data.GetMiddlewareTestCase("/api/register", http.MethodPost),
		data.GetMiddlewareTestCase("/api/login", http.MethodPost),
		data.GetMiddlewareTestCase("/api/posts/", http.MethodGet),
		data.GetMiddlewareTestCase("/api/posts/music", http.MethodGet),
		data.GetMiddlewareTestCase("/api/post/1", http.MethodGet),
		data.GetMiddlewareTestCase("/api/user/test", http.MethodGet),
		data.GetMiddlewareTestCase("/api/posts", http.MethodPost),
		data.GetMiddlewareTestCase("/api/post/1", http.MethodPost),
		data.GetMiddlewareTestCase("/api/post/1/1", http.MethodDelete),
		data.GetMiddlewareTestCase("/api/post/1/upvote", http.MethodGet),
		data.GetMiddlewareTestCase("/api/post/1/downvote", http.MethodGet),
		data.GetMiddlewareTestCase("/api/post/1/unvote", http.MethodGet),
		data.GetMiddlewareTestCase("/api/post/1", http.MethodDelete),
		data.GetMiddlewareTestCaseWithStatus("/api/post/1/1", http.MethodDelete, 1, http.StatusInternalServerError),
		data.GetMiddlewareTestCaseWithStatus("/api/post/1/1", http.MethodDelete, 2, http.StatusUnauthorized),
	}

	for idx, test := range tests {
		log.Printf("TEST CASE %d\n", idx)
		req, err := http.NewRequest(test.Method, test.Handler, nil)
		if err != nil {
			t.Fatalf("TEST %d of TestMiddlewareAuth failed create new request for %s with error: %s", idx, test.Handler, err)
		}
		req.Header.Set("authorization", "Bearer testToken")

		testFunc := func(w http.ResponseWriter, r *http.Request) {
			if idx > 7 {
				var usr user.User
				usr.GetFromContext(r.Context(), auth.TokenKey)
				if usr.Login != testUser.Login && usr.Password != testUser.Password {
					http.Error(w, "Users dont match!", http.StatusInternalServerError)
				}
			}
		}

		db := data.InitAuthDb(test.Int)
		middlewareHandler := middleware.Middleware{
			Sessions: (*auth.Database)(database.InitDB(db)),
		}

		recorder := httptest.NewRecorder()
		handler := middlewareHandler.Auth(http.HandlerFunc(testFunc))
		handler.ServeHTTP(recorder, req)

		status := http.StatusOK
		if test.Status != 0 {
			status = test.Status
		}
		if ok, err := data.CheckStatus(recorder.Code, status); !ok {
			t.Errorf("TEST %d of TestMiddlewareAuth failed with error: %s", idx, *err)
		}
		db.Close()
	}
}

func TestGenerateToken(t *testing.T) {
	testUser := user.User{
		Id:       "1",
		Login:    "Login",
		Password: "Password",
	}

	tests := []TokenTestCase{
		data.GetTokenTestCase(&testUser, false),
		data.GetTokenTestCaseWithAction(
			&testUser,
			true,
			func() {
				auth.SigningToken = nil
			},
			func() {
				auth.SigningToken = auth.SigningTokenValue
			},
		),
	}

	for idx, test := range tests {
		if test.Action != nil {
			test.Action()
		}
		_, err := auth.Generate(test.Data)
		if err != nil && !test.IsErr {
			t.Errorf("TEST %d of TestGenerateToken failed with error: %s", idx, err.Message)
		} else if err == nil && test.IsErr {
			t.Errorf("TEST %d of TestGenerateToken failed with error: expected some error, got ok", idx)
		}
		if test.Action != nil {
			test.Unaction()
		}
	}
}
