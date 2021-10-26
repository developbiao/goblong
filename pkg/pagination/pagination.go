package pagination

import "gorm.io/gorm"

// Define page
type Page struct {
	// Link
	URL string
	// Number
	Number int
}

// Define view data
type ViewData struct {
	// heed pagination
	HasPages bool

	// Next page
	Next    Page
	HasNext bool

	// Pre page
	Prev    Page
	HasPrev bool

	// Current page
	Current Page

	// Total of count
	TotalCount int64

	// Total page
	TotalPage int
}

// Pagination object
type pagination struct {
	BaseURL string
	PerPage int
	Page    int
	Count   int64
	db      *gorm.DB
}
