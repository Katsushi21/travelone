package main

import (
	"io"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Katsushi21/travelone/database"
	"github.com/Katsushi21/travelone/middleware"
	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

func main() {
	db := database.Connect()

	log, _ := os.Create(os.Getenv("SERVER_LOG"))
	gin.DefaultWriter = io.MultiWriter(log, os.Stdout)

	store := gormsessions.NewStore(db, true, []byte("secret"))

	r := gin.Default()
	r.Use(sessions.Sessions("travelone", store))
	r.Use(middleware.CorsConfig())
	r.Use(middleware.GinContextToContextMiddleware())

	r.POST("/graphql", graphqlHandler(db))
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
