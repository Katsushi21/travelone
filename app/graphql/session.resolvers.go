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
	err := r.DB.
		Debug().
		Create(&session).
		Error

	if err != nil {
		return nil, err
	}

	return session, nil
}

func (r *mutationResolver) UpdateSession(ctx context.Context, input models.SessionInput) (*models.Session, error) {
	session := &models.Session{
		UserID: input.UserID,
	}
	err := r.DB.
		Debug().
		Model(&session).
		Update(
			"session",
			&input.Session,
		).
		Error

	if err != nil {
		return nil, err
	}

	return session, nil
}

func (r *mutationResolver) DeleteSession(ctx context.Context, userID int) (*models.Session, error) {
	session := &models.Session{
		UserID: userID,
	}
	err := r.DB.
		Debug().
		Clauses(clause.Returning{}).
		Delete(&session).
		Error

	if err != nil {
		return nil, err
	}

	return session, nil
}
