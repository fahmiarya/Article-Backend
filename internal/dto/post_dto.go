package dto

type CreatePostRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

type UpdatePostRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}
