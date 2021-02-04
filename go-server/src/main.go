package main

import (
	"fmt"
	"log"
	"mainpkg/config"
	"mainpkg/mydb"
	"mainpkg/session"
	"mainpkg/util"
	"net/http"
	"os"
)

type tinyHandler struct{}

func (tinyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if util.EnableCors(&w, r) {
		return
	}
	http.DefaultServeMux.ServeHTTP(w, r)
}

var sessions *session.T
var mongodb *mydb.T

func main() {
	configFile := config.GetConfigFile("config.yaml")

	sessions = session.Init(configFile.CookieDurationInSeconds)
	mongodb = mydb.Init(configFile.MongodbURI)

	defer sessions.Close()

	http.HandleFunc("/login", login)
	http.HandleFunc("/note", note)
	http.HandleFunc("/logout", logout)

	var PORT string
	if value, ok := os.LookupEnv("PORT"); ok {
		PORT = value
	} else {
		PORT = "8090"
	}

	fmt.Printf("Serving at %s", PORT)
	if http.ListenAndServe(fmt.Sprintf(":%s", PORT), tinyHandler{}) == nil {
		log.Fatal("Something happened")
	}
}
