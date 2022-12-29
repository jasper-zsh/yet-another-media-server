module yet-another-media-server

go 1.16

require (
	github.com/fogleman/gg v1.3.0 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/juliangruber/go-intersect v1.1.0
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646 // indirect
	github.com/tebeka/selenium v0.9.9
	github.com/zeromicro/go-queue v1.1.8 // indirect
	github.com/zeromicro/go-zero v1.4.3
	go.etcd.io/etcd/client/pkg/v3 v3.5.5 // indirect
	golang.org/x/image v0.2.0 // indirect
	google.golang.org/grpc v1.50.1
	google.golang.org/protobuf v1.28.1
	gorm.io/driver/mysql v1.4.4
	gorm.io/driver/sqlite v1.4.3
	gorm.io/gorm v1.24.2
)

replace github.com/tebeka/selenium => ../selenium
