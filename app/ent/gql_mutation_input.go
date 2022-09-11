// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Katsushi21/travelone/ent/account"
	"github.com/Katsushi21/travelone/ent/request"
	"github.com/google/uuid"
)

// CreateAccountInput represents a mutation input for creating accounts.
type CreateAccountInput struct {
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	Email        string
	Password     string
	Type         account.Type
	Name         *string
	Age          int
	Gender       account.Gender
	Avatar       string
	Introduction string
	PostIDs      []uuid.UUID
	CommentIDs   []uuid.UUID
	FriendIDs    []uuid.UUID
	MuteIDs      []uuid.UUID
	RequestIDs   []uuid.UUID
	LikeIDs      []uuid.UUID
	SessionIDs   []uuid.UUID
}

// Mutate applies the CreateAccountInput on the AccountMutation builder.
func (i *CreateAccountInput) Mutate(m *AccountMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetEmail(i.Email)
	m.SetPassword(i.Password)
	m.SetType(i.Type)
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	m.SetAge(i.Age)
	m.SetGender(i.Gender)
	m.SetAvatar(i.Avatar)
	m.SetIntroduction(i.Introduction)
	if v := i.PostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
	if v := i.CommentIDs; len(v) > 0 {
		m.AddCommentIDs(v...)
	}
	if v := i.FriendIDs; len(v) > 0 {
		m.AddFriendIDs(v...)
	}
	if v := i.MuteIDs; len(v) > 0 {
		m.AddMuteIDs(v...)
	}
	if v := i.RequestIDs; len(v) > 0 {
		m.AddRequestIDs(v...)
	}
	if v := i.LikeIDs; len(v) > 0 {
		m.AddLikeIDs(v...)
	}
	if v := i.SessionIDs; len(v) > 0 {
		m.AddSessionIDs(v...)
	}
}

// SetInput applies the change-set in the CreateAccountInput on the AccountCreate builder.
func (c *AccountCreate) SetInput(i CreateAccountInput) *AccountCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateAccountInput represents a mutation input for updating accounts.
type UpdateAccountInput struct {
	UpdatedAt        *time.Time
	Email            *string
	Password         *string
	Type             *account.Type
	Name             *string
	Age              *int
	Gender           *account.Gender
	Avatar           *string
	Introduction     *string
	AddPostIDs       []uuid.UUID
	RemovePostIDs    []uuid.UUID
	AddCommentIDs    []uuid.UUID
	RemoveCommentIDs []uuid.UUID
	AddFriendIDs     []uuid.UUID
	RemoveFriendIDs  []uuid.UUID
	AddMuteIDs       []uuid.UUID
	RemoveMuteIDs    []uuid.UUID
	AddRequestIDs    []uuid.UUID
	RemoveRequestIDs []uuid.UUID
	AddLikeIDs       []uuid.UUID
	RemoveLikeIDs    []uuid.UUID
	AddSessionIDs    []uuid.UUID
	RemoveSessionIDs []uuid.UUID
}

