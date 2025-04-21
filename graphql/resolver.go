package graphql

import (
    "context"
    "log"

    "github.com/im-vishalanand/paper-social-backend/grpc/proto/postpb"
    "github.com/im-vishalanand/paper-social-backend/graphql/model"
    "google.golang.org/grpc"
)

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
    return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetTimeline(ctx context.Context, userID string) ([]*model.Post, error) {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("could not connect to gRPC server: %v", err)
    }
    defer conn.Close()

    client := postpb.NewPostServiceClient(conn)

    // Hardcoded for now – fetch posts for 3 followed users
    followedUsers := []string{"user2", "user3", "user4"}
    var allPosts []*model.Post

    for _, uid := range followedUsers {
        res, err := client.ListPostsByUser(ctx, &postpb.ListPostsRequest{UserId: uid})
        if err != nil {
            return nil, err
        }

        for _, p := range res.Posts {
            allPosts = append(allPosts, &model.Post{
                ID:        p.Id,
                AuthorID:  p.AuthorId,
                Content:   p.Content,
                Timestamp: int(p.Timestamp),
            })
        }
    }

    // Optionally: sort by timestamp DESC here

    return allPosts, nil
}
