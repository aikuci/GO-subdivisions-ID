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
