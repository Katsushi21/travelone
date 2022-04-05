package graphql

// Comment returns CommentResolver implementation.
func (r *Resolver) Comment() CommentResolver { return &commentResolver{r} }

// Friend returns FriendResolver implementation.
func (r *Resolver) Friend() FriendResolver { return &friendResolver{r} }

// Like returns LikeResolver implementation.
func (r *Resolver) Like() LikeResolver { return &likeResolver{r} }

// Marker returns MarkerResolver implementation.
func (r *Resolver) Marker() MarkerResolver { return &markerResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type commentResolver struct{ *Resolver }
type friendResolver struct{ *Resolver }
type likeResolver struct{ *Resolver }
type markerResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
