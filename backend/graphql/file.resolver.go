package graphql

import (
	"context"
	"fmt"

	"github.com/Katsushi21/traveling_alone/models"
)

func (r *mutationResolver) UploadFile(ctx context.Context, input models.UploadFile) (*models.File, error) {
	panic(fmt.Errorf("not implemented"))
}
