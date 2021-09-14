package tests

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllPage(t *testing.T) {
	baseURL := "http://localhost:3000"

	// Define all initilization test data
	var tests = []struct {
		method   string
		url      string
		expected int
	}{
		{"GET", "/", 200},
		{"GET", "/about", 200},
		{"GET", "/notfound", 404},
		{"GET", "/articles", 200},
		{"GET", "/articles/create", 200},
		{"GET", "/articles/3", 200},
		{"GET", "/articles/3/edit", 200},
		{"POST", "/articles/3", 200},
		{"POST", "/articles", 200},
		{"POST", "/articles/1/delete", 404},
	}

	for _, test := range tests {
		t.Logf("Current request URL: %v \n", test.url)
		var (
			resp *http.Response
			err  error
		)

		switch {
		case test.method == "POST":
			data := make(map[string][]string)
			resp, err = http.PostForm(baseURL+test.url, data)
		default:
			resp, err = http.Get(baseURL + test.url)
		}

		assert.NoError(t, err, "Request "+test.url+" report error")
		assert.Equal(t, test.expected, resp.StatusCode, test.url+" should response status code "+strconv.Itoa(test.expected))

	}

}
