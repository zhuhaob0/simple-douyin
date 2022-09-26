package logic

import (
	"github.com/MiniDouyin/dao"
	"github.com/MiniDouyin/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// UserInfo
//首先获取token,根据token从内存找出用户的信息
// 然后设置用户为在线状态，最后将用户信息返回给客户端展示出来
func UserInfo(token string, c *gin.Context) {
	user, exist := dao.UsersLoginInfo[token]
	user, _ = model.GetUserByID(user.ID)
	//fmt.Printf("UserInfo: %v", user)
	if exist { // 用户存在
		err := model.SetOnline(user.ID, true)
		if err != nil {
			log.Printf("UserInfo:更新在线状态失败: %s\n", err.Error())
			c.JSON(http.StatusOK, dao.UserResponse{
				Response: dao.Response{StatusCode: 1, StatusMsg: "服务器错误"},
			})
			return
		}
		c.JSON(http.StatusOK, dao.UserResponse{
			Response: dao.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, dao.UserResponse{
			Response: dao.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
