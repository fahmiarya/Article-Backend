package repository

import (
	"github.com/fahmiarya/Article-Backend/internal/model"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *model.Post) error
	FindAll(limit, offset int, status string) ([]model.Post, error)
	FindByID(id uint) (*model.Post, error)
	Update(post *model.Post) error
	Delete(id uint) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(post *model.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) FindAll(limit, offset int, status string) ([]model.Post, error) {
	var posts []model.Post
	query := r.db.Limit(limit).Offset(offset)
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err := query.Find(&posts).Error
	if posts == nil {
		posts = []model.Post{}
	}
	return posts, err
}

func (r *postRepository) FindByID(id uint) (*model.Post, error) {
	var post model.Post
	if err := r.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) Update(post *model.Post) error {
	return r.db.Save(post).Error
}

func (r *postRepository) Delete(id uint) error {
	result := r.db.Model(&model.Post{}).Where("id = ?", id).Update("status", "trash")
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
