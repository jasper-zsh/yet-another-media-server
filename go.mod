module yet-another-media-server

go 1.16

require (
	github.com/juliangruber/go-intersect v1.1.0
	github.com/tebeka/selenium v0.9.9
	github.com/zeromicro/go-zero v1.4.3
	go.etcd.io/etcd/client/pkg/v3 v3.5.5 // indirect
	google.golang.org/grpc v1.50.1
	google.golang.org/protobuf v1.28.1
	gorm.io/driver/mysql v1.4.4
	gorm.io/driver/sqlite v1.4.3
	gorm.io/gorm v1.24.2
)

replace github.com/tebeka/selenium => ../selenium