// Mutate applies the UpdateAccountInput on the AccountMutation builder.
func (i *UpdateAccountInput) Mutate(m *AccountMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Email; v != nil {
		m.SetEmail(*v)
	}
	if v := i.Password; v != nil {
		m.SetPassword(*v)
	}
	if v := i.Type; v != nil {
		m.SetType(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Age; v != nil {
		m.SetAge(*v)
	}
	if v := i.Gender; v != nil {
		m.SetGender(*v)
	}
	if v := i.Avatar; v != nil {
		m.SetAvatar(*v)
	}
	if v := i.Introduction; v != nil {
		m.SetIntroduction(*v)
	}
	if v := i.AddPostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
	if v := i.RemovePostIDs; len(v) > 0 {
		m.RemovePostIDs(v...)
	}
	if v := i.AddCommentIDs; len(v) > 0 {
		m.AddCommentIDs(v...)
	}
	if v := i.RemoveCommentIDs; len(v) > 0 {
		m.RemoveCommentIDs(v...)
	}
	if v := i.AddFriendIDs; len(v) > 0 {
		m.AddFriendIDs(v...)
	}
	if v := i.RemoveFriendIDs; len(v) > 0 {
		m.RemoveFriendIDs(v...)
	}
	if v := i.AddMuteIDs; len(v) > 0 {
		m.AddMuteIDs(v...)
	}
	if v := i.RemoveMuteIDs; len(v) > 0 {
		m.RemoveMuteIDs(v...)
	}
	if v := i.AddRequestIDs; len(v) > 0 {
		m.AddRequestIDs(v...)
	}
	if v := i.RemoveRequestIDs; len(v) > 0 {
		m.RemoveRequestIDs(v...)
	}
	if v := i.AddLikeIDs; len(v) > 0 {
		m.AddLikeIDs(v...)
	}
	if v := i.RemoveLikeIDs; len(v) > 0 {
		m.RemoveLikeIDs(v...)
	}
	if v := i.AddSessionIDs; len(v) > 0 {
		m.AddSessionIDs(v...)
	}
	if v := i.RemoveSessionIDs; len(v) > 0 {
		m.RemoveSessionIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateAccountInput on the AccountUpdate builder.
func (c *AccountUpdate) SetInput(i UpdateAccountInput) *AccountUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateAccountInput on the AccountUpdateOne builder.
func (c *AccountUpdateOne) SetInput(i UpdateAccountInput) *AccountUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateCommentInput represents a mutation input for creating comments.
type CreateCommentInput struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Body      string
	PostID    uuid.UUID
	AccountID uuid.UUID
}

// Mutate applies the CreateCommentInput on the CommentMutation builder.
func (i *CreateCommentInput) Mutate(m *CommentMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetBody(i.Body)
	m.SetPostID(i.PostID)
	m.SetAccountID(i.AccountID)
}

// SetInput applies the change-set in the CreateCommentInput on the CommentCreate builder.
func (c *CommentCreate) SetInput(i CreateCommentInput) *CommentCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateCommentInput represents a mutation input for updating comments.
type UpdateCommentInput struct {
	UpdatedAt    *time.Time
	Body         *string
	ClearPost    bool
	PostID       *uuid.UUID
	ClearAccount bool
	AccountID    *uuid.UUID
}

// Mutate applies the UpdateCommentInput on the CommentMutation builder.
func (i *UpdateCommentInput) Mutate(m *CommentMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Body; v != nil {
		m.SetBody(*v)
	}
	if i.ClearPost {
		m.ClearPost()
	}
	if v := i.PostID; v != nil {
		m.SetPostID(*v)
	}
	if i.ClearAccount {
		m.ClearAccount()
	}
	if v := i.AccountID; v != nil {
		m.SetAccountID(*v)
	}
}

// SetInput applies the change-set in the UpdateCommentInput on the CommentUpdate builder.
func (c *CommentUpdate) SetInput(i UpdateCommentInput) *CommentUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateCommentInput on the CommentUpdateOne builder.
func (c *CommentUpdateOne) SetInput(i UpdateCommentInput) *CommentUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateFriendInput represents a mutation input for creating friends.
type CreateFriendInput struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	AccountID uuid.UUID
	FriendID  uuid.UUID
}

// Mutate applies the CreateFriendInput on the FriendMutation builder.
func (i *CreateFriendInput) Mutate(m *FriendMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetAccountID(i.AccountID)
	m.SetFriendID(i.FriendID)
}

// SetInput applies the change-set in the CreateFriendInput on the FriendCreate builder.
func (c *FriendCreate) SetInput(i CreateFriendInput) *FriendCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateLikeInput represents a mutation input for creating likes.
type CreateLikeInput struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	AccountID uuid.UUID
	PostID    uuid.UUID
}

// Mutate applies the CreateLikeInput on the LikeMutation builder.
func (i *CreateLikeInput) Mutate(m *LikeMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetAccountID(i.AccountID)
	m.SetPostID(i.PostID)
}

// SetInput applies the change-set in the CreateLikeInput on the LikeCreate builder.
func (c *LikeCreate) SetInput(i CreateLikeInput) *LikeCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateMarkerInput represents a mutation input for creating markers.
type CreateMarkerInput struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Title     string
	Lat       string
	Lng       string
	PostID    uuid.UUID
}

