package post

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reddit/pkg/errors"
	"reddit/pkg/user"
	"time"
)

type Vote struct {
	Post bson.ObjectId `json:"-" bson:"post_id"`
	User string        `json:"user" bson:"user"`
	Vote int           `json:"vote" bson:"vote"`
}

type Comment struct {
	Post    bson.ObjectId `json:"-" bson:"post_id"`
	Created string        `json:"created" bson:"created"`
	Author  *user.User    `json:"author" bson:"author"`
	Body    string        `json:"body" bson:"body"`
	Id      bson.ObjectId `json:"id" bson:"id"`
}

type Post struct {
	Score      int           `json:"score" bson:"score"`
	Views      int           `json:"views" bson:"views"`
	Type       string        `json:"type,required" bson:"type"`
	Title      string        `json:"title,required" bson:"title"`
	Url        string        `json:"url,omitempty" bson:"url"`
	Text       string        `json:"text,omitempty" bson:"text"`
	Author     user.User     `json:"author" bson:"author"`
	Category   string        `json:"category,required" bson:"category"`
	Votes      []Vote        `json:"votes"`
	Comments   []Comment     `json:"comments"`
	Created    string        `json:"created" bson:"created"`
	Percentage uint          `json:"upvotePercentage" bson:"upvotePercentage"`
	Id         bson.ObjectId `json:"id" bson:"id"`
}

type UserToPostId struct {
	Username string        `bson:"username"`
	PostId   bson.ObjectId `bson:"post_id"`
}

func (p *Post) Parse(body io.ReadCloser, funcName string) *errors.Error {
	defer body.Close()
	bodyData, _ := ioutil.ReadAll(body)
	// never happens ?
	//if bodyError != nil {
	//	log.Printf("%s error: %s\n", funcName, bodyError.Error())
	//	return errors.New(http.StatusInternalServerError, errors.ReadErr)
	//}

	jsonError := json.Unmarshal(bodyData, p)
	if jsonError != nil {
		log.Printf("%s error: %s\n", funcName, jsonError.Error())
		return errors.New(http.StatusBadRequest, errors.InvalidBody)
	}

	if p.Url == "" && p.Text == "" {
		log.Printf("%s error: %s\n", funcName, "No url or text")
		return errors.New(http.StatusBadRequest, errors.InvalidBody)
	}
	return nil
}

func (p *Post) Init() {
	p.Score = 1
	p.Views = 0
	p.Created = GenerateCreatedTime(time.Now())
	p.Percentage = 100
}
