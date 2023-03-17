package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"server/models"
)

// CheckUserExist 检查指定用户是否存在
func CheckUserExist(username string) (errs error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	fmt.Println("before db.GET")
	if err := db.Get(&count, sqlStr, username); err != nil {
		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
		return err
	}
	fmt.Println("after db.GET")
	return
}

// InsertUser 数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	password := encryptPassword(user.Password)
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, password)
	return err
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte("wocaocoacoa"))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
