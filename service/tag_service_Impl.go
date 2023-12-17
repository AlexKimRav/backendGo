package service

import (
	"backendgo/data/request"
	"backendgo/data/response"
	"backendgo/helper"
	"backendgo/model"
	"backendgo/repository"

	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	validate       *validator.Validate
}

// Create implements TagsService.
func (t *TagsServiceImpl) Create(tags request.CreateTagsRequest) {
	err := t.validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)

}

// Delete implements TagsService.
func (t *TagsServiceImpl) Delete(tagId int) {
	t.TagsRepository.Delete(tagId)
}

// FindAll implements TagsService.
func (t *TagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()
	var tags []response.TagsResponse

	for _, value := range result {
		tag := response.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}

	return tags
}

// FindById implements TagsService.
func (t *TagsServiceImpl) FindById(tagId int) response.TagsResponse {
	tagData, err := t.TagsRepository.FindById(tagId)
	helper.ErrorPanic(err)
	tagResponse := response.TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}
	return tagResponse
}

// Update implements TagsService.
func (t *TagsServiceImpl) Update(tags request.UpdateTagsRequest) {
	tagData, err := t.TagsRepository.FindById(tags.Id)
	helper.ErrorPanic(err)
	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)
}

// func NewTagsServiceImpl(tagRepository repository.TagsRepository, validate *validator.Validate) TagsService {
// 	return &TagsServiceImpl{
// 		TagsRepository: tagRepository,
// 		validate:       validate,
// 	}
// }
