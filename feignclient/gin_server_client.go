// Package feignclient Package feigning /**
package feignclient

import (
	"gin-client/domain"

	"github.com/maotan/go-truffle/feign"
	"github.com/maotan/go-truffle/httpresult"
)

const (
	feignAppName = "gin-server"
)

//type GinServer struct {
//}
//var FeignGinServer = &GinServer{}

func GinServerPing() *httpresult.BaseResult {
	return feign.Get(feignAppName, "/v1/ping")
}

func GinServerPingPost(pingDo domain.PingDo) *httpresult.BaseResult {
	return feign.Post(feignAppName, "/v1/ping", pingDo)
}
