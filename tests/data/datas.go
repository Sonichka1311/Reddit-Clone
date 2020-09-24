package data

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"reddit/pkg/auth"
	"reddit/pkg/post"
	"reddit/pkg/user"
	mock_interfaces "reddit/tests/mock"
	"time"
)

func GetTestUserCtx(ctx context.Context) context.Context {
	usr := user.User{
		Id:       "1",
		Login:    "Sonia",
		Password: "testpassword",
	}
	return context.WithValue(ctx, auth.TokenKey, &usr)
}

func GetTestAccessToken() string {
	token, _ := json.Marshal(auth.AccessToken{Token: "testToken"})
	return string(token)
}

func GetTestAuthor() *user.User {
	return &user.User{
		Id:    "1",
		Login: "Sonia",
	}
}

func GetTestUser() *user.User {
	return &user.User{
		Id:    "2",
		Login: "test name",
	}
}

func GetTestVote(turn int) post.Vote {
	return post.Vote{
		Post: post.GeneratePostId(),
		User: GetTestUser().Id,
		Vote: turn,
	}
}

func GetTestVotes(status int) []post.Vote {
	switch status {
	case 0:
		return []post.Vote{
			{
				Post: post.GeneratePostId(),
				User: GetTestAuthor().Id,
				Vote: 1,
			},
		}
	case 1:
		return []post.Vote{
			{
				Post: post.GeneratePostId(),
				User: GetTestUser().Id,
				Vote: 1,
			},
		}
	}
	return []post.Vote{}
}

func GetTestComments(status int) []post.Comment {
	return []post.Comment{
		{
			Post:    post.GeneratePostId(),
			Created: post.GenerateCreatedTime(time.Now()),
			Author:  GetTestAuthor(),
			Body:    "New comment for new post",
			Id:      post.GenerateCommentId(),
		},
	}
}

func GetTestTextPost() *post.Post {
	return &post.Post{
		Type:       "text",
		Title:      "Test Post!",
		Text:       "New test post",
		Category:   "programming",
		Author:     *GetTestAuthor(),
		Id:         post.GeneratePostId(),
		Score:      1,
		Views:      0,
		Created:    post.GenerateCreatedTime(time.Now()),
		Percentage: 100,
	}
}

func GetTestUserPost() *post.UserToPostId {
	return &post.UserToPostId{
		Username: "Sonia",
		PostId:   post.GeneratePostId(),
	}
}

func GetTestTextPostStr() string {
	// should returns `{"score":1,"views":0,"type":"text","title":"Test Post!","text":"New test post","author":{"id":"1","username":"Sonia"},"category":"programming","votes":null,"comments":null,"created":"2020-04-17T01:31:56Z","upvotePercentage":100,"id":"617531333135316677663536"}`
	pst := GetTestTextPost()
	pst.Votes = []post.Vote{
		{
			User: GetTestAuthor().Id,
			Vote: 1,
		},
	}
	pst.Comments = []post.Comment{}

	jsonPost, jsonError := json.Marshal(pst)
	if jsonError != nil {
		log.Printf("TEST error while marshalling data: %s\n", jsonError.Error())
	}
	return string(jsonPost)
}

func GetTestTextPostWithoutVotesStr() string {
	// should returns `{"score":1,"views":0,"type":"text","title":"Test Post!","text":"New test post","author":{"id":"1","username":"Sonia"},"category":"programming","votes":null,"comments":null,"created":"2020-04-17T01:31:56Z","upvotePercentage":100,"id":"617531333135316677663536"}`
	pst := GetTestTextPost()
	pst.Votes = []post.Vote{}
	pst.Comments = []post.Comment{}

	jsonPost, jsonError := json.Marshal(pst)
	if jsonError != nil {
		log.Printf("TEST error while marshalling data: %s\n", jsonError.Error())
	}
	return string(jsonPost)
}

func GetTestTextPostWithViewStr() string {
	// should returns `{"score":1,"views":1,"type":"text","title":"Test Post!","text":"New test post","author":{"id":"1","username":"Sonia"},"category":"programming","votes":null,"comments":null,"created":"2020-04-17T01:31:56Z","upvotePercentage":100,"id":"617531333135316677663536"}`
	pst := GetTestTextPost()
	pst.Votes = []post.Vote{
		{
			User: GetTestAuthor().Id,
			Vote: 1,
		},
	}
	pst.Comments = []post.Comment{}
	pst.Views = 1

	jsonPost, jsonError := json.Marshal(pst)
	if jsonError != nil {
		log.Printf("TEST error while marshalling data: %s\n", jsonError.Error())
	}
	return string(jsonPost)
}

func GetTestTextPostWithCommentStr() string {
	// should returns `{"score":1,"views":0,"type":"text","title":"Test Post!","text":"New test post","author":{"id":"1","username":"Sonia"},"category":"programming","votes":null,"comments":[вот тут какой-то комментарий]],"created":"2020-04-17T01:31:56Z","upvotePercentage":100,"id":"617531333135316677663536"}`
	pst := GetTestTextPost()
	pst.Votes = []post.Vote{
		{
			User: GetTestAuthor().Id,
			Vote: 1,
		},
	}
	pst.Comments = []post.Comment{
		{
			Created: post.GenerateCreatedTime(time.Now()),
			Author:  GetTestAuthor(),
			Body:    "New comment for new post",
			Id:      post.GenerateCommentId(),
		},
	}

	jsonPost, jsonError := json.Marshal(pst)
	if jsonError != nil {
		log.Printf("TEST error while marshalling data: %s\n", jsonError.Error())
	}
	return string(jsonPost)
}

