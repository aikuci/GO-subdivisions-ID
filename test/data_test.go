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
		log.Fatal("Failed clear province data: %+v", zap.Error(err))
	}
}

func CreateProvinces(total int) {
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
	}
}

func CreateCities(total int, provinceId int) {
	for i := 1; i < total+1; i++ {
		id := provinceId*total + i
		city := &entity.City{
			Base:        entity.Base{ID: id},
			ProvinceID:  provinceId,
			Code:        strconv.Itoa(id),
			Name:        "City " + strconv.Itoa(id),
			PostalCodes: pq.Int64Array{int64(id * 1000)},
		}
		err := db.Create(city).Error
		if err != nil {
			log.Fatal("Failed create city data : %+v", zap.Error(err))
		}
	}
}

func CreateProvincesAndCities(totalProvince int, totalCity int) {
	for i := 1; i < totalProvince+1; i++ {
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

		CreateCities(totalCity, i)
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
