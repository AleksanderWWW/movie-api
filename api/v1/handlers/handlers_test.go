package handlers

import (
	"encoding/json"
	"movie-api/db"
	"movie-api/internal/movie"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetAllMoviesHandler(t *testing.T) {
	server := httptest.NewServer(GetAllMoviesHandler(&db.MockRepo{}))
	resp, err := http.Get(server.URL)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Invalid response status code. Expected: 200. Got: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	var movies []movie.Movie
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&movies)

	if err != nil {
		t.Error(err)
	}

	expected := []movie.Movie{
		{ID: 1, Title: "some title 1", Director: "some director 1", Year: 2021},
		{ID: 2, Title: "some title 2", Director: "some director 2", Year: 2022},
	}

	if !reflect.DeepEqual(movies, expected) {
		t.Errorf("Invalid response body.\n Expected: %v\n Actual: %v", expected, movies)
	}
}
