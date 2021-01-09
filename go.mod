module gin-client

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-resty/resty/v2 v2.3.0
	github.com/maotan/go-truffle v1.1.6
	gopkg.in/resty.v1 v1.12.0 // indirect
)

replace github.com/maotan/go-truffle => ../go-truffle
