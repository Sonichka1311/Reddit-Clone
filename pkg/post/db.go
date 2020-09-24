package post

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"math"
	"net/http"
	"reddit/pkg/errors"
	"reddit/pkg/interfaces"
)

type Database struct {
	Posts    interfaces.MongoCollection
	Votes    interfaces.MongoCollection
	Comments interfaces.MongoCollection
	Users    interfaces.MongoCollection
}

func (d *Database) AddPost(post *Post) *errors.Error {
	insertError := d.Posts.Insert(*post)
	if insertError != nil {
		log.Printf("Add post error: error from database: %s", insertError.Error())
		return errors.New(http.StatusInternalServerError, errors.InternalError)
	}

	vote := Vote{
		Post: post.Id,
		User: post.Author.Id,
		Vote: 1,
	}

	insertError = d.Votes.Insert(vote)
	if insertError != nil {
		log.Printf("Add vote error: error from database: %s", insertError.Error())
		return errors.New(http.StatusInternalServerError, errors.InternalError)
	}

	post.Votes = []Vote{vote}
	post.Comments = []Comment{}

	user := UserToPostId{
		Username: post.Author.Login,
		PostId:   post.Id,
	}

	insertError = d.Users.Insert(user)
	if insertError != nil {
		log.Printf("Add user error: error from database: %s", insertError.Error())
		return errors.New(http.StatusInternalServerError, errors.InternalError)
	}

	return nil
}

func (d *Database) AddComment(comment *Comment) *errors.Error {
	insertError := d.Comments.Insert(comment)
	if insertError != nil {
		log.Printf("Add comment error: error from database: %s", insertError.Error())
		return errors.New(http.StatusInternalServerError, errors.InternalError)
	}
	return nil
}

func (d *Database) GetAll(param bson.M) ([]*Post, *errors.Error) {
	var posts []*Post
	findError := d.Posts.Find(param).All(&posts)
	if findError != nil {
		log.Printf("Get posts error: error from database: %s", findError.Error())
		return nil, errors.New(http.StatusInternalServerError, errors.InternalError)
	}

	for _, post := range posts {
		findError = d.Votes.Find(bson.M{"post_id": post.Id}).All(&post.Votes)
		if findError != nil {
			log.Printf("Get votes for id %s error: error from database: %s", post.Id.Hex(), findError.Error())
			return nil, errors.New(http.StatusInternalServerError, errors.InternalError)
		}

		findError = d.Comments.Find(bson.M{"post_id": post.Id}).All(&post.Comments)
		if findError == mgo.ErrNotFound {
			post.Comments = []Comment{}
		} else if findError != nil {
			log.Printf("Get votes for id %s error: error from database: %s", post.Id.Hex(), findError.Error())
			return nil, errors.New(http.StatusInternalServerError, errors.InternalError)
		}
	}

	return posts, nil
}

func (d *Database) GetPostsByUsername(username string) ([]bson.ObjectId, *errors.Error) {
	var postIds []UserToPostId
	err := d.Users.Find(bson.M{"username": username}).All(&postIds)
	if err == mgo.ErrNotFound {
		log.Printf("No posts for user with login %s\n", username)
		return nil, nil
	} else if err != nil {
		log.Printf("Get posts by user %s error: %s\n", username, err.Error())
		return nil, errors.New(http.StatusInternalServerError, errors.InternalError)
	}
	ids := []bson.ObjectId{}
	for _, post := range postIds {
		ids = append(ids, post.PostId)
	}
	return ids, nil
}

func (d *Database) GetPost(id bson.ObjectId, view bool) (*Post, *errors.Error) {
	var post Post
	findError := d.Posts.Find(bson.M{"id": id}).One(&post)
	if findError != nil {
		log.Printf("Get post for id %s error: error from database: %s", id.Hex(), findError.Error())
		return nil, errors.New(http.StatusInternalServerError, errors.InternalError)
	}

	if view {
		post.Views += 1

		updateError := d.Posts.Update(bson.M{"id": id}, post)
		if updateError != nil {
			log.Printf("Update post for id %s error: error from database: %s", id.Hex(), updateError.Error())
			return nil, errors.New(http.StatusInternalServerError, errors.InternalError)
		}
	}

	findError = d.Votes.Find(bson.M{"post_id": id}).All(&post.Votes)
	if findError != nil {
		log.Printf("Get votes for id %s error: error from database: %s", id.Hex(), findError.Error())
		return nil, errors.New(http.StatusInternalServerError, errors.InternalError)
	}

	findError = d.Comments.Find(bson.M{"post_id": id}).All(&post.Comments)
	if findError == mgo.ErrNotFound {
		post.Comments = []Comment{}
	} else if findError != nil {
		log.Printf("Get votes for id %s error: error from database: %s", id.Hex(), findError.Error())
		return nil, errors.New(http.StatusInternalServerError, errors.InternalError)
	}

	return &post, nil
}

