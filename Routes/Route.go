package Routes

import (
	"Curd-Api-Http/Controllers"
	"Curd-Api-Http/Database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start(){

	Database.ConnectDB()
	
	//Init Router
	r := mux.NewRouter()

  	// arrange our route
	r.HandleFunc("/movie",Controllers.GetMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", Controllers.GetMovie).Methods("GET")
	r.HandleFunc("/movie",Controllers.CreateMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", Controllers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", Controllers.DeleteMovie).Methods("DELETE")

  	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))
}