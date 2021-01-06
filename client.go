/**
* @Author: mo tan
* @Description:
* @Date 2021/1/1 21:58
 */
package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/maotan/go-truffle/feign"
	"github.com/maotan/go-truffle/truffle"
	"github.com/maotan/go-truffle/web"
	"github.com/maotan/go-truffle/yaml_config"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	
	err, registryClient:= web.ConsulInit(map[string]string{"client": "zyn3"})
	if err != nil{
		panic(err)
	}

	router := gin.Default()
	web.RouterInit(router)
	router.GET("/client/ping", func(c *gin.Context) {
		//instances, _  := registryDiscoveryClient.GetInstances("go-user-server")
		//fmt.Print(len(instances))
		res, err := feign.DefaultFeign.App("gin-server").R().SetHeaders(map[string]string{
			"Content-Type": "application/json",
		}).Get("/v1/ping")
		if err != nil{
			panic(truffle.NewWarnError(700, "123"))
		}
		c.String(res.StatusCode(), string(res.Body()))
	})

	router.GET("/client/users", func(c *gin.Context) {
		session := sessions.Default(c)
		u := session.Get("user")
		c.JSON(http.StatusOK, gin.H{"user":u})
	})

	serverConf := yaml_config.YamlConf.ServerConf
	runHostPort := fmt.Sprintf(":%d", serverConf.Port)
	log.Info("app run...")
	err = router.Run(runHostPort)
	if err != nil{
		registryClient.Deregister()
	}
}