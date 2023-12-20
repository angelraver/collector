package entities

type IGDBGameResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	FirstReleaseDate int `json:"first_release_date"`
}

type IGDBCoverResponse struct {
	Id        int       `json:"id"`
	Url      string    `json:"url"`
}
