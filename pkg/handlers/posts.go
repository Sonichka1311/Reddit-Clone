package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"reddit/pkg/auth"
	"reddit/pkg/errors"
	"reddit/pkg/post"
	"reddit/pkg/user"
)

type PostHandler struct {
	Repo *post.Repo
}

func (h *PostHandler) AddPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var usr user.User
	usr.GetFromContext(r.Context(), auth.TokenKey)

	var pst post.Post
	parseError := pst.Parse(r.Body, "Handlers: AddPost")
	if parseError != nil {
		http.Error(w, parseError.Message, parseError.Status)
		return
	}

	dbError := h.Repo.Add(&pst, usr)
	if dbError != nil {
		http.Error(w, dbError.Message, dbError.Status)
		return
	}

	json.NewEncoder(w).Encode(pst)
}

func (h *PostHandler) GetAll(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	posts, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, err.Message, err.Status)
		return
	}
	json.NewEncoder(w).Encode(posts)
}

func (h *PostHandler) GetByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	posts, err := h.Repo.GetByCategory(mux.Vars(r)["category"])
	if err != nil {
		http.Error(w, err.Message, err.Status)
		return
	}
	json.NewEncoder(w).Encode(posts)
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pst, err := h.Repo.Get(mux.Vars(r)["post_id"])
	if err != nil {
		http.Error(w, err.Message, err.Status)
		return
	}

	json.NewEncoder(w).Encode(pst)
}

func (h *PostHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var usr user.User
	usr.GetFromContext(r.Context(), auth.TokenKey)

	defer r.Body.Close()
	bodyData, _ := ioutil.ReadAll(r.Body)
	// never happens ?
	//if bodyError != nil {
	//	log.Printf("Handlers: AddComment error: %s\n", bodyError.Error())
	//	http.Error(w, errors.ReadErr, http.StatusInternalServerError)
	//	return
	//}

	var comment struct {
		Comment string `json:"comment,required"`
	}
	jsonError := json.Unmarshal(bodyData, &comment)
	if jsonError != nil {
		log.Printf("Handlers: AddComment error: %s\n", jsonError.Error())
		http.Error(w, errors.InvalidBody, http.StatusBadRequest)
		return
	}

	pst, addError := h.Repo.AddComment(comment.Comment, mux.Vars(r)["post_id"], usr)
	if addError != nil {
		http.Error(w, addError.Message, addError.Status)
		return
	}
	json.NewEncoder(w).Encode(pst)
}

func (h *PostHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var usr user.User
	usr.GetFromContext(r.Context(), auth.TokenKey)

	pst, deleteError := h.Repo.DeleteComment(mux.Vars(r)["post_id"], mux.Vars(r)["comment_id"])
	if deleteError != nil {
		log.Printf("Handlers: DeleteComment error: %s\n", deleteError.Message)
		http.Error(w, deleteError.Message, deleteError.Status)
		return
	}

	json.NewEncoder(w).Encode(pst)
}

func (h *PostHandler) Upvote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var usr user.User
	usr.GetFromContext(r.Context(), auth.TokenKey)

	pst, err := h.Repo.ChangeVote(mux.Vars(r)["post_id"], &usr, 1)
	if err != nil {
		http.Error(w, err.Message, err.Status)
		return
	}
	json.NewEncoder(w).Encode(pst)
}

func (h *PostHandler) Downvote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var usr user.User
	usr.GetFromContext(r.Context(), auth.TokenKey)

	pst, err := h.Repo.ChangeVote(mux.Vars(r)["post_id"], &usr, -1)
	if err != nil {
		http.Error(w, err.Message, err.Status)
		return
	}
	json.NewEncoder(w).Encode(pst)
}

func (h *PostHandler) Unvote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var usr user.User
	usr.GetFromContext(r.Context(), auth.TokenKey)

	vote, err := h.Repo.GetVote(mux.Vars(r)["post_id"], usr.Id)
	if err != nil {
		log.Printf("Unvote error: can't get vote for user %s\n", usr.Id)
		http.Error(w, err.Message, err.Status)
		return
	}

	pst, err := h.Repo.ChangeVote(mux.Vars(r)["post_id"], &usr, vote.Vote)
	if err != nil {
		http.Error(w, err.Message, err.Status)
		return
	}
	json.NewEncoder(w).Encode(pst)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var usr user.User
	usr.GetFromContext(r.Context(), auth.TokenKey)

	deleteError := h.Repo.Delete(mux.Vars(r)["post_id"])
	if deleteError != nil {
		log.Printf("Handlers: DeleteComment error: %s\n", deleteError.Message)
		http.Error(w, deleteError.Message, deleteError.Status)
		return
	}
	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{"success"})
}

func (h *PostHandler) GetByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	posts, getError := h.Repo.GetByUser(mux.Vars(r)["user_login"])
	if getError != nil {
		http.Error(w, getError.Message, getError.Status)
		return
	}
	json.NewEncoder(w).Encode(posts)
}
