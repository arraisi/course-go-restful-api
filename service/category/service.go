package category

import (
	"context"
	"database/sql"
)

type Service interface {
	Save(ctx context.Context, request CategoryCreateRequest) (tx *sql.Tx, category Category, err error)
	Update(ctx context.Context, request CategoryUpdateRequest) (tx *sql.Tx, category Category, err error)
	Delete(ctx context.Context, id int) (tx *sql.Tx, err error)
	FindById(ctx context.Context, id int) (category Category, err error)
	FindAll(ctx context.Context) (categories []Category, err error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s service) Save(ctx context.Context, request CategoryCreateRequest) (tx *sql.Tx, category Category, err error) {
	category = Category{
		Name: request.Name,
	}
	tx, category, err = s.repository.Save(ctx, category)
	if err != nil {
		return tx, category, err
	}
	return tx, category, err
}

func (s service) Update(ctx context.Context, request CategoryUpdateRequest) (tx *sql.Tx, category Category, err error) {
	category = Category{
		Id:   request.Id,
		Name: request.Name,
	}
	tx, category, err = s.repository.Update(ctx, category)
	if err != nil {
		return tx, category, err
	}
	return tx, category, err
}

func (s service) Delete(ctx context.Context, id int) (tx *sql.Tx, err error) {
	tx, err = s.repository.Delete(ctx, id)
	if err != nil {
		return tx, err
	}
	return tx, err
}

func (s service) FindById(ctx context.Context, id int) (category Category, err error) {
	category, err = s.repository.FindById(ctx, id)
	if err != nil {
		return category, err
	}
	return category, err
}

func (s service) FindAll(ctx context.Context) (categories []Category, err error) {
	categories, err = s.repository.FindAll(ctx)
	if err != nil {
		return categories, err
	}

	for _, category := range categories {
		categories = append(categories, category)
	}
	return categories, err
}