// Mutate applies the CreateMarkerInput on the MarkerMutation builder.
func (i *CreateMarkerInput) Mutate(m *MarkerMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetTitle(i.Title)
	m.SetLat(i.Lat)
	m.SetLng(i.Lng)
	m.SetPostID(i.PostID)
}

// SetInput applies the change-set in the CreateMarkerInput on the MarkerCreate builder.
func (c *MarkerCreate) SetInput(i CreateMarkerInput) *MarkerCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateMarkerInput represents a mutation input for updating markers.
type UpdateMarkerInput struct {
	UpdatedAt *time.Time
	Title     *string
	Lat       *string
	Lng       *string
	ClearPost bool
	PostID    *uuid.UUID
}

// Mutate applies the UpdateMarkerInput on the MarkerMutation builder.
func (i *UpdateMarkerInput) Mutate(m *MarkerMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Title; v != nil {
		m.SetTitle(*v)
	}
	if v := i.Lat; v != nil {
		m.SetLat(*v)
	}
	if v := i.Lng; v != nil {
		m.SetLng(*v)
	}
	if i.ClearPost {
		m.ClearPost()
	}
	if v := i.PostID; v != nil {
		m.SetPostID(*v)
	}
}

// SetInput applies the change-set in the UpdateMarkerInput on the MarkerUpdate builder.
func (c *MarkerUpdate) SetInput(i UpdateMarkerInput) *MarkerUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateMarkerInput on the MarkerUpdateOne builder.
func (c *MarkerUpdateOne) SetInput(i UpdateMarkerInput) *MarkerUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateMuteInput represents a mutation input for creating mutes.
type CreateMuteInput struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	AccountID uuid.UUID
	MuteID    uuid.UUID
}

// Mutate applies the CreateMuteInput on the MuteMutation builder.
func (i *CreateMuteInput) Mutate(m *MuteMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetAccountID(i.AccountID)
	m.SetMuteID(i.MuteID)
}

// SetInput applies the change-set in the CreateMuteInput on the MuteCreate builder.
func (c *MuteCreate) SetInput(i CreateMuteInput) *MuteCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreatePostInput represents a mutation input for creating posts.
type CreatePostInput struct {
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	Title      string
	Body       string
	Img        string
	CommentIDs []uuid.UUID
	MarkerID   *uuid.UUID
	AccountID  uuid.UUID
}

// Mutate applies the CreatePostInput on the PostMutation builder.
func (i *CreatePostInput) Mutate(m *PostMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetTitle(i.Title)
	m.SetBody(i.Body)
	m.SetImg(i.Img)
	if v := i.CommentIDs; len(v) > 0 {
		m.AddCommentIDs(v...)
	}
	if v := i.MarkerID; v != nil {
		m.SetMarkerID(*v)
	}
	m.SetAccountID(i.AccountID)
}

// SetInput applies the change-set in the CreatePostInput on the PostCreate builder.
func (c *PostCreate) SetInput(i CreatePostInput) *PostCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdatePostInput represents a mutation input for updating posts.
type UpdatePostInput struct {
	UpdatedAt        *time.Time
	Title            *string
	Body             *string
	Img              *string
	AddCommentIDs    []uuid.UUID
	RemoveCommentIDs []uuid.UUID
	ClearMarker      bool
	MarkerID         *uuid.UUID
	ClearAccount     bool
	AccountID        *uuid.UUID
}

