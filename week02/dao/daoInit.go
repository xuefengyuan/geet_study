package dao

import "database/sql"

var SqlDB *sql.DB

func InitDB(connInfo string) error {
    db, err := sql.Open("mysql", connInfo)

    if err != nil {
        return err
    }

    err = db.Ping()
    if err != nil {
        return err
    }

    SqlDB = db
    return nil
}
