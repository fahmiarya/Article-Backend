package service

import (
	"errors"
	"strings"

	"github.com/fahmiarya/Article-Backend/internal/dto"
	"github.com/fahmiarya/Article-Backend/internal/model"
	"github.com/fahmiarya/Article-Backend/internal/repository"
	"gorm.io/gorm"
)

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type PostService interface {
	CreatePost(req *dto.CreatePostRequest) (*model.Post, *ValidationError, error)
	GetPosts(limit, offset int, status string) ([]model.Post, error)
	GetPostByID(id uint) (*model.Post, error)
	UpdatePost(id uint, req *dto.UpdatePostRequest) (*model.Post, *ValidationError, error)
	DeletePost(id uint) error
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo: repo}
}

func (s *postService) validateInput(title, content, category, status string) *ValidationError {
	if strings.TrimSpace(title) == "" || len(title) < 20 {
		return &ValidationError{
			Field: "title",
			Error: "title must be at least 20 characters",
		}
	}

	if strings.TrimSpace(content) == "" || len(content) < 200 {
		return &ValidationError{
			Field: "content",
			Error: "content must be at least 200 characters",
		}
	}

	if strings.TrimSpace(category) == "" || len(category) < 3 {
		return &ValidationError{
			Field: "category",
			Error: "category must be at least 3 characters",
		}
	}

	validStatus := map[string]bool{
		"publish": true,
		"draft":   true,
		"trash":  true,
	}

	if !validStatus[status] {
		return &ValidationError{
			Field: "status",
			Error: "status must be one of: publish, draft, thrash",
		}
	}

	return nil
}

func (s *postService) CreatePost(req *dto.CreatePostRequest) (*model.Post, *ValidationError, error) {
	if valErr := s.validateInput(req.Title, req.Content, req.Category, req.Status); valErr != nil {
		return nil, valErr, nil
	}

	post := model.Post{
		Title:    req.Title,
		Content:  req.Content,
		Category: req.Category,
		Status:   req.Status,
	}

	if err := s.repo.Create(&post); err != nil {
		return nil, nil, err
	}

	return &post, nil, nil
}

func (s *postService) GetPosts(limit, offset int, status string) ([]model.Post, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	return s.repo.FindAll(limit, offset, status)
}

func (s *postService) GetPostByID(id uint) (*model.Post, error) {
	return s.repo.FindByID(id)
}

func (s *postService) UpdatePost(id uint, req *dto.UpdatePostRequest) (*model.Post, *ValidationError, error) {
	if valErr := s.validateInput(req.Title, req.Content, req.Category, req.Status); valErr != nil {
		return nil, valErr, nil
	}

	existing, err := s.repo.FindByID(id)
	if err != nil {
		return nil, nil, err
	}

	existing.Title = req.Title
	existing.Content = req.Content
	existing.Category = req.Category
	existing.Status = req.Status

	if err := s.repo.Update(existing); err != nil {
		return nil, nil, err
	}

	return existing, nil, nil
}

func (s *postService) DeletePost(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}

	return s.repo.Delete(id)
}
