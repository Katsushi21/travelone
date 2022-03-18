package graphql

import (
	"github.com/Katsushi21/traveling_alone/models"
)

type Resolver struct {
	posts    []*models.Post
	users    []*models.User
	comments []*models.Comment
	markers  []*models.Marker
	requests []*models.Request
}
