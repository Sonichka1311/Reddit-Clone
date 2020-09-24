package tests

import (
	"github.com/golang/mock/gomock"
	"net/http"
	"reddit/pkg/post"
	"reddit/tests/data"
	"reddit/tests/mock"
	"testing"
	"time"
)

type PostsTestCase = data.PostsTestCase
type Collection = mock_interfaces.MockMongoCollection

func TestGenerateId(t *testing.T) {
	if len(post.GenerateId().Hex()) != 24 {
		t.Errorf("TEST 0 of TestGenerateId failed with error: Unexpected body len: got: %d, expected: %d", len(post.GenerateId().Hex()), 24)
	}
}

func TestGenerateTime(t *testing.T) {
	tm := time.Unix(1587374985, 0)
	timeStr := "2020-04-20T12:29:45Z"
	res := post.GenerateTime(tm)
	if res != timeStr {
		t.Errorf("TEST 0 of TestGenerateId failed with error: Unexpected body len: got: %s, expected: %s", res, timeStr)
	}
}

func TestAdd(t *testing.T) {
	data.ChangePostGenerations()
	defer data.UnchangePostGenerations()

	test := data.PostTest{
		Func:       "add post",
		Handler:    "/api/posts",
		Controller: gomock.NewController(t),
	}

	base := PostsTestCase{
		Method:   http.MethodPost,
		Handler:  "/api/posts",
		NeedAuth: true,
	}

	test.Cases = []PostsTestCase{
		data.SetCollections(
			data.GetPostTestCase(base, 0, data.GetTestTextPostStr(), http.StatusOK),
			[4]*Collection{
				data.SetExpectInsert(data.InitCollection(test.Controller)),
				data.SetExpectInsert(data.InitCollection(test.Controller)),
				data.SetExpectInsert(data.InitCollection(test.Controller)),
				data.SetExpectInsert(data.InitCollection(test.Controller)),
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, 1, data.GetTestUrlPostStr(), http.StatusOK),
			[4]*Collection{
				data.SetExpectInsert(data.InitCollection(test.Controller)),
				data.SetExpectInsert(data.InitCollection(test.Controller)),
				data.SetExpectInsert(data.InitCollection(test.Controller)),
				data.SetExpectInsert(data.InitCollection(test.Controller)),
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, 2, `Invalid body`, http.StatusBadRequest),
			[4]*Collection{
				data.SetExpectInsert(data.InitCollection(test.Controller)),
				data.SetExpectInsert(data.InitCollection(test.Controller)),
				data.SetExpectInsert(data.InitCollection(test.Controller)),
				data.SetExpectInsert(data.InitCollection(test.Controller)),
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, 3, `Invalid body`, http.StatusBadRequest),
			[4]*Collection{
				data.SetExpectInsert(data.InitCollection(test.Controller)),
				data.SetExpectInsert(data.InitCollection(test.Controller)),
				data.SetExpectInsert(data.InitCollection(test.Controller)),
				data.SetExpectInsert(data.InitCollection(test.Controller)),
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, 0, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{data.SetExpectInsertWithError(data.InitCollection(test.Controller)), nil, nil, nil},
		),
		data.SetCollections(
			data.GetPostTestCase(base, 0, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectInsert(data.InitCollection(test.Controller)),
				data.SetExpectInsertWithError(data.InitCollection(test.Controller)),
				nil,
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, 0, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectInsert(data.InitCollection(test.Controller)),
				data.SetExpectInsert(data.InitCollection(test.Controller)),
				nil,
				data.SetExpectInsertWithError(data.InitCollection(test.Controller)),
			},
		),
	}

	test.Test(t)
}

func TestGetAll(t *testing.T) {
	data.ChangePostGenerations()
	defer data.UnchangePostGenerations()

	base := PostsTestCase{
		Method:   http.MethodGet,
		Handler:  "/api/posts",
		NeedAuth: false,
	}

	test := data.PostTest{
		Func:       "get all posts",
		Handler:    "/api/posts",
		Controller: gomock.NewController(t),
	}

	test.Cases = []PostsTestCase{
		data.SetCollections(
			data.GetPostTestCase(base, -1, data.GetTestArrayPosts(), http.StatusOK),
			[4]*Collection{
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), []*post.Post{data.GetTestTextPost()}),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), data.GetTestVotes(0)),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), []post.Comment{}),
				),
				data.SetExpectInsert(data.InitCollection(test.Controller)),
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithDbErr(data.InitQuery(test.Controller)),
				),
				nil,
				nil,
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), []*post.Post{data.GetTestTextPost()}),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithDbErr(data.InitQuery(test.Controller)),
				),
				nil,
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), []*post.Post{data.GetTestTextPost()}),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), data.GetTestVotes(0)),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithDbErr(data.InitQuery(test.Controller)),
				),
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, data.GetTestArrayPosts(), http.StatusOK),
			[4]*Collection{
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), []*post.Post{data.GetTestTextPost()}),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), data.GetTestVotes(0)),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithNotFoundErr(data.InitQuery(test.Controller)),
				),
				nil,
			},
		),
	}

	test.Test(t)
}