func (d *Database) GetVote(postId, userId string) (*Vote, *errors.Error) {
	var vote Vote
	findError := d.Votes.Find(
		bson.M{
			"post_id": bson.ObjectIdHex(postId),
			"user":    userId,
		},
	).One(&vote)
	if findError != nil {
		log.Printf("Get vote for post %s by user %s error: %s\n", postId, userId, findError.Error())
		return nil, errors.New(http.StatusInternalServerError, errors.InternalError)
	}
	return &vote, nil
}

func (d *Database) DeleteComment(postId, commentId bson.ObjectId) *errors.Error {
	deleteError := d.Comments.Remove(bson.M{
		"post_id": postId,
		"id":      commentId,
	})
	if deleteError != nil {
		log.Printf("Add comment error: error from database: %s", deleteError.Error())
		return errors.New(http.StatusInternalServerError, errors.InternalError)
	}
	return nil
}

func (d *Database) DeletePost(id bson.ObjectId) *errors.Error {
	postErr := d.Posts.Remove(bson.M{"id": id})
	votesErr := d.Votes.Remove(bson.M{"post_id": id})
	commentsErr := d.Comments.Remove(bson.M{"id": id})
	userErr := d.Users.Remove(bson.M{"post_id": id})
	if postErr != nil || votesErr != nil || commentsErr != nil || userErr != nil {
		return errors.New(http.StatusInternalServerError, errors.InternalError)
	}
	return nil
}

func (d *Database) ChangeVote(postId bson.ObjectId, userId string, turn int) *errors.Error {
	var vote Vote
	findError := d.Votes.Find(
		bson.M{
			"post_id": postId,
			"user":    userId,
		},
	).One(&vote)

	var post Post
	getErr := d.Posts.Find(bson.M{"id": postId}).One(&post)

	switch findError {
	case mgo.ErrNotFound:
		log.Printf("Add vote for new user %s\n", userId)
		vote := Vote{
			Post: postId,
			User: userId,
			Vote: turn,
		}
		addErr := d.Votes.Insert(vote)
		if addErr != nil {
			log.Printf("Cant insert vote for user %s in post %s: %s\n", userId, postId.Hex(), addErr.Error())
			return errors.New(http.StatusInternalServerError, errors.InternalError)
		}
		post.Score += turn
	case nil:
		log.Printf("Change vote for user %s\n", userId)
		if vote.Vote != turn {
			vote.Vote = turn
			updateErr := d.Votes.Update(
				bson.M{
					"post_id": postId,
					"user":    userId,
				},
				vote,
			)
			if updateErr != nil {
				log.Printf("Cant update vote for user %s in post %s: %s\n", userId, postId.Hex(), updateErr.Error())
				return errors.New(http.StatusInternalServerError, errors.InternalError)
			}
			post.Score += 2 * turn
		} else {
			removeErr := d.Votes.Remove(vote)
			if removeErr != nil {
				log.Printf("Cant delete vote for user %s in post %s: %s\n", userId, postId.Hex(), removeErr.Error())
				return errors.New(http.StatusInternalServerError, errors.InternalError)
			}
			post.Score -= turn
		}
	default:
		return errors.New(http.StatusInternalServerError, errors.InternalError)
	}

	var allVotes []Vote
	var upVotes []Vote
	getErr = d.Votes.Find(bson.M{"post_id": postId}).All(&allVotes)
	if getErr != nil {
		log.Printf("Cant get all votes for post %s: %s\n", postId.Hex(), getErr.Error())
		return errors.New(http.StatusInternalServerError, errors.InternalError)
	}
	getErr = d.Votes.Find(bson.M{"post_id": postId, "vote": 1}).All(&upVotes)
	if getErr != nil {
		log.Printf("Cant get up votes for post %s: %s\n", postId.Hex(), getErr.Error())
		return errors.New(http.StatusInternalServerError, errors.InternalError)
	}
	if len(allVotes) > 0 {
		post.Percentage = uint(math.Floor((float64(len(upVotes)) / float64(len(allVotes))) * 100))
	} else {
		post.Percentage = 100
	}

	updateErr := d.Posts.Update(bson.M{"id": postId}, post)
	if updateErr != nil {
		log.Printf("Cant update post %s: %s\n", postId.Hex(), updateErr.Error())
		return errors.New(http.StatusInternalServerError, errors.InternalError)
	}

	return nil
}
