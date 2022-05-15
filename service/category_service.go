package service

import (
	"context"
	"course-golang-restful-api/model/request"
	"course-golang-restful-api/model/response"
)

type CategoryService interface {
	Create(ctx context.Context, request request.CategoryCreateRequest) response.CategoryResponse
	Update(ctx context.Context, updateRequest request.CategoryUpdateRequest) response.CategoryResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) response.CategoryResponse
	FindAll(ctx context.Context) []response.CategoryResponse
}
