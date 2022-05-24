package domain

type Movie struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"Description"`
	ReleaseDate string `json:"release_date"`
	Image       string `json:"image"`
}
