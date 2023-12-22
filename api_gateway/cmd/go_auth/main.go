package main

import (
	"fmt"
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

	fmt.Println("Connected to MongoDB!")

	r.Use(middleware.NewRateLimitMiddleware().Handler())

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{})
	})

	// Cookie endpoint to retrieve account and jwt
	r.GET("/get-cookie", handlers.GetCookieHandler)

	// Login endpoint
	r.POST("/login", handlers.HandleLogin)

	//registration endpoint
	r.POST("/registration", handlers.HandleRegistration)

	//metamask signature verification endpoint
	r.POST("/verify-signature", handlers.HandleverifySignature)

	//Protected middleware User endpoints
	userRoutes := r.Group("/user")

	userRoutes.Static("/css", "../../templates/css")
	userRoutes.Static("/js", "../../templates/js")

	userRoutes.Use(middleware.Authenticate())
	userRoutes.Use(middleware.RoleAuth("user"))
	{
		userRoutes.GET("/user_home", func(c *gin.Context) {
			c.HTML(http.StatusOK, "user_home.html", gin.H{})
		})

		userRoutes.GET("/app/*proxyPath", middleware.Authenticate(), handlers.ProxyHandler)

	}

	//Protected middleware Admin endpoints
	adminRoutes := r.Group("/admin")

	adminRoutes.Static("/css", "../../templates/css")
	adminRoutes.Static("/js", "../../templates/js")

	adminRoutes.Use(middleware.Authenticate())
	adminRoutes.Use(middleware.RoleAuth("admin"))
	{
		adminRoutes.GET("/admin_home", func(c *gin.Context) {
			c.HTML(http.StatusOK, "admin_home.html", gin.H{})
		})

		adminRoutes.GET("/app/*proxyPath", middleware.Authenticate(), handlers.ProxyHandler)
	}

	//TODO: this route use the proxy + cb launched by button in login
	//r.GET("/app/*proxyPath", middleware.Authenticate(), handlers.ProxyHandler) //handler of the proxyy

	// Run the server
	_ = r.Run(":8080")

}
