package main

import (
	"github.com/MiniDouyin/dao"
	"github.com/gin-gonic/gin"
)

/*
url  -->  router-->  controller -->  logic   -->  models         -->  dao
请求  -->  路由   -->  控制器      -->  业务逻辑  -->  模型层的增删改查  -->  数据库
*/

func main() {
	gin.SetMode(gin.DebugMode)
	if err := dao.Init(); err != nil {
		return
	}
	r := setupRouter()
	r.Run(":8080")
}
