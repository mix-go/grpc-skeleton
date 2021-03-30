module github.com/mix-go/grpc-skeleton

go 1.13

replace (
	github.com/mix-go/dotenv => ../mix/src/dotenv
	github.com/mix-go/xcli => ../mix/src/xcli
	github.com/mix-go/xdi => ../mix/src/xdi
)

require (
	github.com/go-redis/redis/v8 v8.8.0
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/configor v1.2.0
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.4 // indirect
	github.com/mix-go/dotenv v1.0.22
	github.com/mix-go/xcli v1.1.2
	github.com/mix-go/xdi v1.1.1
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.6.0
	google.golang.org/genproto v0.0.0-20191009194640-548a555dbc03 // indirect
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.23.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.3
)
