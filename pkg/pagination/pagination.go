package pagination

import (
	"fmt"
	"goblong/pkg/config"
	"goblong/pkg/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"net/http"
	"strconv"
	"strings"
)

// Define page
type Page struct {
	// Link
	URL string
	// Number
	Number int
}

// Define view data for render
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
type Pagination struct {
	BaseURL string
	PerPage int
	Page    int
	Count   int64
	db      *gorm.DB
}

// New pagination object construct
// r Get pagination parameters, default is page, can modify from config/pagination.go
// db GORM query handle
// baseURL  Get page link
// PerPage // Each page total, when give parameter less than or equal 0 use default, can modify from config/pagination.go
func New(r *http.Request, db *gorm.DB, baseURL string, PerPage int) *Pagination {
	// Default each page length
	if PerPage <= 0 {
		PerPage = config.GetInt("pagination.perPage")
	}

	// New instance
	p := &Pagination{
		db:      db,
		PerPage: PerPage,
		Page:    1,
		Count:   -1,
	}

	if strings.Contains(baseURL, "?") {
		p.BaseURL = baseURL + "&" + config.GetString("pagination.url_query") + "="
	} else {
		p.BaseURL = baseURL + "?" + config.GetString("pagination.url_query") + "="
	}

	// Set current page
	p.SetPage(p.GetPageFromRequest(r))

	return p
}

// Paging return render page data
func (p *Pagination) Paging() ViewData {
	return ViewData{
		HasPages: p.HasPages(),

		Next:    p.NewPage(p.NextPage()),
		HasNext: p.HasNext(),

		Prev:    p.NewPage(p.PrevPage()),
		HasPrev: p.HasPrev(),

		Current:   p.NewPage(p.CurrentPage()),
		TotalPage: p.TotalPage(),

		TotalCount: p.Count,
	}
}

// New Page set current page
func (p Pagination) NewPage(page int) Page {
	return Page{
		Number: page,
		URL:    p.BaseURL + strconv.Itoa(page),
	}
}

// Set current page
func (p *Pagination) SetPage(page int) {
	if page <= 0 {
		page = 1
	}
	p.Page = page
}

// Return current page
func (p *Pagination) CurrentPage() int {
	totalPage := p.TotalPage()
	if totalPage == 0 {
		return 0
	}
	if p.Page > totalPage {
		return totalPage
	}
	return p.Page
}

// Return request pagination data
func (p Pagination) Results(data interface{}) error {
	var err error
	var offset int
	page := p.CurrentPage()
	if page == 0 {
		return err
	}
	if page > 1 {
		offset = (page - 1) * p.PerPage
	}

	return p.db.Preload(clause.Associations).Limit(p.PerPage).Offset(offset).Find(data).Error
}

// Return total count
func (p *Pagination) TotalCount() int64 {
	if p.Count == -1 {
		var count int64
		if err := p.db.Count(&count).Error; err != nil {
			return 0
		}
		p.Count = count
	}
	return p.Count
}

// Return total page
func (p Pagination) TotalPage() int {
	count := p.TotalCount()
	if count == 0 {
		return 0
	}
	nums := int64(math.Ceil(float64(count) / float64(p.PerPage)))
	if nums == 0 {
		return 1
	}
	return int(nums)
}

// HasPages total page granter than 1 return true
func (p *Pagination) HasPages() bool {
	n := p.TotalCount()
	return n > int64(p.PerPage)
}

// Get Prev page
func (p Pagination) PrevPage() int {
	hasPrev := p.HasPrev()

	if !hasPrev {
		fmt.Println("Does'not exists prev page")
		return 0
	}
	page := p.CurrentPage()
	if page == 0 {
		fmt.Println("Prev page is 0")
		return 0
	}
	fmt.Println("Prev page:", page-1)
	return page - 1
}

// Get Next page
func (p Pagination) NextPage() int {
	hasNext := p.HasNext()
	if !hasNext {
		return 0
	}

	page := p.CurrentPage()
	if page == 0 {
		return 0
	}
	return page + 1
}

// HasNext returns true if current page is not the last page
func (p Pagination) HasNext() bool {
	totalPage := p.TotalPage()
	if totalPage == 0 {
		return false
	}
	page := p.CurrentPage()
	if page == 0 {
		return false
	}
	return page < totalPage
}

// Has prev if current page is not first page exists prev page
func (p Pagination) HasPrev() bool {
	page := p.CurrentPage()
	if page == 0 {
		return false
	}
	return page > 1
}

// Get page parameter from request
func (p Pagination) GetPageFromRequest(r *http.Request) int {
	page := r.URL.Query().Get(config.GetString("pagination.url_query"))
	pageInt := types.StringToInt(page)
	if pageInt <= 0 {
		return 1
	}
	return pageInt
}
