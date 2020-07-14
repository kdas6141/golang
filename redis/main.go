package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/go-redis/redis"
	"html/template"
)

var client *redis.Client
var templates *template.Template

func main() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	templates = template.Must(template.ParseGlob("templates/*.html"))
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", indexGetHandler).Methods("GET")
    myRouter.HandleFunc("/", indexPostHandler).Methods("POST")
	http.Handle("/", myRouter)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	comments, err := client.LRange(client.Context(), "comments", 0, 10).Result()
	if err != nil {
		return
	}
	templates.ExecuteTemplate(w, "index.html", comments)
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	comment := r.PostForm.Get("comment")
	client.LPush(client.Context(), "comments", comment)
	http.Redirect(w, r, "/", 302)
}

