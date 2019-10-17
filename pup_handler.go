package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/plush"
)

func pupHandler(tpl string) func(*gin.Context) {
	dogs := []string{
		"/img/pups/1.jpg",
		"/img/pups/2.jpg",
		"/img/pups/3.png",
		"/img/pups/3.png",
		"/img/pups/4.png",
		"/img/pups/5.png",
		"/img/pups/6.jpg",
		"/img/pups/7.jpg",
	}
	return func(c *gin.Context) {
		randomDog := randStr(dogs)
		ctx := plush.NewContext()
		ctx.Set("imgSrc", randomDog)
		rendered, err := plush.Render(tpl, ctx)
		if err != nil {
			c.String(500, "Error rendering template: %s", err)
			return
		}
		renderHTML(c, rendered)
	}
}
