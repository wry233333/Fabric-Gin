# Fabric-Gin
一个使用fabric-go-sdk启动fabric网络并提供Gin框架支持的demo应用

1.注意该项目应该放在 $GOPATH/src 下否则将无法启动

```
cd $GOPATH && mkdir src
cd $GOPATH/src && git clone https://github.com/wry233333/Fabric-Gin.git
```

2.启动命令
`./Start.sh`

3.关闭命令
`./Shutdown.sh`

4.链码存放路径
`chaincode/`

PS:如果启动提示链接终止无法连接到对等节点则进行以下操作
  `sudo vim /etc/hosts`

  将以下域名解析粘贴在hosts里
  ```
  127.0.0.1	orderer.example.com
  127.0.0.1	peer0.org1.example.com
  127.0.0.1	peer0.org2.example.com
  127.0.0.1	ca.org1.example.com
  127.0.0.1	ipfs_host
  ```
  然后进行第二步重试
  
  
  
