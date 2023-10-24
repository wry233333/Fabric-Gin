package main

import (
	"fabric-sdk-go/fabric"
	"fabric-sdk-go/server"
	"fabric-sdk-go/redis"
)



func main() {
	// init fabric network and install chaincode
	App := fabric.InitFabric()
	// init redis client
	rc := redis.InitRedis()
	// test redis
	rc.Set("111","111")
	// start server
	server.ServerInit(App)
}

