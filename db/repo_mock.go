package db

import (
	"errors"
	"movie-api/internal/movie"
)

var storage = []movie.Movie{
	{ID: 1, Title: "some title 1", Director: "some director 1", Year: 2021},
	{ID: 2, Title: "some title 2", Director: "some director 2", Year: 2022},
}

type MockRepo struct{}

func (m *MockRepo) GetAllMovies() []movie.Movie {
	return storage
}

func (m *MockRepo) GetMovieByID(id int) movie.Movie {
	for _, mv := range storage {
		if mv.ID == id {
			return mv
		}
	}
	return movie.Movie{}
}

func (m *MockRepo) CreateMovie(movie movie.Movie) error {
	return nil
}

func (m *MockRepo) UpdateMovie(movie movie.Movie) error {
	for _, mv := range storage {
		if mv.ID == movie.ID {
			return nil
		}
	}
	return errors.New("Error updating movie.")
}

func (m *MockRepo) DeleteMovie(id int) error {
	return nil
}
