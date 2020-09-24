package data

import (
	"bytes"
	"database/sql"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
	"reddit/pkg/user"
	mock_interfaces "reddit/tests/mock"
	"testing"
)

type AuthTestCase struct {
	Data     bytes.Reader
	Type     string
	Handler  string
	DB       *sql.DB
	Status   int
	Response string
	Action   func()
	Unaction func()
}

type TokenTestCase struct {
	Data     user.User
	IsErr    bool
	Action   func()
	Unaction func()
}

type MiddlewareTestCase struct {
	Handler string
	Method  string
	Int     int
	Status  int
}

type PostsTestCase struct {
	Method      string
	Handler     string
	NeedAuth    bool
	Data        bytes.Reader
	Collections [4]*mock_interfaces.MockMongoCollection
	Response    string
	Status      int
}

type PostTest struct {
	Cases      []PostsTestCase
	Func       string
	Handler    string
	Controller *gomock.Controller
}

func (pt PostTest) Test(t *testing.T) {
	for idx, test := range pt.Cases {
		log.Printf("TEST CASE %d\n", idx)
		req, err := http.NewRequest(test.Method, test.Handler, &test.Data)
		if err != nil {
			t.Fatalf("TEST %d of TestAdd failed create new request for %s with error %s", idx, test.Handler, err)
		}

		postHandler := GetHandler(pt.Controller, test.Collections)

		var handler http.HandlerFunc
		switch pt.Func {
		case "add post":
			handler = postHandler.AddPost
		case "get all posts":
			handler = postHandler.GetAll
		case "get by category":
			handler = postHandler.GetByCategory
		case "get post":
			handler = postHandler.GetPost
		case "add comment":
			handler = postHandler.AddComment
		case "delete comment":
			handler = postHandler.DeleteComment
		case "upvote":
			handler = postHandler.Upvote
		case "downvote":
			handler = postHandler.Downvote
		case "unvote":
			handler = postHandler.Unvote
		case "delete post":
			handler = postHandler.DeletePost
		case "get by user":
			handler = postHandler.GetByUser
		}

		router := mux.NewRouter()
		router.HandleFunc(pt.Handler, handler)
		recorder := httptest.NewRecorder()

		if test.NeedAuth {
			router.ServeHTTP(recorder, req.WithContext(GetTestUserCtx(req.Context())))
		} else {
			router.ServeHTTP(recorder, req)
		}

		if ok, err := CheckStatus(recorder.Code, test.Status); !ok {
			t.Errorf("TEST %d of TestAdd failed with error: %s", idx, *err)
		}
		if ok, err := CheckResponse(recorder.Body.String(), test.Response); !ok {
			t.Errorf("TEST %d of TestAdd failed with error: %s", idx, *err)
		}
	}
}
