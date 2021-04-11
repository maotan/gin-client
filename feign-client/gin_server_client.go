/**
* @Author: mo tan
* @Description:
* @Date 2021/1/9 23:10
 */
package feign_client

import (
	"gin-client/domain"

	"github.com/maotan/go-truffle/feign"
	"github.com/maotan/go-truffle/httpresult"
)

const (
	feignAppName = "gin-server"
)

type GinServer struct {
}

var FeignGinServer = &GinServer{}

func (f *GinServer) GinServerPing() *httpresult.BaseResult {
	return feign.Get(feignAppName, "/v1/ping")
}

func (f *GinServer) GinServerPingPost(pingDo domain.PingDo) *httpresult.BaseResult {
	//res, err := feign.GetRequest(feignAppName).SetBody(pingDo).SetResult(&httpresult.BaseResult{}).Post("/v1/ping")
	//if err != nil {
	//	panic(httpresult.NewWarnError(500, err.Error()))
	//}
	//base := res.Result().(*httpresult.BaseResult)
	return feign.Post(feignAppName, "/v1/ping", pingDo)
}
