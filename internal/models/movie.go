package models

type Movie struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	ReleaseDate string  `json:"release_date"`
	Overview    string  `json:"overview"`
	Rating      float64 `json:"vote_average"`
	PosterUrl   string  `json:"poster_path"`
	TrailerUrl  string  `json:"trailer_path"`
	RawData     string  `json:"-"`
}
