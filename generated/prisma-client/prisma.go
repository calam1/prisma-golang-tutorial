// Code generated by Prisma CLI (https://github.com/prisma/prisma). DO NOT EDIT.

package prisma

import (
	"context"
	"errors"

	"github.com/prisma/prisma-client-lib-go"

	"github.com/machinebox/graphql"
)

var ErrNoResult = errors.New("query returned no result")

func Str(v string) *string { return &v }
func Int32(v int32) *int32 { return &v }
func Bool(v bool) *bool    { return &v }

type BatchPayloadExec struct {
	exec *prisma.BatchPayloadExec
}

func (exec *BatchPayloadExec) Exec(ctx context.Context) (BatchPayload, error) {
	bp, err := exec.exec.Exec(ctx)
	return BatchPayload(bp), err
}

type BatchPayload struct {
	Count int64 `json:"count"`
}

type Aggregate struct {
	Count int64 `json:"count"`
}

type Client struct {
	Client *prisma.Client
}

type Options struct {
	Endpoint string
	Secret   string
}

func New(options *Options, opts ...graphql.ClientOption) *Client {
	endpoint := DefaultEndpoint
	secret := Secret
	if options != nil {
		endpoint = options.Endpoint
		secret = options.Secret
	}
	return &Client{
		Client: prisma.New(endpoint, secret, opts...),
	}
}

func (client *Client) GraphQL(ctx context.Context, query string, variables map[string]interface{}) (map[string]interface{}, error) {
	return client.Client.GraphQL(ctx, query, variables)
}

var DefaultEndpoint = "http://localhost:4466"
var Secret = ""

func (client *Client) Post(params PostWhereUniqueInput) *PostExec {
	ret := client.Client.GetOne(
		nil,
		params,
		[2]string{"PostWhereUniqueInput!", "Post"},
		"post",
		[]string{"id", "title", "published"})

	return &PostExec{ret}
}

type PostsParams struct {
	Where   *PostWhereInput   `json:"where,omitempty"`
	OrderBy *PostOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32            `json:"skip,omitempty"`
	After   *string           `json:"after,omitempty"`
	Before  *string           `json:"before,omitempty"`
	First   *int32            `json:"first,omitempty"`
	Last    *int32            `json:"last,omitempty"`
}

func (client *Client) Posts(params *PostsParams) *PostExecArray {
	var wparams *prisma.WhereParams
	if params != nil {
		wparams = &prisma.WhereParams{
			Where:   params.Where,
			OrderBy: (*string)(params.OrderBy),
			Skip:    params.Skip,
			After:   params.After,
			Before:  params.Before,
			First:   params.First,
			Last:    params.Last,
		}
	}

	ret := client.Client.GetMany(
		nil,
		wparams,
		[3]string{"PostWhereInput", "PostOrderByInput", "Post"},
		"posts",
		[]string{"id", "title", "published"})

	return &PostExecArray{ret}
}

type PostsConnectionParams struct {
	Where   *PostWhereInput   `json:"where,omitempty"`
	OrderBy *PostOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32            `json:"skip,omitempty"`
	After   *string           `json:"after,omitempty"`
	Before  *string           `json:"before,omitempty"`
	First   *int32            `json:"first,omitempty"`
	Last    *int32            `json:"last,omitempty"`
}

func (client *Client) PostsConnection(params *PostsConnectionParams) *PostConnectionExec {
	var wparams *prisma.WhereParams
	if params != nil {
		wparams = &prisma.WhereParams{
			Where:   params.Where,
			OrderBy: (*string)(params.OrderBy),
			Skip:    params.Skip,
			After:   params.After,
			Before:  params.Before,
			First:   params.First,
			Last:    params.Last,
		}
	}

	ret := client.Client.GetMany(
		nil,
		wparams,
		[3]string{"PostWhereInput", "PostOrderByInput", "Post"},
		"postsConnection",
		[]string{"edges", "pageInfo"})

	return &PostConnectionExec{ret}
}

func (client *Client) User(params UserWhereUniqueInput) *UserExec {
	ret := client.Client.GetOne(
		nil,
		params,
		[2]string{"UserWhereUniqueInput!", "User"},
		"user",
		[]string{"id", "email", "name"})

	return &UserExec{ret}
}

type UsersParams struct {
	Where   *UserWhereInput   `json:"where,omitempty"`
	OrderBy *UserOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32            `json:"skip,omitempty"`
	After   *string           `json:"after,omitempty"`
	Before  *string           `json:"before,omitempty"`
	First   *int32            `json:"first,omitempty"`
	Last    *int32            `json:"last,omitempty"`
}

