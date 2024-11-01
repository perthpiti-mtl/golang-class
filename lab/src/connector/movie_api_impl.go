package connector

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/golang-class/lab/config"
	"github.com/golang-class/lab/model"
)

type RealMovieAPIConnector struct {
	client  *http.Client
	baseURL string
}

type MovieSearchResponse struct {
	Ok          bool `json:"ok"`
	Description []struct {
		Title   string  `json:"title"`
		Year    int     `json:"year"`
		IMDBID  string  `json:"imdb_id"`
		Rank    int     `json:"rank"`
		Actors  string  `json:"actors"`
		IMDBURL string  `json:"imdb_url"`
		Rating  float32 `json:"rating"`
	} `json:"description"`
}

type MovieSearchDetail struct {
	Ok          bool `json:"ok"`
	Description struct {
		Title   string  `json:"title"`
		Year    int     `json:"year"`
		IMDBID  string  `json:"imdb_id"`
		Rank    int     `json:"rank"`
		Actors  string  `json:"actors"`
		IMDBURL string  `json:"imdb_url"`
		Rating  float32 `json:"rating"`
	} `json:"description"`
}

func (r *RealMovieAPIConnector) ListMovie(ctx context.Context) ([]model.Movie, error) {
	fullUrl := r.baseURL + "/list"
	method := "GET"
	req, err := http.NewRequestWithContext(ctx, method, fullUrl, nil)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable to search movie with API: %s", res.Status)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var movieSearchResponse MovieSearchResponse
	err = json.Unmarshal(body, &movieSearchResponse)
	if err != nil {
		return nil, err
	}
	var movies []model.Movie
	for _, movie := range movieSearchResponse.Description {
		movies = append(movies, model.Movie{
			MovieID: movie.IMDBID,
			Title:   movie.Title,
			Year:    movie.Year,
			Rating:  movie.Rating,
		})
	}
	return movies, nil
}

func (r *RealMovieAPIConnector) GetMovieDetail(ctx context.Context, movieId string) (*model.Movie, error) {
	fullUrl := fmt.Sprintf("%s/%s", r.baseURL, movieId)
	method := "GET"
	req, err := http.NewRequestWithContext(ctx, method, fullUrl, nil)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("movie not found")
		} else if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("unable to get movie detail with API: %s", res.Status)
		}
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var movieDetailResponse MovieSearchDetail
	err = json.Unmarshal(body, &movieDetailResponse)
	if err != nil {
		return nil, err
	}
	movie := &model.Movie{
		MovieID: movieDetailResponse.Description.IMDBID,
		Title:   movieDetailResponse.Description.Title,
		Year:    movieDetailResponse.Description.Year,
		Rating:  movieDetailResponse.Description.Rating,
	}
	if movie.Title == "" || movie.Year == 0 {
		return nil, fmt.Errorf("movie not found")
	}
	return movie, nil
}

func NewRealMovieAPI(c *config.Config) MovieAPIConnector {
	return &RealMovieAPIConnector{
		client:  &http.Client{},
		baseURL: c.MovieAPIConnector.URL,
	}
}
