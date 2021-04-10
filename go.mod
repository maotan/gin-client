module gin-client

go 1.16

require (
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.6.3
	github.com/go-resty/resty/v2 v2.3.0
	github.com/ilibs/gosql/v2 v2.1.0
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/maotan/go-truffle v1.1.6
	github.com/sirupsen/logrus v1.7.0
	github.com/tebeka/strftime v0.1.5 // indirect
)

replace github.com/maotan/go-truffle => ../go-truffle
