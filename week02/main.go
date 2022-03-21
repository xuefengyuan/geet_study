package main

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.geet_study.week02/dao"
)

func main() {

    DbConnInfo := "root:mysql@tcp(192.168.159.133:3306)/geek?charset=utf8"
    //DbConnInfo := "root:mysql@tcp(127.0.0.1:3306)/geek?charset=utf8"
    dao.InitDB(DbConnInfo)
    userID := 1
    name, err := dao.UserQueryToName(userID)
    if err != nil || name == "" {
        fmt.Printf("query id : %d, name error : %v\n or not find user info ", userID, err)
        return
    }

    fmt.Printf("user %d = %s\n", userID, name)

    userID = 2
    name, err = dao.UserQueryToName(userID)

    if err != nil || name == "" {
        fmt.Printf("query id : %d, name error : %v\n or not find user info ", userID, err)
        return
    }
    fmt.Printf("user %d = %s\n", userID, name)
}
