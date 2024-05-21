package routes

import (
	"log"
	"movie-card/controller"
	"net/http"

	"github.com/gorilla/mux"
)

const port = ":8000"

func Web() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", controller.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", controller.GetMovie).Methods("GET")
	r.HandleFunc("/movies", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", controller.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", controller.DeleteMovie).Methods("DELETE")


	log.Fatal(http.ListenAndServe(port, r))


}