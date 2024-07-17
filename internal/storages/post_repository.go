package storages

import (
	"github.com/RomanGolovanov/go-chat-playground/internal/types"
)

type InMemoryPostRepository struct {
	posts []types.Post
}

func NewInMemoryPostRepository() *InMemoryPostRepository {
	return &InMemoryPostRepository{
		posts: make([]types.Post, 0),
	}
}

func (s *InMemoryPostRepository) AddPost(post types.Post) error {
	s.posts = append(s.posts, post)
	return nil
}

func (s *InMemoryPostRepository) GetPosts() ([]types.Post, error) {
	return s.posts, nil
}
