package main

import (
	"templ-air/views"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		posts := []views.Post{
			{
				ID:      1,
				Title:   "First Post",
				Content: "This is the first post",
			},
			{
				ID:      2,
				Title:   "Second Post",
				Content: "This is the second post",
			},
			{
				ID:      3,
				Title:   "Third Post",
				Content: "This is the third post",
			},
		}
		render(c, 200, views.MakeHome(posts))

	})

	r.Run()
}
