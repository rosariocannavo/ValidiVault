package main

import (
	"log"
	"net/http"

	"github.com/rosariocannavo/api_gateway/internal/db"
	"github.com/rosariocannavo/api_gateway/internal/handlers"
	"github.com/rosariocannavo/api_gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// Load Static files
	r.LoadHTMLGlob("./templates/html/*.html")
	r.Static("/css", "./templates/css")
	r.Static("/js", "./templates/js")

	// Connect to DB
	err := db.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	r.Use(middleware.NewRateLimitMiddleware().Handler())

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{})
	})

	r.GET("/get-cookie", handlers.GetCookieHandler)

	r.POST("/login", handlers.HandleLogin)

	r.POST("/registration", handlers.HandleRegistration)

	r.POST("/verify-signature", handlers.HandleverifySignature)

	userRoutes := r.Group("/user")
	userRoutes.Static("/css", "../../templates/css")
	userRoutes.Static("/js", "../../templates/js")

	userRoutes.Use(middleware.NewRateLimitMiddleware().Handler())
	userRoutes.Use(middleware.Authenticate())
	userRoutes.Use(middleware.RoleAuth("user"))
	{
		userRoutes.GET("/user_home", func(c *gin.Context) {
			c.HTML(http.StatusOK, "user_home.html", gin.H{})
		})

		userRoutes.GET("/app/*proxyPath", handlers.ProxyHandler)

	}

	adminRoutes := r.Group("/admin")
	adminRoutes.Static("/css", "../../templates/css")
	adminRoutes.Static("/js", "../../templates/js")

	adminRoutes.Use(middleware.NewRateLimitMiddleware().Handler())
	adminRoutes.Use(middleware.Authenticate())
	adminRoutes.Use(middleware.RoleAuth("admin"))
	{
		adminRoutes.GET("/admin_home", func(c *gin.Context) {
			c.HTML(http.StatusOK, "admin_home.html", gin.H{})
		})

		adminRoutes.GET("/app/*proxyPath", handlers.ProxyHandler)
	}

	_ = r.Run(":8080")

}
