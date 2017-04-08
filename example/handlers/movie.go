package handlers

import (
	"encoding/json"
	"github.com/zpatrick/go-plugin-swagger/example/movie"
	"net/http"
)

func ListMovies(w http.ResponseWriter, r *http.Request, store *movie.MovieStore) {
	movies := store.Movies()
	bytes, err := movies.JSON()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(bytes)
	return
}

func AddMovie(w http.ResponseWriter, r *http.Request, store *movie.MovieStore) {
	var movie movie.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	store.Insert(movie)

	bytes, err := movie.JSON()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(bytes)
	return
}