func TestGetByCategory(t *testing.T) {
	data.ChangePostGenerations()
	defer data.UnchangePostGenerations()

	base := PostsTestCase{
		Method:   http.MethodGet,
		Handler:  "/api/posts/programming",
		NeedAuth: false,
	}

	test := data.PostTest{
		Func:       "get by category",
		Handler:    "/api/posts/programming",
		Controller: gomock.NewController(t),
	}

	test.Cases = []PostsTestCase{
		data.SetCollections(
			data.GetPostTestCase(base, -1, data.GetTestArrayPosts(), http.StatusOK),
			[4]*Collection{
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), []*post.Post{data.GetTestTextPost()}),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), data.GetTestVotes(0)),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), []post.Comment{}),
				),
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithNotFoundErr(data.InitQuery(test.Controller)),
				),
				nil,
				nil,
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), []*post.Post{data.GetTestTextPost()}),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithNotFoundErr(data.InitQuery(test.Controller)),
				),
				nil,
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), []*post.Post{data.GetTestTextPost()}),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), data.GetTestVotes(0)),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithDbErr(data.InitQuery(test.Controller)),
				),
				nil,
			},
		),
	}

	test.Test(t)
}

func TestGet(t *testing.T) {
	data.ChangePostGenerations()
	defer data.UnchangePostGenerations()

	base := PostsTestCase{
		Method:   http.MethodGet,
		Handler:  "/api/post/617531333135316677663536",
		NeedAuth: false,
	}

	test := data.PostTest{
		Func:       "get post",
		Handler:    "/api/post/{post_id}",
		Controller: gomock.NewController(t),
	}

	test.Cases = []PostsTestCase{
		data.SetCollections(
			data.GetPostTestCase(base, -1, data.GetTestTextPostWithViewStr(), http.StatusOK),
			[4]*Collection{
				data.SetExpectUpdate(
					data.SetExpectFindReturnQuery(
						data.InitCollection(test.Controller),
						data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
					),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), data.GetTestVotes(0)),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), []post.Comment{}),
				),
				nil,
			},
		),
		data.SetHandler(
			data.GetPostTestCase(base, -1, `Invalid url parameters`, http.StatusBadRequest),
			"/api/post/6175313331",
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectOneWithNotFoundErr(data.InitQuery(test.Controller)),
				),
				nil,
				nil,
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectUpdate(
					data.SetExpectFindReturnQuery(
						data.InitCollection(test.Controller),
						data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
					),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithDbErr(data.InitQuery(test.Controller)),
				),
				nil,
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectUpdateWithError(
					data.SetExpectFindReturnQuery(
						data.InitCollection(test.Controller),
						data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
					),
				),
				nil,
				nil,
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectUpdate(
					data.SetExpectFindReturnQuery(
						data.InitCollection(test.Controller),
						data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
					),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), data.GetTestVotes(0)),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithDbErr(data.InitQuery(test.Controller)),
				),
				nil,
			},
		),
	}

	test.Test(t)
}

