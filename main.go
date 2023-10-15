package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

var movies []Movie

func main(){
	r := mux.NewRouter()
	movieHandler := MovieHandler{}
	r.HandleFunc("/movies", movieHandler.getMovies).Methods("GET")
	r.HandleFunc("movies/{id}", movieHandler.getMovie).Methods("GET")
	r.HandleFunc("/movies/create", movieHandler.createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", movieHandler.updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", movieHandler.deleteMovie).Methods("DELETE")

	fmt.Print("Starting server at port 8000 \n")
	log.Fatal(http.ListenAndServe(":8000",r))
}