package db

import "movie-api/models"

type Repo interface {
	GetAllMovies() []models.Movie
	GetMovieByID(id int) models.Movie
	CreateMovie(models.Movie) error
	UpdateMovie(models.Movie) error
	DeleteMovie(id int) error
}
