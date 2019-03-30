package models

//Session .
type Session struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Depth       int       `json:"depth"`
	Children    []Session `json:"children"`
	Status      string    `json:"status"`
	ContentType string    `json:"contentType"`
	IsEnabled   bool      `json:"isEnabled"`
	IsCurrent   bool      `json:"isCurrent"`
}
