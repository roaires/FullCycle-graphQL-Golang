package model

type Category struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`

	// Será implementado via Lazy loading
	// Courses     []*Course `json:"courses"`
}
