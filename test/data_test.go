package test

import (
	"strconv"
	"testing"

	"github.com/aikuci/go-subdivisions-id/internal/entity"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func ClearAll() {
	ClearProvinces()
	ClearCities()
	ClearDistricts()
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

type TotalProvinceRelations struct {
	totalCity     int
	totalDistrict int
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

		CreateCities(totalRelations.totalCity, i, totalRelations.totalDistrict)
	}
}

func CreateCities(total int, provinceId int, totalDistrict int) {
	for i := 1; i < total+1; i++ {
		city := &entity.City{
			Base:        entity.Base{ID: i},
			ProvinceID:  provinceId,
			Code:        strconv.Itoa(i),
			Name:        "City " + strconv.Itoa(i),
			PostalCodes: pq.Int64Array{int64(provinceId*1000 + i*100)},
		}
		err := db.Create(city).Error
		if err != nil {
			log.Fatal("Failed create city data : %+v", zap.Error(err))
		}

		CreateDistricts(totalDistrict, i, provinceId)
	}
}

func CreateDistricts(total int, cityId int, provinceId int) {
	for i := 1; i < total+1; i++ {
		district := &entity.District{
			Base:        entity.Base{ID: i},
			ProvinceID:  provinceId,
			CityID:      cityId,
			Code:        strconv.Itoa(i),
			Name:        "District " + strconv.Itoa(i),
			PostalCodes: pq.Int64Array{int64(provinceId*1000 + cityId*100 + i*10)},
		}
		err := db.Create(district).Error
		if err != nil {
			log.Fatal("Failed create district data : %+v", zap.Error(err))
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
