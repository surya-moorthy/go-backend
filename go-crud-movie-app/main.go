package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
     ID string 	`json:"id"`
	 Isbn string `json:"isbn"`
	 Title string `json:"title"`
	 Director *Director `json:"director"`
}

type Director struct{
     Firstname string `json:"firstname"`
	 Lastname string  `json:"lastname"`

}

var movies[] Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type","application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index],movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) { 
	w.Header().Set("Content-type","application/json")
	params := mux.Vars(r)

	for _ , item := range movies {

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type","appication/json")
	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.IntN(4151152562))
	movies = append(movies,movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type","application/json")
	params := mux.Vars(r)

	for index , item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index],movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)

			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}   
}

func main() {
	r := mux.NewRouter()


    movies = append(movies,Movie{ID: "1", Isbn: "526254125232",Title: "Vallavan",Director: &Director{Firstname: "Yaaro", Lastname: "Oruthan"}})  
    movies = append(movies,Movie{ID: "2", Isbn: "5262242455232",Title: "Vallavanaa",Director: &Director{Firstname: "Yavano", Lastname: "Oruthan"}})

	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("the Server is getting started in port 8080")

	if err := http.ListenAndServe(":8080",r); err != nil {
		log.Fatal(err)
	}
}