package db

import (
	"database/sql"
	"fmt"
	"movie-api/internal/movie"

	_ "github.com/mattn/go-sqlite3"
)

const DB_INIT_STRING string = "CREATE TABLE IF NOT EXISTS movies (id INTEGER PRIMARY KEY, title TEXT, director TEXT, year INTEGER)" 

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

func (s *SqliteRepo) CreateMovie(singleMovie movie.Movie) error {
	statement, err := s.conn.Prepare("INSERT INTO movies VALUES (?, ?, ?, ?)")

	if err != nil {
		return err
	}

    res, err := statement.Exec(singleMovie.ID, singleMovie.Title, singleMovie.Director, singleMovie.Year)

	fmt.Println(res)

	return err
}

func (s *SqliteRepo) UpdateMovie(movie.Movie) error {
	return nil
}

func (s *SqliteRepo) DeleteMovie(id int) error {
	 return nil
}

func NewSqliteRepo() (Repo, error) {
	db, err := sql.Open("sqlite3", "./movie.db")
	if err != nil {
		return nil, err
	}

	statement, err := db.Prepare(DB_INIT_STRING)
	if err != nil {
		return nil, err
	}

    statement.Exec()

	return &SqliteRepo{db}, nil
}
