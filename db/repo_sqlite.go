package db

import (
	"database/sql"
	"movie-api/internal/movie"
)



type SqliteRepo struct {
	conn *sql.DB
}

func (s *SqliteRepo) getAllMovies() []movie.Movie {

	res, err := s.conn.Query("SELECT * FROM MOVIES")

	var results []movie.Movie

	if err != nil {
		return results
	}

	for res.Next() {
		var singleMovie movie.Movie

		err = res.Scan(&singleMovie.ID, &singleMovie.Title, &singleMovie.Director, &singleMovie.Year)

		if err == nil {
			results = append(results, singleMovie)
		}
	}

	return results

}

func (s *SqliteRepo) getMovieByID(id int) movie.Movie {
	var singleMovie movie.Movie

	res, err := s.conn.Query("SELECT * FROM MOVIES WHERE ID = " + string(rune(id)))

	if err != nil {
		return singleMovie
	}

	_ = res.Scan(&singleMovie.ID, &singleMovie.Title, &singleMovie.Director, &singleMovie.Year)

	return singleMovie
}

func (s *SqliteRepo) createMovie(movie movie.Movie) error

func (s *SqliteRepo) updateMovie(movie movie.Movie) error

func (s *SqliteRepo) DeleteMovie(id int) error
