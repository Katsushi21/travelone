package main

import (
	"io"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Katsushi21/travelone/database"
	"github.com/Katsushi21/travelone/ent"
	"github.com/Katsushi21/travelone/middleware"
	"github.com/Katsushi21/travelone/resolvers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func graphqlHandler(client *ent.Client) gin.HandlerFunc {
	h := handler.NewDefaultServer(
		resolvers.NewSchema(client),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {

	log, _ := os.Create(os.Getenv("SERVER_LOG"))
	gin.DefaultWriter = io.MultiWriter(log, os.Stdout)

	r := gin.Default()
	r.Use(middleware.CorsConfig())
	r.Use(middleware.GinContextToContextMiddleware())

	client := database.Connect()
	r.POST("/graphql", graphqlHandler(client))
	r.GET("/increment", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8000")
}
