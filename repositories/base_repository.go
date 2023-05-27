package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type BaseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) BaseRepository {
	return BaseRepository{db: db}
}

func (r *BaseRepository) Get(id string, out interface{}) error {
	fmt.Print(id)
	return r.db.First(out, id).Unscoped().Error
}

func (r *BaseRepository) Create(data interface{}) error {
	return r.db.Create(data).Error
}

func (r *BaseRepository) Update(id string, data interface{}) error {
	return r.db.Model(data).Where("id = ?", id).Updates(data).Error
}

func (r *BaseRepository) Delete(id string, model interface{}) error {
	return r.db.Where("id = ?", id).Delete(model).Error
}
