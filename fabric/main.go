package fabric

import (
	"fabric-sdk-go/sdkInit"
	"fmt"
	"os"
)

const (
	cc_name    = "simplecc"
	cc_version = "1.0.0"
)

var App sdkInit.Application


func InitFabric ()  sdkInit.Application{
	orgs := []*sdkInit.OrgInfo{
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org1",
			OrgMspId:      "Org1MSP",
			OrgUser:       "User1",
			OrgPeerNum:    1,
			OrgAnchorFile: os.Getenv("GOPATH") + "/src/fabric-sdk-go/fixtures/channel-artifacts/Org1MSPanchors.tx",
		},
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org2",
			OrgMspId:      "Org2MSP",
			OrgUser:       "User2",
			OrgPeerNum:    1,
			OrgAnchorFile: os.Getenv("GOPATH") + "/src/fabric-sdk-go/fixtures/channel-artifacts/Org2MSPanchors.tx",
		},
	}

	// init sdk env info
	info := sdkInit.SdkEnvInfo{
		ChannelID:        "mychannel",
		ChannelConfig:    os.Getenv("GOPATH") + "/src/fabric-sdk-go/fixtures/channel-artifacts/channel.tx",
		Orgs:             orgs,
		OrdererAdminUser: "Admin",
		OrdererOrgName:   "OrdererOrg",
		OrdererEndpoint:  "orderer.example.com",
		ChaincodeID:      cc_name,
		ChaincodePath:    os.Getenv("GOPATH") + "/src/fabric-sdk-go/chaincode/",
		ChaincodeVersion: cc_version,
	}
	// content, err := os.ReadFile("/etc/hosts")
	// if err != nil {
	// 	fmt.Println(">> SDK setup error:", err)
	// 	os.Exit(-1)
	// }	
	// lines := strings.Split(string(content), "\n")
	// found := false 
	// for _, line := range lines {
  	// 	if strings.Contains(line, "orderer.example.com") {
    // 		found = true
    // 		break
  	// 	}
	// }
	// if !found {
	// 	f, err := os.OpenFile("/etc/hosts", os.O_APPEND|os.O_WRONLY, 0644)
	// 	if err != nil {
	// 		fmt.Println(">> SDK setup error:", err)
	// 		os.Exit(-1)
	// 	}	
	// 	defer f.Close()
		
	// 	_, err = f.WriteString("127.0.0.1 orderer.example.com \n 127.0.0.1 peer0.org1.example.com \n 127.0.0.1 peer0.org2.example.com \n 127.0.0.1 ca.org1.example.com \n 127.0.0.1 ipfs_host ")
	// 	if err != nil {
	// 		fmt.Println(">> SDK setup error:", err)
	// 		os.Exit(-1)
	// 	}
	// }

	// sdk setup
	sdk, err := sdkInit.Setup("config.yaml", &info)
	if err != nil {
		fmt.Println(">> SDK setup error:", err)
		os.Exit(-1)
	}

	// create channel and join
	if err := sdkInit.CreateAndJoinChannel(&info); err != nil {
		fmt.Println(">> Create channel and join error:", err)
		os.Exit(-1)
	}

	// create chaincode lifecycle
	if err := sdkInit.CreateCCLifecycle(&info, 1, false, sdk); err != nil {
		fmt.Println(">> create chaincode lifecycle error %v", err)
		os.Exit(-1)
	}

	// invoke chaincode set status
	fmt.Println(">> 通过链码外部服务设置链码状态......")

	if err := info.InitService(info.ChaincodeID, info.ChannelID, info.Orgs[0], sdk); err != nil {

		fmt.Println("InitService successful")
		os.Exit(-1)
	}

	App = sdkInit.Application{
		SdkEnvInfo: &info,
	}
	fmt.Println(">> 设置链码状态完成")
	
	defer info.EvClient.Unregister(sdkInit.BlockListener(info.EvClient))
	defer info.EvClient.Unregister(sdkInit.ChainCodeEventListener(info.EvClient, info.ChaincodeID))
	
	return App
}