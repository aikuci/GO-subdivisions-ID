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
			name:          "Not found: Get province by unregistered ID",
			route:         "/v1/provinces/0",
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
	}

	ExecTestRequest(t, tests)
}

func TestGetProvincesWithItsCities(t *testing.T) {
	ClearAll()
	CreateProvincesAndCities(30, 1)

	tests := []TestStruct{
		{
			name:          "Successful request: Get provinces include its cities",
			route:         "/v1/provinces?include=cities",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Bad request: Invalid 'include' query parameter",
			route:         "/v1/provinces?include=relation",
			expectedError: false,
			expectedCode:  StatusBadRequest,
		},

		{
			name:          "Successful request: Get province by valid ID include its cities",
			route:         "/v1/provinces/1?include=cities",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Bad request: Invalid 'include' query parameter",
			route:         "/v1/provinces/1?include=relation",
			expectedError: false,
			expectedCode:  StatusBadRequest,
		},
	}

	ExecTestRequest(t, tests)
}
