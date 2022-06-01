package main

import (
	"encoding/json" //to encode data into json while sending to postman
	//for generating movie-id
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//struct of movies
type Movie struct {
	ID       string    `json: "id"`
	Isbn     string    `json: "isbn"`
	title    string    `json: "title"`
	Director *Director `json: director`
}

//struct of directors
type Director struct {
	FirstName string `json: "firstname"`
	LastName  string `json: "lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies) //encoding the response to sent
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for idx, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)
			break
		}
	}
	// reutnr all the remaining movies after removing
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)

		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // setting content type as json from our custom struct that we created
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(1000000))

	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for idx, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)

			var movie Movie

			_ = json.NewDecoder(r.Body).Decode(&movie)

			movie.ID = params["id"]

			movies = append(movies, movie)

			json.NewEncoder(w).Encode(movies)

			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "123456", title: "Movie 1", Director: &Director{FirstName: "Sean", LastName: "Paul"}})
	movies = append(movies, Movie{ID: "2", Isbn: "121312", title: "Movie 2", Director: &Director{FirstName: "Alex", LastName: "Lo"}})

	//ROUTES
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
