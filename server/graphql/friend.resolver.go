package graphql

import (
	"context"
	"fmt"

	"github.com/Katsushi21/traveling_alone/models"
	"gorm.io/gorm"
)

func (r *mutationResolver) CreateFriend(ctx context.Context, input models.FriendInput) (*models.Friend, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteFriend(ctx context.Context, input models.FriendInput) (*models.Friend, error) {
	friend := models.Friend{
		UserID:   input.UserID,
		FriendID: input.FriendID,
	}
	r.DB.First(&friend)
	r.DB.Model(&friend).Update( // Whereが必要か要検証
		"friends", gorm.Expr("array_remove(?, ?)", "friends", input.UserID),
	)

	return &friend, nil
}