func GetTestTextPostUpvoteStr() string {
	pst := GetTestTextPost()
	pst.Votes = []post.Vote{
		{
			User: GetTestAuthor().Id,
			Vote: 1,
		},
		{
			User: GetTestUser().Id,
			Vote: 1,
		},
	}
	pst.Comments = []post.Comment{}

	jsonPost, jsonError := json.Marshal(pst)
	if jsonError != nil {
		log.Printf("TEST error while marshalling data: %s\n", jsonError.Error())
	}
	return string(jsonPost)
}

func GetTestTextPostDownvoteStr() string {
	pst := GetTestTextPost()
	pst.Votes = []post.Vote{
		{
			User: GetTestAuthor().Id,
			Vote: 1,
		},
		{
			User: GetTestUser().Id,
			Vote: -1,
		},
	}
	pst.Comments = []post.Comment{}

	jsonPost, jsonError := json.Marshal(pst)
	if jsonError != nil {
		log.Printf("TEST error while marshalling data: %s\n", jsonError.Error())
	}
	return string(jsonPost)
}

func GetTestArrayPosts() string {
	// `[{"score":1,"views":0,"type":"text","title":"Test Post!","text":"New test post","author":{"id":"1","username":"Sonia"},"category":"programming","votes":[{"user":"1","vote":1}],"comments":[],"created":"2020-04-17T01:31:56Z","upvotePercentage":100,"id":"617531333135316677663536"}]`
	return "[" + GetTestTextPostStr() + "]"
}

func GetTestUrlPostStr() string {
	// should returns `{"score":1,"views":0,"type":"url","title":"Test Post!","url":"https://newTestUrl.test","author":{"id":"1","username":"Sonia"},"category":"programming","votes":null,"comments":null,"created":"2020-04-17T01:31:56Z","upvotePercentage":100,"id":"617531333135316677663536"}`
	pst := post.Post{
		Type:     "url",
		Title:    "Test Post!",
		Url:      "https://newTestUrl.test",
		Category: "programming",
		Author:   *GetTestAuthor(),
		Id:       post.GeneratePostId(),
		Votes: []post.Vote{
			{
				User: GetTestAuthor().Id,
				Vote: 1,
			},
		},
		Comments:   []post.Comment{},
		Score:      1,
		Views:      0,
		Created:    post.GenerateCreatedTime(time.Now()),
		Percentage: 100,
	}

	jsonPost, jsonError := json.Marshal(pst)
	if jsonError != nil {
		log.Printf("TEST error while marshalling data: %s\n", jsonError.Error())
	}
	return string(jsonPost)
}

func GetSuccessMessage() string {
	data, jsonError := json.Marshal(
		struct {
			Message string `json:"message"`
		}{
			Message: "success",
		},
	)
	if jsonError != nil {
		log.Printf("TEST error while marshalling data: %s\n", jsonError.Error())
	}
	return string(data)
}

func GetAuthTestCase(base AuthTestCase, data *bytes.Reader, dbType string, dbStatus int, status int, resp string) AuthTestCase {
	if data != nil {
		base.Data = *data
	}

	switch dbType {
	case "register":
		base.DB = InitRegisterDb(dbStatus)
	case "login":
		base.DB = InitLoginDb(dbStatus)
	case "auth":
		base.DB = InitAuthDb(dbStatus)
	}

	base.Status = status
	base.Response = resp

	return base
}

func GetMiddlewareTestCase(handler, method string) MiddlewareTestCase {
	return MiddlewareTestCase{
		Handler: handler,
		Method:  method,
	}
}

func GetMiddlewareTestCaseWithStatus(handler, method string, param, status int) MiddlewareTestCase {
	return MiddlewareTestCase{
		Handler: handler,
		Method:  method,
		Int:     param,
		Status:  status,
	}
}

func GetAuthTestCaseWithAction(base AuthTestCase, action, unaction func()) AuthTestCase {
	base.Action = action
	base.Unaction = unaction
	return base
}

func GetTokenTestCase(data *user.User, isErr bool) TokenTestCase {
	return TokenTestCase{
		Data:  *data,
		IsErr: isErr,
	}
}

func GetTokenTestCaseWithAction(data *user.User, isErr bool, action, unaction func()) TokenTestCase {
	return TokenTestCase{
		Data:     *data,
		IsErr:    isErr,
		Action:   action,
		Unaction: unaction,
	}
}

type Collection = mock_interfaces.MockMongoCollection

func GetPostTestCase(base PostsTestCase, dataStatus int, resp string, status int) PostsTestCase {
	if dataStatus != -1 {
		base.Data = *GetTestPostDataReader(dataStatus)
	}
	base.Response = resp
	base.Status = status
	return base
}

func SetCollections(base PostsTestCase, collections [4]*Collection) PostsTestCase {
	base.Collections = collections
	return base
}

func SetHandler(base PostsTestCase, handler string) PostsTestCase {
	base.Handler = handler
	return base
}
