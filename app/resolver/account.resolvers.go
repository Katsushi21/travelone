package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Katsushi21/travelone"
	"github.com/Katsushi21/travelone/ent"
)

func (r *accountResolver) ID(ctx context.Context, obj *ent.Account) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountResolver) Type(ctx context.Context, obj *ent.Account) (travelone.AccountType, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountResolver) Gender(ctx context.Context, obj *ent.Account) (travelone.AccountGender, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountResolver) Friend(ctx context.Context, obj *ent.Account) ([]*ent.Friend, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountResolver) Mute(ctx context.Context, obj *ent.Account) ([]*ent.Mute, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountResolver) Post(ctx context.Context, obj *ent.Account) ([]*ent.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountResolver) Like(ctx context.Context, obj *ent.Account) ([]*ent.Like, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountResolver) Comment(ctx context.Context, obj *ent.Account) ([]*ent.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// Account returns travelone.AccountResolver implementation.
func (r *Resolver) Account() travelone.AccountResolver { return &accountResolver{r} }

type accountResolver struct{ *Resolver }
