package main

import (
	"context"

	prisma "github.com/hello-world/generated/prisma-client"
	//main "github.com/hello-world/server"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

//type Resolver struct{}

type Resolver struct {
	Prisma *prisma.Client
}

// Mutation mutation
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

//Post post
func (r *Resolver) Post() PostResolver {
	return &postResolver{r}
}

//Query query
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

//User user
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, name string) (*prisma.User, error) {
	return r.Prisma.CreateUser(prisma.UserCreateInput{
		Name: name,
	}).Exec(ctx)
}
func (r *mutationResolver) CreateDraft(ctx context.Context, title string, userID string) (*prisma.Post, error) {
	return r.Prisma.CreatePost(prisma.PostCreateInput{
		Title: title,
		Author: &prisma.UserCreateOneWithoutPostsInput{
			Connect: &prisma.UserWhereUniqueInput{ID: &userID},
		},
	}).Exec(ctx)
}
func (r *mutationResolver) Publish(ctx context.Context, postID string) (*prisma.Post, error) {
	published := true
	return r.Prisma.UpdatePost(prisma.PostUpdateParams{
		Where: prisma.PostWhereUniqueInput{ID: &postID},
		Data:  prisma.PostUpdateInput{Published: &published},
	}).Exec(ctx)
}

type postResolver struct{ *Resolver }

func (r *postResolver) Author(ctx context.Context, obj *prisma.Post) (*prisma.User, error) {
	return r.Prisma.Post(prisma.PostWhereUniqueInput{ID: &obj.ID}).Author().Exec(ctx)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) PublishedPosts(ctx context.Context) ([]*prisma.Post, error) {
	published := true
	vals, err := r.Prisma.Posts(&prisma.PostsParams{
		Where: &prisma.PostWhereInput{Published: &published},
	}).Exec(ctx)

	valPointers := make([]*prisma.Post, 0)
	for _, v := range vals {
		valPointers = append(valPointers, &v)
	}

	return valPointers, err
}
func (r *queryResolver) Post(ctx context.Context, postID string) (*prisma.Post, error) {
	return r.Prisma.Post(prisma.PostWhereUniqueInput{ID: &postID}).Exec(ctx)
}
func (r *queryResolver) PostsByUser(ctx context.Context, userID string) ([]prisma.Post, error) {
	return r.Prisma.Posts(&prisma.PostsParams{
		Where: &prisma.PostWhereInput{
			Author: &prisma.UserWhereInput{
				ID: &userID,
			},
		},
	}).Exec(ctx)
}

type userResolver struct{ *Resolver }

func (r *userResolver) Posts(ctx context.Context, obj *prisma.User) ([]prisma.Post, error) {
	return r.Prisma.User(prisma.UserWhereUniqueInput{ID: &obj.ID}).Posts(nil).Exec(ctx)

}