func TestAddComment(t *testing.T) {
	data.ChangePostGenerations()
	defer data.UnchangePostGenerations()

	base := PostsTestCase{
		Method:   http.MethodPost,
		Handler:  "/api/post/617531333135316677663536",
		NeedAuth: true,
	}

	test := data.PostTest{
		Func:       "add comment",
		Handler:    "/api/post/{post_id}",
		Controller: gomock.NewController(t),
	}

	test.Cases = []PostsTestCase{
		data.SetCollections(
			data.GetPostTestCase(base, 4, data.GetTestTextPostWithCommentStr(), http.StatusOK),
			[4]*Collection{
				data.SetExpectUpdate(
					data.SetExpectFindReturnQuery(
						data.InitCollection(test.Controller),
						data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
					),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), data.GetTestVotes(0)),
				),
				data.SetExpectFindReturnQuery(
					data.SetExpectInsert(data.InitCollection(test.Controller)),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), data.GetTestComments(0)),
				),
				data.SetExpectInsert(data.InitCollection(test.Controller)),
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, 4, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectInsert(data.InitCollection(test.Controller)),
				data.SetExpectInsert(data.InitCollection(test.Controller)),
				data.SetExpectInsertWithError(data.InitCollection(test.Controller)),
				data.SetExpectInsert(data.InitCollection(test.Controller)),
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, 5, `Invalid body`, http.StatusBadRequest),
			[4]*Collection{
				data.SetExpectInsert(data.InitCollection(test.Controller)),
				data.SetExpectInsert(data.InitCollection(test.Controller)),
				data.SetExpectInsertWithError(data.InitCollection(test.Controller)),
				data.SetExpectInsert(data.InitCollection(test.Controller)),
			},
		),
	}

	test.Test(t)
}

func TestDeleteComment(t *testing.T) {
	data.ChangePostGenerations()
	defer data.UnchangePostGenerations()

	base := PostsTestCase{
		Method:   http.MethodDelete,
		Handler:  "/api/post/617531333135316677663536/617531333135316677663536",
		NeedAuth: true,
	}

	test := data.PostTest{
		Func:       "delete comment",
		Handler:    "/api/post/{post_id}/{comment_id}",
		Controller: gomock.NewController(t),
	}

	test.Cases = []PostsTestCase{
		data.SetCollections(
			data.GetPostTestCase(base, 4, data.GetTestTextPostStr(), http.StatusOK),
			[4]*Collection{
				data.SetExpectUpdate(
					data.SetExpectFindReturnQuery(
						data.InitCollection(test.Controller),
						data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
					),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), data.GetTestVotes(0)),
				),
				data.SetExpectRemove(
					data.SetExpectFindReturnQuery(
						data.InitCollection(test.Controller),
						data.SetExpectAllWithNotFoundErr(data.InitQuery(test.Controller)),
					),
				),
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, 4, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				nil,
				nil,
				data.SetExpectRemoveWithError(data.SetExpectInsert(data.InitCollection(test.Controller))),
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, 4, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectOneWithNotFoundErr(data.InitQuery(test.Controller)),
				),
				nil,
				data.SetExpectRemoveWithError(data.SetExpectInsert(data.InitCollection(test.Controller))),
				nil,
			},
		),
	}

	test.Test(t)
}

