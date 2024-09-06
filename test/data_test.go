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
}

func ClearProvinces() {
	err := db.Where("id is not null").Delete(&entity.Province{}).Error
	if err != nil {
		log.Fatal("Failed clear province data: %+v", zap.Error(err))
	}
}

func CreateProvinces(total int) {
	for i := 0; i < total; i++ {
		id := i + 1
		province := &entity.Province{
			Base:        entity.Base{ID: id},
			Code:        strconv.Itoa(id),
			Name:        "Province " + strconv.Itoa(id),
			PostalCodes: pq.Int64Array{int64(id * 1000)},
		}
		err := db.Create(province).Error
		if err != nil {
			log.Fatal("Failed create province data : %+v", zap.Error(err))
		}
	}
}

func CreateCities(total int, provinceId int) {
	for i := 0; i < total; i++ {
		id := provinceId*total + i + 1
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
	for i := 0; i < totalProvince; i++ {
		id := i + 1
		province := &entity.Province{
			Base:        entity.Base{ID: id},
			Code:        strconv.Itoa(id),
			Name:        "Province " + strconv.Itoa(id),
			PostalCodes: pq.Int64Array{int64(id * 1000)},
		}
		err := db.Create(province).Error
		if err != nil {
			log.Fatal("Failed create province data : %+v", zap.Error(err))
		}

		CreateCities(totalCity, id)
	}
}

func GetFirstProvince(t *testing.T) *entity.Province {
	province := new(entity.Province)
	err := db.First(province).Error
	assert.Nil(t, err)
	return province
}
