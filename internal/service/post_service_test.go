package service

import (
	"strings"
	"testing"

	"github.com/fahmiarya/Article-Backend/internal/dto"
	"github.com/fahmiarya/Article-Backend/internal/model"
)

type mockRepo struct {
	posts []model.Post
}

func (m *mockRepo) Create(post *model.Post) error {
	post.ID = uint(len(m.posts) + 1)
	m.posts = append(m.posts, *post)
	return nil
}

func (m *mockRepo) FindAll(limit, offset int, status string) ([]model.Post, error) {
	var filtered []model.Post
	for _, p := range m.posts {
		if status == "" || p.Status == status {
			filtered = append(filtered, p)
		}
	}
	if offset >= len(filtered) {
		return []model.Post{}, nil
	}
	end := offset + limit
	if end > len(filtered) {
		end = len(filtered)
	}
	return filtered[offset:end], nil
}

func (m *mockRepo) FindByID(id uint) (*model.Post, error) {
	for _, p := range m.posts {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, nil
}

func (m *mockRepo) Update(post *model.Post) error {
	for i, p := range m.posts {
		if p.ID == post.ID {
			m.posts[i] = *post
			return nil
		}
	}
	return nil
}

func (m *mockRepo) Delete(id uint) error {
	for i, p := range m.posts {
		if p.ID == id {
			m.posts = append(m.posts[:i], m.posts[i+1:]...)
			return nil
		}
	}
	return nil
}

func TestValidationRules(t *testing.T) {
	repo := &mockRepo{}
	svc := NewPostService(repo)

	validContent := strings.Repeat("A", 200)

	// Test 1: Title too short
	req := &dto.CreatePostRequest{
		Title:    "Too short",
		Content:  validContent,
		Category: "Tech",
		Status:   "draft",
	}
	_, valErr, _ := svc.CreatePost(req)
	if valErr == nil || valErr.Field != "title" {
		t.Fatalf("expected title validation error, got %+v", valErr)
	}

	// Test 2: Content too short
	req = &dto.CreatePostRequest{
		Title:    "Valid Title With Enough Characters",
		Content:  "Short content",
		Category: "Tech",
		Status:   "draft",
	}
	_, valErr, _ = svc.CreatePost(req)
	if valErr == nil || valErr.Field != "content" {
		t.Fatalf("expected content validation error, got %+v", valErr)
	}

	// Test 3: Category too short
	req = &dto.CreatePostRequest{
		Title:    "Valid Title With Enough Characters",
		Content:  validContent,
		Category: "Te",
		Status:   "draft",
	}
	_, valErr, _ = svc.CreatePost(req)
	if valErr == nil || valErr.Field != "category" {
		t.Fatalf("expected category validation error, got %+v", valErr)
	}

	// Test 4: Invalid status (e.g., trash instead of thrash)
	req = &dto.CreatePostRequest{
		Title:    "Valid Title With Enough Characters",
		Content:  validContent,
		Category: "Tech",
		Status:   "trash",
	}
	_, valErr, _ = svc.CreatePost(req)
	if valErr == nil || valErr.Field != "status" {
		t.Fatalf("expected status validation error for trash, got %+v", valErr)
	}

	// Test 5: Valid status 'thrash'
	req = &dto.CreatePostRequest{
		Title:    "Valid Title With Enough Characters",
		Content:  validContent,
		Category: "Tech",
		Status:   "thrash",
	}
	post, valErr, err := svc.CreatePost(req)
	if valErr != nil || err != nil || post == nil {
		t.Fatalf("expected valid creation with thrash, got valErr: %+v, err: %v", valErr, err)
	}
}
