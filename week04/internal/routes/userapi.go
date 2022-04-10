package routes

import (
    "github.com/gin-gonic/gin"
    "week04/internal/app/http/controllers/api/v1/login"
)

func RegisterAPIRoutes(r *gin.Engine) {
    v1 := r.Group("/api/v1")
    {
        v1.POST("/login",login.UserLogin)
        v1.GET("/info",login.UserInfo)
    }

}
