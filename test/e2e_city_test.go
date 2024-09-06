package test

import "testing"

func TestGetCity(t *testing.T) {
	ClearAll()
	CreateProvincesAndCities(1, 1)

	tests := []TestStruct{
		{
			name:          "Successful request: Get city by valid ID",
			route:         "/v1/cities/1",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Not found: Get city by unregistered ID",
			route:         "/v1/cities/0",
			expectedError: false,
			expectedCode:  StatusNotFound,
		},
		{
			name:          "Not found: Invalid ID format in route",
			route:         "/v1/cities/city",
			expectedError: false,
			expectedCode:  StatusNotFound,
		},
	}

	ExecTestRequest(t, tests)
}

func TestGetCities(t *testing.T) {
	ClearAll()
	CreateProvincesAndCities(1, 20)

	tests := []TestStruct{
		{
			name:          "Successful request: Get cities",
			route:         "/v1/cities",
			expectedError: false,
			expectedCode:  StatusOK,
		},
	}

	ExecTestRequest(t, tests)
}

func TestGetCitiesWithItsProvince(t *testing.T) {
	ClearAll()
	CreateProvincesAndCities(1, 30)

	tests := []TestStruct{
		{
			name:          "Successful request: Get cities include its province",
			route:         "/v1/provinces/1/cities?include=province",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Bad request: Invalid 'include' query parameter",
			route:         "/v1/provinces/1/cities?include=relation",
			expectedError: false,
			expectedCode:  StatusBadRequest,
		},

		{
			name:          "Successful request: Get province by valid ID include its province",
			route:         "/v1/provinces/1/cities/1?include=province",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Bad request: Invalid 'include' query parameter",
			route:         "/v1/provinces/1/cities/1?include=relation",
			expectedError: false,
			expectedCode:  StatusBadRequest,
		},
	}

	ExecTestRequest(t, tests)
}
