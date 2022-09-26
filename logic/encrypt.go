package logic

import (
	"errors"
	"github.com/MiniDouyin/dao"
	"github.com/MiniDouyin/model"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

// 秘钥
var jwtKey = []byte("Serein404")

type MyClaims struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// SetToken 颁发token
func SetToken(name string, id int64) (string, error) {
	expireTime := time.Now().Add(24 * time.Hour)
	claims := &MyClaims{
		Id:       id,
		Username: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(), //开始时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		log.Printf("SetToken -> SignedString: %s", err.Error())
		return "", err
	}
	return tokenStr, err
}

//GetToken 解析token
func GetToken(tokenStr string) *MyClaims {
	claims, err := parseToken(tokenStr)
	if err != nil {
		log.Println(err)
	}
	return claims
}

func parseToken(tokenStr string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		log.Println("解析失败", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// FlushToken
// 为用户重新生成token
// 将用户原token映射信息从userInfoLogin删除
// 更改token映射的用户信息, 更新token结束时间
func FlushToken(username string, id int64) (string, error) {
	newToken, err := SetToken(username, id) //为用户重新生成token
	oldToken := dao.TokenEndTime[username].Token
	userInfo := dao.UsersLoginInfo[oldToken]
	// fmt.Println("FlushToken: id=", userInfo.ID)

	// 更新UsersLoginInfo 和 TokenEndTime
	delete(dao.UsersLoginInfo, oldToken)
	dao.UsersLoginInfo[newToken] = userInfo
	dao.TokenEndTime[username] = dao.UsrToken{
		Token:   newToken,
		EndTime: time.Now().Add(dao.TokenValidTime),
	}
	return newToken, err
}

// TokenIsValid token有效返回true，否则false
func TokenIsValid(token string) bool {
	/*
		此处会有一个潜在问题，如果用户在刷视频过程中突然token失效怎么办，其实在实际生产中，这几乎不会出现，
		但是在这里如果出现了，我并不能解决。原因在于我判断token过期是在进入Feed函数时判断的，但是用户在刷
		视频时突然token过期，也将会在Feed函数里判断，并不能识别出是那种情况，进而无法知道是否要强制用户重
		新登陆。我的解决方案是，将token过期时间设置为很长，也就是说刷视频中token突然过期的情况几乎不能出现
	*/
	claims := GetToken(token)
	if time.Now().After(dao.TokenEndTime[claims.Username].EndTime) {
		// 用户登录状态改为下线
		err := model.SetOnline(dao.UsersLoginInfo[token].ID, false)
		if err != nil {
			log.Println(err)
		}
		return false
	}
	return true
}
