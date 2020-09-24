package data

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"reddit/pkg/post"
	"time"
)

func CheckStatus(respCode, expCode int) (bool, *string) {
	if respCode != expCode {
		err := fmt.Sprintf("Wrong status: got: %v, expected: %v", respCode, expCode)
		return false, &err
	}
	return true, nil
}

func CheckResponse(resp, exp string) (bool, *string) {
	if resp != exp+"\n" {
		err := fmt.Sprintf("Unexpected body: got: %v, expected: %v", resp, exp+"\n")
		return false, &err
	}
	return true, nil
}

func ChangePostGenerations() {
	post.GeneratePostId = func() bson.ObjectId {
		return "au13151fwf56" // 617531333135316677663536
	}
	post.GenerateCommentId = func() bson.ObjectId {
		return "au13151fwf56"
	}
	post.GenerateCreatedTime = func(tm time.Time) string {
		return "2020-04-17T01:31:56Z"
	}
}

func UnchangePostGenerations() {
	post.GeneratePostId = post.GenerateId
	post.GenerateCommentId = post.GenerateId
	post.GenerateCommentId = post.GenerateId
}
