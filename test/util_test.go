package test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestStruct represents a single test case for an HTTP request.
type TestStruct struct {
	name          string
	route         string
	expectedError bool
	expectedCode  int
}

const (
	// Status codes
	StatusOK = 200

	StatusBadRequest = 400
	StatusNotFound   = 404
)

func ExecTestRequest(t *testing.T, tests []TestStruct) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", test.route, nil)
			req.Header.Set("Content-Type", "application/json")

			// Perform the request using the app.
			resp, err := app.Test(req, -1) // the -1 disables request latency

			if test.expectedError {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.NoError(t, err, "Unexpected error occurred")
				assert.Equal(t, test.expectedCode, resp.StatusCode)
			}
		})
	}
}
