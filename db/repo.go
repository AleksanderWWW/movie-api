package db

import "movie-api/internal/movie"


type Repo interface {
	getAllMovies() []movie.Movie
	getMovieByID(id int) movie.Movie
	createMovie(movie.Movie)
	UpdateMovie(movie.Movie)
	DeleteMovie(id int)
}
