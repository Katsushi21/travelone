package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Katsushi21/travelone"
)

// AccountPageInfo is the resolver for the accountPageInfo field.
func (r *queryResolver) AccountPageInfo(ctx context.Context, id string) (*travelone.Account, error) {
	panic(fmt.Errorf("not implemented: AccountPageInfo - accountPageInfo"))
}

// MyPageInfo is the resolver for the myPageInfo field.
func (r *queryResolver) MyPageInfo(ctx context.Context, id string) (*travelone.Account, error) {
	panic(fmt.Errorf("not implemented: MyPageInfo - myPageInfo"))
}

// LikesByPost is the resolver for the likesByPost field.
func (r *queryResolver) LikesByPost(ctx context.Context, postID string) ([]*travelone.Like, error) {
	panic(fmt.Errorf("not implemented: LikesByPost - likesByPost"))
}

// AllMarkers is the resolver for the allMarkers field.
func (r *queryResolver) AllMarkers(ctx context.Context) ([]*travelone.Marker, error) {
	panic(fmt.Errorf("not implemented: AllMarkers - allMarkers"))
}

// AllPosts is the resolver for the allPosts field.
func (r *queryResolver) AllPosts(ctx context.Context) ([]*travelone.Post, error) {
	panic(fmt.Errorf("not implemented: AllPosts - allPosts"))
}

// RequestsByAccountID is the resolver for the requestsByAccountId field.
func (r *queryResolver) RequestsByAccountID(ctx context.Context, accountID string) ([]*travelone.Request, error) {
	panic(fmt.Errorf("not implemented: RequestsByAccountID - requestsByAccountId"))
}

// RequestsByTargetID is the resolver for the requestsByTargetId field.
func (r *queryResolver) RequestsByTargetID(ctx context.Context, targetAccountID string) ([]*travelone.Request, error) {
	panic(fmt.Errorf("not implemented: RequestsByTargetID - requestsByTargetId"))
}

// SessionByAccountID is the resolver for the sessionByAccountId field.
func (r *queryResolver) SessionByAccountID(ctx context.Context, accountID string) (*travelone.Session, error) {
	panic(fmt.Errorf("not implemented: SessionByAccountID - sessionByAccountId"))
}

// Query returns travelone.QueryResolver implementation.
func (r *Resolver) Query() travelone.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
