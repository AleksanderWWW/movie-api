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

	if err != nil {
		return []movie.Movie{}
	}

	var results []movie.Movie

	for res.Next() {
		var singleMovie movie.Movie

		err = res.Scan(&singleMovie.ID, &singleMovie.Title, &singleMovie.Director, &singleMovie.Year)

		if err == nil {
			results = append(results, singleMovie)
		}
	}

	return results

}
