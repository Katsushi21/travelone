package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Katsushi21/travelone"
	"github.com/Katsushi21/travelone/ent"
)

func (r *queryResolver) AccountPageInfo(ctx context.Context, id string) (*ent.Account, error) {
	panic(fmt.Errorf("not implemented: AccountPageInfo - accountPageInfo"))
}

func (r *queryResolver) MyPageInfo(ctx context.Context, id string) (*ent.Account, error) {
	panic(fmt.Errorf("not implemented: MyPageInfo - myPageInfo"))
}

func (r *queryResolver) LikesByPost(ctx context.Context, postID string) ([]*ent.Like, error) {
	panic(fmt.Errorf("not implemented: LikesByPost - likesByPost"))
}

func (r *queryResolver) AllMarkers(ctx context.Context) ([]*ent.Marker, error) {
	panic(fmt.Errorf("not implemented: AllMarkers - allMarkers"))
}

func (r *queryResolver) AllPosts(ctx context.Context) ([]*ent.Post, error) {
	panic(fmt.Errorf("not implemented: AllPosts - allPosts"))
}

func (r *queryResolver) RequestsByAccountID(ctx context.Context, accountID string) ([]*ent.Request, error) {
	panic(fmt.Errorf("not implemented: RequestsByAccountID - requestsByAccountId"))
}

func (r *queryResolver) RequestsByTargetID(ctx context.Context, targetAccountID string) ([]*ent.Request, error) {
	panic(fmt.Errorf("not implemented: RequestsByTargetID - requestsByTargetId"))
}

func (r *queryResolver) SessionByAccountID(ctx context.Context, accountID string) (*ent.Session, error) {
	panic(fmt.Errorf("not implemented: SessionByAccountID - sessionByAccountId"))
}

func (r *queryResolver) Node(ctx context.Context, id string) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns travelone.QueryResolver implementation.
func (r *Resolver) Query() travelone.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
