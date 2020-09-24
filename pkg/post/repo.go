package post

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"reddit/pkg/errors"
	"reddit/pkg/interfaces"
	"reddit/pkg/user"
	"sync"
	"time"
)

type Repo struct {
	Storage *Database
	Mutex   *sync.RWMutex
}

var (
	GeneratePostId      = GenerateId
	GenerateCommentId   = GenerateId
	GenerateCreatedTime = GenerateTime
)

func GenerateId() bson.ObjectId {
	return bson.NewObjectId()
}

func GenerateTime(tm time.Time) string {
	return tm.Format("2006-01-02T15:04:05Z")
}

func NewRepo(db interfaces.MongoSession) *Repo {
	return &Repo{
		Storage: &Database{
			Posts:    db.DB("reddit").C("posts"),
			Votes:    db.DB("reddit").C("votes"),
			Comments: db.DB("reddit").C("comments"),
			Users:    db.DB("reddit").C("users"),
		},
		Mutex: &sync.RWMutex{},
	}
}

func (r *Repo) Add(post *Post, user user.User) *errors.Error {
	user.Password = ""
	post.Init()
	post.Author = user
	post.Id = GeneratePostId()

	log.Printf("Trying to add post with id %s\n", post.Id.Hex())

	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	insertError := r.Storage.AddPost(post)
	return insertError
}

func (r *Repo) GetAll() ([]*Post, *errors.Error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	log.Println("Trying to get all posts")
	return r.Storage.GetAll(bson.M{})
}

func (r *Repo) Get(id string) (*Post, *errors.Error) {
	log.Printf("Trying to get post with id %s\n", id)

	if !bson.IsObjectIdHex(id) {
		log.Printf("Id %s is not hex\n", id)
		return nil, errors.New(http.StatusBadRequest, errors.InvalidParams)
	}

	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.Storage.GetPost(bson.ObjectIdHex(id), true)
}

func (r *Repo) GetByCategory(category string) ([]*Post, *errors.Error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	log.Printf("Trying to get all posts by category %s\n", category)
	return r.Storage.GetAll(bson.M{"category": category})
}

func (r *Repo) GetByUser(user string) ([]*Post, *errors.Error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	log.Printf("Trying to get all posts by user %s\n", user)
	ids, err := r.Storage.GetPostsByUsername(user)
	if err != nil {
		return nil, err
	}
	log.Printf("All posts by user %s is %v\n", user, ids)
	return r.Storage.GetAll(
		bson.M{
			"id": bson.M{
				"$in": ids,
			},
		},
	)
}

func (r *Repo) AddComment(comment string, postId string, author user.User) (*Post, *errors.Error) {
	author.Password = ""
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	newComment := Comment{
		Post:    bson.ObjectIdHex(postId),
		Created: GenerateCreatedTime(time.Now()),
		Author:  &author,
		Body:    comment,
		Id:      GenerateCommentId(),
	}
	addErr := r.Storage.AddComment(&newComment)
	if addErr != nil {
		return nil, addErr
	}

	return r.Storage.GetPost(bson.ObjectIdHex(postId), false)
}

func (r *Repo) DeleteComment(postId, commentId string) (*Post, *errors.Error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	deleteErr := r.Storage.DeleteComment(bson.ObjectIdHex(postId), bson.ObjectIdHex(commentId))
	if deleteErr != nil {
		return nil, deleteErr
	}

	return r.Storage.GetPost(bson.ObjectIdHex(postId), false)
}

func (r *Repo) GetVote(postId, userId string) (*Vote, *errors.Error) {
	return r.Storage.GetVote(postId, userId)
}

func (r *Repo) ChangeVote(postId string, user *user.User, turn int) (*Post, *errors.Error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	changeErr := r.Storage.ChangeVote(bson.ObjectIdHex(postId), user.Id, turn)
	if changeErr != nil {
		return nil, changeErr
	}

	return r.Storage.GetPost(bson.ObjectIdHex(postId), false)
}

func (r *Repo) Delete(postId string) *errors.Error {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.Storage.DeletePost(bson.ObjectIdHex(postId))
}
