package repository

import (
	"github.com/aikuci/go-subdivisions-id/pkg/model"

	"gorm.io/gorm"
)

type Repository[T any, TId model.IdSingular, TIds model.IdPlural] struct{}

// TODO:
// Refer to the GORM documentation for advanced query examples: https://gorm.io/docs/advanced_query.html#Find-To-Map
// Issues:
// 1. Known bug with unsupported data type `&[]`, affecting `pq.Int64Array`.
// 2. Known issue with unsupported data type `&[]` for `[]CityResponse` when processing `provinceResponse`.

// Retrieve First Collection
func (r *Repository[T, TId, TIds]) First(db *gorm.DB) (T, error) {
	var collection T
	err := db.First(&collection).Error
	return collection, err
}
func (r *Repository[T, TId, TIds]) FirstBy(db *gorm.DB, where map[string]interface{}) (T, error) {
	var collection T
	err := db.Where(where).First(&collection).Error
	return collection, err
}
func (r *Repository[T, TId, TIds]) FirstById(db *gorm.DB, id TId) (T, error) {
	return r.FirstBy(db, map[string]interface{}{"id": id})
}

// Retrieve Collections
func (r *Repository[T, TId, TIds]) Find(db *gorm.DB) ([]T, error) {
	var collections []T
	err := db.Find(&collections).Error
	return collections, err
}
func (r *Repository[T, TId, TIds]) FindBy(db *gorm.DB, where map[string]interface{}) ([]T, error) {
	var collections []T
	err := db.Where(where).Find(&collections).Error
	return collections, err
}
func (r *Repository[T, TId, TIds]) FindById(db *gorm.DB, id TId) ([]T, error) {
	return r.FindBy(db, map[string]interface{}{"id": id})
}
func (r *Repository[T, TId, TIds]) FindByIds(db *gorm.DB, ids TIds) ([]T, error) {
	return r.FindBy(db, map[string]interface{}{"id": ids})
}

// Retrieve Collections and Count
func (r *Repository[T, TId, TIds]) FindAndCount(db *gorm.DB) ([]T, int64, error) {
	var collections []T
	var total int64
	err := db.Find(&collections).Count(&total).Error
	return collections, total, err
}
func (r *Repository[T, TId, TIds]) FindAndCountBy(db *gorm.DB, where map[string]interface{}) ([]T, int64, error) {
	var collections []T
	var total int64
	err := db.Where(where).Find(&collections).Count(&total).Error
	return collections, total, err
}
func (r *Repository[T, TId, TIds]) FindAndCountById(db *gorm.DB, id TId) ([]T, int64, error) {
	return r.FindAndCountBy(db, map[string]interface{}{"id": id})
}
func (r *Repository[T, TId, TIds]) FindAndCountByIds(db *gorm.DB, ids TIds) ([]T, int64, error) {
	return r.FindAndCountBy(db, map[string]interface{}{"id": ids})
}

// Collection Action
func (r *Repository[T, TId, TIds]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}
func (r *Repository[T, TId, TIds]) Update(db *gorm.DB, entity *T) error {
	return db.Save(entity).Error
}
func (r *Repository[T, TId, TIds]) Delete(db *gorm.DB, entity *T) error {
	return db.Delete(entity).Error
}
