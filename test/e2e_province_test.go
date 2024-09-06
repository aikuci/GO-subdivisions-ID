package test

import "testing"

func TestGetProvince(t *testing.T) {
	ClearAll()
	CreateProvinces(1)

	tests := []TestStruct{
		{
			name:          "Successful request: Get province by valid ID",
			route:         "/v1/provinces/1",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Bad request: Invalid 'include' query parameter",
			route:         "/v1/provinces/1?include=relation", // Assuming 'relation' is an invalid include parameter
			expectedError: false,
			expectedCode:  StatusBadRequest,
		},
		{
			name:          "Not found: Get province by unregistered ID",
			route:         "/v1/provinces/0", // Assuming there is no ID registered with 0 value
			expectedError: false,
			expectedCode:  StatusNotFound,
		},
		{
			name:          "Not found: Invalid ID format in route",
			route:         "/v1/provinces/province",
			expectedError: false,
			expectedCode:  StatusNotFound,
		},
	}

	ExecTestRequest(t, tests)
}

func TestGetProvinces(t *testing.T) {
	ClearAll()
	CreateProvinces(20)

	tests := []TestStruct{
		{
			name:          "Successful request: Get provinces",
			route:         "/v1/provinces",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Bad request: Invalid 'include' query parameter",
			route:         "/v1/provinces?include=relation", // Assuming 'relation' is an invalid include parameter
			expectedError: false,
			expectedCode:  StatusBadRequest,
		},
	}

	ExecTestRequest(t, tests)
}