// Mutate applies the UpdatePostInput on the PostMutation builder.
func (i *UpdatePostInput) Mutate(m *PostMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Title; v != nil {
		m.SetTitle(*v)
	}
	if v := i.Body; v != nil {
		m.SetBody(*v)
	}
	if v := i.Img; v != nil {
		m.SetImg(*v)
	}
	if v := i.AddCommentIDs; len(v) > 0 {
		m.AddCommentIDs(v...)
	}
	if v := i.RemoveCommentIDs; len(v) > 0 {
		m.RemoveCommentIDs(v...)
	}
	if i.ClearMarker {
		m.ClearMarker()
	}
	if v := i.MarkerID; v != nil {
		m.SetMarkerID(*v)
	}
	if i.ClearAccount {
		m.ClearAccount()
	}
	if v := i.AccountID; v != nil {
		m.SetAccountID(*v)
	}
}

// SetInput applies the change-set in the UpdatePostInput on the PostUpdate builder.
func (c *PostUpdate) SetInput(i UpdatePostInput) *PostUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdatePostInput on the PostUpdateOne builder.
func (c *PostUpdateOne) SetInput(i UpdatePostInput) *PostUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateRequestInput represents a mutation input for creating requests.
type CreateRequestInput struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Status    request.Status
	AccountID uuid.UUID
	RequestID uuid.UUID
}

// Mutate applies the CreateRequestInput on the RequestMutation builder.
func (i *CreateRequestInput) Mutate(m *RequestMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetStatus(i.Status)
	m.SetAccountID(i.AccountID)
	m.SetRequestID(i.RequestID)
}

// SetInput applies the change-set in the CreateRequestInput on the RequestCreate builder.
func (c *RequestCreate) SetInput(i CreateRequestInput) *RequestCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateRequestInput represents a mutation input for updating requests.
type UpdateRequestInput struct {
	UpdatedAt    *time.Time
	Status       *request.Status
	ClearAccount bool
	AccountID    *uuid.UUID
	ClearRequest bool
	RequestID    *uuid.UUID
}

// Mutate applies the UpdateRequestInput on the RequestMutation builder.
func (i *UpdateRequestInput) Mutate(m *RequestMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if i.ClearAccount {
		m.ClearAccount()
	}
	if v := i.AccountID; v != nil {
		m.SetAccountID(*v)
	}
	if i.ClearRequest {
		m.ClearRequest()
	}
	if v := i.RequestID; v != nil {
		m.SetRequestID(*v)
	}
}

// SetInput applies the change-set in the UpdateRequestInput on the RequestUpdate builder.
func (c *RequestUpdate) SetInput(i UpdateRequestInput) *RequestUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateRequestInput on the RequestUpdateOne builder.
func (c *RequestUpdateOne) SetInput(i UpdateRequestInput) *RequestUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateSessionInput represents a mutation input for creating sessions.
type CreateSessionInput struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Session   string
	AccountID uuid.UUID
}

// Mutate applies the CreateSessionInput on the SessionMutation builder.
func (i *CreateSessionInput) Mutate(m *SessionMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetSession(i.Session)
	m.SetAccountID(i.AccountID)
}

// SetInput applies the change-set in the CreateSessionInput on the SessionCreate builder.
func (c *SessionCreate) SetInput(i CreateSessionInput) *SessionCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateSessionInput represents a mutation input for updating sessions.
type UpdateSessionInput struct {
	UpdatedAt    *time.Time
	Session      *string
	ClearAccount bool
	AccountID    *uuid.UUID
}

// Mutate applies the UpdateSessionInput on the SessionMutation builder.
func (i *UpdateSessionInput) Mutate(m *SessionMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Session; v != nil {
		m.SetSession(*v)
	}
	if i.ClearAccount {
		m.ClearAccount()
	}
	if v := i.AccountID; v != nil {
		m.SetAccountID(*v)
	}
}

// SetInput applies the change-set in the UpdateSessionInput on the SessionUpdate builder.
func (c *SessionUpdate) SetInput(i UpdateSessionInput) *SessionUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateSessionInput on the SessionUpdateOne builder.
func (c *SessionUpdateOne) SetInput(i UpdateSessionInput) *SessionUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