func (client *Client) Users(params *UsersParams) *UserExecArray {
	var wparams *prisma.WhereParams
	if params != nil {
		wparams = &prisma.WhereParams{
			Where:   params.Where,
			OrderBy: (*string)(params.OrderBy),
			Skip:    params.Skip,
			After:   params.After,
			Before:  params.Before,
			First:   params.First,
			Last:    params.Last,
		}
	}

	ret := client.Client.GetMany(
		nil,
		wparams,
		[3]string{"UserWhereInput", "UserOrderByInput", "User"},
		"users",
		[]string{"id", "email", "name"})

	return &UserExecArray{ret}
}

type UsersConnectionParams struct {
	Where   *UserWhereInput   `json:"where,omitempty"`
	OrderBy *UserOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32            `json:"skip,omitempty"`
	After   *string           `json:"after,omitempty"`
	Before  *string           `json:"before,omitempty"`
	First   *int32            `json:"first,omitempty"`
	Last    *int32            `json:"last,omitempty"`
}

func (client *Client) UsersConnection(params *UsersConnectionParams) *UserConnectionExec {
	var wparams *prisma.WhereParams
	if params != nil {
		wparams = &prisma.WhereParams{
			Where:   params.Where,
			OrderBy: (*string)(params.OrderBy),
			Skip:    params.Skip,
			After:   params.After,
			Before:  params.Before,
			First:   params.First,
			Last:    params.Last,
		}
	}

	ret := client.Client.GetMany(
		nil,
		wparams,
		[3]string{"UserWhereInput", "UserOrderByInput", "User"},
		"usersConnection",
		[]string{"edges", "pageInfo"})

	return &UserConnectionExec{ret}
}

func (client *Client) CreatePost(params PostCreateInput) *PostExec {
	ret := client.Client.Create(
		params,
		[2]string{"PostCreateInput!", "Post"},
		"createPost",
		[]string{"id", "title", "published"})

	return &PostExec{ret}
}

type PostUpdateParams struct {
	Data  PostUpdateInput      `json:"data"`
	Where PostWhereUniqueInput `json:"where"`
}

func (client *Client) UpdatePost(params PostUpdateParams) *PostExec {
	ret := client.Client.Update(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[3]string{"PostUpdateInput!", "PostWhereUniqueInput!", "Post"},
		"updatePost",
		[]string{"id", "title", "published"})

	return &PostExec{ret}
}

type PostUpdateManyParams struct {
	Data  PostUpdateManyMutationInput `json:"data"`
	Where *PostWhereInput             `json:"where,omitempty"`
}

func (client *Client) UpdateManyPosts(params PostUpdateManyParams) *BatchPayloadExec {
	exec := client.Client.UpdateMany(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[2]string{"PostUpdateManyMutationInput!", "PostWhereInput"},
		"updateManyPosts")
	return &BatchPayloadExec{exec}
}

type PostUpsertParams struct {
	Where  PostWhereUniqueInput `json:"where"`
	Create PostCreateInput      `json:"create"`
	Update PostUpdateInput      `json:"update"`
}

func (client *Client) UpsertPost(params PostUpsertParams) *PostExec {
	uparams := &prisma.UpsertParams{
		Where:  params.Where,
		Create: params.Create,
		Update: params.Update,
	}
	ret := client.Client.Upsert(
		uparams,
		[4]string{"PostWhereUniqueInput!", "PostCreateInput!", "PostUpdateInput!", "Post"},
		"upsertPost",
		[]string{"id", "title", "published"})

	return &PostExec{ret}
}

func (client *Client) DeletePost(params PostWhereUniqueInput) *PostExec {
	ret := client.Client.Delete(
		params,
		[2]string{"PostWhereUniqueInput!", "Post"},
		"deletePost",
		[]string{"id", "title", "published"})

	return &PostExec{ret}
}

func (client *Client) DeleteManyPosts(params *PostWhereInput) *BatchPayloadExec {
	exec := client.Client.DeleteMany(params, "PostWhereInput", "deleteManyPosts")
	return &BatchPayloadExec{exec}
}

func (client *Client) CreateUser(params UserCreateInput) *UserExec {
	ret := client.Client.Create(
		params,
		[2]string{"UserCreateInput!", "User"},
		"createUser",
		[]string{"id", "email", "name"})

	return &UserExec{ret}
}

type UserUpdateParams struct {
	Data  UserUpdateInput      `json:"data"`
	Where UserWhereUniqueInput `json:"where"`
}

func (client *Client) UpdateUser(params UserUpdateParams) *UserExec {
	ret := client.Client.Update(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[3]string{"UserUpdateInput!", "UserWhereUniqueInput!", "User"},
		"updateUser",
		[]string{"id", "email", "name"})

	return &UserExec{ret}
}

