package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Katsushi21/travelone"
	"github.com/Katsushi21/travelone/ent"
)

func (r *accountResolver) ID(ctx context.Context, obj *ent.Account) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentResolver) ID(ctx context.Context, obj *ent.Comment) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentResolver) AccountID(ctx context.Context, obj *ent.Comment) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentResolver) PostID(ctx context.Context, obj *ent.Comment) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *friendResolver) ID(ctx context.Context, obj *ent.Friend) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *friendResolver) AccountID(ctx context.Context, obj *ent.Friend) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *friendResolver) FriendID(ctx context.Context, obj *ent.Friend) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeResolver) ID(ctx context.Context, obj *ent.Like) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeResolver) AccountID(ctx context.Context, obj *ent.Like) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeResolver) PostID(ctx context.Context, obj *ent.Like) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *markerResolver) ID(ctx context.Context, obj *ent.Marker) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *markerResolver) PostID(ctx context.Context, obj *ent.Marker) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *muteResolver) ID(ctx context.Context, obj *ent.Mute) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *muteResolver) AccountID(ctx context.Context, obj *ent.Mute) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *muteResolver) MuteID(ctx context.Context, obj *ent.Mute) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postResolver) ID(ctx context.Context, obj *ent.Post) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postResolver) AccountID(ctx context.Context, obj *ent.Post) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Accounts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AccountOrder, where *ent.AccountWhereInput) (*ent.AccountConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Comments(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.CommentOrder, where *ent.CommentWhereInput) (*ent.CommentConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Friends(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.FriendOrder, where *ent.FriendWhereInput) (*ent.FriendConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Likes(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.LikeOrder, where *ent.LikeWhereInput) (*ent.LikeConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Markers(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.MarkerOrder, where *ent.MarkerWhereInput) (*ent.MarkerConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Mutes(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.MuteOrder, where *ent.MuteWhereInput) (*ent.MuteConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Posts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.PostOrder, where *ent.PostWhereInput) (*ent.PostConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Requests(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.RequestOrder, where *ent.RequestWhereInput) (*ent.RequestConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Sessions(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.SessionOrder, where *ent.SessionWhereInput) (*ent.SessionConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *requestResolver) ID(ctx context.Context, obj *ent.Request) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *requestResolver) AccountID(ctx context.Context, obj *ent.Request) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *requestResolver) RequestID(ctx context.Context, obj *ent.Request) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *sessionResolver) ID(ctx context.Context, obj *ent.Session) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *sessionResolver) AccountID(ctx context.Context, obj *ent.Session) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountWhereInputResolver) ID(ctx context.Context, obj *ent.AccountWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountWhereInputResolver) IDNeq(ctx context.Context, obj *ent.AccountWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountWhereInputResolver) IDIn(ctx context.Context, obj *ent.AccountWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.AccountWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountWhereInputResolver) IDGt(ctx context.Context, obj *ent.AccountWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountWhereInputResolver) IDGte(ctx context.Context, obj *ent.AccountWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountWhereInputResolver) IDLt(ctx context.Context, obj *ent.AccountWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountWhereInputResolver) IDLte(ctx context.Context, obj *ent.AccountWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentWhereInputResolver) ID(ctx context.Context, obj *ent.CommentWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentWhereInputResolver) IDNeq(ctx context.Context, obj *ent.CommentWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentWhereInputResolver) IDIn(ctx context.Context, obj *ent.CommentWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.CommentWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentWhereInputResolver) IDGt(ctx context.Context, obj *ent.CommentWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentWhereInputResolver) IDGte(ctx context.Context, obj *ent.CommentWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentWhereInputResolver) IDLt(ctx context.Context, obj *ent.CommentWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentWhereInputResolver) IDLte(ctx context.Context, obj *ent.CommentWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentWhereInputResolver) AccountID(ctx context.Context, obj *ent.CommentWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentWhereInputResolver) AccountIDNeq(ctx context.Context, obj *ent.CommentWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentWhereInputResolver) AccountIDIn(ctx context.Context, obj *ent.CommentWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentWhereInputResolver) AccountIDNotIn(ctx context.Context, obj *ent.CommentWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentWhereInputResolver) PostID(ctx context.Context, obj *ent.CommentWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentWhereInputResolver) PostIDNeq(ctx context.Context, obj *ent.CommentWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentWhereInputResolver) PostIDIn(ctx context.Context, obj *ent.CommentWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentWhereInputResolver) PostIDNotIn(ctx context.Context, obj *ent.CommentWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createAccountInputResolver) PostIDs(ctx context.Context, obj *ent.CreateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createAccountInputResolver) CommentIDs(ctx context.Context, obj *ent.CreateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createAccountInputResolver) FriendIDs(ctx context.Context, obj *ent.CreateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createAccountInputResolver) MuteIDs(ctx context.Context, obj *ent.CreateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createAccountInputResolver) RequestIDs(ctx context.Context, obj *ent.CreateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createAccountInputResolver) LikeIDs(ctx context.Context, obj *ent.CreateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createAccountInputResolver) SessionIDs(ctx context.Context, obj *ent.CreateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createCommentInputResolver) PostID(ctx context.Context, obj *ent.CreateCommentInput, data int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createCommentInputResolver) AccountID(ctx context.Context, obj *ent.CreateCommentInput, data int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createFriendInputResolver) AccountID(ctx context.Context, obj *ent.CreateFriendInput, data int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createFriendInputResolver) FriendID(ctx context.Context, obj *ent.CreateFriendInput, data int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createLikeInputResolver) AccountID(ctx context.Context, obj *ent.CreateLikeInput, data int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createLikeInputResolver) PostID(ctx context.Context, obj *ent.CreateLikeInput, data int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createMarkerInputResolver) PostID(ctx context.Context, obj *ent.CreateMarkerInput, data int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createMuteInputResolver) AccountID(ctx context.Context, obj *ent.CreateMuteInput, data int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createMuteInputResolver) MuteID(ctx context.Context, obj *ent.CreateMuteInput, data int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createPostInputResolver) CommentIDs(ctx context.Context, obj *ent.CreatePostInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createPostInputResolver) MarkerID(ctx context.Context, obj *ent.CreatePostInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createPostInputResolver) AccountID(ctx context.Context, obj *ent.CreatePostInput, data int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createRequestInputResolver) AccountID(ctx context.Context, obj *ent.CreateRequestInput, data int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createRequestInputResolver) RequestID(ctx context.Context, obj *ent.CreateRequestInput, data int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createSessionInputResolver) AccountID(ctx context.Context, obj *ent.CreateSessionInput, data int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *friendWhereInputResolver) ID(ctx context.Context, obj *ent.FriendWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *friendWhereInputResolver) IDNeq(ctx context.Context, obj *ent.FriendWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *friendWhereInputResolver) IDIn(ctx context.Context, obj *ent.FriendWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *friendWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.FriendWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *friendWhereInputResolver) IDGt(ctx context.Context, obj *ent.FriendWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *friendWhereInputResolver) IDGte(ctx context.Context, obj *ent.FriendWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *friendWhereInputResolver) IDLt(ctx context.Context, obj *ent.FriendWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *friendWhereInputResolver) IDLte(ctx context.Context, obj *ent.FriendWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeWhereInputResolver) ID(ctx context.Context, obj *ent.LikeWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeWhereInputResolver) IDNeq(ctx context.Context, obj *ent.LikeWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeWhereInputResolver) IDIn(ctx context.Context, obj *ent.LikeWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.LikeWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeWhereInputResolver) IDGt(ctx context.Context, obj *ent.LikeWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeWhereInputResolver) IDGte(ctx context.Context, obj *ent.LikeWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeWhereInputResolver) IDLt(ctx context.Context, obj *ent.LikeWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeWhereInputResolver) IDLte(ctx context.Context, obj *ent.LikeWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeWhereInputResolver) AccountID(ctx context.Context, obj *ent.LikeWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeWhereInputResolver) AccountIDNeq(ctx context.Context, obj *ent.LikeWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeWhereInputResolver) AccountIDIn(ctx context.Context, obj *ent.LikeWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeWhereInputResolver) AccountIDNotIn(ctx context.Context, obj *ent.LikeWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeWhereInputResolver) PostID(ctx context.Context, obj *ent.LikeWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeWhereInputResolver) PostIDNeq(ctx context.Context, obj *ent.LikeWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeWhereInputResolver) PostIDIn(ctx context.Context, obj *ent.LikeWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeWhereInputResolver) PostIDNotIn(ctx context.Context, obj *ent.LikeWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *markerWhereInputResolver) ID(ctx context.Context, obj *ent.MarkerWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *markerWhereInputResolver) IDNeq(ctx context.Context, obj *ent.MarkerWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *markerWhereInputResolver) IDIn(ctx context.Context, obj *ent.MarkerWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *markerWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.MarkerWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *markerWhereInputResolver) IDGt(ctx context.Context, obj *ent.MarkerWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *markerWhereInputResolver) IDGte(ctx context.Context, obj *ent.MarkerWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *markerWhereInputResolver) IDLt(ctx context.Context, obj *ent.MarkerWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *markerWhereInputResolver) IDLte(ctx context.Context, obj *ent.MarkerWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *markerWhereInputResolver) PostID(ctx context.Context, obj *ent.MarkerWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *markerWhereInputResolver) PostIDNeq(ctx context.Context, obj *ent.MarkerWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *markerWhereInputResolver) PostIDIn(ctx context.Context, obj *ent.MarkerWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *markerWhereInputResolver) PostIDNotIn(ctx context.Context, obj *ent.MarkerWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *muteWhereInputResolver) ID(ctx context.Context, obj *ent.MuteWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *muteWhereInputResolver) IDNeq(ctx context.Context, obj *ent.MuteWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *muteWhereInputResolver) IDIn(ctx context.Context, obj *ent.MuteWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *muteWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.MuteWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *muteWhereInputResolver) IDGt(ctx context.Context, obj *ent.MuteWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *muteWhereInputResolver) IDGte(ctx context.Context, obj *ent.MuteWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *muteWhereInputResolver) IDLt(ctx context.Context, obj *ent.MuteWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *muteWhereInputResolver) IDLte(ctx context.Context, obj *ent.MuteWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *muteWhereInputResolver) AccountID(ctx context.Context, obj *ent.MuteWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *muteWhereInputResolver) AccountIDNeq(ctx context.Context, obj *ent.MuteWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *muteWhereInputResolver) AccountIDIn(ctx context.Context, obj *ent.MuteWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *muteWhereInputResolver) AccountIDNotIn(ctx context.Context, obj *ent.MuteWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *muteWhereInputResolver) MuteID(ctx context.Context, obj *ent.MuteWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *muteWhereInputResolver) MuteIDNeq(ctx context.Context, obj *ent.MuteWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *muteWhereInputResolver) MuteIDIn(ctx context.Context, obj *ent.MuteWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *muteWhereInputResolver) MuteIDNotIn(ctx context.Context, obj *ent.MuteWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *postWhereInputResolver) ID(ctx context.Context, obj *ent.PostWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *postWhereInputResolver) IDNeq(ctx context.Context, obj *ent.PostWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *postWhereInputResolver) IDIn(ctx context.Context, obj *ent.PostWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *postWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.PostWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *postWhereInputResolver) IDGt(ctx context.Context, obj *ent.PostWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *postWhereInputResolver) IDGte(ctx context.Context, obj *ent.PostWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *postWhereInputResolver) IDLt(ctx context.Context, obj *ent.PostWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *postWhereInputResolver) IDLte(ctx context.Context, obj *ent.PostWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *postWhereInputResolver) AccountID(ctx context.Context, obj *ent.PostWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *postWhereInputResolver) AccountIDNeq(ctx context.Context, obj *ent.PostWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *postWhereInputResolver) AccountIDIn(ctx context.Context, obj *ent.PostWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *postWhereInputResolver) AccountIDNotIn(ctx context.Context, obj *ent.PostWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *requestWhereInputResolver) ID(ctx context.Context, obj *ent.RequestWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *requestWhereInputResolver) IDNeq(ctx context.Context, obj *ent.RequestWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *requestWhereInputResolver) IDIn(ctx context.Context, obj *ent.RequestWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *requestWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.RequestWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *requestWhereInputResolver) IDGt(ctx context.Context, obj *ent.RequestWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *requestWhereInputResolver) IDGte(ctx context.Context, obj *ent.RequestWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *requestWhereInputResolver) IDLt(ctx context.Context, obj *ent.RequestWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *requestWhereInputResolver) IDLte(ctx context.Context, obj *ent.RequestWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *sessionWhereInputResolver) ID(ctx context.Context, obj *ent.SessionWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *sessionWhereInputResolver) IDNeq(ctx context.Context, obj *ent.SessionWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *sessionWhereInputResolver) IDIn(ctx context.Context, obj *ent.SessionWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *sessionWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.SessionWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *sessionWhereInputResolver) IDGt(ctx context.Context, obj *ent.SessionWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *sessionWhereInputResolver) IDGte(ctx context.Context, obj *ent.SessionWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *sessionWhereInputResolver) IDLt(ctx context.Context, obj *ent.SessionWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *sessionWhereInputResolver) IDLte(ctx context.Context, obj *ent.SessionWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *sessionWhereInputResolver) AccountID(ctx context.Context, obj *ent.SessionWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *sessionWhereInputResolver) AccountIDNeq(ctx context.Context, obj *ent.SessionWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *sessionWhereInputResolver) AccountIDIn(ctx context.Context, obj *ent.SessionWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *sessionWhereInputResolver) AccountIDNotIn(ctx context.Context, obj *ent.SessionWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateAccountInputResolver) AddPostIDs(ctx context.Context, obj *ent.UpdateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateAccountInputResolver) RemovePostIDs(ctx context.Context, obj *ent.UpdateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateAccountInputResolver) AddCommentIDs(ctx context.Context, obj *ent.UpdateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateAccountInputResolver) RemoveCommentIDs(ctx context.Context, obj *ent.UpdateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateAccountInputResolver) AddFriendIDs(ctx context.Context, obj *ent.UpdateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateAccountInputResolver) RemoveFriendIDs(ctx context.Context, obj *ent.UpdateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateAccountInputResolver) AddMuteIDs(ctx context.Context, obj *ent.UpdateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateAccountInputResolver) RemoveMuteIDs(ctx context.Context, obj *ent.UpdateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateAccountInputResolver) AddRequestIDs(ctx context.Context, obj *ent.UpdateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateAccountInputResolver) RemoveRequestIDs(ctx context.Context, obj *ent.UpdateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateAccountInputResolver) AddLikeIDs(ctx context.Context, obj *ent.UpdateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateAccountInputResolver) RemoveLikeIDs(ctx context.Context, obj *ent.UpdateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateAccountInputResolver) AddSessionIDs(ctx context.Context, obj *ent.UpdateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateAccountInputResolver) RemoveSessionIDs(ctx context.Context, obj *ent.UpdateAccountInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateCommentInputResolver) PostID(ctx context.Context, obj *ent.UpdateCommentInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateCommentInputResolver) AccountID(ctx context.Context, obj *ent.UpdateCommentInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateMarkerInputResolver) PostID(ctx context.Context, obj *ent.UpdateMarkerInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updatePostInputResolver) AddCommentIDs(ctx context.Context, obj *ent.UpdatePostInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updatePostInputResolver) RemoveCommentIDs(ctx context.Context, obj *ent.UpdatePostInput, data []int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updatePostInputResolver) MarkerID(ctx context.Context, obj *ent.UpdatePostInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updatePostInputResolver) AccountID(ctx context.Context, obj *ent.UpdatePostInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateRequestInputResolver) AccountID(ctx context.Context, obj *ent.UpdateRequestInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateRequestInputResolver) RequestID(ctx context.Context, obj *ent.UpdateRequestInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateSessionInputResolver) AccountID(ctx context.Context, obj *ent.UpdateSessionInput, data *int) error {
	panic(fmt.Errorf("not implemented"))
}

// Account returns travelone.AccountResolver implementation.
func (r *Resolver) Account() travelone.AccountResolver { return &accountResolver{r} }

// Comment returns travelone.CommentResolver implementation.
func (r *Resolver) Comment() travelone.CommentResolver { return &commentResolver{r} }

// Friend returns travelone.FriendResolver implementation.
func (r *Resolver) Friend() travelone.FriendResolver { return &friendResolver{r} }

// Like returns travelone.LikeResolver implementation.
func (r *Resolver) Like() travelone.LikeResolver { return &likeResolver{r} }

// Marker returns travelone.MarkerResolver implementation.
func (r *Resolver) Marker() travelone.MarkerResolver { return &markerResolver{r} }

// Mute returns travelone.MuteResolver implementation.
func (r *Resolver) Mute() travelone.MuteResolver { return &muteResolver{r} }

// Post returns travelone.PostResolver implementation.
func (r *Resolver) Post() travelone.PostResolver { return &postResolver{r} }

// Query returns travelone.QueryResolver implementation.
func (r *Resolver) Query() travelone.QueryResolver { return &queryResolver{r} }

// Request returns travelone.RequestResolver implementation.
func (r *Resolver) Request() travelone.RequestResolver { return &requestResolver{r} }

// Session returns travelone.SessionResolver implementation.
func (r *Resolver) Session() travelone.SessionResolver { return &sessionResolver{r} }

// AccountWhereInput returns travelone.AccountWhereInputResolver implementation.
func (r *Resolver) AccountWhereInput() travelone.AccountWhereInputResolver {
	return &accountWhereInputResolver{r}
}

// CommentWhereInput returns travelone.CommentWhereInputResolver implementation.
func (r *Resolver) CommentWhereInput() travelone.CommentWhereInputResolver {
	return &commentWhereInputResolver{r}
}

// CreateAccountInput returns travelone.CreateAccountInputResolver implementation.
func (r *Resolver) CreateAccountInput() travelone.CreateAccountInputResolver {
	return &createAccountInputResolver{r}
}

// CreateCommentInput returns travelone.CreateCommentInputResolver implementation.
func (r *Resolver) CreateCommentInput() travelone.CreateCommentInputResolver {
	return &createCommentInputResolver{r}
}

// CreateFriendInput returns travelone.CreateFriendInputResolver implementation.
func (r *Resolver) CreateFriendInput() travelone.CreateFriendInputResolver {
	return &createFriendInputResolver{r}
}

// CreateLikeInput returns travelone.CreateLikeInputResolver implementation.
func (r *Resolver) CreateLikeInput() travelone.CreateLikeInputResolver {
	return &createLikeInputResolver{r}
}

// CreateMarkerInput returns travelone.CreateMarkerInputResolver implementation.
func (r *Resolver) CreateMarkerInput() travelone.CreateMarkerInputResolver {
	return &createMarkerInputResolver{r}
}

// CreateMuteInput returns travelone.CreateMuteInputResolver implementation.
func (r *Resolver) CreateMuteInput() travelone.CreateMuteInputResolver {
	return &createMuteInputResolver{r}
}

// CreatePostInput returns travelone.CreatePostInputResolver implementation.
func (r *Resolver) CreatePostInput() travelone.CreatePostInputResolver {
	return &createPostInputResolver{r}
}

// CreateRequestInput returns travelone.CreateRequestInputResolver implementation.
func (r *Resolver) CreateRequestInput() travelone.CreateRequestInputResolver {
	return &createRequestInputResolver{r}
}

// CreateSessionInput returns travelone.CreateSessionInputResolver implementation.
func (r *Resolver) CreateSessionInput() travelone.CreateSessionInputResolver {
	return &createSessionInputResolver{r}
}

// FriendWhereInput returns travelone.FriendWhereInputResolver implementation.
func (r *Resolver) FriendWhereInput() travelone.FriendWhereInputResolver {
	return &friendWhereInputResolver{r}
}

// LikeWhereInput returns travelone.LikeWhereInputResolver implementation.
func (r *Resolver) LikeWhereInput() travelone.LikeWhereInputResolver {
	return &likeWhereInputResolver{r}
}

// MarkerWhereInput returns travelone.MarkerWhereInputResolver implementation.
func (r *Resolver) MarkerWhereInput() travelone.MarkerWhereInputResolver {
	return &markerWhereInputResolver{r}
}

// MuteWhereInput returns travelone.MuteWhereInputResolver implementation.
func (r *Resolver) MuteWhereInput() travelone.MuteWhereInputResolver {
	return &muteWhereInputResolver{r}
}

// PostWhereInput returns travelone.PostWhereInputResolver implementation.
func (r *Resolver) PostWhereInput() travelone.PostWhereInputResolver {
	return &postWhereInputResolver{r}
}

// RequestWhereInput returns travelone.RequestWhereInputResolver implementation.
func (r *Resolver) RequestWhereInput() travelone.RequestWhereInputResolver {
	return &requestWhereInputResolver{r}
}

// SessionWhereInput returns travelone.SessionWhereInputResolver implementation.
func (r *Resolver) SessionWhereInput() travelone.SessionWhereInputResolver {
	return &sessionWhereInputResolver{r}
}

// UpdateAccountInput returns travelone.UpdateAccountInputResolver implementation.
func (r *Resolver) UpdateAccountInput() travelone.UpdateAccountInputResolver {
	return &updateAccountInputResolver{r}
}

// UpdateCommentInput returns travelone.UpdateCommentInputResolver implementation.
func (r *Resolver) UpdateCommentInput() travelone.UpdateCommentInputResolver {
	return &updateCommentInputResolver{r}
}

// UpdateMarkerInput returns travelone.UpdateMarkerInputResolver implementation.
func (r *Resolver) UpdateMarkerInput() travelone.UpdateMarkerInputResolver {
	return &updateMarkerInputResolver{r}
}

// UpdatePostInput returns travelone.UpdatePostInputResolver implementation.
func (r *Resolver) UpdatePostInput() travelone.UpdatePostInputResolver {
	return &updatePostInputResolver{r}
}

// UpdateRequestInput returns travelone.UpdateRequestInputResolver implementation.
func (r *Resolver) UpdateRequestInput() travelone.UpdateRequestInputResolver {
	return &updateRequestInputResolver{r}
}

// UpdateSessionInput returns travelone.UpdateSessionInputResolver implementation.
func (r *Resolver) UpdateSessionInput() travelone.UpdateSessionInputResolver {
	return &updateSessionInputResolver{r}
}

type accountResolver struct{ *Resolver }
type commentResolver struct{ *Resolver }
type friendResolver struct{ *Resolver }
type likeResolver struct{ *Resolver }
type markerResolver struct{ *Resolver }
type muteResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type requestResolver struct{ *Resolver }
type sessionResolver struct{ *Resolver }
type accountWhereInputResolver struct{ *Resolver }
type commentWhereInputResolver struct{ *Resolver }
type createAccountInputResolver struct{ *Resolver }
type createCommentInputResolver struct{ *Resolver }
type createFriendInputResolver struct{ *Resolver }
type createLikeInputResolver struct{ *Resolver }
type createMarkerInputResolver struct{ *Resolver }
type createMuteInputResolver struct{ *Resolver }
type createPostInputResolver struct{ *Resolver }
type createRequestInputResolver struct{ *Resolver }
type createSessionInputResolver struct{ *Resolver }
type friendWhereInputResolver struct{ *Resolver }
type likeWhereInputResolver struct{ *Resolver }
type markerWhereInputResolver struct{ *Resolver }
type muteWhereInputResolver struct{ *Resolver }
type postWhereInputResolver struct{ *Resolver }
type requestWhereInputResolver struct{ *Resolver }
type sessionWhereInputResolver struct{ *Resolver }
type updateAccountInputResolver struct{ *Resolver }
type updateCommentInputResolver struct{ *Resolver }
type updateMarkerInputResolver struct{ *Resolver }
type updatePostInputResolver struct{ *Resolver }
type updateRequestInputResolver struct{ *Resolver }
type updateSessionInputResolver struct{ *Resolver }
