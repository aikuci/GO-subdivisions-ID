package test

import "testing"

func TestGetDistrict(t *testing.T) {
	ClearAll()
	CreateProvincesAndItsRelations(1, TotalProvinceRelations{totalCity: 1, TotalCityRelations: TotalCityRelations{totalDistrict: 1}})

	tests := []TestStruct{
		{
			name:          "Successful request: Get district by valid ID",
			route:         "/v1/districts/1",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Not found: Get district by unregistered ID",
			route:         "/v1/districts/0",
			expectedError: false,
			expectedCode:  StatusNotFound,
		},
		{
			name:          "Not found: Invalid ID format in route",
			route:         "/v1/districts/district",
			expectedError: false,
			expectedCode:  StatusNotFound,
		},
	}

	ExecTestRequest(t, tests)
}

func TestGetDistricts(t *testing.T) {
	ClearAll()
	CreateProvincesAndItsRelations(1, TotalProvinceRelations{totalCity: 1, TotalCityRelations: TotalCityRelations{totalDistrict: 20}})

	tests := []TestStruct{
		{
			name:          "Successful request: Get districts",
			route:         "/v1/districts",
			expectedError: false,
			expectedCode:  StatusOK,
		},
	}

	ExecTestRequest(t, tests)
}

func TestGetDistrictsWithItsRelations(t *testing.T) {
	ClearAll()
	CreateProvincesAndItsRelations(1,
		TotalProvinceRelations{totalCity: 1,
			TotalCityRelations: TotalCityRelations{totalDistrict: 30,
				TotalDistrictRelations: TotalDistrictRelations{totalVillage: 1},
			},
		},
	)

	tests := []TestStruct{
		{
			name:          "Successful request: Get districts include its province",
			route:         "/v1/provinces/1/cities/1/districts?include=province",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Successful request: Get districts include its city",
			route:         "/v1/provinces/1/cities/1/districts?include=city",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Successful request: Get districts include its villages",
			route:         "/v1/provinces/1/cities/1/districts?include=villages",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Bad request: Invalid 'include' query parameter",
			route:         "/v1/provinces/1/cities/1/districts?include=relation",
			expectedError: false,
			expectedCode:  StatusBadRequest,
		},

		{
			name:          "Successful request: Get districts include its province, city and villages",
			route:         "/v1/provinces/1/cities/1/districts?include=province,city,villages",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Successful request: Get district by valid ID include its province, city and villages",
			route:         "/v1/provinces/1/cities/1/districts/1?include=province,city,villages",
			expectedError: false,
			expectedCode:  StatusOK,
		},

		{
			name:          "Successful request: Get district by valid ID include its province",
			route:         "/v1/provinces/1/cities/1/districts/1?include=province",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Successful request: Get district by valid ID include its city",
			route:         "/v1/provinces/1/cities/1/districts/1?include=city",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Successful request: Get district by valid ID include its villages",
			route:         "/v1/provinces/1/cities/1/districts/1?include=villages",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Bad request: Invalid 'include' query parameter",
			route:         "/v1/provinces/1/cities/1/districts/1?include=relation",
			expectedError: false,
			expectedCode:  StatusBadRequest,
		},
	}

	ExecTestRequest(t, tests)
}
