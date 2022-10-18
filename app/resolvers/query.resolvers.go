package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Katsushi21/travelone/ent"
	"github.com/google/uuid"
)

func (r *queryResolver) QueryAccountByID(ctx context.Context, id uuid.UUID) (*ent.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QueryAccountByMyID(ctx context.Context, id uuid.UUID) (*ent.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QueryLikesByPostID(ctx context.Context, postID uuid.UUID) ([]*ent.Like, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QueryMarkers(ctx context.Context) ([]*ent.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QueryPosts(ctx context.Context) ([]*ent.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QueryRequestsByAccountID(ctx context.Context, accountID uuid.UUID) ([]*ent.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QueryRequestsByRequestID(ctx context.Context, requestID uuid.UUID) ([]*ent.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QuerySessionByAccountID(ctx context.Context, id uuid.UUID) (*ent.Session, error) {
	panic(fmt.Errorf("not implemented"))
}
