package test

import "testing"

func TestGetVillage(t *testing.T) {
	ClearAll()
	CreateProvincesAndItsRelations(1,
		TotalProvinceRelations{totalCity: 1,
			TotalCityRelations: TotalCityRelations{totalDistrict: 1,
				TotalDistrictRelations: TotalDistrictRelations{totalVillage: 1},
			},
		},
	)

	tests := []TestStruct{
		{
			name:          "Successful request: Get village by valid ID",
			route:         "/v1/villages/1",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Not found: Get village by unregistered ID",
			route:         "/v1/villages/0",
			expectedError: false,
			expectedCode:  StatusNotFound,
		},
		{
			name:          "Not found: Invalid ID format in route",
			route:         "/v1/villages/village",
			expectedError: false,
			expectedCode:  StatusNotFound,
		},
	}

	ExecTestRequest(t, tests)
}

func TestGetVillages(t *testing.T) {
	ClearAll()
	CreateProvincesAndItsRelations(1,
		TotalProvinceRelations{totalCity: 1,
			TotalCityRelations: TotalCityRelations{totalDistrict: 1,
				TotalDistrictRelations: TotalDistrictRelations{totalVillage: 20},
			},
		},
	)

	tests := []TestStruct{
		{
			name:          "Successful request: Get villages",
			route:         "/v1/villages",
			expectedError: false,
			expectedCode:  StatusOK,
		},
	}

	ExecTestRequest(t, tests)
}
func TestGetVillagesWithItsRelations(t *testing.T) {
	ClearAll()
	CreateProvincesAndItsRelations(1,
		TotalProvinceRelations{totalCity: 1,
			TotalCityRelations: TotalCityRelations{totalDistrict: 1,
				TotalDistrictRelations: TotalDistrictRelations{totalVillage: 30},
			},
		},
	)

	tests := []TestStruct{
		{
			name:          "Successful request: Get villages include its province",
			route:         "/v1/provinces/1/cities/1/districts/1/villages?include=province",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Successful request: Get villages include its city",
			route:         "/v1/provinces/1/cities/1/districts/1/villages?include=city",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Successful request: Get villages include its district",
			route:         "/v1/provinces/1/cities/1/districts/1/villages?include=district",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Bad request: Invalid 'include' query parameter",
			route:         "/v1/provinces/1/cities/1/districts/1/villages?include=relation",
			expectedError: false,
			expectedCode:  StatusBadRequest,
		},

		{
			name:          "Successful request: Get villages include its province, city and district",
			route:         "/v1/provinces/1/cities/1/districts/1/villages?include=province,city,district",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Successful request: Get district by valid ID include its province, city and district",
			route:         "/v1/provinces/1/cities/1/districts/1/villages/1?include=province,city,district",
			expectedError: false,
			expectedCode:  StatusOK,
		},

		{
			name:          "Successful request: Get district by valid ID include its province",
			route:         "/v1/provinces/1/cities/1/districts/1/villages/1?include=province",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Successful request: Get district by valid ID include its city",
			route:         "/v1/provinces/1/cities/1/districts/1/villages/1?include=city",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Successful request: Get district by valid ID include its district",
			route:         "/v1/provinces/1/cities/1/districts/1/villages/1?include=district",
			expectedError: false,
			expectedCode:  StatusOK,
		},
		{
			name:          "Bad request: Invalid 'include' query parameter",
			route:         "/v1/provinces/1/cities/1/districts/1/villages/1?include=relation",
			expectedError: false,
			expectedCode:  StatusBadRequest,
		},
	}

	ExecTestRequest(t, tests)
}
