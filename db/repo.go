package db

import "movie-api/internal/movie"


type Repo interface {
	GetAllMovies() []movie.Movie
	GetMovieByID(id int) movie.Movie
	CreateMovie(movie.Movie) error
	UpdateMovie(movie.Movie) error
	DeleteMovie(id int) error
}