type UserUpdateManyParams struct {
	Data  UserUpdateManyMutationInput `json:"data"`
	Where *UserWhereInput             `json:"where,omitempty"`
}

func (client *Client) UpdateManyUsers(params UserUpdateManyParams) *BatchPayloadExec {
	exec := client.Client.UpdateMany(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[2]string{"UserUpdateManyMutationInput!", "UserWhereInput"},
		"updateManyUsers")
	return &BatchPayloadExec{exec}
}

type UserUpsertParams struct {
	Where  UserWhereUniqueInput `json:"where"`
	Create UserCreateInput      `json:"create"`
	Update UserUpdateInput      `json:"update"`
}

func (client *Client) UpsertUser(params UserUpsertParams) *UserExec {
	uparams := &prisma.UpsertParams{
		Where:  params.Where,
		Create: params.Create,
		Update: params.Update,
	}
	ret := client.Client.Upsert(
		uparams,
		[4]string{"UserWhereUniqueInput!", "UserCreateInput!", "UserUpdateInput!", "User"},
		"upsertUser",
		[]string{"id", "email", "name"})

	return &UserExec{ret}
}

func (client *Client) DeleteUser(params UserWhereUniqueInput) *UserExec {
	ret := client.Client.Delete(
		params,
		[2]string{"UserWhereUniqueInput!", "User"},
		"deleteUser",
		[]string{"id", "email", "name"})

	return &UserExec{ret}
}

func (client *Client) DeleteManyUsers(params *UserWhereInput) *BatchPayloadExec {
	exec := client.Client.DeleteMany(params, "UserWhereInput", "deleteManyUsers")
	return &BatchPayloadExec{exec}
}

type PostOrderByInput string

const (
	PostOrderByInputIDAsc         PostOrderByInput = "id_ASC"
	PostOrderByInputIDDesc        PostOrderByInput = "id_DESC"
	PostOrderByInputTitleAsc      PostOrderByInput = "title_ASC"
	PostOrderByInputTitleDesc     PostOrderByInput = "title_DESC"
	PostOrderByInputPublishedAsc  PostOrderByInput = "published_ASC"
	PostOrderByInputPublishedDesc PostOrderByInput = "published_DESC"
)

type UserOrderByInput string

const (
	UserOrderByInputIDAsc     UserOrderByInput = "id_ASC"
	UserOrderByInputIDDesc    UserOrderByInput = "id_DESC"
	UserOrderByInputEmailAsc  UserOrderByInput = "email_ASC"
	UserOrderByInputEmailDesc UserOrderByInput = "email_DESC"
	UserOrderByInputNameAsc   UserOrderByInput = "name_ASC"
	UserOrderByInputNameDesc  UserOrderByInput = "name_DESC"
)

type MutationType string

const (
	MutationTypeCreated MutationType = "CREATED"
	MutationTypeUpdated MutationType = "UPDATED"
	MutationTypeDeleted MutationType = "DELETED"
)

type UserUpdateOneWithoutPostsInput struct {
	Create     *UserCreateWithoutPostsInput     `json:"create,omitempty"`
	Update     *UserUpdateWithoutPostsDataInput `json:"update,omitempty"`
	Upsert     *UserUpsertWithoutPostsInput     `json:"upsert,omitempty"`
	Delete     *bool                            `json:"delete,omitempty"`
	Disconnect *bool                            `json:"disconnect,omitempty"`
	Connect    *UserWhereUniqueInput            `json:"connect,omitempty"`
}

type PostWhereUniqueInput struct {
	ID *string `json:"id,omitempty"`
}

type PostUpdateManyWithoutAuthorInput struct {
	Create     []PostCreateWithoutAuthorInput                `json:"create,omitempty"`
	Delete     []PostWhereUniqueInput                        `json:"delete,omitempty"`
	Connect    []PostWhereUniqueInput                        `json:"connect,omitempty"`
	Set        []PostWhereUniqueInput                        `json:"set,omitempty"`
	Disconnect []PostWhereUniqueInput                        `json:"disconnect,omitempty"`
	Update     []PostUpdateWithWhereUniqueWithoutAuthorInput `json:"update,omitempty"`
	Upsert     []PostUpsertWithWhereUniqueWithoutAuthorInput `json:"upsert,omitempty"`
	DeleteMany []PostScalarWhereInput                        `json:"deleteMany,omitempty"`
	UpdateMany []PostUpdateManyWithWhereNestedInput          `json:"updateMany,omitempty"`
}

