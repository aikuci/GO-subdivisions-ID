package repository

import (
	"github.com/aikuci/go-subdivisions-id/internal/model"
	"gorm.io/gorm"
)

type Repository[T any, TId model.IdSingular, TIds model.IdPlural] struct{}

// Retrieve Collections
func (r *Repository[T, TId, TIds]) Find(db *gorm.DB) ([]T, error) {
	var collections []T
	if err := db.Find(&collections).Error; err != nil {
		return nil, err
	}
	return collections, nil
}
func (r *Repository[T, TId, TIds]) FindBy(db *gorm.DB, where map[string]interface{}) ([]T, error) {
	var collections []T
	if err := db.Where(where).Find(&collections).Error; err != nil {
		return nil, err
	}
	return collections, nil
}
func (r *Repository[T, TId, TIds]) FindById(db *gorm.DB, id TId) ([]T, error) {
	return r.FindBy(db, map[string]interface{}{"id": id})
}
func (r *Repository[T, TId, TIds]) FindByIds(db *gorm.DB, ids TIds) ([]T, error) {
	return r.FindBy(db, map[string]interface{}{"id": ids})
}

// Retrieve First Collection
func (r *Repository[T, TId, TIds]) First(db *gorm.DB) (*T, error) {
	var collection T
	if err := db.First(&collection).Error; err != nil {
		return nil, err
	}
	return &collection, nil
}
func (r *Repository[T, TId, TIds]) FirstBy(db *gorm.DB, where map[string]interface{}) (*T, error) {
	var collection T
	if err := db.Where(where).First(&collection).Error; err != nil {
		return nil, err
	}
	return &collection, nil
}
func (r *Repository[T, TId, TIds]) FirstById(db *gorm.DB, id TId) (*T, error) {
	return r.FirstBy(db, map[string]interface{}{"id": id})
}

// Retrieve Meta Collection
func (r *Repository[T, TId, TIds]) CountById(db *gorm.DB, id TId) (int64, error) {
	var total int64
	err := db.Model(new(T)).Where("id = ?", id).Count(&total).Error
	return total, err
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
