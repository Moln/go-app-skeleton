package app

import (
	"github.com/gin-gonic/gin"
	"go-demo/module/app/controller"
	"net/http"
)

type Route struct {
	app  *controller.AppController
	auth *controller.AuthController
}

func (e *Route) Register(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("/users", e.app.GetUsers)
	r.GET("/users/:id", e.app.GetUser)
	r.POST("/auth/login", e.auth.Login)
	r.POST("/auth/logout", e.auth.Logout)
	//r.GET("/users/:id", e.app.GetUser)

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	//authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	//	"foo":  "bar", // user:foo password:bar
	//	"manu": "123", // user:manu password:123
	//}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	//authorized.POST("admin", func(c *gin.Context) {
	//	user := c.MustGet(gin.AuthUserKey).(string)
	//
	//	// Parse JSON
	//	var json struct {
	//		Value string `json:"value" binding:"required"`
	//	}
	//
	//	if c.Bind(&json) == nil {
	//		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	//	}
	//})
}

func NewAppRoute(
	app *controller.AppController,
	auth *controller.AuthController,
) *Route {
	return &Route{app, auth}
}
