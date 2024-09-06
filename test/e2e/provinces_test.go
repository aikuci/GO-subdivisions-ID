package e2e

import (
	"testing"

	"github.com/aikuci/go-subdivisions-id/test"
)

func TestGetProvince(t *testing.T) {
	test.ClearAll()
	test.CreateProvinces(1)

	tests := []test.TestStruct{
		{
			name:          "Successful request: Get province by valid ID",
			route:         "/v1/provinces/1",
			expectedError: false,
			expectedCode:  test.StatusOK,
		},
	}

	test.ExecTestRequest(t, tests)
}

func TestGetProvinces(t *testing.T) {
	test.ClearAll()
	test.CreateProvinces(20)

	tests := []test.TestStruct{
		{
			name:          "Successful request: Get provinces",
			route:         "/v1/provinces",
			expectedError: false,
			expectedCode:  test.StatusOK,
		},
	}

	test.ExecTestRequest(t, tests)
}
