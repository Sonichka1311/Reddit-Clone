package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"reddit/pkg/auth"
	"reddit/pkg/errors"
	"reddit/pkg/user"
)

type UserHandler struct {
	Repo     *user.Repo
	Sessions *auth.Database
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var usr user.User
	parseError := usr.Parse(&r.Body, "Handlers: Register")
	if parseError != nil {
		http.Error(w, parseError.Message, parseError.Status)
		return
	}
	log.Printf("Trying to register with user %s\n", usr.Login)

	userData, addError := h.Repo.Add(usr.Login, usr.Password)
	if addError != nil {
		http.Error(w, addError.Message, addError.Status)
		return
	}
	log.Printf("User %s has been added to database\n", usr.Login)

	token, tokenError := auth.GenerateTokenFunc(*userData)
	if tokenError != nil {
		http.Error(w, tokenError.Message, tokenError.Status)
		return
	}
	log.Printf("Generated token %s for user %s\n", token.Token, usr.Login)

	_, _, getError := h.Sessions.GetByToken(token.Token)
	if getError != sql.ErrNoRows {
		if getError != nil {
			log.Printf("Login: Error from DB: %s\n", getError.Error())
		} else {
			log.Printf("Session with token %s already exists\n", token.Token)
		}
		http.Error(w, errors.InternalError, http.StatusInternalServerError)
		return
	}

	dbError := h.Sessions.AddNewSession(token.Token, userData)
	if dbError != nil {
		log.Printf("Add new session for user %s error: %s", usr.Login, dbError.Error())
		http.Error(w, errors.InternalError, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(token)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var usr user.User
	parseError := usr.Parse(&r.Body, "Handlers: Login")
	if parseError != nil {
		http.Error(w, parseError.Message, parseError.Status)
		return
	}

	userData, authError := h.Repo.Authorize(usr.Login, usr.Password)
	if authError != nil {
		http.Error(w, authError.Message, authError.Status)
		return
	}

	token, tokenError := auth.GenerateTokenFunc(*userData)
	if tokenError != nil {
		http.Error(w, tokenError.Message, tokenError.Status)
		return
	}

	_, _, getError := h.Sessions.GetByToken(token.Token)
	if getError != sql.ErrNoRows {
		if getError != nil {
			log.Printf("Login: Error from DB: %s\n", getError.Error())
		} else {
			log.Printf("Session with token %s already exists\n", token)
		}
		http.Error(w, errors.InternalError, http.StatusInternalServerError)
		return
	}

	_, _, getError = h.Sessions.GetByUser(usr.Login)
	if getError != nil && getError != sql.ErrNoRows {
		log.Printf("Login: Error from DB: %s\n", getError.Error())
		http.Error(w, errors.InternalError, http.StatusInternalServerError)
		return
	} else if getError == nil {
		log.Printf("Try to update session with user %s\n", usr.Login)
		dbError := h.Sessions.UpdateSession(token.Token, userData)
		if dbError != nil {
			log.Printf("Login: Error from DB: %s\n", dbError.Error())
			http.Error(w, errors.InternalError, http.StatusInternalServerError)
			return
		}
	} else if getError == sql.ErrNoRows {
		log.Printf("Try to add session for user %s\n", usr.Login)

		dbError := h.Sessions.AddNewSession(token.Token, userData)
		if dbError != nil {
			log.Printf("Login: Error from DB: %s\n", dbError.Error())
			http.Error(w, errors.InternalError, http.StatusInternalServerError)
			return
		}
	}

	json.NewEncoder(w).Encode(token)
}
