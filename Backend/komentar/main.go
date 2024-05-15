package main

import (
	"fmt"
	"komentar/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterCommentRoutes(r)
	http.Handle("/", r)
	fmt.Print("Starting Server localhost:8003")
	log.Fatal(http.ListenAndServe("localhost:8003", r))
}
