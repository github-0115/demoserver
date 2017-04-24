package login

import (
	cfg "CoreSystemBase/config"
	manager "CoreSystemBase/service/model/managermodel"
	vars "CoreSystemBase/service/vars"
	security "CoreSystemBase/utils/security"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type LoginParams struct {
	Managername string `json:"managername"`
	Password    string `json:"password"`
}

func Login(c *gin.Context) {

	var loginParams LoginParams
	if err := c.BindJSON(&loginParams); err != nil {
		log.Error(fmt.Sprintf("bind json error:%s", err.Error()))
		vars.FailReturn(c, vars.ErrBindJSON)
		return
	}
	managername := loginParams.Managername
	password := loginParams.Password

	ManagerColl, err := manager.QueryUser(managername)
	if err != nil {
		log.Error(fmt.Sprintf("find user error:%s", err.Error()))
		vars.FailReturn(c, vars.ErrLoginParams)
		return
	}

	if !security.CheckPasswordHash(password, ManagerColl.Password) {
		log.Error(fmt.Sprintf("password error.managername=%s, password=%s, saved_password=%s", managername, password, ManagerColl.Password))
		vars.FailReturn(c, vars.ErrLoginParams)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * time.Duration(cfg.Cfg.LoginTokenExpire)).Unix(),
		"username": ManagerColl.Managername,
	})

	tokenStr, err := token.SignedString([]byte(cfg.Cfg.LoginSecret))
	if err != nil {
		log.Error("gen token failed. err=" + err.Error())
		vars.FailReturn(c, &vars.Err{400, "gen token failed."})
		return
	}

	c.JSON(200, gin.H{
		"code":  0,
		"token": tokenStr,
	})
}
