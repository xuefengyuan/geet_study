package bootstarap

import (
    "github.com/gin-gonic/gin"
    "week04/internal/routes"
)

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {

    // 注册全局中间件
    registerGlobalMiddleWare(router)

    //  注册 API 路由
    routes.RegisterAPIRoutes(router)


}

func registerGlobalMiddleWare(router *gin.Engine) {
    //router.Use(
    //
    //)
}