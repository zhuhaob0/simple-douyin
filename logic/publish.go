package logic

import (
	"fmt"
	"github.com/MiniDouyin/dao"
	"github.com/MiniDouyin/model"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func Publish(token string, c *gin.Context) {
	data, err := c.FormFile("data")
	title := c.PostForm("title")
	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	saveIntoVideoDB(token, title, data, c)
}

func saveIntoVideoDB(token, title string, data *multipart.FileHeader, c *gin.Context) {
	filename := filepath.Base(data.Filename)
	fmt.Printf("data.header = %v\n", data.Header)
	user := dao.UsersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.ID, filename)
	playUrl := saveVideo(finalName, data, c)
	coverUrl := saveCover(finalName, playUrl, data)
	// 生成video，放入数据库
	video := dao.Video{
		AuthorID:      user.ID,
		Author:        user,
		PlayUrl:       playUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         title,
	}
	model.AddIntoVideo(video)

	c.JSON(http.StatusOK, dao.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

func saveVideo(finalName string, data *multipart.FileHeader, c *gin.Context) string {
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return "ErrorUrl"
	}
	return getUrlPrefix() + finalName
}
func saveCover(finalName, playUrl string, data *multipart.FileHeader) string {
	//获取视频第一帧作为封面保存在本地
	filename := GetFirstFrame(finalName, playUrl)
	//返回封面的名字
	return getUrlPrefix() + filename
}

// ffmpeg -i http://video.pearvideo.com/head/20180301/cont-1288289-11630613.mp4
// -r 1 -t 4 -f image2 image-%05d.jpeg

func GetFirstFrame(finaName, filepath string) string {
	curPath := GetCurrentPath()
	filename := strings.Split(finaName, ".")[0] + ".jpg"
	dstPath := curPath + "/public/" + filename
	fmt.Println("filepath:", filepath)
	fmt.Println("dst:", dstPath)
	cmd := exec.Command("ffmpeg", "-i", filepath, "-ss", "1", "-frames:v", "1", "-f", "image2", dstPath)

	if err := cmd.Run(); err != nil {
		log.Printf("cmd.Run: %s", err.Error())
	}
	return filename
}

func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// PublishList 发布的作品
func PublishList(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	// 用userID 找作品
	videoList, err := model.GetVideosByAuthorID(int64(userID))
	if err != nil {
		return
	}
	// 返回客户端
	c.JSON(http.StatusOK, dao.VideoListResponse{
		Response: dao.Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}

func getUrlPrefix() string {
	ip := "http://192.168.0.4:8080/static/"
	return ip
}
