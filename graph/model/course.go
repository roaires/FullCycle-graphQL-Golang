package model

type Course struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Category    *Category `json:"category"`

	// Será implementado via Lazy loading
	//Chapters    []*Chapter `json:"chapters"`
}
