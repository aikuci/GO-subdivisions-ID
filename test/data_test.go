package test

import (
	"strconv"
	"testing"

	"github.com/aikuci/go-subdivisions-id/internal/entity"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type TotalProvinceRelations struct {
	totalCity int
	TotalCityRelations
}
type TotalCityRelations struct {
	totalDistrict int
	TotalDistrictRelations
}
type TotalDistrictRelations struct {
	totalVillage int
}

type VillageRelations struct {
	DistrictRelations
	districtId int
}
type DistrictRelations struct {
	CityRelations
	cityId int
}
type CityRelations struct {
	provinceId int
}

func CreateProvincesAndItsRelations(total int, totalRelations TotalProvinceRelations) {
	for i := 1; i < total+1; i++ {
		province := &entity.Province{
			Base:        entity.Base{ID: i},
			Code:        strconv.Itoa(i),
			Name:        "Province " + strconv.Itoa(i),
			PostalCodes: pq.Int64Array{int64(i * 1000)},
		}
		err := db.Create(province).Error
		if err != nil {
			log.Fatal("Failed create province data : %+v", zap.Error(err))
		}

		CreateCities(
			totalRelations.totalCity,
			CityRelations{provinceId: i},
			TotalCityRelations{
				totalDistrict: totalRelations.totalDistrict,
				TotalDistrictRelations: TotalDistrictRelations{
					totalVillage: totalRelations.totalVillage,
				},
			},
		)
	}
}

func CreateCities(total int, relations CityRelations, totalRelations TotalCityRelations) {
	for i := 1; i < total+1; i++ {
		city := &entity.City{
			Base:        entity.Base{ID: i},
			ProvinceID:  relations.provinceId,
			Code:        strconv.Itoa(i),
			Name:        "City " + strconv.Itoa(i),
			PostalCodes: pq.Int64Array{int64(relations.provinceId*1000 + i*100)},
		}
		err := db.Create(city).Error
		if err != nil {
			log.Fatal("Failed create city data : %+v", zap.Error(err))
		}

		CreateDistricts(
			totalRelations.totalDistrict,
			DistrictRelations{cityId: i, CityRelations: relations},
			TotalDistrictRelations{totalVillage: totalRelations.totalVillage},
		)
	}
}

func CreateDistricts(total int, relations DistrictRelations, totalRelations TotalDistrictRelations) {
	for i := 1; i < total+1; i++ {
		district := &entity.District{
			Base:        entity.Base{ID: i},
			ProvinceID:  relations.provinceId,
			CityID:      relations.cityId,
			Code:        strconv.Itoa(i),
			Name:        "District " + strconv.Itoa(i),
			PostalCodes: pq.Int64Array{int64(relations.provinceId*1000 + relations.cityId*100 + i*10)},
		}
		err := db.Create(district).Error
		if err != nil {
			log.Fatal("Failed create district data : %+v", zap.Error(err))
		}

		CreateVillages(totalRelations.totalVillage, VillageRelations{districtId: i, DistrictRelations: relations})
	}
}

func CreateVillages(total int, relations VillageRelations) {
	for i := 1; i < total+1; i++ {
		village := &entity.Village{
			Base:        entity.Base{ID: i},
			ProvinceID:  relations.provinceId,
			CityID:      relations.cityId,
			DistrictID:  relations.districtId,
			Code:        strconv.Itoa(i),
			Name:        "Village " + strconv.Itoa(i),
			PostalCodes: pq.Int64Array{int64(relations.provinceId*1000 + relations.cityId*100 + relations.districtId*10 + i)},
		}
		err := db.Create(village).Error
		if err != nil {
			log.Fatal("Failed create village data : %+v", zap.Error(err))
		}
	}
}

func GetFirstProvince(t *testing.T) *entity.Province {
	province := new(entity.Province)
	err := db.First(province).Error
	assert.Nil(t, err)
	return province
}

func GetFirstCity(t *testing.T) *entity.City {
	city := new(entity.City)
	err := db.First(city).Error
	assert.Nil(t, err)
	return city
}

func GetFirstDistrict(t *testing.T) *entity.District {
	district := new(entity.District)
	err := db.First(district).Error
	assert.Nil(t, err)
	return district
}

func GetFirstVillage(t *testing.T) *entity.Village {
	village := new(entity.Village)
	err := db.First(village).Error
	assert.Nil(t, err)
	return village
}

func ClearAll() {
	ClearProvinces()
	ClearCities()
	ClearDistricts()
	ClearVillages()
}

func ClearProvinces() {
	err := db.Where("id is not null").Delete(&entity.Province{}).Error
	if err != nil {
		log.Fatal("Failed clear province data: %+v", zap.Error(err))
	}
}

func ClearCities() {
	err := db.Where("id is not null").Delete(&entity.City{}).Error
	if err != nil {
		log.Fatal("Failed clear city data: %+v", zap.Error(err))
	}
}

func ClearDistricts() {
	err := db.Where("id is not null").Delete(&entity.District{}).Error
	if err != nil {
		log.Fatal("Failed clear district data: %+v", zap.Error(err))
	}
}

func ClearVillages() {
	err := db.Where("id is not null").Delete(&entity.Village{}).Error
	if err != nil {
		log.Fatal("Failed clear village data: %+v", zap.Error(err))
	}
}
