package repository

import (
	"backendgo/helper"
	"backendgo/model"

	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements TagsRepository.
func (t *TagsRepositoryImpl) Delete(tagId int) {
	var tags model.Tags
	result := t.Db.Where("id = ?", tagId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

// FindAll implements TagsRepository.
func (*TagsRepositoryImpl) FindAll() []model.Tags {
	panic("unimplemented")
}

// FindById implements TagsRepository.
func (*TagsRepositoryImpl) FindById(tagId int) (tags model.Tags, err error) {
	panic("unimplemented")
}

// Save implements TagsRepository.
func (*TagsRepositoryImpl) Save(tags model.Tags) {
	panic("unimplemented")
}

// Update implements TagsRepository.
func (*TagsRepositoryImpl) Update(tags model.Tags) {
	panic("unimplemented")
}
