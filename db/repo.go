package db

import "movie-api/types"

type Repo interface {
	GetAllMovies() []movie.Movie
	GetMovieByID(id int) movie.Movie
	CreateMovie(movie.Movie) error
	UpdateMovie(movie.Movie) error
	DeleteMovie(id int) error
}
