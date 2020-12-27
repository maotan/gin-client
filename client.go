package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maotan/go-truffle/cloud"
	"github.com/maotan/go-truffle/cloud/serviceregistry"
	"github.com/maotan/go-truffle/truffle"
	"github.com/maotan/go-truffle/util"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	host := "127.0.0.1"
	port := 8500
	token := ""
	registryDiscoveryClient, err := serviceregistry.NewConsulServiceRegistry(host, port, token)
	if err != nil {
		panic(err)
	}

	ip, err := util.GetLocalIP()
	if err != nil {
		panic(err)
	}

	fmt.Println(ip)
	rand.Seed(time.Now().UnixNano())

	si, _ := cloud.NewDefaultServiceInstance("go-client-server", ip, 9000,
		false, map[string]string{"client": "zyn3"}, "")

	registryDiscoveryClient.Register(si)

	r := gin.Default()
	r.GET("/actuator/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/client/ping", func(c *gin.Context) {
		instances, err  := registryDiscoveryClient.GetInstances("go-user-server")
		if err != nil{
			panic(truffle.NewWarnError(600 ,"找不到server"))
		}
		instance := instances[0]
		url := fmt.Sprintf("http://%s:%d/%s",instance.GetHost(), instance.GetPort(), "v2/ping")

		fmt.Println("url:", url)
		resp, err := http.Get (url)
		if err!=nil{

		}
		dd, _:= truffle.ParseResponse(resp)
		c.JSON(resp.StatusCode, dd)
	})
	r.Use(truffle.Recover)
	r.Run(":9001")
}