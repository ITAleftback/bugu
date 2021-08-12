package routers

import (
	"bugu/internal/middleware"
	chip2 "bugu/internal/routers/api/chip"
	module2 "bugu/internal/routers/api/module"
	overload2 "bugu/internal/routers/api/overload"
	plugin2 "bugu/internal/routers/api/plugin"
	user2 "bugu/internal/routers/api/user"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CrosHandler())
	r.Use(middleware.Translations())

	user := user2.NewUser()
	apiUser := r.Group("/api/user")
	{
		apiUser.POST("/register", user.Register)
		apiUser.POST("/login", user.Login)
		apiUser.PUT("/updateCode", user.UpdateCode)
	}

	chip := chip2.NewChip()
	apiChip := r.Group("/api/chips")
	apiChip.Use(middleware.JWT())
	{
		apiChip.POST("/", chip.Create)
		apiChip.PUT("/:id", chip.Update)
		apiChip.DELETE("/:id", chip.Delete)
		apiChip.GET("/:id", chip.Get)
		apiChip.GET("/list", chip.List)
	}

	plugin := plugin2.NewPlugin()
	apiPlugin := r.Group("/api/plugins")
	apiPlugin.Use(middleware.JWT())
	{
		apiPlugin.POST("/", plugin.Create)
		apiPlugin.PUT("/:id", plugin.Update)
		apiPlugin.DELETE("/:id", plugin.Delete)
		apiPlugin.GET("/:id", plugin.Get)
		apiPlugin.GET("/list/:chip_id", plugin.List)
	}

	module := module2.NewModule()
	apiModule := r.Group("api/modules")
	apiModule.Use(middleware.JWT())
	{
		apiModule.POST("/", module.Create)
		apiModule.PUT("/:id", module.Update)
		apiModule.DELETE("/:id", module.Delete)
		apiModule.GET("/:id", module.Get)
		apiModule.GET("/list/:plugin_id", module.List)
	}

	overload := overload2.NewOverload()
	apiOverload := r.Group("api/overloads")
	apiOverload.Use(middleware.JWT())
	{
		apiOverload.POST("/", overload.Create)
		apiOverload.PUT("/:id", overload.Update)
		apiOverload.DELETE("/:id", overload.Delete)
		apiOverload.GET("/:id", overload.Get)
		apiOverload.GET("/list/:module_id", overload.List)
	}
	return r
}
