package main

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Katsushi21/traveling_alone/database"
	"github.com/Katsushi21/traveling_alone/graphql"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		"travel-db.com", 5432, "traveling1145", "traveling_alone", "a4a1274c79f54a3c976f495b04c021c0",
	)
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		database.Comments{},
		database.Markers{},
		database.Posts{},
		database.Requests{},
		database.Users{},
	)

	r := gin.Default()
	r.POST("/query", graphqlHandler(db))
	r.GET("/", playgroundHandler())
	r.Run()
}

// Defining the Graphql handler
func graphqlHandler(db *gorm.DB) gin.HandlerFunc {
	h := handler.NewDefaultServer(
		graphql.NewExecutableSchema(
			graphql.Config{Resolvers: &graphql.Resolver{DB: db}},
		),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}
	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}
