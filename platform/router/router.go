// platform/router/router.go

package router

import (
	"encoding/gob"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"todoapp/controllers"
	"todoapp/platform/authenticator"
	"todoapp/platform/middleware"
	"todoapp/web/app/callback"
	"todoapp/web/app/login"
	"todoapp/web/app/logout"
	"todoapp/web/app/user"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))

	//Tanpa Auth
	//router.GET("/user", user.Handler)
	//Dengan Auth
	router.GET("/user", middleware.IsAuthenticated, user.Handler)

	router.GET("/logout", logout.Handler)

	return router
}

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/posts", controllers.FindBlogPosts)
	r.POST("/posts", controllers.CreateBlogPost)
	r.GET("/posts/:id", controllers.FindPost)
	r.PATCH("/posts/:id", controllers.UpdatePost)
	r.DELETE("posts/:id", controllers.DeletePost)
	return r
}
