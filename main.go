package main

import (
	"fabric-sdk-go/fabric"
	"fabric-sdk-go/server"
)



func main() {
	// init fabric network and install chaincode
	App := fabric.InitFabric()
	// start server
	server.ServerInit(App)
}

