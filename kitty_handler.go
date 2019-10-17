package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/plush"
)

func kittyHandler(tpl string) func(*gin.Context) {
	cats := []string{
		"/img/kitties/1.jpg",
		"/img/kitties/2.jpg",
		"/img/kitties/3.jpg",
		"/img/kitties/4.jpg",
		"/img/kitties/5.jpg",
		"/img/kitties/6.jpg",
	}
	return func(c *gin.Context) {
		ctx := plush.NewContext()
		ctx.Set("imgSrc", randStr(cats))
		str, err := plush.Render(tpl, ctx)
		if err != nil {
			c.String(500, "Error rendering template: %s", err)
			return
		}
		renderHTML(c, str)
	}
}
