package main

import (
	"fmt"
	"log"
	"mainpkg/config"
	"mainpkg/mydb"
	"mainpkg/session"
	"net/http"
	"os"
)

type tinyHandler struct{}

func (tinyHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	http.DefaultServeMux.ServeHTTP(writer, req)
}

var sessions *session.T
var mongodb *mydb.T

func main() {
	configFile := config.GetConfigFile("config.yaml")

	sessions = session.Init(configFile.CookieDurationInSeconds)
	mongodb = mydb.Init(configFile.MongodbURI)

	defer sessions.Close()

	http.HandleFunc("/api/login", login)
	http.HandleFunc("/api/note", note)
	http.HandleFunc("/api/logout", logout)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./js-build/static"))))
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		http.ServeFile(writer, req, "./js-build/index.html")
	})

	var PORT string
	if value, ok := os.LookupEnv("PORT"); ok {
		PORT = value
	} else {
		PORT = "3000"
	}

	fmt.Printf("Serving at %s", PORT)

	if http.ListenAndServe(fmt.Sprintf(":%s", PORT), tinyHandler{}) == nil {
		log.Fatal("Something happened")
	}
}
