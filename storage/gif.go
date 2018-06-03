package storage

import (
	"fmt"
)

// Gif is a structure holds gif's info.
type Gif struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Validate is a method to validate gif
//
// Rules are next: Name and Slug are not empty, Width and Height
// are greater than zero. In case of validation fail, ValidationError
// is returned.
func (r *Gif) Validate() error {
	if r.Name == "" {
		return ValidationError("name")
	}

	if r.Slug == "" {
		return ValidationError("slug")
	}

	if r.Width <= 0 {
		return ValidationError("width")
	}

	if r.Height <= 0 {
		return ValidationError("height")
	}

	return nil
}

// Update is a method which updates gif's fields using
// provided item. It copies all non-empty fields except of ID.
func (r *Gif) Update(item Gif) {
	if item.Name != "" {
		r.Name = item.Name
	}

	if item.Slug != "" {
		r.Slug = item.Slug
	}

	if item.Width != 0 {
		r.Width = item.Width
	}

	if item.Height != 0 {
		r.Height = item.Height
	}
}

func (r *Gif) String() string {
	return fmt.Sprintf("GIF<%d: %s>", r.ID, r.Name)
}
