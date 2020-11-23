package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/mahendrafathan/golang_simple_registration_login/util"
)

func main() {
	util.InitConfig()
	util.ConnectDB()

	router := mux.NewRouter().StrictSlash(true)
	fileServer := http.FileServer(http.Dir("static/"))
	router.Handle("/static", http.NotFoundHandler())
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		util.Neuter(fileServer)))
	router.HandleFunc("/register", util.RegisterHandler)
	router.HandleFunc("/login", util.LoginHandler)
	router.HandleFunc("/logout", util.LogoutHandler)
	router.HandleFunc("/", util.HomeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}
	log.Println("starting serve on " + port)
	http.ListenAndServe(":"+port, router)
}
