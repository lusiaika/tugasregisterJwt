package main

import (
	"log"
	"net/http"
	"time"

	"tugasregisterjwt/database"
	"tugasregisterjwt/handler"

	"github.com/gorilla/mux"
)

func main() {

	handler.ParseConfig()
	sql := database.NewSqlConnection(handler.GetConnectionString())
	handler.SqlConnect = sql
	r := mux.NewRouter()

	handler.InstallUserAPI(r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	//http.ListenAndServe(PORT, nil)
	log.Fatal(srv.ListenAndServe())
}
