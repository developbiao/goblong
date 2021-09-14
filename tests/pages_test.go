package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomePage(t *testing.T) {
	baseURL := "http://localhost:3000"

	// Mock user request broswer

	var (
		resp *http.Response
		err  error
	)

	resp, err = http.Get(baseURL + "/")

	// 2. Check status is 200
	assert.NoError(t, err, "Has error happen, err is not empty")
	assert.Equal(t, 200, resp.StatusCode, "Should be return status is 200")

}

func TestAboutPage(t *testing.T) {
	baseURL := "http://localhsot:3000"

	var (
		resp *http.Response
		err  error
	)

	resp, err = http.Get(baseURL + "/about")

	// Assert check status
	assert.NoError(t, err, "has error happen,  err is not empty")
	assert.Equal(t, 200, resp.StatusCode, "Should be return status is  200")
}
