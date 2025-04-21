package service

import (
	"context"
	"time"

	"github.com/im-vishalanand/paper-social-backend/grpc/proto/postpb"
)

// PostServiceServerImpl implements the gRPC server
type PostServiceServerImpl struct {
	postpb.UnimplementedPostServiceServer
}

// mockPosts simulates post data for users
var mockPosts = map[string][]*postpb.Post{
	"user1": {
		{Id: "p1", AuthorId: "user1", Content: "Post 1 from user1", Timestamp: time.Now().Add(-10 * time.Minute).Unix()},
		{Id: "p2", AuthorId: "user1", Content: "Post 2 from user1", Timestamp: time.Now().Add(-5 * time.Minute).Unix()},
	},
	"user2": {
		{Id: "p3", AuthorId: "user2", Content: "Post 1 from user2", Timestamp: time.Now().Add(-8 * time.Minute).Unix()},
		{Id: "p4", AuthorId: "user2", Content: "Post 2 from user2", Timestamp: time.Now().Add(-2 * time.Minute).Unix()},
	},
	"user3": {
		{Id: "p5", AuthorId: "user3", Content: "Hello from user3", Timestamp: time.Now().Add(-20 * time.Minute).Unix()},
	},
	// Add more users as needed
}

func (s *PostServiceServerImpl) ListPostsByUser(ctx context.Context, req *postpb.ListPostsRequest) (*postpb.ListPostsResponse, error) {
	posts := mockPosts[req.UserId]
	return &postpb.ListPostsResponse{
		Posts: posts,
	}, nil
}