type UserCreateInput struct {
	ID    *string                           `json:"id,omitempty"`
	Email *string                           `json:"email,omitempty"`
	Name  string                            `json:"name"`
	Posts *PostCreateManyWithoutAuthorInput `json:"posts,omitempty"`
}

type PostUpdateManyMutationInput struct {
	Title     *string `json:"title,omitempty"`
	Published *bool   `json:"published,omitempty"`
}

type UserSubscriptionWhereInput struct {
	MutationIn                 []MutationType               `json:"mutation_in,omitempty"`
	UpdatedFieldsContains      *string                      `json:"updatedFields_contains,omitempty"`
	UpdatedFieldsContainsEvery []string                     `json:"updatedFields_contains_every,omitempty"`
	UpdatedFieldsContainsSome  []string                     `json:"updatedFields_contains_some,omitempty"`
	Node                       *UserWhereInput              `json:"node,omitempty"`
	And                        []UserSubscriptionWhereInput `json:"AND,omitempty"`
	Or                         []UserSubscriptionWhereInput `json:"OR,omitempty"`
	Not                        []UserSubscriptionWhereInput `json:"NOT,omitempty"`
}

type PostCreateInput struct {
	ID        *string                         `json:"id,omitempty"`
	Title     string                          `json:"title"`
	Published *bool                           `json:"published,omitempty"`
	Author    *UserCreateOneWithoutPostsInput `json:"author,omitempty"`
}

type UserUpdateManyMutationInput struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
}

type UserCreateOneWithoutPostsInput struct {
	Create  *UserCreateWithoutPostsInput `json:"create,omitempty"`
	Connect *UserWhereUniqueInput        `json:"connect,omitempty"`
}

type PostUpdateManyWithWhereNestedInput struct {
	Where PostScalarWhereInput    `json:"where"`
	Data  PostUpdateManyDataInput `json:"data"`
}

type UserCreateWithoutPostsInput struct {
	ID    *string `json:"id,omitempty"`
	Email *string `json:"email,omitempty"`
	Name  string  `json:"name"`
}

type UserWhereUniqueInput struct {
	ID    *string `json:"id,omitempty"`
	Email *string `json:"email,omitempty"`
}

type PostUpdateInput struct {
	Title     *string                         `json:"title,omitempty"`
	Published *bool                           `json:"published,omitempty"`
	Author    *UserUpdateOneWithoutPostsInput `json:"author,omitempty"`
}

type PostUpdateWithoutAuthorDataInput struct {
	Title     *string `json:"title,omitempty"`
	Published *bool   `json:"published,omitempty"`
}

type UserUpdateInput struct {
	Email *string                           `json:"email,omitempty"`
	Name  *string                           `json:"name,omitempty"`
	Posts *PostUpdateManyWithoutAuthorInput `json:"posts,omitempty"`
}

type UserWhereInput struct {
	ID                 *string          `json:"id,omitempty"`
	IDNot              *string          `json:"id_not,omitempty"`
	IDIn               []string         `json:"id_in,omitempty"`
	IDNotIn            []string         `json:"id_not_in,omitempty"`
	IDLt               *string          `json:"id_lt,omitempty"`
	IDLte              *string          `json:"id_lte,omitempty"`
	IDGt               *string          `json:"id_gt,omitempty"`
	IDGte              *string          `json:"id_gte,omitempty"`
	IDContains         *string          `json:"id_contains,omitempty"`
	IDNotContains      *string          `json:"id_not_contains,omitempty"`
	IDStartsWith       *string          `json:"id_starts_with,omitempty"`
	IDNotStartsWith    *string          `json:"id_not_starts_with,omitempty"`
	IDEndsWith         *string          `json:"id_ends_with,omitempty"`
	IDNotEndsWith      *string          `json:"id_not_ends_with,omitempty"`
	Email              *string          `json:"email,omitempty"`
	EmailNot           *string          `json:"email_not,omitempty"`
	EmailIn            []string         `json:"email_in,omitempty"`
	EmailNotIn         []string         `json:"email_not_in,omitempty"`
	EmailLt            *string          `json:"email_lt,omitempty"`
	EmailLte           *string          `json:"email_lte,omitempty"`
	EmailGt            *string          `json:"email_gt,omitempty"`
	EmailGte           *string          `json:"email_gte,omitempty"`
	EmailContains      *string          `json:"email_contains,omitempty"`
	EmailNotContains   *string          `json:"email_not_contains,omitempty"`
	EmailStartsWith    *string          `json:"email_starts_with,omitempty"`
	EmailNotStartsWith *string          `json:"email_not_starts_with,omitempty"`
	EmailEndsWith      *string          `json:"email_ends_with,omitempty"`
	EmailNotEndsWith   *string          `json:"email_not_ends_with,omitempty"`
	Name               *string          `json:"name,omitempty"`
	NameNot            *string          `json:"name_not,omitempty"`
	NameIn             []string         `json:"name_in,omitempty"`
	NameNotIn          []string         `json:"name_not_in,omitempty"`
	NameLt             *string          `json:"name_lt,omitempty"`
	NameLte            *string          `json:"name_lte,omitempty"`
	NameGt             *string          `json:"name_gt,omitempty"`
	NameGte            *string          `json:"name_gte,omitempty"`
	NameContains       *string          `json:"name_contains,omitempty"`
	NameNotContains    *string          `json:"name_not_contains,omitempty"`
	NameStartsWith     *string          `json:"name_starts_with,omitempty"`
	NameNotStartsWith  *string          `json:"name_not_starts_with,omitempty"`
	NameEndsWith       *string          `json:"name_ends_with,omitempty"`
	NameNotEndsWith    *string          `json:"name_not_ends_with,omitempty"`
	PostsEvery         *PostWhereInput  `json:"posts_every,omitempty"`
	PostsSome          *PostWhereInput  `json:"posts_some,omitempty"`
	PostsNone          *PostWhereInput  `json:"posts_none,omitempty"`
	And                []UserWhereInput `json:"AND,omitempty"`
	Or                 []UserWhereInput `json:"OR,omitempty"`
	Not                []UserWhereInput `json:"NOT,omitempty"`
}

