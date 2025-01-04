package main

import ()

type AnimeResponse struct {
	Data AnimeData `json:"data"`
}

type ImageURL struct {
	ImageURL string `json:"image_url"`
}

type AnimeData struct {
	MalID      int     `json:"mal_id"`
	URL        string  `json:"url"`
	Title      string  `json:"title"`
	Score      float64 `json:"score"`
	Synopsis   string  `json:"synopsis"`
	Episodes   int     `json:"episodes"`
	Duration   string  `json:"duration"`
	Popularity int     `json:"popularity"`
	Image      struct {
		Jpg ImageURL `json:"jpg"`
	} `json:"images"`
}
