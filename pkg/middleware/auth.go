package middleware

import (
	"context"
	"log"
	"net/http"
	"reddit/pkg/auth"
	"reddit/pkg/errors"
	"strings"
	"time"
)

type Middleware struct {
	Sessions *auth.Database
}

var (
	All          = "/"
	Static       = "/static/"
	Register     = "/api/register"
	Login        = "/api/login"
	Posts        = "/api/posts"
	Post         = "/api/post/"
	User         = "/api/user"
	Method       = "method"
	DontContains = "dontContains"
	IsString     = "isString"
	Conditions   = map[string]map[string]interface{}{
		All: {
			IsString: All,
		},
		Static:   {},
		Register: {},
		Login:    {},
		Posts: {
			Method: http.MethodGet,
		},
		Post: {
			Method:       http.MethodGet,
			DontContains: []string{"upvote", "downvote", "unvote"},
		},
		User: {},
	}
)

func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for path, cond := range Conditions {
			if strings.Contains(r.URL.Path, path) {
				needToAuth := false
				for condition, value := range cond {
					switch condition {
					case Method:
						if r.Method != value {
							needToAuth = true
						}
					case DontContains:
						patterns := value.([]string)
						for _, pat := range patterns {
							if strings.Contains(r.URL.Path, pat) {
								needToAuth = true
							}
						}
					case IsString:
						if r.URL.Path != path {
							needToAuth = true
						}
					}
				}
				if !needToAuth {
					next.ServeHTTP(w, r)
					return
				}
			}
		}

		authorization := strings.TrimPrefix(r.Header.Get("authorization"), "Bearer ")
		log.Printf("Trying to get user by token %s", authorization)

		usr, expire, err := m.Sessions.GetByToken(authorization)
		if err != nil {
			log.Printf("Auth middleware: error from DB: %s\n", err.Error())
			http.Error(w, errors.InternalError, http.StatusInternalServerError)
			return
		}
		if expire < time.Now().Unix() {
			http.Error(w, errors.InvalidToken, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), auth.TokenKey, usr)
		log.Printf("Login with user id %s\n", usr.Id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
