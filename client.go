package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maotan/go-truffle/cloud"
	"github.com/maotan/go-truffle/cloud/serviceregistry"
	"github.com/maotan/go-truffle/util"
	"log"
	"math/rand"
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

	si, _ := cloud.NewDefaultServiceInstance("go-client-server", ip, 6000,
		false, map[string]string{"client": "zyn3"}, "")

	registryDiscoveryClient.Register(si)

	r := gin.Default()
	r.GET("/actuator/health", func(c *gin.Context) {
		log.Print("1111111")
		log.Print( registryDiscoveryClient.GetServices())
		log.Print("222222")
		log.Print(registryDiscoveryClient.GetInstances("consul"))
		log.Print(registryDiscoveryClient.GetInstances("go-user-server"))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err = r.Run(":6000")
	if err != nil {
		registryDiscoveryClient.Deregister()
	}
}