type UserUpdateWithoutPostsDataInput struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
}

type PostUpdateManyDataInput struct {
	Title     *string `json:"title,omitempty"`
	Published *bool   `json:"published,omitempty"`
}

type PostCreateManyWithoutAuthorInput struct {
	Create  []PostCreateWithoutAuthorInput `json:"create,omitempty"`
	Connect []PostWhereUniqueInput         `json:"connect,omitempty"`
}

type PostCreateWithoutAuthorInput struct {
	ID        *string `json:"id,omitempty"`
	Title     string  `json:"title"`
	Published *bool   `json:"published,omitempty"`
}

type PostWhereInput struct {
	ID                 *string          `json:"id,omitempty"`
	IDNot              *string          `json:"id_not,omitempty"`
	IDIn               []string         `json:"id_in,omitempty"`
	IDNotIn            []string         `json:"id_not_in,omitempty"`
	IDLt               *string          `json:"id_lt,omitempty"`
	IDLte              *string          `json:"id_lte,omitempty"`
	IDGt               *string          `json:"id_gt,omitempty"`
	IDGte              *string          `json:"id_gte,omitempty"`
	IDContains         *string          `json:"id_contains,omitempty"`
	IDNotContains      *string          `json:"id_not_contains,omitempty"`
	IDStartsWith       *string          `json:"id_starts_with,omitempty"`
	IDNotStartsWith    *string          `json:"id_not_starts_with,omitempty"`
	IDEndsWith         *string          `json:"id_ends_with,omitempty"`
	IDNotEndsWith      *string          `json:"id_not_ends_with,omitempty"`
	Title              *string          `json:"title,omitempty"`
	TitleNot           *string          `json:"title_not,omitempty"`
	TitleIn            []string         `json:"title_in,omitempty"`
	TitleNotIn         []string         `json:"title_not_in,omitempty"`
	TitleLt            *string          `json:"title_lt,omitempty"`
	TitleLte           *string          `json:"title_lte,omitempty"`
	TitleGt            *string          `json:"title_gt,omitempty"`
	TitleGte           *string          `json:"title_gte,omitempty"`
	TitleContains      *string          `json:"title_contains,omitempty"`
	TitleNotContains   *string          `json:"title_not_contains,omitempty"`
	TitleStartsWith    *string          `json:"title_starts_with,omitempty"`
	TitleNotStartsWith *string          `json:"title_not_starts_with,omitempty"`
	TitleEndsWith      *string          `json:"title_ends_with,omitempty"`
	TitleNotEndsWith   *string          `json:"title_not_ends_with,omitempty"`
	Published          *bool            `json:"published,omitempty"`
	PublishedNot       *bool            `json:"published_not,omitempty"`
	Author             *UserWhereInput  `json:"author,omitempty"`
	And                []PostWhereInput `json:"AND,omitempty"`
	Or                 []PostWhereInput `json:"OR,omitempty"`
	Not                []PostWhereInput `json:"NOT,omitempty"`
}

type UserUpsertWithoutPostsInput struct {
	Update UserUpdateWithoutPostsDataInput `json:"update"`
	Create UserCreateWithoutPostsInput     `json:"create"`
}

