package login

import (
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    user2 "week04/internal/models/user"
)

func UserLogin(c *gin.Context)  {

    // TODO 这里获取输入参数进行跟db交互的数据转换
    user := user2.NewUser()
    user.Name = "save name"
    user.Password= "password"
    err := user.Save()
    log.Fatal(err)
    c.String(http.StatusOK,"ok")

}

func UserInfo(c *gin.Context)  {
    user := user2.NewUser()
    err := user.GetUserInfo(10)
    log.Fatal(user.Name)
    log.Fatal(err)

    c.String(http.StatusOK,user.Name)
}