func TestUpvote(t *testing.T) {
	data.ChangePostGenerations()
	defer data.UnchangePostGenerations()

	base := PostsTestCase{
		Method:   http.MethodGet,
		Handler:  "/api/post/617531333135316677663536/upvote",
		NeedAuth: true,
	}

	test := data.PostTest{
		Func:       "upvote",
		Handler:    "/api/post/{post_id}/upvote",
		Controller: gomock.NewController(t),
	}

	test.Cases = []PostsTestCase{
		data.SetCollections(
			data.GetPostTestCase(base, -1, data.GetTestTextPostUpvoteStr(), http.StatusOK),
			[4]*Collection{
				data.SetExpectUpdate(
					data.SetExpectFindReturnQuery(
						data.SetExpectUpdate(
							data.SetExpectFindReturnQuery(
								data.InitCollection(test.Controller),
								data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
							),
						),
						data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
					),
				),
				data.SetExpectFindReturnQuery(
					data.SetExpectFindReturnQuery(
						data.SetExpectFindReturnQuery(
							data.SetExpectInsert(
								data.SetExpectFindReturnQuery(
									data.InitCollection(test.Controller),
									data.SetExpectOneWithNotFoundErr(data.InitQuery(test.Controller)),
								),
							),
							data.SetExpectAllWithArg(
								data.InitQuery(test.Controller),
								append(data.GetTestVotes(0), data.GetTestVote(1)),
							),
						),
						data.SetExpectAllWithArg(
							data.InitQuery(test.Controller),
							append(data.GetTestVotes(0), data.GetTestVote(1)),
						),
					),
					data.SetExpectAllWithArg(
						data.InitQuery(test.Controller),
						append(data.GetTestVotes(0), data.GetTestVote(1)),
					),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(
						data.InitQuery(test.Controller),
						[]post.Comment{},
					),
				),
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
				),
				data.SetExpectInsertWithError(
					data.SetExpectFindReturnQuery(
						data.InitCollection(test.Controller),
						data.SetExpectOneWithNotFoundErr(data.InitQuery(test.Controller)),
					),
				),
				nil,
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectOneWithError(data.InitQuery(test.Controller)),
				),
				nil,
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, data.GetTestTextPostUpvoteStr(), http.StatusOK),
			[4]*Collection{
				data.SetExpectUpdate(
					data.SetExpectFindReturnQuery(
						data.SetExpectUpdate(
							data.SetExpectFindReturnQuery(
								data.InitCollection(test.Controller),
								data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
							),
						),
						data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
					),
				),
				data.SetExpectFindReturnQuery(
					data.SetExpectFindReturnQuery(
						data.SetExpectFindReturnQuery(
							data.SetExpectUpdate(
								data.SetExpectFindReturnQuery(
									data.InitCollection(test.Controller),
									data.SetExpectOneWithArg(data.InitQuery(test.Controller), data.GetTestVote(-1)),
								),
							),
							data.SetExpectAllWithArg(
								data.InitQuery(test.Controller),
								append(data.GetTestVotes(0), data.GetTestVote(1)),
							),
						),
						data.SetExpectAllWithArg(
							data.InitQuery(test.Controller),
							append(data.GetTestVotes(0), data.GetTestVote(1)),
						),
					),
					data.SetExpectAllWithArg(
						data.InitQuery(test.Controller),
						append(data.GetTestVotes(0), data.GetTestVote(1)),
					),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(
						data.InitQuery(test.Controller),
						[]post.Comment{},
					),
				),
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
				),
				data.SetExpectUpdateWithError(
					data.SetExpectFindReturnQuery(
						data.InitCollection(test.Controller),
						data.SetExpectOneWithArg(data.InitQuery(test.Controller), data.GetTestVote(-1)),
					),
				),
				nil,
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
				),
				data.SetExpectFindReturnQuery(
					data.SetExpectInsert(
						data.SetExpectFindReturnQuery(
							data.InitCollection(test.Controller),
							data.SetExpectOneWithNotFoundErr(data.InitQuery(test.Controller)),
						),
					),
					data.SetExpectAllWithDbErr(data.InitQuery(test.Controller)),
				),
				nil,
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
				),
				data.SetExpectFindReturnQuery(
					data.SetExpectFindReturnQuery(
						data.SetExpectInsert(
							data.SetExpectFindReturnQuery(
								data.InitCollection(test.Controller),
								data.SetExpectOneWithNotFoundErr(data.InitQuery(test.Controller)),
							),
						),
						data.SetExpectAllWithArg(
							data.InitQuery(test.Controller),
							append(data.GetTestVotes(0), data.GetTestVote(1)),
						),
					),
					data.SetExpectAllWithDbErr(data.InitQuery(test.Controller)),
				),
				nil,
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, data.GetTestTextPostWithoutVotesStr(), http.StatusOK),
			[4]*Collection{
				data.SetExpectUpdate(
					data.SetExpectFindReturnQuery(
						data.SetExpectUpdate(
							data.SetExpectFindReturnQuery(
								data.InitCollection(test.Controller),
								data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
							),
						),
						data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
					),
				),
				data.SetExpectFindReturnQuery(
					data.SetExpectFindReturnQuery(
						data.SetExpectFindReturnQuery(
							data.SetExpectInsert(
								data.SetExpectFindReturnQuery(
									data.InitCollection(test.Controller),
									data.SetExpectOneWithNotFoundErr(data.InitQuery(test.Controller)),
								),
							),
							data.SetExpectAllWithArg(
								data.InitQuery(test.Controller),
								[]post.Vote{},
							),
						),
						data.SetExpectAllWithArg(
							data.InitQuery(test.Controller),
							[]post.Vote{},
						),
					),
					data.SetExpectAllWithArg(
						data.InitQuery(test.Controller),
						[]post.Vote{},
					),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(
						data.InitQuery(test.Controller),
						[]post.Comment{},
					),
				),
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectUpdateWithError(
					data.SetExpectFindReturnQuery(
						data.InitCollection(test.Controller),
						data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
					),
				),
				data.SetExpectFindReturnQuery(
					data.SetExpectFindReturnQuery(
						data.SetExpectInsert(
							data.SetExpectFindReturnQuery(
								data.InitCollection(test.Controller),
								data.SetExpectOneWithNotFoundErr(data.InitQuery(test.Controller)),
							),
						),
						data.SetExpectAllWithArg(
							data.InitQuery(test.Controller),
							append(data.GetTestVotes(0), data.GetTestVote(1)),
						),
					),
					data.SetExpectAllWithArg(
						data.InitQuery(test.Controller),
						append(data.GetTestVotes(0), data.GetTestVote(1)),
					),
				),
				nil,
				nil,
			},
		),
	}

	test.Test(t)
}

