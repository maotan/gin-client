/**
* @Author: mo tan
* @Description:
* @Date 2021/1/1 21:58
 */
package main

import (
	"fmt"
	"gin-client/domain"
	feign_client "gin-client/feign-client"
	"gin-client/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ilibs/gosql/v2"
	"github.com/maotan/go-truffle/truffle"
	"github.com/maotan/go-truffle/web"
	"github.com/maotan/go-truffle/yaml_config"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {

	registryClient, err:= web.ConsulInit(map[string]string{"client": "zyn3"})
	if err != nil{
		panic(err)
	}

	router := gin.Default()
	web.RouterInit(router)
	web.DatabaseInit()  // db
	router.GET("/client/ping", func(c *gin.Context) {

		pingDo := domain.PingDo{Name:"ping", Age: 12, Email: "gk@126.com"}
		res, err := feign_client.FeignGinServer.GinServerPingPost(pingDo)
		if err != nil{
			panic(truffle.NewWarnError(700, "123"))
		}
		c.String(res.StatusCode(), string(res.Body()))
	})

	router.GET("/client/users", func(c *gin.Context) {
		session := sessions.Default(c)
		userId := session.Get("userId")
		userDb := &model.User{}
		gosql.Model(userDb).Where("id=?", userId).Get()
		if userDb.Id == 0{
			panic(truffle.NewWarnError(40400, "不存在该用户"))
		}
		c.JSON(http.StatusOK, truffle.Success(userDb))
	})

	serverConf := yaml_config.YamlConf.ServerConf
	runHostPort := fmt.Sprintf(":%d", serverConf.Port)
	log.Info("app run...")
	err = router.Run(runHostPort)
	if err != nil{
		registryClient.Deregister()
	}
}