type PostScalarWhereInput struct {
	ID                 *string                `json:"id,omitempty"`
	IDNot              *string                `json:"id_not,omitempty"`
	IDIn               []string               `json:"id_in,omitempty"`
	IDNotIn            []string               `json:"id_not_in,omitempty"`
	IDLt               *string                `json:"id_lt,omitempty"`
	IDLte              *string                `json:"id_lte,omitempty"`
	IDGt               *string                `json:"id_gt,omitempty"`
	IDGte              *string                `json:"id_gte,omitempty"`
	IDContains         *string                `json:"id_contains,omitempty"`
	IDNotContains      *string                `json:"id_not_contains,omitempty"`
	IDStartsWith       *string                `json:"id_starts_with,omitempty"`
	IDNotStartsWith    *string                `json:"id_not_starts_with,omitempty"`
	IDEndsWith         *string                `json:"id_ends_with,omitempty"`
	IDNotEndsWith      *string                `json:"id_not_ends_with,omitempty"`
	Title              *string                `json:"title,omitempty"`
	TitleNot           *string                `json:"title_not,omitempty"`
	TitleIn            []string               `json:"title_in,omitempty"`
	TitleNotIn         []string               `json:"title_not_in,omitempty"`
	TitleLt            *string                `json:"title_lt,omitempty"`
	TitleLte           *string                `json:"title_lte,omitempty"`
	TitleGt            *string                `json:"title_gt,omitempty"`
	TitleGte           *string                `json:"title_gte,omitempty"`
	TitleContains      *string                `json:"title_contains,omitempty"`
	TitleNotContains   *string                `json:"title_not_contains,omitempty"`
	TitleStartsWith    *string                `json:"title_starts_with,omitempty"`
	TitleNotStartsWith *string                `json:"title_not_starts_with,omitempty"`
	TitleEndsWith      *string                `json:"title_ends_with,omitempty"`
	TitleNotEndsWith   *string                `json:"title_not_ends_with,omitempty"`
	Published          *bool                  `json:"published,omitempty"`
	PublishedNot       *bool                  `json:"published_not,omitempty"`
	And                []PostScalarWhereInput `json:"AND,omitempty"`
	Or                 []PostScalarWhereInput `json:"OR,omitempty"`
	Not                []PostScalarWhereInput `json:"NOT,omitempty"`
}

type PostSubscriptionWhereInput struct {
	MutationIn                 []MutationType               `json:"mutation_in,omitempty"`
	UpdatedFieldsContains      *string                      `json:"updatedFields_contains,omitempty"`
	UpdatedFieldsContainsEvery []string                     `json:"updatedFields_contains_every,omitempty"`
	UpdatedFieldsContainsSome  []string                     `json:"updatedFields_contains_some,omitempty"`
	Node                       *PostWhereInput              `json:"node,omitempty"`
	And                        []PostSubscriptionWhereInput `json:"AND,omitempty"`
	Or                         []PostSubscriptionWhereInput `json:"OR,omitempty"`
	Not                        []PostSubscriptionWhereInput `json:"NOT,omitempty"`
}

type PostUpdateWithWhereUniqueWithoutAuthorInput struct {
	Where PostWhereUniqueInput             `json:"where"`
	Data  PostUpdateWithoutAuthorDataInput `json:"data"`
}

type PostUpsertWithWhereUniqueWithoutAuthorInput struct {
	Where  PostWhereUniqueInput             `json:"where"`
	Update PostUpdateWithoutAuthorDataInput `json:"update"`
	Create PostCreateWithoutAuthorInput     `json:"create"`
}

type UserPreviousValuesExec struct {
	exec *prisma.Exec
}

