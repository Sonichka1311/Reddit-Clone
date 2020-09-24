package data

import (
	"database/sql"
	"github.com/golang/mock/gomock"
	"reddit/pkg/auth"
	"reddit/pkg/database"
	"reddit/pkg/errors"
	"reddit/pkg/handlers"
	"reddit/pkg/interfaces"
	"reddit/pkg/post"
	"reddit/pkg/user"
	"reddit/tests/mock"
)

func GetHandler(ctrl *gomock.Controller, collections [4]*mock_interfaces.MockMongoCollection) *handlers.PostHandler {
	if len(collections) < 4 {
		return nil
	}

	db := mock_interfaces.NewMockMongoDatabase(ctrl)
	fields := []string{"posts", "votes", "comments", "users"}
	for idx, field := range fields {
		db.EXPECT().C(field).Return(collections[idx])
	}

	session := mock_interfaces.NewMockMongoSession(ctrl)
	session.EXPECT().DB("reddit").Return(db).MaxTimes(4)
	return GetTestPostHandler(session)
}

func GetTestPostHandler(db interfaces.MongoSession) *handlers.PostHandler {
	postHandler := handlers.PostHandler{
		Repo: post.NewRepo(db),
	}
	return &postHandler
}

func GetTestUserHandler(db *sql.DB) *handlers.UserHandler {
	userHandler := handlers.UserHandler{
		Repo:     user.NewRepo((*user.Database)(database.InitDB(db))),
		Sessions: (*auth.Database)(database.InitDB(db)),
	}

	auth.GenerateTokenFunc = func(userData user.User) (*auth.AccessToken, *errors.Error) {
		return &auth.AccessToken{Token: "testToken"}, nil
	}

	expire := auth.GenerateTime(12)

	auth.GenerateTime = func(hours int) string {
		return expire
	}
	return &userHandler
}
