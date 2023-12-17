package service

import (
	"backendgo/data/request"
	"backendgo/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagId int)
	FindById(tagId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