func (instance UserPreviousValuesExec) Exec(ctx context.Context) (*UserPreviousValues, error) {
	var v UserPreviousValues
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserPreviousValuesExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserPreviousValuesExecArray struct {
	exec *prisma.Exec
}

func (instance UserPreviousValuesExecArray) Exec(ctx context.Context) ([]UserPreviousValues, error) {
	var v []UserPreviousValues
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type UserPreviousValues struct {
	ID    string  `json:"id"`
	Email *string `json:"email,omitempty"`
	Name  string  `json:"name"`
}

type PageInfoExec struct {
	exec *prisma.Exec
}

func (instance PageInfoExec) Exec(ctx context.Context) (*PageInfo, error) {
	var v PageInfo
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance PageInfoExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type PageInfoExecArray struct {
	exec *prisma.Exec
}

func (instance PageInfoExecArray) Exec(ctx context.Context) ([]PageInfo, error) {
	var v []PageInfo
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor,omitempty"`
	EndCursor       *string `json:"endCursor,omitempty"`
}

type UserExec struct {
	exec *prisma.Exec
}

type PostsParamsExec struct {
	Where   *PostWhereInput
	OrderBy *PostOrderByInput
	Skip    *int32
	After   *string
	Before  *string
	First   *int32
	Last    *int32
}

func (instance *UserExec) Posts(params *PostsParamsExec) *PostExecArray {
	var wparams *prisma.WhereParams
	if params != nil {
		wparams = &prisma.WhereParams{
			Where:   params.Where,
			OrderBy: (*string)(params.OrderBy),
			Skip:    params.Skip,
			After:   params.After,
			Before:  params.Before,
			First:   params.First,
			Last:    params.Last,
		}
	}

	ret := instance.exec.Client.GetMany(
		instance.exec,
		wparams,
		[3]string{"PostWhereInput", "PostOrderByInput", "Post"},
		"posts",
		[]string{"id", "title", "published"})

	return &PostExecArray{ret}
}

func (instance UserExec) Exec(ctx context.Context) (*User, error) {
	var v User
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserExecArray struct {
	exec *prisma.Exec
}

func (instance UserExecArray) Exec(ctx context.Context) ([]User, error) {
	var v []User
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type User struct {
	ID    string  `json:"id"`
	Email *string `json:"email,omitempty"`
	Name  string  `json:"name"`
}

type PostEdgeExec struct {
	exec *prisma.Exec
}

func (instance *PostEdgeExec) Node() *PostExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "Post"},
		"node",
		[]string{"id", "title", "published"})

	return &PostExec{ret}
}

func (instance PostEdgeExec) Exec(ctx context.Context) (*PostEdge, error) {
	var v PostEdge
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance PostEdgeExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type PostEdgeExecArray struct {
	exec *prisma.Exec
}

func (instance PostEdgeExecArray) Exec(ctx context.Context) ([]PostEdge, error) {
	var v []PostEdge
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type PostEdge struct {
	Node   Post   `json:"node"`
	Cursor string `json:"cursor"`
}

type PostSubscriptionPayloadExec struct {
	exec *prisma.Exec
}

func (instance *PostSubscriptionPayloadExec) Node() *PostExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "Post"},
		"node",
		[]string{"id", "title", "published"})

	return &PostExec{ret}
}

func (instance *PostSubscriptionPayloadExec) PreviousValues() *PostPreviousValuesExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "PostPreviousValues"},
		"previousValues",
		[]string{"id", "title", "published"})

	return &PostPreviousValuesExec{ret}
}

func (instance PostSubscriptionPayloadExec) Exec(ctx context.Context) (*PostSubscriptionPayload, error) {
	var v PostSubscriptionPayload
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance PostSubscriptionPayloadExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type PostSubscriptionPayloadExecArray struct {
	exec *prisma.Exec
}

func (instance PostSubscriptionPayloadExecArray) Exec(ctx context.Context) ([]PostSubscriptionPayload, error) {
	var v []PostSubscriptionPayload
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type PostSubscriptionPayload struct {
	Mutation      MutationType `json:"mutation"`
	Node          *Post        `json:"node,omitempty"`
	UpdatedFields []string     `json:"updatedFields,omitempty"`
}

type PostConnectionExec struct {
	exec *prisma.Exec
}

func (instance *PostConnectionExec) PageInfo() *PageInfoExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "PageInfo"},
		"pageInfo",
		[]string{"hasNextPage", "hasPreviousPage", "startCursor", "endCursor"})

	return &PageInfoExec{ret}
}

func (instance *PostConnectionExec) Edges() *PostEdgeExecArray {
	edges := instance.exec.Client.GetMany(
		instance.exec,
		nil,
		[3]string{"PostWhereInput", "PostOrderByInput", "PostEdge"},
		"edges",
		[]string{"cursor"})

	nodes := edges.Client.GetMany(
		edges,
		nil,
		[3]string{"", "", "Post"},
		"node",
		[]string{"id", "createdAt", "updatedAt", "name", "desc"})

	return &PostEdgeExecArray{nodes}
}

func (instance *PostConnectionExec) Aggregate(ctx context.Context) (*Aggregate, error) {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "AggregatePost"},
		"aggregate",
		[]string{"count"})

	var v Aggregate
	_, err := ret.Exec(ctx, &v)
	return &v, err
}

func (instance PostConnectionExec) Exec(ctx context.Context) (*PostConnection, error) {
	var v PostConnection
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance PostConnectionExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type PostConnectionExecArray struct {
	exec *prisma.Exec
}

func (instance PostConnectionExecArray) Exec(ctx context.Context) ([]PostConnection, error) {
	var v []PostConnection
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type PostConnection struct {
	PageInfo PageInfo   `json:"pageInfo"`
	Edges    []PostEdge `json:"edges"`
}

type PostExec struct {
	exec *prisma.Exec
}

func (instance *PostExec) Author() *UserExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "User"},
		"author",
		[]string{"id", "email", "name"})

	return &UserExec{ret}
}

