package tests

//
//import (
//	"bytes"
//	"context"
//	"database/sql"
//	"encoding/json"
//	errors2 "errors"
//	"fmt"
//	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/golang/mock/gomock"
//	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
//	"log"
//	"reddit/pkg/auth"
//	"reddit/pkg/database"
//	"reddit/pkg/errors"
//	"reddit/pkg/handlers"
//	"reddit/pkg/interfaces"
//	"reddit/pkg/post"
//	"reddit/pkg/user"
//	"reddit/tests/mock"
//	"time"
//)
//
//// DBs -----------------------------------------------------------------------------------------------------------------
//func InitRegisterDb(status int) *sql.DB {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
//	}
//
//	expire := auth.GenerateTime(12)
//
//	if status == 1 {
//		mock.
//			ExpectQuery("[SELECT (.+) FROM users]").
//			WithArgs("Sonia").
//			WillReturnRows(sqlmock.NewRows([]string{"HEX(id)"}).AddRow("1"))
//	} else if status == 4 {
//		mock.
//			ExpectQuery("[SELECT (.+) FROM users]").
//			WithArgs("Sonia").
//			WillReturnError(sql.ErrConnDone)
//	}
//	mock.
//		ExpectQuery("[SELECT (.+) FROM users]").
//		WithArgs("Sonia").
//		WillReturnRows(sqlmock.NewRows([]string{"HEX(id)"}))
//
//	if status == 2 {
//		mock.
//			ExpectExec("[INSERT INTO users]").
//			WithArgs("Sonia", "testpassword").
//			WillReturnError(sql.ErrConnDone)
//	}
//	mock.
//		ExpectExec("[INSERT INTO users]").
//		WithArgs("Sonia", "testpassword").
//		WillReturnResult(sqlmock.NewResult(1, 1))
//
//	if status == 3 {
//		mock.
//			ExpectQuery("[SELECT (.+) FROM users]").
//			WithArgs("Sonia").
//			WillReturnRows(sqlmock.NewRows([]string{"HEX(id)"}))
//	}
//	mock.
//		ExpectQuery("[SELECT (.+) FROM users]").
//		WithArgs("Sonia").
//		WillReturnRows(sqlmock.NewRows([]string{"HEX(id)"}).AddRow("1"))
//
//	if status == 5 {
//		mock.
//			ExpectQuery("[SELECT (.+) FROM sessions]").
//			WithArgs("testToken").
//			WillReturnRows(sqlmock.NewRows([]string{"id", "login", "expire"}).AddRow("1", "Sonia", expire))
//	} else if status == 6 {
//		mock.
//			ExpectQuery("[SELECT (.+) FROM sessions]").
//			WithArgs("testToken").
//			WillReturnError(sql.ErrConnDone)
//	}
//	mock.
//		ExpectQuery("[SELECT (.+) FROM sessions]").
//		WithArgs("testToken").
//		WillReturnRows(sqlmock.NewRows([]string{"id", "login", expire}))
//
//	if status == 7 {
//		mock.
//			ExpectExec("[INSERT INTO sessions]").WithArgs("testToken", "1", "Sonia", expire).
//			WillReturnError(sql.ErrConnDone)
//	}
//	mock.
//		ExpectExec("[INSERT INTO sessions]").WithArgs("testToken", "1", "Sonia", expire).
//		WillReturnResult(sqlmock.NewResult(1, 1))
//
//	return db
//}
//
//func InitLoginDb(status int) *sql.DB {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
//	}
//
//	expire := auth.GenerateTime(12)
//
//	if status == 1 {
//		mock.
//			ExpectQuery("[SELECT (.+) FROM users]").
//			WithArgs("Sonia").
//			WillReturnRows(sqlmock.NewRows([]string{"HEX(id)", "login", "password"}))
//	} else if status == 2 {
//		mock.
//			ExpectQuery("[SELECT (.+) FROM users]").
//			WithArgs("Sonia").
//			WillReturnError(sql.ErrConnDone)
//	} else if status == 3 {
//		mock.
//			ExpectQuery("[SELECT (.+) FROM users]").
//			WithArgs("Sonia").
//			WillReturnRows(sqlmock.NewRows([]string{"HEX(id)", "login", "password"}).AddRow("1", "Sonia", "wrongpassword"))
//	}
//	mock.
//		ExpectQuery("[SELECT (.+) FROM users]").
//		WithArgs("Sonia").
//		WillReturnRows(sqlmock.NewRows([]string{"HEX(id)", "login", "password"}).AddRow("1", "Sonia", "testpassword"))
//
//	if status == 4 {
//		mock.
//			ExpectQuery("[SELECT (.+) FROM sessions]").
//			WithArgs("testToken").
//			WillReturnRows(sqlmock.NewRows([]string{"id", "login", "expire"}).AddRow("1", "Sonia", expire))
//	} else if status == 5 {
//		mock.
//			ExpectQuery("[SELECT (.+) FROM sessions]").
//			WithArgs("testToken").
//			WillReturnError(sql.ErrConnDone)
//	}
//	mock.
//		ExpectQuery("[SELECT (.+) FROM sessions]").
//		WithArgs("testToken").
//		WillReturnRows(sqlmock.NewRows([]string{"id", "login", "expire"}))
//
//	if status == 6 {
//		mock.
//			ExpectQuery("[SELECT (.+) FROM sessions]").
//			WithArgs("Sonia").
//			WillReturnError(sql.ErrConnDone)
//	} else if status == 7 || status == 8 {
//		mock.
//			ExpectQuery("[SELECT (.+) FROM sessions]").
//			WithArgs("Sonia").
//			WillReturnError(sql.ErrNoRows)
//
//		if status == 8 {
//			mock.
//				ExpectExec("[INSERT INTO sessions]").
//				WithArgs("testToken", "1", "Sonia", expire).
//				WillReturnError(sql.ErrConnDone)
//		}
//		mock.
//			ExpectExec("[INSERT INTO sessions]").
//			WithArgs("testToken", "1", "Sonia", expire).
//			WillReturnResult(sqlmock.NewResult(1, 1))
//	}
//	mock.
//		ExpectQuery("[SELECT (.+) FROM sessions]").
//		WithArgs("Sonia").
//		WillReturnRows(sqlmock.NewRows([]string{"token", "expire"}).AddRow("testToken", expire))
//
//	if status == 9 {
//		mock.
//			ExpectExec("[UPDATE sessions]").
//			WithArgs("testToken", expire, "Sonia").
//			WillReturnError(sql.ErrConnDone)
//	}
//	mock.
//		ExpectExec("[UPDATE sessions]").
//		WithArgs("testToken", expire, "Sonia").
//		WillReturnResult(sqlmock.NewResult(1, 1))
//
//	return db
//}
//
//func InitAuthDb(status int) *sql.DB {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
//	}
//
//	expire := auth.GenerateTime(12)
//
//	if status == 1 {
//		mock.
//			ExpectQuery("[SELECT (.+) FROM sessions]").
//			WithArgs("testToken").
//			WillReturnError(sql.ErrNoRows)
//	} else if status == 2 {
//		mock.
//			ExpectQuery("[SELECT (.+) FROM sessions]").
//			WithArgs("testToken").
//			WillReturnRows(sqlmock.NewRows([]string{"id", "login", "expire"}).AddRow("1", "Sonia", auth.GenerateTime(-5)))
//	}
//	mock.
//		ExpectQuery("[SELECT (.+) FROM sessions]").
//		WithArgs("testToken").
//		WillReturnRows(sqlmock.NewRows([]string{"id", "login", "expire"}).AddRow("1", "Sonia", expire))
//
//	return db
//}
//
//// Bodys ---------------------------------------------------------------------------------------------------------------
//
//func GetTestUserReader() *bytes.Reader {
//	usr := user.User{
//		Login:    "Sonia",
//		Password: "testpassword",
//	}
//	usrJson, err := json.Marshal(usr)
//	if err != nil {
//		log.Fatalf("GetTestUserReader() failed parse user json with error %s", err)
//	}
//	return bytes.NewReader(usrJson)
//}
//
//func GetWrongTestUserReader() *bytes.Reader {
//	usr := struct {
//		Username int
//	}{
//		Username: 1234,
//	}
//	usrJson, err := json.Marshal(usr)
//	if err != nil {
//		log.Fatalf("GetWrongTestUserReader() failed parse user json with error %s", err)
//	}
//	return bytes.NewReader(usrJson)
//}
//
//func GetTestUserHandler(db *sql.DB) *handlers.UserHandler {
//	userHandler := handlers.UserHandler{
//		Repo:     user.NewRepo((*user.Database)(database.InitDB(db))),
//		Sessions: (*auth.Database)(database.InitDB(db)),
//	}
//
//	auth.GenerateTokenFunc = func(userData user.User) (*auth.AccessToken, *errors.Error) {
//		return &auth.AccessToken{Token: "testToken"}, nil
//	}
//
//	expire := auth.GenerateTime(12)
//
//	auth.GenerateTime = func(hours int) string {
//		return expire
//	}
//	return &userHandler
//}
//
//func GetTestUserCtx(ctx context.Context) context.Context {
//	usr := user.User{
//		Id:       "1",
//		Login:    "Sonia",
//		Password: "testpassword",
//	}
//	return context.WithValue(ctx, auth.TokenKey, &usr)
//}
//
//func GetTestPostReader(status int) *bytes.Reader {
//	var pst post.Post
//	switch status {
//	case 0:
//		pst = post.Post{
//			Type:     "text",
//			Title:    "Test Post!",
//			Text:     "New test post",
//			Category: "programming",
//		}
//	case 1:
//		pst = post.Post{
//			Type:     "url",
//			Title:    "Test Post!",
//			Url:      "https://newTestUrl.test",
//			Category: "programming",
//		}
//	case 2:
//		pst = post.Post{
//			Type:     "text",
//			Title:    "Test Post!",
//			Category: "programming",
//		}
//	}
//	pstJson, err := json.Marshal(pst)
//	if err != nil {
//		log.Fatalf("GetTestTextPostReader() failed parse user json with error %s", err)
//	}
//	return bytes.NewReader(pstJson)
//}
//
//func GetWrongTestPostReader() *bytes.Reader {
//	usr := struct {
//		Type int
//	}{
//		Type: 1234,
//	}
//	usrJson, err := json.Marshal(usr)
//	if err != nil {
//		log.Fatalf("GetWrongTesPostReader() failed parse user json with error %s", err)
//	}
//	return bytes.NewReader(usrJson)
//}
//
//func GetTestCommentReader() *bytes.Reader {
//	comment := struct {
//		Comment string `json:"comment"`
//	}{
//		Comment: "New comment",
//	}
//	commentJson, err := json.Marshal(comment)
//	if err != nil {
//		log.Fatalf("GetTestCommentReader() failed parse user json with error %s", err)
//	}
//	return bytes.NewReader(commentJson)
//}
//
//func GetWrongTestCommentReader() *bytes.Reader {
//	comment := struct {
//		Comment string `json:"comment"`
//	}{
//		Comment: "New comment",
//	}
//	commentJson, err := json.Marshal(comment)
//	if err != nil {
//		log.Fatalf("GetTestCommentReader() failed parse user json with error %s", err)
//	}
//	return bytes.NewReader(commentJson)
//}
//
//func GetTestAuthor() *user.User {
//	return &user.User{
//		Id:    "1",
//		Login: "Sonia",
//	}
//}
//
//func GetTestVotes(status int) []post.Vote {
//	return []post.Vote{
//		{
//			Post: post.GeneratePostId(),
//			User: GetTestAuthor().Id,
//			Vote: 1,
//		},
//	}
//}
//
//func GetTestComments(status int) []post.Comment {
//	return []post.Comment{
//		{
//			Post:    post.GeneratePostId(),
//			Created: post.GenerateCreatedTime(time.Now()),
//			Author:  GetTestAuthor(),
//			Body:    "New comment for new post",
//			Id:      post.GenerateCommentId(),
//		},
//	}
//}
//
//func GetTestUserToPostId() *post.UserToPostId {
//	return &post.UserToPostId{
//		Username: "Sonia",
//		PostId:   "1",
//	}
//}
//
//func GetTestTextPost() *post.Post {
//	return &post.Post{
//		Type:       "text",
//		Title:      "Test Post!",
//		Text:       "New test post",
//		Category:   "programming",
//		Author:     *GetTestAuthor(),
//		Id:         post.GeneratePostId(),
//		Score:      1,
//		Views:      0,
//		Votes:      []post.Vote{},
//		Comments:   []post.Comment{},
//		Created:    post.GenerateCreatedTime(time.Now()),
//		Percentage: 100,
//	}
//}
//
//func GetTestUrlPost() *post.Post {
//	return &post.Post{
//		Type:       "url",
//		Title:      "Test Post!",
//		Url:        "New test post",
//		Category:   "programming",
//		Author:     *GetTestAuthor(),
//		Id:         post.GeneratePostId(),
//		Score:      1,
//		Views:      0,
//		Votes:      []post.Vote{},
//		Comments:   []post.Comment{},
//		Created:    post.GenerateCreatedTime(time.Now()),
//		Percentage: 100,
//	}
//}
//
//func GetTestBrokenPost() *post.Post {
//	return &post.Post{
//		Type:       "text",
//		Title:      "Test Post!",
//		Category:   "programming",
//		Author:     *GetTestAuthor(),
//		Id:         post.GeneratePostId(),
//		Score:      1,
//		Views:      0,
//		Votes:      []post.Vote{},
//		Comments:   []post.Comment{},
//		Created:    post.GenerateCreatedTime(time.Now()),
//		Percentage: 100,
//	}
//}
//
//func GetTestPostHandler(db interfaces.MongoSession) *handlers.PostHandler {
//	postHandler := handlers.PostHandler{
//		Repo: post.NewRepo(db),
//	}
//	return &postHandler
//}
//
//func GetQuery(ctrl *gomock.Controller, status int, data interface{}) *mock_interfaces.MockMongoQuery {
//	switch status {
//	case 2:
//		q := mock_interfaces.NewMockMongoQuery(ctrl)
//		q.EXPECT().All(gomock.Any()).SetArg(0, data)
//		return q
//	case 3:
//		q := mock_interfaces.NewMockMongoQuery(ctrl)
//		q.EXPECT().All(gomock.Any()).Return(errors2.New("DB err"))
//		return q
//	case 4:
//		q := mock_interfaces.NewMockMongoQuery(ctrl)
//		q.EXPECT().All(gomock.Any()).Return(mgo.ErrNotFound)
//		return q
//	case 5, 6:
//		q := mock_interfaces.NewMockMongoQuery(ctrl)
//		q.EXPECT().One(gomock.Any()).SetArg(0, data)
//		return q
//	case 7:
//		q := mock_interfaces.NewMockMongoQuery(ctrl)
//		q.EXPECT().One(gomock.Any()).Return(mgo.ErrNotFound)
//		return q
//	}
//	return nil
//}
//
//func GetCollection(ctrl *gomock.Controller, status int, data interface{}) *mock_interfaces.MockMongoCollection {
//	switch status {
//	case 0:
//		col := mock_interfaces.NewMockMongoCollection(ctrl)
//		col.EXPECT().Insert(gomock.Any()).Return(nil)
//		return col
//	case 1:
//		col := mock_interfaces.NewMockMongoCollection(ctrl)
//		col.EXPECT().Insert(gomock.Any()).Return(errors2.New("DB err"))
//		return col
//	case 2, 3, 4, 5:
//		col := mock_interfaces.NewMockMongoCollection(ctrl)
//		col.EXPECT().Find(gomock.Any()).Return(GetQuery(ctrl, status, data))
//		return col
//	case 6:
//		col := mock_interfaces.NewMockMongoCollection(ctrl)
//		col.EXPECT().Find(gomock.Any()).Return(GetQuery(ctrl, status, data))
//		col.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
//		return col
//	case 7:
//		col := mock_interfaces.NewMockMongoCollection(ctrl)
//		col.EXPECT().Find(gomock.Any()).Return(GetQuery(ctrl, status, data))
//		return col
//	case 8:
//		col := mock_interfaces.NewMockMongoCollection(ctrl)
//		col.EXPECT().Find(gomock.Any()).Return(GetQuery(ctrl, 5, data))
//		col.EXPECT().Update(gomock.Any(), gomock.Any()).Return(mgo.ErrNotFound)
//		return col
//	case 9:
//		col := mock_interfaces.NewMockMongoCollection(ctrl)
//		col.EXPECT().Insert(gomock.Any()).Return(nil)
//		col.EXPECT().Find(gomock.Any()).Return(GetQuery(ctrl, 2, data))
//		return col
//	}
//	return nil
//}
//
//func GetHandler(ctrl *gomock.Controller, statuses []int, data []interface{}) *handlers.PostHandler {
//	if len(statuses) < 4 {
//		return nil
//	}
//
//	db := mock_interfaces.NewMockMongoDatabase(ctrl)
//	fields := []string{"posts", "votes", "comments", "users"}
//	for idx, field := range fields {
//		db.EXPECT().C(field).Return(GetCollection(ctrl, statuses[idx], data[idx]))
//	}
//
//	session := mock_interfaces.NewMockMongoSession(ctrl)
//	session.EXPECT().DB("reddit").Return(db).MaxTimes(4)
//	return GetTestPostHandler(session)
//}
//
//// checks --------------------------------------------------------------------------------------------------------------
//func CheckStatus(respCode, expCode int) (bool, *string) {
//	if respCode != expCode {
//		err := fmt.Sprintf("Wrong status: got: %v, expected: %v", respCode, expCode)
//		return false, &err
//	}
//	return true, nil
//}
//
//func CheckResponse(resp, exp string) (bool, *string) {
//	if resp != exp+"\n" {
//		err := fmt.Sprintf("handler returned unexpected body: got: %v, expected: %v", resp, exp+"\n")
//		return false, &err
//	}
//	return true, nil
//}
//
//func ChangePostGenerations() {
//	post.GeneratePostId = func() bson.ObjectId {
//		return "au13151fwf56" // 617531333135316677663536
//	}
//	post.GenerateCommentId = func() bson.ObjectId {
//		return "au13151fwf56"
//	}
//	post.GenerateCreatedTime = func(tm time.Time) string {
//		return "2020-04-17T01:31:56Z"
//	}
//}
//
//func UnchangePostGenerations() {
//	post.GeneratePostId = post.GenerateId
//	post.GenerateCommentId = post.GenerateId
//	post.GenerateCommentId = post.GenerateId
//}
//
//// structs -------------------------------------------------------------------------------------------------------------
//type AuthTestCase struct {
//	Data     *bytes.Reader
//	Type     string
//	Handler  string
//	DB       *sql.DB
//	Status   int
//	Response string
//	Action   func()
//	Unaction func()
//}
//
//type TokenTestCase struct {
//	Data     user.User
//	IsErr    bool
//	Action   func()
//	Unaction func()
//}
//
//type MiddlewareTestCase struct {
//	Handler string
//	Method  string
//	Int     int
//	Status  int
//}
//
//type PostsTestCase struct {
//	Method    string
//	Handler   string
//	NeedAuth  bool
//	Data      *bytes.Reader
//	PostSt    int
//	VoteSt    int
//	CommentSt int
//	UserSt    int
//	Params    []interface{}
//	DB        *mgo.Session
//	Response  string
//	Status    int
//}