func TestDownvote(t *testing.T) {
	data.ChangePostGenerations()
	defer data.UnchangePostGenerations()

	base := PostsTestCase{
		Method:   http.MethodGet,
		Handler:  "/api/post/617531333135316677663536/downvote",
		NeedAuth: true,
	}

	test := data.PostTest{
		Func:       "downvote",
		Handler:    "/api/post/{post_id}/downvote",
		Controller: gomock.NewController(t),
	}

	test.Cases = []PostsTestCase{
		data.SetCollections(
			data.GetPostTestCase(base, -1, data.GetTestTextPostDownvoteStr(), http.StatusOK),
			[4]*Collection{
				data.SetExpectUpdate(
					data.SetExpectFindReturnQuery(
						data.SetExpectUpdate(
							data.SetExpectFindReturnQuery(
								data.InitCollection(test.Controller),
								data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
							),
						),
						data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
					),
				),
				data.SetExpectFindReturnQuery(
					data.SetExpectFindReturnQuery(
						data.SetExpectFindReturnQuery(
							data.SetExpectUpdate(
								data.SetExpectFindReturnQuery(
									data.InitCollection(test.Controller),
									data.SetExpectOneWithArg(data.InitQuery(test.Controller), data.GetTestVote(1)),
								),
							),
							data.SetExpectAllWithArg(
								data.InitQuery(test.Controller),
								append(data.GetTestVotes(0), data.GetTestVote(-1)),
							),
						),
						data.SetExpectAllWithArg(
							data.InitQuery(test.Controller),
							data.GetTestVotes(0),
						),
					),
					data.SetExpectAllWithArg(
						data.InitQuery(test.Controller),
						append(data.GetTestVotes(0), data.GetTestVote(-1)),
					),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(
						data.InitQuery(test.Controller),
						[]post.Comment{},
					),
				),
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
				),
				data.SetExpectUpdateWithError(
					data.SetExpectFindReturnQuery(
						data.InitCollection(test.Controller),
						data.SetExpectOneWithArg(data.InitQuery(test.Controller), data.GetTestVote(1)),
					),
				),
				nil,
				nil,
			},
		),
	}

	test.Test(t)
}

func TestUnvote(t *testing.T) {
	data.ChangePostGenerations()
	defer data.UnchangePostGenerations()

	base := PostsTestCase{
		Method:   http.MethodGet,
		Handler:  "/api/post/617531333135316677663536/unvote",
		NeedAuth: true,
	}

	test := data.PostTest{
		Func:       "unvote",
		Handler:    "/api/post/{post_id}/unvote",
		Controller: gomock.NewController(t),
	}

	test.Cases = []PostsTestCase{
		data.SetCollections(
			data.GetPostTestCase(base, -1, data.GetTestTextPostStr(), http.StatusOK),
			[4]*Collection{
				data.SetExpectUpdate(
					data.SetExpectFindReturnQuery(
						data.SetExpectUpdate(
							data.SetExpectFindReturnQuery(
								data.InitCollection(test.Controller),
								data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
							),
						),
						data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
					),
				),
				data.SetExpectFindReturnQuery(
					data.SetExpectFindReturnQuery(
						data.SetExpectFindReturnQuery(
							data.SetExpectRemove(
								data.SetExpectFindReturnQuery(
									data.SetExpectFindReturnQuery(
										data.InitCollection(test.Controller),
										data.SetExpectOneWithArg(data.InitQuery(test.Controller), data.GetTestVote(1)),
									),
									data.SetExpectOneWithArg(data.InitQuery(test.Controller), data.GetTestVote(1)),
								),
							),
							data.SetExpectAllWithArg(
								data.InitQuery(test.Controller),
								data.GetTestVotes(0),
							),
						),
						data.SetExpectAllWithArg(
							data.InitQuery(test.Controller),
							data.GetTestVotes(0),
						),
					),
					data.SetExpectAllWithArg(
						data.InitQuery(test.Controller),
						data.GetTestVotes(0),
					),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(
						data.InitQuery(test.Controller),
						[]post.Comment{},
					),
				),
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				nil,
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectOneWithError(data.InitQuery(test.Controller)),
				),
				nil,
				nil,
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectUpdate(
					data.SetExpectFindReturnQuery(
						data.InitCollection(test.Controller),
						data.SetExpectOneWithArg(data.InitQuery(test.Controller), *data.GetTestTextPost()),
					),
				),
				data.SetExpectRemoveWithError(
					data.SetExpectFindReturnQuery(
						data.SetExpectFindReturnQuery(
							data.InitCollection(test.Controller),
							data.SetExpectOneWithArg(data.InitQuery(test.Controller), data.GetTestVote(1)),
						),
						data.SetExpectOneWithArg(data.InitQuery(test.Controller), data.GetTestVote(1)),
					),
				),
				nil,
				nil,
			},
		),
	}

	test.Test(t)
}

