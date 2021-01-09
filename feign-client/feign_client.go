/**
* @Author: mo tan
* @Description:
* @Date 2021/1/9 23:10
 */
package feign_client

import "github.com/maotan/go-truffle/feign"
import "github.com/go-resty/resty/v2"

const (
	feignAppName = "gin-server"
)

func GinServerPing() (res *resty.Response, e error){
	res, err := feign.GetRequest(feignAppName).Get("/v1/ping")
	return res, err
}