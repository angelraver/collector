package entities

type IGDBGameResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
}

type IGDBCoverResponse struct {
	Id        int       `json:"id"`
	Url      string    `json:"url"`
}
