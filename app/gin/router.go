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

	gin.GET("/list_ipl_recurring", func(ctx *g.Context) {
		ctx.HTML(200, "list_ipl_recurring.html", nil)
	})

	gin.GET("/transaksi", func(ctx *g.Context) {
		ctx.HTML(200, "kelola-transaksi.html", nil)
	})

	gin.GET("/kas", func(ctx *g.Context) {
		ctx.HTML(200, "kas.html", nil)
	})

	gin.GET("/login", func(ctx *g.Context) {
		ctx.HTML(200, "login.html", nil)
	})

	gin.GET("/profile", func(ctx *g.Context) {
		ctx.HTML(200, "account-profile.html", nil)
	})

	return gin
}
