package main

import (
	"github.com/gorilla/mux"
	"github.com/moatas19m/GoLang-FirstDive/pkg/routes"
	_ "gorm.io/driver/sqlserver"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9000", r))

}
