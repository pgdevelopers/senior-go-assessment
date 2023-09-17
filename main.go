package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	router := gin.New()
	conn, err := pgx.Connect(ctx, "host=localhost user=postgres password=password database=postgres")
	if err != nil {
		log.Fatal(err)
	}

	router.GET("/welcome", func(c *gin.Context) {
		lastname := c.Query("lastname")
		var name string
		q := fmt.Sprintf("SELECT first_name FROM users WHERE last_name = '%s' LIMIT 1", lastname)
		err = conn.QueryRow(ctx, q).Scan(&name)
		if err != nil {
			c.String(404, "User not found")
		}
		c.String(200, "Hello %s %s", name, lastname)
	})

	router.Run(":8080")
}
