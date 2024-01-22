package usecase

import (
	"enigmacamp.com/blog-apps/model"
	"enigmacamp.com/blog-apps/repository"
	"enigmacamp.com/blog-apps/shared/shared_model"
)

type BlogUseCase interface {
	CreateNewBlog(payload model.Blog) (model.Blog, error)
	FindAllBlog(page, size int) ([]model.Blog, shared_model.Paging, error)
}

type blogUseCase struct {
	repo repository.BlogRepository
}

func (b *blogUseCase) CreateNewBlog(payload model.Blog) (model.Blog, error) {
	return b.repo.Create(payload)
}

func (b *blogUseCase) FindAllBlog(page int, size int) ([]model.Blog, shared_model.Paging, error) {
	return b.repo.List(page, size)
}

func NewBlogUseCase(repo repository.BlogRepository) BlogUseCase {
	return &blogUseCase{repo: repo}
}
