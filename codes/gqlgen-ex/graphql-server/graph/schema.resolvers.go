package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math/rand"
	"strconv"

	"github.com/redstone-wb/gqlgen-todos/graph/generated"
	"github.com/redstone-wb/gqlgen-todos/graph/model"
	"github.com/redstone-wb/gqlgen-todos/repository"
)

var videoRepo repository.VideoRepository = repository.New()

func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	video := &model.Video{
		ID:     strconv.Itoa(rand.Intn(100)),
		Title:  input.Title,
		URL:    input.URL,
		Author: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	videoRepo.Save(video)
	return video, nil
}

func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	return videoRepo.FindAll(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }