package db

import "movie-api/internal/movie"


type Repo interface {
	getAllMovies() []movie.Movie
	getMovieByID(id int) movie.Movie
	createMovie(movie.Movie) error
	UpdateMovie(movie.Movie) error
	DeleteMovie(id int) error
}
