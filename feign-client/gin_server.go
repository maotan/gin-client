/**
* @Author: mo tan
* @Description:
* @Date 2021/1/9 23:10
 */
package feign_client

import (
	"gin-client/domain"
	"github.com/maotan/go-truffle/feign"
)
import "github.com/go-resty/resty/v2"

const (
	feignAppName = "gin-server"
)

type GinServer struct {

}

var FeignGinServer = &GinServer{}

func (f *GinServer) GinServerPing() (res *resty.Response, e error){
	res, err := feign.GetRequest(feignAppName).Get("/v1/ping")
	return res, err
}

func (f *GinServer) GinServerPingPost(pingDo domain.PingDo) (res *resty.Response, e error){
	res, err := feign.GetRequest(feignAppName).SetBody(pingDo).Post("/v1/ping/post")
	return res, err
}