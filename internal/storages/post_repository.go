package storages

import (
	"context"
	"sync"

	"github.com/RomanGolovanov/go-chat-playground/internal/types"
)

type InMemoryPostRepository struct {
	posts []types.Post
	mu    sync.Mutex
}

func NewInMemoryPostRepository() *InMemoryPostRepository {
	return &InMemoryPostRepository{
		posts: make([]types.Post, 0),
	}
}

func (s *InMemoryPostRepository) AddPost(ctx context.Context, post types.Post) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.posts = append(s.posts, post)
	return nil
}

func (s *InMemoryPostRepository) GetPosts(ctx context.Context) ([]types.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.posts, nil
}
