package main

import "gorm.io/gorm"

type RomanRepository[T any] interface {
	FindAll() ([]T, error)
	FindByID(id uint) (*T, error)
	Save(entity *T) error
	DeleteByID(id uint) error
}

type romanRepository[T any] struct {
	db *gorm.DB
}

func NewRomanRepository[T any](db *gorm.DB) RomanRepository[T] {
	return &romanRepository[T]{db: db}
}

func (r *romanRepository[T]) FindAll() ([]T, error) {
	var entities []T
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *romanRepository[T]) FindByID(id uint) (*T, error) {
	var entity T
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *romanRepository[T]) Save(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *romanRepository[T]) DeleteByID(id uint) error {
	return r.db.Delete(new(T), id).Error
}
