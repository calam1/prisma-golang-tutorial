package tmp

import (
	"context"

	prisma "github.com/hello-world/generated/prisma-client"
	main "github.com/hello-world/server"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() main.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Post() main.PostResolver {
	return &postResolver{r}
}
func (r *Resolver) Query() main.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() main.UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, name string) (*prisma.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateDraft(ctx context.Context, title string, userID string) (*prisma.Post, error) {
	panic("not implemented")
}
func (r *mutationResolver) Publish(ctx context.Context, postID string) (*prisma.Post, error) {
	panic("not implemented")
}

type postResolver struct{ *Resolver }

func (r *postResolver) Author(ctx context.Context, obj *prisma.Post) (*prisma.User, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) PublishedPosts(ctx context.Context) ([]*prisma.Post, error) {
	panic("not implemented")
}
func (r *queryResolver) Post(ctx context.Context, postID string) (*prisma.Post, error) {
	panic("not implemented")
}
func (r *queryResolver) PostsByUser(ctx context.Context, userID string) ([]*prisma.Post, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) Posts(ctx context.Context, obj *prisma.User) ([]*prisma.Post, error) {
	panic("not implemented")
}
