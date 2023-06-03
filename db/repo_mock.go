package db

import "movie-api/internal/movie"

type MockRepo struct {}

func (m *MockRepo) GetAllMovies() []movie.Movie {
	return []movie.Movie{
		{ID: 1, Title: "some title 1", Director: "some director 1", Year: 2021},
		{ID: 2, Title: "some title 2", Director: "some director 2", Year: 2022},
	}
}

func (m *MockRepo) GetMovieByID(id int) movie.Movie {
	return movie.Movie{
		ID: id,
		Title: "some title " + string(rune(id)),
		Director: "some director " + string(rune(id)),
		Year: 2020 + id,
	}
}

func (m *MockRepo) CreateMovie(movie movie.Movie) error {
	return nil
}

func (m *MockRepo) UpdateMovie(movie movie.Movie) error {
	return nil
}

func (m *MockRepo) DeleteMovie(id int) error {
	return nil
}
