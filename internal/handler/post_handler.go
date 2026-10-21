package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/fahmiarya/Article-Backend/internal/dto"
	"github.com/fahmiarya/Article-Backend/internal/response"
	"github.com/fahmiarya/Article-Backend/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"gorm.io/gorm"
)

type PostHandler struct {
	postService service.PostService
}

func NewPostHandler(postService service.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var req dto.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
        // Log the binding error for debugging
        log.Printf("Failed to bind JSON: %v", err)

		response.Error(c, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	post, valErr, err := h.postService.CreatePost(&req)
	if valErr != nil {
		response.JSON(c, http.StatusBadRequest, "error", "Validation failed", valErr)
		return
	}
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create article: "+err.Error(), nil)
		return
	}

	response.Success(c, http.StatusCreated, "Article created successfully", post)
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	limitStr := c.Param("limit")
	if limitStr == "" {
		limitStr = c.Param("id")
	}
	offsetStr := c.Param("offset")
	status := c.Query("status") // optional: publish | draft | thrash

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		response.Error(c, http.StatusBadRequest, "limit must be a positive integer", nil)
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		response.Error(c, http.StatusBadRequest, "offset must be a non-negative integer", nil)
		return
	}

	posts, err := h.postService.GetPosts(limit, offset, status)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to retrieve articles: "+err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "Articles retrieved successfully", posts)
}

func (h *PostHandler) GetPostByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid article ID", nil)
		return
	}

	post, err := h.postService.GetPostByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Article not found", nil)
			return
		}
		response.Error(c, http.StatusInternalServerError, "Failed to retrieve article: "+err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "Article retrieved successfully", post)
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid article ID", nil)
		return
	}

	var req dto.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
        // Log the binding error for debugging
        log.Printf("Failed to bind JSON: %v", err)

		response.Error(c, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	post, valErr, err := h.postService.UpdatePost(uint(id), &req)
	if valErr != nil {
		response.JSON(c, http.StatusBadRequest, "error", "Validation failed", valErr)
		return
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Article not found", nil)
			return
		}
		response.Error(c, http.StatusInternalServerError, "Failed to update article: "+err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "Article updated successfully", post)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid article ID", nil)
		return
	}

	if err := h.postService.DeletePost(uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Article not found", nil)
			return
		}
		response.Error(c, http.StatusInternalServerError, "Failed to delete article: "+err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "Article deleted successfully", nil)
}
