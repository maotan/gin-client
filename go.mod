module gin-client

go 1.16

require (
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.6.3
	github.com/go-resty/resty/v2 v2.3.0
	github.com/goinggo/mapstructure v0.0.0-20140717182941-194205d9b4a9
	github.com/ilibs/gosql/v2 v2.1.0
	github.com/maotan/go-truffle v1.1.7-0.20210411082126-1014448b8445
	github.com/sirupsen/logrus v1.7.0
)

replace github.com/maotan/go-truffle => ../go-truffle
