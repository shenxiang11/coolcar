module github.com/shenxiang11/coolcar/gateway

go 1.18

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.0
	github.com/shenxiang11/coolcar/auth-service v0.0.0-00010101000000-000000000000
	github.com/shenxiang11/coolcar/rental-service v0.0.0-00010101000000-000000000000
	github.com/shenxiang11/coolcar/shared v0.0.0-20220729003453-d71abf7487e4
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220719170305-83ca9fad585f // indirect
)

replace github.com/shenxiang11/coolcar/auth-service => ../auth

replace github.com/shenxiang11/coolcar/rental-service => ../rental

replace github.com/shenxiang11/coolcar/shared => ../shared

replace github.com/shenxiang11/coolcar/blob => ../blob
