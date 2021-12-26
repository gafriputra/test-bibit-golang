package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Response struct {
	Movie []Movie `json:"Search"`
}

type Movie struct {
	ImdbID string `json:"imdbID"`
	Title  string `json:"Title"`
	Year   int    `json:"Year"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type MovieDetail struct {
	ImdbID     string `json:"imdbID"`
	Title      string `json:"Title"`
	Year       int    `json:"Year"`
	Type       string `json:"Type"`
	Poster     string `json:"Poster"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Country    string `json:"Country"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
}

type ResponseDataMovieList struct {
	Data  []Movie
	Total int
}

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	log.Printf("Running")

	//Routing
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/movie", movie).Methods("GET")
	router.HandleFunc("/movie/{id}", movieDetail).Methods("GET")

	//Logging
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Logging to a file in Go!")

	//Serve
	srv := &http.Server{
		Handler:      router,
		Addr:         ":5000",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  40 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Printf("/home")
	fmt.Println("Hello")
}

func movie(w http.ResponseWriter, r *http.Request) {
	log.Printf("/movie")
	search_query := r.URL.Query().Get("searchword")
	page_query := r.URL.Query().Get("pagination")
	if len(search_query) == 0 || len(page_query) == 0 {
		log.Printf("Search Query or Page is not define")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	api_url := goDotEnvVariable("API_URL") + "?apikey=" + goDotEnvVariable("API_KEY") + "&s=" + search_query + "&page=" + page_query

	response, err := http.Get(api_url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	log.Printf("go to %s", api_url)

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	movielist := ResponseDataMovieList{responseObject.Movie, len(responseObject.Movie)}
	data, err := json.Marshal(movielist)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func movieDetail(w http.ResponseWriter, r *http.Request) {
	log.Printf("/movieDetail")
	movie_id := mux.Vars(r)["id"]
	api_url := goDotEnvVariable("API_URL") + "?apikey=" + goDotEnvVariable("API_KEY") + "&i=" + movie_id

	response, err := http.Get(api_url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	log.Printf("go to %s", api_url)

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject MovieDetail
	json.Unmarshal(responseData, &responseObject)
	json.NewEncoder(w).Encode(responseObject)

}
