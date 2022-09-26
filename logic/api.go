package logic

import (
	"github.com/MiniDouyin/dao"
	"github.com/MiniDouyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func usernameIsValid(username string, c *gin.Context) bool {
	if len(username) > 32 || len(username) < 6 {
		c.JSON(http.StatusOK, dao.UserLoginResponse{
			Response: dao.Response{StatusCode: 1, StatusMsg: "用户名超过32字符或者小于6个字符"},
		})
		return false
	}
	return true
}

func passwordIsValid(password string, c *gin.Context) bool {
	if len(password) > 32 || len(password) < 6 {
		c.JSON(http.StatusOK, dao.UserLoginResponse{
			Response: dao.Response{StatusCode: 1, StatusMsg: "密码超过32字符或者小于6个字符"},
		})
		return false
	}
	return true
}

// CheckLogin 返回值：用户是否存在 密码是否正确 userinfo
func CheckLogin(username, password string) (bool, bool, dao.User) {
	user, err := model.GetUserByName(username)
	if err != nil {
		return false, false, dao.User{}
	}
	if user.Password == password {
		return true, true, user
	}
	return true, false, dao.User{}
}