func (instance PostExec) Exec(ctx context.Context) (*Post, error) {
	var v Post
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance PostExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type PostExecArray struct {
	exec *prisma.Exec
}

func (instance PostExecArray) Exec(ctx context.Context) ([]Post, error) {
	var v []Post
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type Post struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Published bool   `json:"published"`
}

type PostPreviousValuesExec struct {
	exec *prisma.Exec
}

func (instance PostPreviousValuesExec) Exec(ctx context.Context) (*PostPreviousValues, error) {
	var v PostPreviousValues
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance PostPreviousValuesExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type PostPreviousValuesExecArray struct {
	exec *prisma.Exec
}

func (instance PostPreviousValuesExecArray) Exec(ctx context.Context) ([]PostPreviousValues, error) {
	var v []PostPreviousValues
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type PostPreviousValues struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Published bool   `json:"published"`
}

type UserEdgeExec struct {
	exec *prisma.Exec
}

func (instance *UserEdgeExec) Node() *UserExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "User"},
		"node",
		[]string{"id", "email", "name"})

	return &UserExec{ret}
}

func (instance UserEdgeExec) Exec(ctx context.Context) (*UserEdge, error) {
	var v UserEdge
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserEdgeExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserEdgeExecArray struct {
	exec *prisma.Exec
}

func (instance UserEdgeExecArray) Exec(ctx context.Context) ([]UserEdge, error) {
	var v []UserEdge
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type UserEdge struct {
	Node   User   `json:"node"`
	Cursor string `json:"cursor"`
}

type UserSubscriptionPayloadExec struct {
	exec *prisma.Exec
}

func (instance *UserSubscriptionPayloadExec) Node() *UserExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "User"},
		"node",
		[]string{"id", "email", "name"})

	return &UserExec{ret}
}

func (instance *UserSubscriptionPayloadExec) PreviousValues() *UserPreviousValuesExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "UserPreviousValues"},
		"previousValues",
		[]string{"id", "email", "name"})

	return &UserPreviousValuesExec{ret}
}

func (instance UserSubscriptionPayloadExec) Exec(ctx context.Context) (*UserSubscriptionPayload, error) {
	var v UserSubscriptionPayload
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserSubscriptionPayloadExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserSubscriptionPayloadExecArray struct {
	exec *prisma.Exec
}

func (instance UserSubscriptionPayloadExecArray) Exec(ctx context.Context) ([]UserSubscriptionPayload, error) {
	var v []UserSubscriptionPayload
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type UserSubscriptionPayload struct {
	Mutation      MutationType `json:"mutation"`
	Node          *User        `json:"node,omitempty"`
	UpdatedFields []string     `json:"updatedFields,omitempty"`
}

type UserConnectionExec struct {
	exec *prisma.Exec
}

func (instance *UserConnectionExec) PageInfo() *PageInfoExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "PageInfo"},
		"pageInfo",
		[]string{"hasNextPage", "hasPreviousPage", "startCursor", "endCursor"})

	return &PageInfoExec{ret}
}

func (instance *UserConnectionExec) Edges() *UserEdgeExecArray {
	edges := instance.exec.Client.GetMany(
		instance.exec,
		nil,
		[3]string{"UserWhereInput", "UserOrderByInput", "UserEdge"},
		"edges",
		[]string{"cursor"})

	nodes := edges.Client.GetMany(
		edges,
		nil,
		[3]string{"", "", "User"},
		"node",
		[]string{"id", "createdAt", "updatedAt", "name", "desc"})

	return &UserEdgeExecArray{nodes}
}

func (instance *UserConnectionExec) Aggregate(ctx context.Context) (*Aggregate, error) {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "AggregateUser"},
		"aggregate",
		[]string{"count"})

	var v Aggregate
	_, err := ret.Exec(ctx, &v)
	return &v, err
}

func (instance UserConnectionExec) Exec(ctx context.Context) (*UserConnection, error) {
	var v UserConnection
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserConnectionExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserConnectionExecArray struct {
	exec *prisma.Exec
}

func (instance UserConnectionExecArray) Exec(ctx context.Context) ([]UserConnection, error) {
	var v []UserConnection
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type UserConnection struct {
	PageInfo PageInfo   `json:"pageInfo"`
	Edges    []UserEdge `json:"edges"`
}
