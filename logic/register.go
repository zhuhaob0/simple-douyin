package logic

import (
	"github.com/MiniDouyin/dao"
	"github.com/MiniDouyin/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// Register
// 获取username,password;
// 根据username判断用户是否存在;
// 生成User结构体,生成token序列;
// 更新token结束时间;
// 更新userLoginInfo,token映射User{};
// 新用户信息添加进数据库中

func Register(username, password string, c *gin.Context) {
	// 用户是否存在
	if model.IsUserExist(username) {
		c.JSON(http.StatusOK, dao.UserLoginResponse{
			Response: dao.Response{
				StatusCode: 1,
				StatusMsg:  "User already exist",
			}})
		return
	}

	// 检查账号密码合法性
	if !usernameIsValid(username, c) || !passwordIsValid(password, c) {
		return
	}

	// 将用户信息存入数据库
	user := dao.User{
		Name:          username,
		Password:      password,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
		Online:        false,
		Model:         gorm.Model{},
	}

	model.AddIntoUserDB(user)

	// 生成token，更新token映射信息
	newUserID := model.GetUserIDByName(username)
	token, _ := SetToken(username, newUserID) // 给新用户生成token序列
	dao.TokenEndTime[username] = dao.UsrToken{
		Token:   token,
		EndTime: time.Now().Add(dao.TokenValidTime),
	}

	// 更新UserLoginInfo
	dao.UsersLoginInfo[token] = dao.User{
		ID:            newUserID,
		Name:          username,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      true,
	}

	c.JSON(http.StatusOK, dao.UserLoginResponse{
		Response: dao.Response{StatusCode: 0},
		UserId:   newUserID,
		Token:    token,
	})
}