func TestDeletePost(t *testing.T) {
	data.ChangePostGenerations()
	defer data.UnchangePostGenerations()

	base := PostsTestCase{
		Method:   http.MethodDelete,
		Handler:  "/api/post/617531333135316677663536",
		NeedAuth: true,
	}

	test := data.PostTest{
		Func:       "delete post",
		Handler:    "/api/post/{post_id}",
		Controller: gomock.NewController(t),
	}

	test.Cases = []PostsTestCase{
		data.SetCollections(
			data.GetPostTestCase(base, -1, data.GetSuccessMessage(), http.StatusOK),
			[4]*Collection{
				data.SetExpectRemove(data.InitCollection(test.Controller)),
				data.SetExpectRemove(data.InitCollection(test.Controller)),
				data.SetExpectRemove(data.InitCollection(test.Controller)),
				data.SetExpectRemove(data.InitCollection(test.Controller)),
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectRemoveWithError(data.InitCollection(test.Controller)),
				data.SetExpectRemove(data.InitCollection(test.Controller)),
				data.SetExpectRemove(data.InitCollection(test.Controller)),
				data.SetExpectRemove(data.InitCollection(test.Controller)),
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectRemove(data.InitCollection(test.Controller)),
				data.SetExpectRemoveWithError(data.InitCollection(test.Controller)),
				data.SetExpectRemove(data.InitCollection(test.Controller)),
				data.SetExpectRemove(data.InitCollection(test.Controller)),
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectRemove(data.InitCollection(test.Controller)),
				data.SetExpectRemove(data.InitCollection(test.Controller)),
				data.SetExpectRemoveWithError(data.InitCollection(test.Controller)),
				data.SetExpectRemove(data.InitCollection(test.Controller)),
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				data.SetExpectRemove(data.InitCollection(test.Controller)),
				data.SetExpectRemove(data.InitCollection(test.Controller)),
				data.SetExpectRemove(data.InitCollection(test.Controller)),
				data.SetExpectRemoveWithError(data.InitCollection(test.Controller)),
			},
		),
	}

	test.Test(t)
}

func TestGetByUser(t *testing.T) {
	data.ChangePostGenerations()
	defer data.UnchangePostGenerations()

	base := PostsTestCase{
		Method:   http.MethodGet,
		Handler:  "/api/posts/sonia",
		NeedAuth: false,
	}

	test := data.PostTest{
		Func:       "get by user",
		Handler:    "/api/posts/sonia",
		Controller: gomock.NewController(t),
	}

	test.Cases = []PostsTestCase{
		data.SetCollections(
			data.GetPostTestCase(base, -1, data.GetTestArrayPosts(), http.StatusOK),
			[4]*Collection{
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), []*post.Post{data.GetTestTextPost()}),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), data.GetTestVotes(0)),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), []post.Comment{}),
				),
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), []post.UserToPostId{*data.GetTestUserPost()}),
				),
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `[]`, http.StatusOK),
			[4]*Collection{
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithArg(data.InitQuery(test.Controller), []*post.Post{}),
				),
				nil,
				nil,
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithNotFoundErr(data.InitQuery(test.Controller)),
				),
			},
		),
		data.SetCollections(
			data.GetPostTestCase(base, -1, `Internal server error`, http.StatusInternalServerError),
			[4]*Collection{
				nil,
				nil,
				nil,
				data.SetExpectFindReturnQuery(
					data.InitCollection(test.Controller),
					data.SetExpectAllWithDbErr(data.InitQuery(test.Controller)),
				),
			},
		),
	}

	test.Test(t)
}
