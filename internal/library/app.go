package library

import (
	"github.com/akshay/libraryAssign/internal/library/api"
	"github.com/akshay/libraryAssign/internal/library/middleware/auth"
	"github.com/akshay/libraryAssign/internal/library/service"
	"github.com/gin-gonic/gin"
)

type App struct {
	router    *gin.Engine
	myService *service.Service
}

func NewApp(router *gin.Engine, myService *service.Service) *App {
	app := &App{
		router:    router,
		myService: myService,
	}
	app.registerRoutes()

	return app
}

func (app *App) Run(addr string) error {
	return app.router.Run(addr)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (app *App) registerRoutes() {
	handler := api.NewHandler(app.myService)

	router := app.router
	router.Use(CORSMiddleware())
	HealthCheckRouter := router.Group("api/v1/healthCheck")
	HealthCheckRouter.GET("/status", handler.HealthCheck)

	v1AuthRouter := router.Group("/api/v1/auth")
	v1AuthRouter.POST("log-in", handler.Login)

	v1BookRouter := router.Group("/api/v1/library")
	v1BookRouter.Use(auth.Auth)
	v1BookRouter.GET("/home", handler.GetBooks)
	v1BookRouter.POST("/add-book", handler.AddBooks)
	v1BookRouter.DELETE("/home", handler.DeleteBook)
}

func InitializeApp() *App {
	router := gin.Default()
	service := service.NewService()
	app := NewApp(router, service)
	return app
}
