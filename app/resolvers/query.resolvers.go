package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Katsushi21/travelone"
	"github.com/Katsushi21/travelone/ent"
	"github.com/google/uuid"
)

func (r *queryResolver) AccountPageInfo(ctx context.Context, id uuid.UUID) (*ent.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MyPageInfo(ctx context.Context, id uuid.UUID) (*ent.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) LikesByPost(ctx context.Context, postID uuid.UUID) ([]*ent.Like, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllMarkers(ctx context.Context) ([]*ent.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllPosts(ctx context.Context) ([]*ent.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) RequestsByAccountID(ctx context.Context, accountID uuid.UUID) ([]*ent.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) RequestsByTargetID(ctx context.Context, targetAccountID uuid.UUID) ([]*ent.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SessionByAccountID(ctx context.Context, accountID uuid.UUID) (*ent.Session, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns travelone.QueryResolver implementation.
func (r *Resolver) Query() travelone.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
