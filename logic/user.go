package logic

import (
	"server/dao/mysql"
	"server/models"
	"server/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	//生成UID
	userId := snowflake.GenID()
	u := models.User{
		UserID:   userId,
		UserName: p.Username,
		Password: p.Password,
	}
	//保存进数据库
	snowflake.GenID()
	return mysql.InsertUser(&u)
}
