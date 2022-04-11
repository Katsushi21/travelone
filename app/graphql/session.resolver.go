package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateSession(ctx context.Context, input models.SessionInput) (*models.Session, error) {
	session := &models.Session{
		UserID: input.UserID,
	}
	r.DB.Create(&session)

	return session, nil
}

func (r *mutationResolver) UpdateSession(ctx context.Context, input models.SessionInput) (*models.Session, error) {
	session := &models.Session{
		UserID: input.UserID,
	}
	r.DB.First(&session)
	r.DB.Model(&session).Update(
		"session",
		&input.Session,
	)

	return session, nil
}

func (r *mutationResolver) DeleteSession(ctx context.Context, userID int) (*models.Session, error) {
	session := &models.Session{
		UserID: userID,
	}
	r.DB.Clauses(clause.Returning{}).Delete(&session)

	return session, nil
}
