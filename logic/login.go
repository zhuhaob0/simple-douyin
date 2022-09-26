package logic

import (
	"github.com/MiniDouyin/dao"
	"github.com/MiniDouyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Login 获取username, password,
// 然后判断用户是否存在，同时获取用户信息
// 判断用户是否存在 or 密码是否正确
// 判断用户是否已经在线
// 更新token

func Login(username, password string, c *gin.Context) {
	exist, ok, user := CheckLogin(username, password)
	// 用户是否存在  密码是否错误
	if !exist || !ok {
		notExistOrOk(exist, ok, c)
		return
	}
	//用户已在线
	if user.Online {
		c.JSON(http.StatusOK, dao.UserLoginResponse{
			Response: dao.Response{StatusCode: 1, StatusMsg: "User is online, repeat login"},
		})
		return
	}
	var token string
	if _, exist := dao.TokenEndTime[user.Name]; exist { // 存在内存
		// 刷新token
		token, _ = FlushToken(user.Name, user.ID)
	} else { //因某些原因，不在内存中，重新生成token，然后放入
		token, _ = SetToken(user.Name, user.ID)
		// 更新UserLoginInfo 和 TokenEndTime
		dao.TokenEndTime[user.Name] = dao.UsrToken{
			Token:   token,
			EndTime: time.Now().Add(dao.TokenValidTime),
		}
		var err error
		user, err = model.GetUserByID(user.ID)
		if err == nil {
			dao.UsersLoginInfo[token] = dao.User{
				ID:            user.ID,
				Name:          user.Name,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      false,
			}
		}
	}

	c.JSON(http.StatusOK, dao.UserLoginResponse{
		Response: dao.Response{StatusCode: 0},
		UserId:   user.ID,
		Token:    token,
	})
}

func notExistOrOk(exist, ok bool, c *gin.Context) {
	var respond string
	if exist == false {
		respond = "User doesn't exist"
	} else if ok == false {
		respond = "Password error"
	}
	c.JSON(http.StatusOK, dao.UserLoginResponse{
		Response: dao.Response{StatusCode: 1, StatusMsg: respond},
	})
}
