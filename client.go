/**
* @Author: mo tan
* @Description:
* @Date 2021/1/1 21:58
 */
package main

import (
	"fmt"
	"gin-client/domain"
	"gin-client/feignclient"
	"gin-client/model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ilibs/gosql/v2"
	"github.com/maotan/go-truffle/httpresult"
	"github.com/maotan/go-truffle/web"
	"github.com/maotan/go-truffle/yamlconf"
	log "github.com/sirupsen/logrus"
)

func main() {

	registryClient, err := web.ConsulInit(map[string]string{"client": "zyn3"})
	if err != nil {
		panic(err)
	}

	router := gin.New()
	web.RouterInit(router)
	web.DatabaseInit() // db
	router.GET("/client/ping", func(ctx *gin.Context) {

		pingDo := domain.PingDo{Name: "ping", Age: 12, Email: "gk@126.com"}
		base := feignclient.GinServerPingPost(pingDo)
		if base.Code != httpresult.SuccessCode {
			panic(httpresult.NewWarnError(base.Code, base.Msg))
		}
		var pingRes domain.PingDo
		httpresult.Decode(base, &pingRes)
		ctx.JSON(200, httpresult.Success(pingRes))
		//ctx.String(res.StatusCode(), res.String())
	})

	router.GET("/client/users", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		userId := session.Get("userId")
		userDb := &model.User{}
		gosql.Model(userDb).Where("id=?", userId).Get()
		if userDb.Id == 0 {
			panic(httpresult.NewWarnError(40400, "不存在该用户"))
		}
		ctx.JSON(http.StatusOK, httpresult.Success(userDb))
	})

	serverConf := yamlconf.YamlConf.ServerConf
	runHostPort := fmt.Sprintf(":%d", serverConf.Port)
	log.Info("app run...")
	err = router.Run(runHostPort)
	if err != nil {
		registryClient.Deregister()
	}
}
