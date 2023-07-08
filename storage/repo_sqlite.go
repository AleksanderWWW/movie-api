package storage

import (
	"movie-api/models"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type SqliteRepo struct {
	db *gorm.DB
}

func (s *SqliteRepo) GetAllMovies() []models.Movie {
	var movies []models.Movie
	s.db.Find(&movies)

	return movies
}

func (s *SqliteRepo) GetMovieByID(id int) models.Movie {
	singleMovie := models.Movie{ID: id}

	s.db.First(&singleMovie)

	return singleMovie
}

func (s *SqliteRepo) CreateMovie(singleMovie models.Movie) error {
	s.db.Create(&singleMovie)
	return nil
}

func (s *SqliteRepo) UpdateMovie(singleMovie models.Movie) error {
	oldMovie := s.GetMovieByID(singleMovie.ID)

	s.db.Model(&oldMovie).Updates(singleMovie)
	return nil
}

func (s *SqliteRepo) DeleteMovie(id int) error {
	s.db.Delete(&models.Movie{}, id)
	return nil
}

func NewSqliteRepo(dbPath string) (Repo, error) {
	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	database.AutoMigrate(&models.Movie{})

	return &SqliteRepo{database}, nil
}
