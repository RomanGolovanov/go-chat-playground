package services

import (
	"time"

	"github.com/RomanGolovanov/go-chat-playground/internal/types"
)

type PostService struct {
	repository types.PostRepository
}

func NewPostService(repository types.PostRepository) *PostService {
	return &PostService{repository: repository}
}

type CreatePostRequest struct {
	From string
	Text string
}

type PostResponse struct {
	From string
	Text string
}

func (s *PostService) AddPost(r CreatePostRequest) error {
	p := types.Post{
		Time: time.Now(),
		From: r.From,
		Text: r.Text,
	}
	_ = p
	return nil
}

func (s *PostService) GetPosts() ([]PostResponse, error) {
	return nil, nil
}