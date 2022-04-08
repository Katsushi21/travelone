package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
)

func (r *mutationResolver) UpdateSession(ctx context.Context, input models.SessionInput) (*models.Session, error) {
	session := models.Session{
		UserID: input.UserID,
	}
	r.DB.First(&session)
	r.DB.Model(&session).Update( // Whereが必要か要検証
		"session",
		input.Session,
	)

	return &session, nil
}
