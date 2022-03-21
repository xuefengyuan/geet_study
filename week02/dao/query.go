package dao

import (
    "database/sql"
    "errors"
)

// create db table sql
//  CREATE TABLE `users` (
//  `id` int(11) NOT NULL,
//  `name` varchar(255) DEFAULT NULL,
//  `age` int(11) DEFAULT NULL,
//  PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8;

// 根据id查询用户名
// db层是不是要处理异常，根据实际需求来做判断
// dao层如果处理不了异常信息，就往上抛
// 根据实际业务看上层怎么处理
func UserQueryToName(id int) (string, error) {
    var name string
    err := SqlDB.QueryRow("select name from users where id = ?", id).Scan(&name)
    // 查找不到数据，能处理，降级error，业务层自己再针对返回值进行校验
    if err != nil && errors.Is(err,sql.ErrNoRows){
        err = nil
    }
    return name, err
}
