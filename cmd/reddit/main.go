package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"html/template"
	"log"
	"net/http"
	"reddit/pkg/auth"
	"reddit/pkg/database"
	"reddit/pkg/handlers"
	"reddit/pkg/interfaces"
	"reddit/pkg/middleware"
	"reddit/pkg/post"
	"reddit/pkg/user"
)

type MgoSession struct {
	session *mgo.Session
}

func (m MgoSession) DB(name string) interfaces.MongoDatabase {
	return MgoDatabase{m.session.DB(name)}
}

type MgoDatabase struct {
	database *mgo.Database
}

func (m MgoDatabase) C(name string) interfaces.MongoCollection {
	return MgoCollection{m.database.C(name)}
}

type MgoCollection struct {
	collection *mgo.Collection
}

func (m MgoCollection) Find(data interface{}) interfaces.MongoQuery {
	return m.collection.Find(data)
}

func (m MgoCollection) Insert(data ...interface{}) error {
	return m.collection.Insert(data)
}

func (m MgoCollection) Update(req interface{}, data interface{}) error {
	return m.collection.Update(req, data)
}

func (m MgoCollection) Remove(data interface{}) error {
	return m.collection.Remove(data)
}

func main() {
	db, dbError := sql.Open("mysql", "root:guest@tcp(localhost:3306)/reddit?charset=utf8&interpolateParams=true")
	if dbError != nil {
		log.Fatalf("Cannot open database: %s", dbError.Error())
	}

	dbError = db.Ping()
	if dbError != nil {
		log.Fatalf("Cannot connect to database: %s", dbError.Error())
	}

	session, mgoError := mgo.Dial("mongodb://localhost")
	if mgoError != nil {
		log.Fatalf("Cannot connect to database: %s", mgoError.Error())
	}

	userHandler := handlers.UserHandler{
		Repo:     user.NewRepo((*user.Database)(database.InitDB(db))),
		Sessions: (*auth.Database)(database.InitDB(db)),
	}

	postHandler := handlers.PostHandler{
		Repo: post.NewRepo(MgoSession{session}),
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/register", userHandler.Register).Methods(http.MethodPost)
	router.HandleFunc("/api/login", userHandler.Login).Methods(http.MethodPost)

	router.HandleFunc("/api/posts/", postHandler.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/api/posts", postHandler.AddPost).Methods(http.MethodPost)
	router.HandleFunc("/api/posts/{category}", postHandler.GetByCategory).Methods(http.MethodGet)

	router.HandleFunc("/api/post/{post_id}", postHandler.GetPost).Methods(http.MethodGet)
	router.HandleFunc("/api/post/{post_id}", postHandler.AddComment).Methods(http.MethodPost)
	router.HandleFunc("/api/post/{post_id}/{comment_id}", postHandler.DeleteComment).Methods(http.MethodDelete)
	router.HandleFunc("/api/post/{post_id}/upvote", postHandler.Upvote).Methods(http.MethodGet)
	router.HandleFunc("/api/post/{post_id}/downvote", postHandler.Downvote).Methods(http.MethodGet)
	router.HandleFunc("/api/post/{post_id}/unvote", postHandler.Unvote).Methods(http.MethodGet)
	router.HandleFunc("/api/post/{post_id}", postHandler.DeletePost).Methods(http.MethodDelete)

	router.HandleFunc("/api/user/{user_login}", postHandler.GetByUser).Methods(http.MethodGet)

	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("redditclone/template/static")))
	router.PathPrefix("/static/").Handler(staticHandler)

	tmpl := template.Must(template.ParseFiles("./redditclone/template/index.html"))
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	middlewareHandler := middleware.Middleware{
		Sessions: (*auth.Database)(database.InitDB(db)),
	}

	middlewares := middlewareHandler.Auth(router)

	err := http.ListenAndServe(":8080", middlewares)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
