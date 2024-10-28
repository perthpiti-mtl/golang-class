package connector

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-class/lab/config"
	"github.com/golang-class/lab/model"
	"io"
	"net/http"
	"net/url"
)

type RealMovieAPIConnector struct {
	client  *http.Client
	baseURL string
}

type MovieSearchResponse struct {
	Ok          bool `json:"ok"`
	Description []struct {
		TITLE       string `json:"#TITLE"`
		YEAR        int    `json:"#YEAR,omitempty"`
		IMDBID      string `json:"#IMDB_ID"`
		RANK        int    `json:"#RANK"`
		ACTORS      string `json:"#ACTORS"`
		AKA         string `json:"#AKA"`
		IMDBURL     string `json:"#IMDB_URL"`
		IMDBIV      string `json:"#IMDB_IV"`
		IMGPOSTER   string `json:"#IMG_POSTER"`
		PhotoWidth  int    `json:"photo_width"`
		PhotoHeight int    `json:"photo_height"`
	} `json:"description"`
	ErrorCode int `json:"error_code"`
}

type MovieDetailResponse struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
	ImdbId      string `json:"imdbId"`
	Top         struct {
		Id        string `json:"id"`
		TitleText struct {
			Text     string `json:"text"`
			Typename string `json:"__typename"`
		} `json:"titleText"`
		ReleaseYear struct {
			Year     int    `json:"year"`
			EndYear  int    `json:"endYear"`
			Typename string `json:"__typename"`
		} `json:"releaseYear"`
		PrimaryImage struct {
			Id     string `json:"id"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
			Url    string `json:"url"`
		} `json:"primaryImage"`
	} `json:"top"`
}

func (r *RealMovieAPIConnector) SearchMovie(ctx context.Context, keyword string) ([]model.Movie, error) {
	fullUrl := r.baseURL + "/search"
	method := "GET"
	req, err := http.NewRequestWithContext(ctx, method, fullUrl, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("q", url.QueryEscape(keyword))
	req.URL.RawQuery = q.Encode()
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
			Title:   movie.TITLE,
			Year:    movie.YEAR,
			Image:   movie.IMGPOSTER,
		})
	}
	return movies, nil
}

func (r *RealMovieAPIConnector) GetMovieDetail(ctx context.Context, movieId string) (*model.Movie, error) {
	fullUrl := r.baseURL + "/search"
	method := "GET"
	req, err := http.NewRequestWithContext(ctx, method, fullUrl, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("tt", url.QueryEscape(movieId))
	req.URL.RawQuery = q.Encode()
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
	var movieDetailResponse MovieDetailResponse
	err = json.Unmarshal(body, &movieDetailResponse)
	if err != nil {
		return nil, err
	}
	movie := &model.Movie{
		MovieID: movieDetailResponse.ImdbId,
		Title:   movieDetailResponse.Top.TitleText.Text,
		Year:    movieDetailResponse.Top.ReleaseYear.Year,
		Image:   movieDetailResponse.Top.PrimaryImage.Url,
	}
	if movie.Title == "" || movie.Year == 0 {
		return nil, fmt.Errorf("movie not found")
	}
	return movie, nil
}

func NewRealMovieAPI(config *config.Config) MovieAPIConnector {
	return &RealMovieAPIConnector{
		client:  &http.Client{},
		baseURL: config.MovieAPI.Url,
	}
}
