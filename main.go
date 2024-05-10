package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type User struct {
	Username string `json:"username"`
	Gender   string `json:"gender"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func getWelcome(r *gin.Engine, conn *pgx.Conn) *gin.Engine {
	r.GET("/welcome", func(c *gin.Context) {
		lastname := c.Query("lastname")
		var name string
		q := fmt.Sprintf("SELECT first_name FROM users WHERE last_name = '%s' LIMIT 1", lastname)
		err := conn.QueryRow(c, q).Scan(&name)
		if err != nil {
			c.String(404, "User not found")
		}
		c.String(200, "Hello %s %s", name, lastname)
	})
	return r
}

func main() {
	ctx := context.Background()
	conn, _ := pgx.Connect(ctx, "host=localhost user=postgres password=password database=postgres")

	r := setupRouter()
	r = getWelcome(r, conn)

	r.Run(":8081")
}
