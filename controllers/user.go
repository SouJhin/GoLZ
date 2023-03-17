package controllers

import (
	"net/http"

	"server/logic"
	"server/models"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("请求参数有误", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "注册失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":      "sign up success",
		"userinfo": p,
	})
}
