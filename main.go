package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	firstName string `json:"firstName"`
	lastName string `json:"lastName"`
}

var movies []Movie

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/create", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	fmt.Print("Starting server at port 8000 \n")
	log.Fatal(http.ListenAndServe(":8000",r))
}