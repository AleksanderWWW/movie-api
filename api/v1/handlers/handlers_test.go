package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"movie-api/db"
	"movie-api/internal/movie"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
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

func TestUpdateMovieHandler(t *testing.T) {
	server := httptest.NewServer(UpdateMovieHandler(&db.MockRepo{}))

	type test struct {
		singleMovie         movie.Movie
		expectedStatusCode  int
		expectedResponseMsg string
	}

	tests := []test{
		{
			singleMovie: movie.Movie{
				ID:       1,
				Title:    "new title",
				Director: "new director",
				Year:     2004,
			},
			expectedStatusCode:  200,
			expectedResponseMsg: "Movie with ID '1' updated",
		},
		{
			singleMovie: movie.Movie{
				ID:       1234,
				Title:    "new title",
				Director: "new director",
				Year:     2004,
			},
			expectedStatusCode:  400,
			expectedResponseMsg: "Movie with ID '1234' does not exist",
		},
	}

	type responseBody struct {
		Msg string `json:"msg"`
	}

	for _, tc := range tests {
		jsonData, err := json.Marshal(tc.singleMovie)

		if err != nil {
			t.Error(err)
		}

		req, err := http.NewRequest(http.MethodPut, server.URL, bytes.NewReader(jsonData))

		if err != nil {
			t.Error(err)
		}

		req.Header.Set("Content-Type", "application/json")
		client := http.Client{Timeout: 10 * time.Second}
		resp, err := client.Do(req)

		if resp.StatusCode != tc.expectedStatusCode {
			t.Errorf("Invalid response status code. Expected: %d. Got: %d",
				tc.expectedStatusCode, resp.StatusCode)
		}

		var body responseBody
		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&body)
		if err != nil {
			t.Error(err)
		}

		if body.Msg != tc.expectedResponseMsg {
			t.Errorf("Wrong response message\n Expected: %s\n Got: %s",
				tc.expectedResponseMsg, body.Msg)
		}
	}
}

func TestDeleteMovieHandler(t *testing.T) {
	server := httptest.NewServer(DeleteMovieHandler(&db.MockRepo{}))
	type responseBody struct {
		Msg string `json:"msg"`
	}

	type test struct {
		movieID         	int
		expectedStatusCode  int
		expectedResponseMsg string
	}

	tests := []test{
		{
			movieID: 1,
			expectedStatusCode:  200,
			expectedResponseMsg: "Movie with ID '1' deleted",
		},
		{
			movieID: 1234,
			expectedStatusCode:  404,
			expectedResponseMsg: "Movie with ID '1234' does not exist",
		},
	}

	for _, tc := range tests {
		url := fmt.Sprintf("%s?movieID=%d", server.URL, tc.movieID)
		req, err := http.NewRequest(http.MethodDelete, url, nil)

		if err != nil {
			t.Error(err)
		}

		req.Header.Set("Content-Type", "application/json")
		client := http.Client{Timeout: 10 * time.Second}
		resp, err := client.Do(req)

		if resp.StatusCode != tc.expectedStatusCode {
			t.Errorf("Invalid response status code. Expected: %d. Got: %d",
				tc.expectedStatusCode, resp.StatusCode)
		}

		var body responseBody
		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&body)
		if err != nil {
			t.Error(err)
		}

		if body.Msg != tc.expectedResponseMsg {
			t.Errorf("Wrong response message\n Expected: %s\n Got: %s",
				tc.expectedResponseMsg, body.Msg)
		}
	}
}