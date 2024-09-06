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
		contact := &entity.Province{
			Base:        entity.Base{ID: id},
			Code:        strconv.Itoa(id),
			Name:        "Province " + strconv.Itoa(id),
			PostalCodes: pq.Int64Array{int64(id * 1000)},
		}
		err := db.Create(contact).Error
		if err != nil {
			log.Fatal("Failed create contact data : %+v", zap.Error(err))
		}
	}
}

func GetFirstProvince(t *testing.T) *entity.Province {
	province := new(entity.Province)
	err := db.First(province).Error
	assert.Nil(t, err)
	return province
}
