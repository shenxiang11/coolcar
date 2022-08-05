module github.com/shenxiang11/coolcar/blob

go 1.18

require (
	github.com/shenxiang11/coolcar/shared v0.0.0-00010101000000-000000000000
	github.com/tencentyun/cos-go-sdk-v5 v0.7.36
	go.mongodb.org/mongo-driver v1.10.1
	go.uber.org/zap v1.21.0
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/mozillazg/go-httpheader v0.2.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220719170305-83ca9fad585f // indirect
)

replace github.com/shenxiang11/coolcar/shared => ../shared

replace github.com/shenxiang11/coolcar/rental-service => ../rental
