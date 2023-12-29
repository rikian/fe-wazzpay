package gin

import (
	g "github.com/gin-gonic/gin"
)

func Router() *g.Engine {
	gin := g.Default()
	gin.LoadHTMLGlob("./views/*.html")

	// middleware
	gin.Use(func(ctx *g.Context) {
		defer func() {
			r := recover()
			if r != nil {
				ctx.Writer.Write([]byte("sorry we are under maintaince"))
			}
		}()

		ctx.Next()
	})

	// static
	gin.Static("/assets/", "./assets/")
	gin.Static("/static/", "./static/")

	// controller
	gin.GET("/", func(ctx *g.Context) {
		ctx.HTML(200, "index.html", nil)
	})

	gin.GET("/login", func(ctx *g.Context) {
		ctx.HTML(200, "login.html", nil)
	})

	return gin
}
