package main

import (
	"io"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Katsushi21/traveling_alone/database"
	"github.com/Katsushi21/traveling_alone/graphql"
	"github.com/Katsushi21/traveling_alone/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

func main() {
	db := database.Connect()

	log, _ := os.Create(os.Getenv("SERVER_LOG"))
	gin.DefaultWriter = io.MultiWriter(log, os.Stdout)

	r := gin.Default()
	r.Use(middleware.GinContextToContextMiddleware())

	r.POST("/query", graphqlHandler(db))
	r.GET("/", playgroundHandler())
	r.Run(":8000")
}
