package routes

import (
	"TodoList/api"
	"TodoList/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret" /* encryption key */))
	r.Use(sessions.Sessions("mysession", store)) // 存储？what
	// 基础路由
	v1 := r.Group("api/v1")
	{
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		authed := v1.Group("/")
		authed.Use(middleware.AuthenticateJWT())
		{
			authed.POST("task/create", api.CreateTask)
			authed.GET("task/:id", api.ShowTask)
			authed.PUT("task/:id", api.UpdateTask)
			authed.DELETE("task/:id", api.DeleteTask)
			authed.POST("task/search", api.SearchTask)
			authed.GET("task/list", api.TaskList)
		}
	}

	return r
}
