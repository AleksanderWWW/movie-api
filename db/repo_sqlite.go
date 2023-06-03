package db

import (
	"database/sql"
	"movie-api/internal/movie"
)



type SqliteRepo struct {
	conn *sql.DB
}

func (s *SqliteRepo) GetAllMovies() []movie.Movie {

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

func (s *SqliteRepo) GetMovieByID(id int) movie.Movie {
	var singleMovie movie.Movie

	res, err := s.conn.Query("SELECT * FROM MOVIES WHERE ID = " + string(rune(id)))

	if err != nil {
		return singleMovie
	}

	_ = res.Scan(&singleMovie.ID, &singleMovie.Title, &singleMovie.Director, &singleMovie.Year)

	return singleMovie
}
