package types

import (
	"context"
	"time"
)

type Post struct {
	Time time.Time
	From string
	Text string
}

type PostRepository interface {
	AddPost(ctx context.Context, post Post) error
	GetPosts(ctx context.Context) ([]Post, error)
}
