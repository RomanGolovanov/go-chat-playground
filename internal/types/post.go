package types

import "time"

type Post struct {
	Time time.Time
	From string
	Text string
}

type PostRepository interface {
	AddPost(post Post) error
	GetPosts() ([]Post, error)
}
