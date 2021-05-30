package test

import (
	"fmt"
	"testing"

	resty "github.com/go-resty/resty/v2"
	assert "github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	fmt.Println("Running E2E test for health check endpoint")

	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/health")
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 200, resp.StatusCode())
}
