package data

import (
	"bytes"
	"encoding/json"
	"log"
	"reddit/pkg/post"
	"reddit/pkg/user"
)

func GetTestUserReader() *bytes.Reader {
	usr := user.User{
		Login:    "Sonia",
		Password: "testpassword",
	}
	usrJson, err := json.Marshal(usr)
	if err != nil {
		log.Fatalf("GetTestUserReader() failed parse user json with error %s", err)
	}
	return bytes.NewReader(usrJson)
}

func GetWrongTestUserReader() *bytes.Reader {
	usr := struct {
		Username int
	}{
		Username: 1234,
	}
	usrJson, err := json.Marshal(usr)
	if err != nil {
		log.Fatalf("GetWrongTestUserReader() failed parse user json with error %s", err)
	}
	return bytes.NewReader(usrJson)
}

func GetTestPostDataReader(status int) *bytes.Reader {
	var data interface{}
	switch status {
	case 0:
		data = post.Post{
			Type:     "text",
			Title:    "Test Post!",
			Text:     "New test post",
			Category: "programming",
		}
	case 1:
		data = post.Post{
			Type:     "url",
			Title:    "Test Post!",
			Url:      "https://newTestUrl.test",
			Category: "programming",
		}
	case 2:
		data = post.Post{
			Type:     "text",
			Title:    "Test Post!",
			Category: "programming",
		}
	case 3:
		data = struct {
			Type int
		}{
			Type: 1234,
		}
	case 4:
		data = struct {
			Comment string `json:"comment"`
		}{
			Comment: "New comment",
		}
	case 5:
		data = struct {
			Comment int
		}{
			Comment: 1234,
		}
	}
	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("GetTestTextPostReader() failed parse user json with error %s", err)
	}
	return bytes.NewReader(dataJson)
}

func GetTestCommentReader() *bytes.Reader {
	comment := struct {
		Comment string `json:"comment"`
	}{
		Comment: "New comment",
	}
	commentJson, err := json.Marshal(comment)
	if err != nil {
		log.Fatalf("GetTestCommentReader() failed parse user json with error %s", err)
	}
	return bytes.NewReader(commentJson)
}

func GetWrongTestCommentReader() *bytes.Reader {
	comment := struct {
		Comment int
	}{
		Comment: 1234,
	}
	commentJson, err := json.Marshal(comment)
	if err != nil {
		log.Fatalf("GetWrongTestCommentReader() failed parse user json with error %s", err)
	}
	return bytes.NewReader(commentJson)
}
