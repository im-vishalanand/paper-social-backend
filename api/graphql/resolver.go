package graphql

import (
    "context"
    "github.com/im-vishalanand/paper-social-backend/grpc/proto/postpb"
    "github.com/im-vishalanand/paper-social-backend/grpc/service"
    "github.com/99designs/gqlgen/graphql"
)

type Resolver struct{}

func (r *Resolver) GetTimeline(ctx context.Context, userId string) ([]*Post, error) {
    // Call the gRPC service to fetch posts for the user
    client := service.NewPostServiceClient("localhost:50051")
    response, err := client.ListPostsByUser(ctx, &postpb.ListPostsRequest{UserId: userId})
    if err != nil {
        return nil, err
    }

    // Convert the gRPC posts to GraphQL posts
    var posts []*Post
    for _, p := range response.Posts {
        posts = append(posts, &Post{
            ID:        p.GetId(),
            AuthorID:  p.GetAuthorId(),
            Content:   p.GetContent(),
            Timestamp: int(p.GetTimestamp()),
        })
    }

    return posts, nil
}
