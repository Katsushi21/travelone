package graphql

import (
	"context"
	"fmt"

	"github.com/Katsushi21/traveling_alone/models"
)

func (r *mutationResolver) CreateSession(ctx context.Context, input models.SessionInput) (*models.Session, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSession(ctx context.Context, input models.SessionInput) (*models.Session, error) {
	session := models.Session{
		UserID: input.UserID,
	}
	r.DB.First(&session)
	r.DB.Model(&session).Update(
		"session",
		input.Session,
	)

	return &session, nil
}

func (r *mutationResolver) DeleteSession(ctx context.Context, id int) (*models.Session, error) {
	panic(fmt.Errorf("not implemented"))
}
